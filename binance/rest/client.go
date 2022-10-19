package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/buger/jsonparser"
)

const (
	BaseEndpoint  = "https://api.binance.com"
	BaseEndpoint1 = "https://api1.binance.com"
	BaseEndpoint2 = "https://api2.binance.com"
	BaseEndpoint3 = "https://api3.binance.com"
)

type Configuration struct {
	Addr      string `json:"addr"`
	ApiKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
	DebugMode bool   `json:"debug_mode"`
}

type BinanceRest struct {
	cfg *Configuration
}

func (b *BinanceRest) GetPair(args ...string) string {
	pair := args[0] + args[1]

	return strings.ToUpper(pair)
}
func New(config *Configuration) *BinanceRest {

	// 	потом тут добавятся различные другие настройки
	b := &BinanceRest{
		cfg: config,
	}
	return b
}

func (ex *BinanceRest) GetSign(parm string) string {
	secretKey := ex.cfg.SecretKey

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(parm))
	sign := hex.EncodeToString(mac.Sum(nil))
	return sign
}

func (ex *BinanceRest) ConnWithHeader(method string, endpoint string, parms string) []byte {

	apiKey := ex.cfg.ApiKey
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := ex.cfg.Addr + endpoint + "?" + parms

	req, err := http.NewRequest(method, url, nil)

	req.Header.Set("X-MBX-APIKEY", apiKey)
	response, err := client.Do(req)
	if err != nil {
		log.Printf(`
			{
				"Status" : "Error",
				"Path to file" : "CCXT_beYANG_Binance/binance/rest",
				"File": "client.go",
				"Functions" : "(ex *BinanceRest) withHeader(method string, endpoint string, parms string) ",
				"Function where err" : "client.Do",
				"Data": [%v],
				"Exchange" : "Binance",
				"Error" : %s
			}`, req, err)
		log.Fatal()
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf(`
			{
				"Status" : "Error",
				"Path to file" : "CCXT_beYANG_Binance/binance/rest",
				"File": "client.go",
				"Functions" : "(ex *BinanceRest) withHeader(method string, endpoint string, parms string) ",
				"Function where err" : "io.ReadAll",
				"Data": [%v],
				"Exchange" : "Binance",
				"Error" : %s
			}`, response.Body, err)
		log.Fatal()
	}
	if ex.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\tBinanceWalletBalance %v", string(data))
	}
	return data
}

func (ex *BinanceRest) isErr(data []byte) bool {
	code, _ := jsonparser.GetInt(data, "code")
	if code != 0 {
		return true
	} else {
		return false
	}
}
