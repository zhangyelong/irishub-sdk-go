package nets

import (
	"io/ioutil"
	"net/http"
	"time"
)

// HttpProvider
type HttpProvider struct {
	httpClient *http.Client
}

// NewHttpProvider initialize a http provider with specifying the timeout in seconds
func NewHttpProvider(timeout int) *HttpProvider {
	t := time.Duration(timeout)
	return &HttpProvider{
		httpClient: &http.Client{Timeout: t * time.Second},
	}
}

// HttpGet execute http Get request, returns []byte or error
func (provider HttpProvider) HttpGet(url string) ([]byte, error)  {
	resp, err := provider.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bz, nil
}