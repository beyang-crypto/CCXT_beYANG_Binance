package v3

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

// для создания собственных json файлов и преобразования json в структуру

func (ex *BinanceWS) GetBalance() WalletBalance {
	//	https://binance-exchange.github.io/docs/spot/?python--pybit#t-wallet
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
	url := "https://api.binance.com/api/v3/account?" + parms
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
