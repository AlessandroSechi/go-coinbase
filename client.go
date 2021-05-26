package coinbase

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

func NewClient(APIKey string, APISecret string) (*Client) {

	return &Client{
		HTTPClient:   &http.Client{},
		APIKey: APIKey,
		APISecret:   APISecret,
		APIBase:  APIBase,
	}
}

// SetLog will set/change the output destination.
func (c *Client) SetLog(log io.Writer) {
	c.Log = log
}

// Send makes a request to the API, the response body will be
// unmarshaled into v, or if v is an io.Writer, the response will
// be written to it without decoding
func (c *Client) Send(req *http.Request, v ...interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// Set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")

	// Default values for headers
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	resp, err = c.HTTPClient.Do(req)
	c.log(req, resp)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &ErrorResponse{Response: resp}
		data, err = ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}

		return errResp
	}
	if v == nil {
		return nil
	}

	if w, ok := v[0].(io.Writer); ok {
		io.Copy(w, resp.Body)
		return nil
	}

	r := Response{}

	json.NewDecoder(resp.Body).Decode(&r)

	dataByte, _ := json.Marshal(&r.Data)
	_ = json.Unmarshal(dataByte, &v[0])

	if len(v) > 1{
		paginationByte, _ := json.Marshal(&r.Pagination)
		_ = json.Unmarshal(paginationByte, &v[1])
	}

	return nil
}

// SendWithAuth makes a request to the API and apply Authentication headers automatically.
func (c *Client) SendWithAuth(req *http.Request, v ...interface{}) error {

	timestamp := time.Now().Unix()

	req.Header.Set("CB-ACCESS-KEY", c.APIKey)

	req.Header.Set("CB-ACCESS-SIGN", c.generateSignature(req, timestamp))

	req.Header.Set("CB-ACCESS-TIMESTAMP", fmt.Sprintf("%d", timestamp))

	return c.Send(req, v...)
}

// NewRequest constructs a request
// Convert payload to a JSON
func (c *Client) NewRequest(ctx context.Context, method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequestWithContext(ctx, method, url, buf)
}

// log will dump request and response to the log file
func (c *Client) log(r *http.Request, resp *http.Response) {
	if c.Log != nil {
		var (
			reqDump  string
			respDump []byte
		)

		if r != nil {
			reqDump = fmt.Sprintf("%s %s. Data: %s", r.Method, r.URL.String(), r.Form.Encode())
		}
		if resp != nil {
			respDump, _ = httputil.DumpResponse(resp, true)
		}

		c.Log.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
	}
}

// generateSignature will generate proper signature https://developers.coinbase.com/api/v2#api-key
func (c *Client) generateSignature(req *http.Request, timestamp int64) string {

	nonce := fmt.Sprintf("%d", timestamp)

	var body string

	if req.Body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		body = buf.String()
	}

	message := nonce + req.Method + req.URL.Path + body //As per Coinbase Documentation

	h := hmac.New(sha256.New, []byte(c.APISecret))
	h.Write([]byte(message))

	signature := hex.EncodeToString(h.Sum(nil))

	return signature
}
