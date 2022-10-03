package httpClient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	TimeOffset int64
}

// NewClient create nre http client
func NewClient(baseURL string) *Client {
	return &Client{
		UserAgent:  "simplegrpc/frontend", //target_1
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
	}
}

// get the request and return it like []byte
func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() error {
		return resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// MakeReq HTTP request helper
func (c *Client) MakeReq(apiUrl string, params url.Values, method string, target ...string) ([]byte, error) {
	myURL := fmt.Sprintf("%s/%s", c.BaseURL, apiUrl)
	for _, v := range params {
		{
			myURL = fmt.Sprintf("%s/%s", myURL, v[0])
		}
	}
	fmt.Println("REQUEST TO ", myURL, " server: ", target[0])
	req, err := http.NewRequest(method, myURL, nil)
	if err != nil {
		return nil, err
	}
	if len(target) > 0 {
		req.Header.Set(target[0], target[0])
	}
	resp, err := doReq(req, c.HTTPClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}
