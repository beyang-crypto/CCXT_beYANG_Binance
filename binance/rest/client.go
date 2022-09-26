package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

const (
	RestBaseEndpoint  = "https://api.binance.com"
	RestBaseEndpoint1 = "https://api1.binance.com"
	RestBaseEndpoint2 = "https://api2.binance.com"
	RestBaseEndpoint3 = "https://api3.binance.com"
)

type Configuration struct {
	Addr           string `json:"addr"`
	ApiKey         string `json:"api_key"`
	SecretKey      string `json:"secret_key"`
	APIKeyPassword string `json:"passphrase"`
	DebugMode      bool   `json:"debug_mode"`
}

type BinanceRest struct {
	cfg *Configuration
}

func (b *BinanceRest) GetPair(coin1 string, coin2 string) string {
	return coin1 + coin2
}

func New(config *Configuration) *BinanceRest {

	// 	потом тут добавятся различные другие настройки
	b := &BinanceRest{
		cfg: config,
	}
	return b
}

func (ex *BinanceRest) GetBalance() WalletBalance {
	//	получение времяни
	ts := time.Now().UTC().Unix() * 1000
	apiKey := ex.cfg.ApiKey
	secretKey := ex.cfg.SecretKey

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	parms := fmt.Sprintf("timestamp=%d", ts)
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(fmt.Sprintf("timestamp=%d", ts)))
	parms += "&signature=" + hex.EncodeToString(mac.Sum(nil))
	//	реализация метода GET
	url := ex.cfg.Addr + "/api/v3/account?" + parms
	log.Printf(url)
	req, err := http.NewRequest("GET", url, nil)

	//	код для вывода полученных данных
	if err != nil {
		log.Fatalln(err)
	}
	//	у бинанса апи ключ надо передавать через заголовки
	req.Header.Set("X-MBX-APIKEY", apiKey)
	response, err := client.Do(req)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if ex.cfg.DebugMode {
		log.Printf("BinanceWalletBalance %v", string(data))
	}

	// {
	// 	"ret_code": 0,
	// 	"ret_msg": "",
	// 	"ext_code": null,
	// 	"ext_info": null,
	// 	"result": {
	// 		"balances": [
	// 			{
	// 				"coin": "USDT",
	// 				"coinId": "USDT",
	// 				"coinName": "USDT",
	// 				"total": "10",
	// 				"free": "10",
	// 				"locked": "0"
	// 			}
	// 		]
	// 	}
	// }

	var walletBalance WalletBalance
	err = json.Unmarshal(data, &walletBalance)
	if err != nil {
		log.Printf(`
			{
				"Status" : "Error",
				"Path to file" : "CCXT_BEYANG_BYBIT/spot/v3",
				"File": "api.go",
				"Functions" : "(ex *BinanceWS) GetBalance() (WalletBalance)",
				"Function where err" : "json.Unmarshal",
				"Exchange" : "Binance",
				"Comment" : %s to WalletBalance struct,
				"Error" : %s
			}`, string(data), err)
		log.Fatal()
	}

	return walletBalance

}
