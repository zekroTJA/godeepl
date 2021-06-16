package godeepl

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

func (c *Client) jsonRpcRequest(endpoint, method string, reqBody interface{}, respBody interface{}) (err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if endpoint == "" {
		endpoint = c.options.Endpoint
	}

	req.Header.SetMethod("POST")
	req.SetRequestURI(fmt.Sprintf("%s?method=%s", endpoint, method))
	req.Header.SetContentType("application/json")

	reqBodyB, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	req.SetBodyRaw(reqBodyB)

	if c.options.SessionID == "" && c.options.Email != "" && c.options.Password != "" && method != "login" {
		if _, err = c.Login(c.options.Email, c.options.Password, true); err != nil {
			return
		}
	}

	if c.options.SessionID != "" {
		req.Header.SetCookie("dl_session", c.options.SessionID)
	}

	if err = c.client.Do(req, resp); err != nil {
		return
	}

	if resp.StatusCode() >= 400 {
		return ResponseError{resp.StatusCode()}
	}

	err = json.Unmarshal(resp.Body(), respBody)
	return
}
