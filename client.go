package godeepl

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

type ClientOptions struct {
	Endpoint  string `json:"endpoint"`
	SessionID string `json:"session_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

var defaultConfig = ClientOptions{
	Endpoint: EndpointPublic,
}

type Client struct {
	options *ClientOptions
	client  *fasthttp.Client
}

func New(options ClientOptions) *Client {
	if options.Endpoint == "" {
		options.Endpoint = EndpointPublic
	}

	return &Client{
		options: &options,
		client: &fasthttp.Client{
			Name: "godeepl",
		},
	}
}

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

func (c *Client) Login(email, password string, keepLogin bool) (res *LoginResult, err error) {
	res = &LoginResult{}

	err = c.jsonRpcRequest(
		EndpointLogin, methodLogin,
		NewJsonRpcRequest(methodLogin, &LoginParams{
			Email:     email,
			Password:  password,
			KeepLogin: keepLogin,
		}),
		&JsonRpcResponse{Result: res})

	if err == nil {
		c.options.SessionID = res.Token
	}

	return
}

func (c *Client) SplitSentence(lang LangSpec, text ...string) (res *SplitSentenceResult, err error) {
	res = &SplitSentenceResult{}

	err = c.jsonRpcRequest(
		"", methodSplitSentence,
		NewJsonRpcRequest(methodSplitSentence, &SplitSentenceParams{
			Texts: text,
			Lang: &Lang{
				LangUserSelected: lang,
			},
		}),
		&JsonRpcResponse{Result: res})

	return
}

func (c *Client) Translate(sourceLang, targetLang LangSpec, text string) (res *TranslationResult, err error) {
	res = &TranslationResult{}

	splitRes, err := c.SplitSentence(sourceLang, text)
	if err != nil {
		return
	}

	if splitRes.LangIsConfident {
		sourceLang = splitRes.Lang
	}

	if len(splitRes.SplittedTexts) == 0 || len(splitRes.SplittedTexts[0]) == 0 {
		err = errors.New("empty sentence")
		return
	}

	sentences := splitRes.SplittedTexts[0]
	jobs := make([]interface{}, len(sentences))
	for i, s := range sentences {
		jobs[i] = &TranslateJob{
			Kind:               "default",
			RawEnSentence:      s,
			RawEnContextBefore: sentences[0:i],
			RawEnContextAfter:  sentences[i+1:],
		}
	}

	err = c.jsonRpcRequest(
		"", methodTranslate,
		NewJsonRpcRequest(methodTranslate, &TranslateParams{
			Priority: -1,
			Lang: &Lang{
				SourceLangComputed: sourceLang,
				TargetLang:         targetLang,
				UserPreferredLangs: []LangSpec{},
			},
			Jobs: jobs,
			CommonJobParams: &CommonJobParams{
				Formality: "formal",
			},
		}),
		&JsonRpcResponse{Result: res})

	return
}
