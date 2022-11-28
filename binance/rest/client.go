package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	Addr      string
	ApiKey    string
	SecretKey string
	DebugMode bool
}

type BinanceRest struct {
	cfg *Configuration
	c   *http.Client
}

func (c *BinanceRest) debug(format string, v ...interface{}) {
	if c.cfg.DebugMode {
		log.Printf(format, v...)
	}
}

func (c *BinanceRest) parseRequest(r *request, opts ...RequestOption) (err error) {
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.cfg.Addr, r.endpoint)

	if r.secType == secTypeSigned {
		r.setParam("timestamp", time.Now().UTC().Unix()*1000)
	}
	queryString := r.query.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}

	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set("X-MBX-APIKEY", c.cfg.ApiKey)
	}

	if r.secType == secTypeSigned {
		sign := c.GetSign(queryString)

		v := url.Values{}
		v.Set("signature", sign)
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s", fullURL)

	r.fullURL = fullURL
	r.header = header
	return nil
}

func (c *BinanceRest) callAPI(r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, nil)
	if err != nil {
		return []byte{}, err
	}
	//req = req.WithContext(ctx)
	req.Header = r.header
	//c.debug("request: %#v", req)

	f := c.c.Do

	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	//c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		// apiErr := new(common.APIError)
		// e := json.Unmarshal(data, apiErr)
		// if e != nil {
		// 	c.debug("failed to unmarshal json: %s", e)
		// }
		// return nil, apiErr
		log.Fatal()
	}
	return data, nil
}

func (b *BinanceRest) GetPair(args ...string) string {
	pair := args[0] + args[1]

	return strings.ToUpper(pair)
}
func New(config *Configuration) *BinanceRest {

	// 	потом тут добавятся различные другие настройки
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	b := &BinanceRest{
		cfg: config,
		c:   client,
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
	url := ex.cfg.Addr + endpoint + "?" + parms

	req, err := http.NewRequest(method, url, nil)
	if ex.cfg.DebugMode {
		log.Printf(url)
	}

	req.Header.Set("X-MBX-APIKEY", apiKey)
	response, err := ex.c.Do(req)
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
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\t %v", string(data))
	}
	return data
}

func (ex *BinanceRest) ConnWithoutHeader(method string, endpoint string, parms string) []byte {
	url := ex.cfg.Addr + endpoint + "?" + parms
	fmt.Print(method + " " + url)

	req, err := http.NewRequest(method, url, nil)
	response, err := ex.c.Do(req)
	if err != nil {
		log.Printf(`
			{
				"Status" : "Error",
				"Path to file" : "CCXT_beYANG_Binance/binance/rest",
				"File": "client.go",
				"Functions" : "(ex *BinanceRest) ConnWithoutHeader(method string, endpoint string, parms string) []byte",
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
				"Functions" : "(ex *BinanceRest) ConnWithoutHeader(method string, endpoint string, parms string) []byte",
				"Function where err" : "io.ReadAll",
				"Data": [%v],
				"Exchange" : "Binance",
				"Error" : %s
			}`, response.Body, err)
		log.Fatal()
	}
	if ex.cfg.DebugMode {
		log.Printf("STATUS: DEBUG\tEXCHANGE: Binance\tAPI: Rest\t %v", string(data))
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
