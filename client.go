package gonhl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseAddress = "http://statsapi.web.nhl.com/api/v1/"

// A Client is required for api calls
type Client struct {
	baseURL string

	httpClient *http.Client
}

func NewClient() *Client {
	customHttp := &http.Client{
		Timeout:   time.Second * 10,
		Transport: buildCustomTransport(),
	}

	return &Client{
		baseURL:    baseAddress,
		httpClient: customHttp,
	}
}

func (c *Client) makeRequest(endpoint string, params map[string]string, schema interface{}) int {
	body, status := c.makeRequestWithoutJson(endpoint, params)
	json.Unmarshal(body, schema)
	return status
}

func (c *Client) makeRequestWithoutJson(endpoint string, params map[string]string) ([]byte, int) {
	request, _ := http.NewRequest("GET", c.baseURL+endpoint, nil)
	request.Header.Set("Content-Type", "application/json")
	query := request.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
	//fmt.Println(request.URL)
	response, _ := c.httpClient.Do(request)
	//check(err)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body, response.StatusCode
}

func buildCustomTransport() *http.Transport {
	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, success := defaultRoundTripper.(*http.Transport)
	if !success {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}
	defaultTransport := *defaultTransportPointer
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	return &defaultTransport
}
