package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BearerTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	DoCallForBearerToken(req *http.Request) (*BearerTokenResponse, error)
	DoGet(url string, headers map[string]string) (*http.Response, error)
	DoPost(url string, headers map[string]string, payload any) (*http.Response, error)
}

type httpClient struct {
	client     *http.Client
	postMethod string
	getMethod  string
}

func NewHTTPClient(cl *http.Client) HTTPClient {
	return &httpClient{client: cl, postMethod: http.MethodPost, getMethod: http.MethodGet}
}

func (hc *httpClient) Do(req *http.Request) (*http.Response, error) {
	return hc.client.Do(req)
}

func (hc *httpClient) DoCallForBearerToken(req *http.Request) (*BearerTokenResponse, error) {
	res, err := hc.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("recieved %d code from client", res.StatusCode)
	}

	var tRes BearerTokenResponse
	err = json.NewDecoder(res.Body).Decode(&tRes)
	if err != nil {
		return nil, err
	}

	return &tRes, nil
}

func (hc *httpClient) DoGet(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(hc.getMethod, url, nil)
	if err != nil {
		return nil, err
	}

	for key, element := range headers {
		req.Header.Set(key, element)
	}
	return hc.client.Do(req)
}

func (hc *httpClient) DoPost(url string, headers map[string]string, payload any) (*http.Response, error) {
	jsonValue, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(hc.postMethod, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	for key, element := range headers {
		req.Header.Set(key, element)
	}
	return hc.client.Do(req)
}
