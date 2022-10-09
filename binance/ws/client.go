package ws

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/buger/jsonparser"      //  Для вытаскивания одного значения из файла json
	"github.com/chuckpreslar/emission" // Эмитер необходим для удобного выполнения функции в какой-то момент
	"github.com/goccy/go-json"         // для создания собственных json файлов и преобразования json в структуру
	"github.com/gorilla/websocket"
)

const (
	HostMainnetPublicTopics = "wss://stream.binance.com:9443/ws"
)

const (
	ChannelTicker = "@bookTicker"
)

type Configuration struct {
	Addr      string `json:"addr"`
	ApiKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
	DebugMode bool   `json:"debug_mode"`
}

type BinanceWS struct {
	cfg  *Configuration
	conn *websocket.Conn

	mu            sync.RWMutex
	subscribeCmds []Cmd //	сохраняем все подписки у данной биржи, чтоб при переподключении можно было к ним повторно подключиться

	emitter *emission.Emitter
}

func (b *BinanceWS) GetPair(args ...string) string {
	pair := args[0] + args[1]

	return strings.ToLower(pair)
}

func New(config *Configuration) *BinanceWS {

	// 	потом тут добавятся различные другие настройки
	b := &BinanceWS{
		cfg:     config,
		emitter: emission.NewEmitter(),
	}
	return b
}

func (b *BinanceWS) Subscribe(channel string, coins []string) {
	var parms []string
	for _, value := range coins {
		parms = append(parms, value+channel)
	}
	cmd := Cmd{
		Method: "SUBSCRIBE",
		Params: parms,
		ID:     804218,
	}
	b.subscribeCmds = append(b.subscribeCmds, cmd)
	if b.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: WS\tСоздание json сообщения на подписку part 1")
	}
	b.SendCmd(cmd)
}

//	отправка команды на сервер в отдельной функции для того, чтобы при переподключении быстро подписаться на все предыдущие каналы
func (b *BinanceWS) SendCmd(cmd Cmd) {
	data, err := json.Marshal(cmd)
	if err != nil {
		log.Printf(`
			{
				"Status" : "Error",
				"Path to file" : "CCXT_BEYANG_BINANCE/binance",
				"File": "client.go",
				"Functions" : "(b *BinanceWS) sendCmd(cmd Cmd)",
				"Function where err" : "json.Marshal",
				"Exchange" : "Binance",
				"Data" : [%v],
				"Error" : %s
			}`, cmd, err)
		log.Fatal()
	}
	if b.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: WS\tСоздание json сообщения на подписку part 2")
	}
	b.Send(string(data))
}

func (b *BinanceWS) Send(msg string) (err error) {
	defer func() {
		// recover необходим для корректной обработки паники
		if r := recover(); r != nil {
			if err != nil {
				log.Printf(`
					{
						"Status" : "Error",
						"Path to file" : "CCXT_BEYANG_BINANCE/binance",
						"File": "client.go",
						"Functions" : "(b *BinanceWS) Send(msg string) (err error) ",
						"Function where err" : "b.conn.WriteMessage",
						"Exchange" : "Binance",
						"Data" : [websocket.TextMessage, %s],
						"Error" : %s,
						"Recover" : %v
					}`, msg, err, r)
				log.Fatal()
			}
			err = errors.New(fmt.Sprintf("BinanceWs send error: %v", r))
		}

	}()
	if b.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: WS\tОтправка сообщения на сервер. текст сообщения:%s", msg)
	}

	err = b.conn.WriteMessage(websocket.TextMessage, []byte(msg))
	return
}

// подключение к серверу и постоянное чтение приходящих ответов
func (b *BinanceWS) Start() error {
	if b.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: WS\tНачало подключения к серверу")
	}
	b.connect()

	cancel := make(chan struct{})

	go func() {
		t := time.NewTicker(time.Second * 15)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				// b.ping()
			case <-cancel:
				return
			}
		}
	}()

	go func() {
		defer close(cancel)

		for {
			_, data, err := b.conn.ReadMessage()
			if err != nil {

				if websocket.IsCloseError(err, 1006) {
					b.closeAndReconnect()
					//Необходим вызв SubscribeToTicker в отдельной горутине, рекурсия, думаю, тут неуместна
					log.Printf("Status: INFO	ошибка 1006 начинается переподключение к серверу")

				} else {
					b.conn.Close()
					log.Printf(`
						{
							"Status" : "Error",
							"Path to file" : "CCXT_BEYANG_BINANCE/binance",
							"File": "client.go",
							"Functions" : "(b *BinanceWS) Start() error",
							"Function where err" : "b.conn.ReadMessage",
							"Exchange" : "Binance",
							"Error" : %s
						}`, err)
					log.Fatal()
				}
			} else {
				b.messageHandler(data)
			}
		}
	}()

	return nil
}

func (b *BinanceWS) connect() {

	c, _, err := websocket.DefaultDialer.Dial(b.cfg.Addr, nil)
	if err != nil {
		log.Printf(`{
						"Status" : "Error",
						"Path to file" : "CCXT_BEYANG_BINANCE/binance",
						"File": "client.go",
						"Functions" : "(b *BinanceWS) connect()",
						"Function where err" : "websocket.DefaultDialer.Dial",
						"Exchange" : "Binance",
						"Data" : [%s, nil],
						"Error" : %s
					}`, b.cfg.Addr, err)
		log.Fatal()
	}
	b.conn = c
	for _, cmd := range b.subscribeCmds {
		b.SendCmd(cmd)
	}
}

func (b *BinanceWS) closeAndReconnect() {
	b.conn.Close()
	b.connect()
}

func (b *BinanceWS) ping() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("BinanceWs ping error: %v", r)
		}
	}()

	//	https://docs.binance.com/?python#websocket-api
	err := b.conn.WriteMessage(websocket.TextMessage, []byte(`pong`))
	if err != nil {
		log.Printf("BinanceWs ping error: %v", err)
	}
}

func (b *BinanceWS) messageHandler(data []byte) {

	if b.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: WS\tBinanceWs %v", string(data))
	}

	//	в ошибке нет необходимости, т.к. она выходит каждый раз, когда не найдет элемент
	eventType, _ := jsonparser.GetString(data, "e")

	switch eventType {
	//	принимаемые запросы от тикера и от подписок не имеют eventType, однако все остальные имеют
	default:
		u, _ := jsonparser.GetInt(data, "u")
		if u > 0 {
			var bookticker BookTicker
			err := json.Unmarshal(data, &bookticker)
			if err != nil {
				log.Printf(`
					{
						"Status" : "Error",
						"Path to file" : "CCXT_BEYANG_BINANCE/binance",
						"File": "client.go",
						"Functions" : "(b *BinanceWS) messageHandler(data []byte)",
						"Function where err" : "json.Unmarshal",
						"Exchange" : "Binance",
						"Comment" : %s to BookTicker struct,
						"Error" : %s
					}`, string(data), err)
				log.Fatal()
			}
			b.processBookTicker("Binance", bookticker.S, bookticker)
		} else {
			id, _ := jsonparser.GetInt(data, "id")
			switch id {
			case 804218:

			case 804219: //	получение нынешних подписок - использую вместо ping

			default:
				log.Printf(`
				{
					"Status" : "INFO",
					"Path to file" : "CCXT_BEYANG_BINANCE/binance",
					"File": "client.go",
					"Functions" : "(b *BinanceWS) messageHandler(data []byte)",
					"Exchange" : "Binance",
					"Comment" : "Ответ от неизвестного канала"
					"Message" : %s
				}`, string(data))
				log.Fatal()
			}

		}
	}
}
