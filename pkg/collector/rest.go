package collector

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)


type Сollector struct{
	url string
	api_key string
	apiSecret string
	recv_window  string
	client *http.Client
}

func CreateСollector(url string, api_key string, apiSecret string, recv_window string) *Сollector{
	return &Сollector{
		url: url,
		api_key: api_key,
		apiSecret: apiSecret,
		recv_window: recv_window,
	}
}

func (c *Сollector) HttpClient() {
	c.client = &http.Client{Timeout: 10 * time.Second}
}

func (c *Сollector) GetRequest(params string, endPoint string) []byte {
	now := time.Now()
	unixNano := now.UnixNano()
	time_stamp := unixNano / 1000000
	hmac256 := hmac.New(sha256.New, []byte(c.apiSecret))
	hmac256.Write([]byte(strconv.FormatInt(time_stamp, 10) + c.api_key + c.recv_window + params))
	signature := hex.EncodeToString(hmac256.Sum(nil))
	request, err := http.NewRequest("GET", c.url + endPoint + "?" + params, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-BAPI-API-KEY", c.api_key)
	request.Header.Set("X-BAPI-SIGN", signature)
	request.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(time_stamp, 10))
	request.Header.Set("X-BAPI-SIGN-TYPE", "2")
	request.Header.Set("X-BAPI-RECV-WINDOW", c.recv_window)
	reqDump, err := httputil.DumpRequestOut(request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Request Dump:\n%s", string(reqDump))
	response, error := c.client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	elapsed := time.Since(now).Seconds()
	fmt.Printf("\n%s took %v seconds \n", endPoint, elapsed)
	body, _ := ioutil.ReadAll(response.Body)
	return body
}

func (c *Сollector) PostRequest(params interface{}, endPoint string) []byte {
	now := time.Now()
	unixNano := now.UnixNano()
	time_stamp := unixNano / 1000000
	jsonData, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
	}
	hmac256 := hmac.New(sha256.New, []byte(c.apiSecret))
	hmac256.Write([]byte(strconv.FormatInt(time_stamp, 10) + c.api_key + c.recv_window + string(jsonData[:])))
	signature := hex.EncodeToString(hmac256.Sum(nil))
	request, err := http.NewRequest("POST", c.url + endPoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-BAPI-API-KEY", c.api_key)
	request.Header.Set("X-BAPI-SIGN", signature)
	request.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(time_stamp, 10))
	request.Header.Set("X-BAPI-SIGN-TYPE", "2")
	request.Header.Set("X-BAPI-RECV-WINDOW", c.recv_window)
	reqDump, err := httputil.DumpRequestOut(request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Request Dump:\n%s", string(reqDump))
	response, error := c.client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	elapsed := time.Since(now).Seconds()
	fmt.Printf("\n%s took %v seconds \n", endPoint, elapsed)
	body, _ := ioutil.ReadAll(response.Body)
	return body
}