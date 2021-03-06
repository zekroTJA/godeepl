package godeepl

import (
	"errors"

	"github.com/valyala/fasthttp"
)

// ClientOptions wraps a set of options passed to the
// client.
type ClientOptions struct {
	// The API endpoint which is used for all requests instead of the
	// login request, which always goes to EndpointLogin.
	//
	// When no value is specified, EndpointPublic is set as default
	// endpoint.
	Endpoint string `json:"endpoint"`

	// You can pass a session ID from a pre-logged-in
	// browser session.
	SessionID string `json:"session_id"`

	// When specified with a Password, a Login is performed and the
	// obtained session token is stored in the client for subsequent
	// requests.
	Email string `json:"email"`

	// When specified with an Email, a Login is performed and the
	// obtained session token is stored in the client for subsequent
	// requests.
	Password string `json:"password"`
}

var defaultConfig = ClientOptions{
	Endpoint: EndpointPublic,
}

// Client is used to perform reqeusts to the Deepl API.
type Client struct {
	options *ClientOptions
	client  *fasthttp.Client
}

// New creates a new instance of Client with
// the passed options, if passed.
//
// Defaultly, when no options are passed,
// EndpointPublic is used as API endpoint.
func New(options ...ClientOptions) *Client {
	var opt ClientOptions
	if len(options) > 0 {
		opt = options[0]
	}

	if opt.Endpoint == "" {
		opt.Endpoint = EndpointPublic
	}

	return &Client{
		options: &opt,
		client: &fasthttp.Client{
			Name: "godeepl",
		},
	}
}

// Login performs an email-password authentication using the passed email and
// password. If the authentication was successful, the obtained session token
// is stored in the Client instance so subsequent requests can be authenticated
// using the session token.
func (c *Client) Login(email, password string, keepLogin bool) (res *LoginResult, err error) {
	res = &LoginResult{}

	err = c.jsonRpcRequest(
		EndpointLogin, methodLogin,
		newJsonRpcRequest(methodLogin, &loginParams{
			Email:     email,
			Password:  password,
			KeepLogin: keepLogin,
		}),
		&jsonRpcResponse{Result: res})

	if err == nil {
		c.options.SessionID = res.Token
	}

	return
}

// SplitSentence separates the passed text into sentences respecting
// the passed lang using the API.
func (c *Client) SplitSentence(lang LangSpec, text ...string) (res *SplitSentenceResult, err error) {
	res = &SplitSentenceResult{}

	err = c.jsonRpcRequest(
		"", methodSplitSentence,
		newJsonRpcRequest(methodSplitSentence, &splitSentenceParams{
			Texts: text,
			Lang: &language{
				LangUserSelected: lang,
			},
		}),
		&jsonRpcResponse{Result: res})

	return
}

// TranslationOptions wraps additional options
// for the translation endpoint.
type TranslationOptions struct {
	Formality        Formality `json:"formality"`
	PreferedNumBeams int       `json:"prefered_num_beams"`
}

func defaultTranslationOptions(options []TranslationOptions) (opt TranslationOptions) {
	if len(options) > 0 {
		opt = options[0]
	}

	if opt.PreferedNumBeams < 1 {
		opt.PreferedNumBeams = 4
	}

	return
}

// Translate performs a translation request of the passed text respecting the
// passed sourceLang and targetLang.
//
// When the passed text consists of multiple sentences, each sentence is
// translated separately respecting the context of the sentences around.
// The result will contain Translation object for each translated sentence
// with their associated translation beams.
//
// You can also pass some options if you want to customize the formality or
// the prefered number of beams, for example.
func (c *Client) Translate(sourceLang, targetLang LangSpec, text string, options ...TranslationOptions) (res *TranslationResult, err error) {
	res = &TranslationResult{}

	opt := defaultTranslationOptions(options)

	if sourceLang == "" {
		sourceLang = LangAuto
	}

	if targetLang == "" {
		err = errors.New("no target lang specified")
		return
	}

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
		jobs[i] = &translateJob{
			Kind:               "default",
			RawEnSentence:      s,
			RawEnContextBefore: sentences[0:i],
			RawEnContextAfter:  sentences[i+1:],
			PreferredNumBeams:  opt.PreferedNumBeams,
		}
	}

	var cjp *commonJobParams
	if opt.Formality != "" {
		cjp = &commonJobParams{
			Formality: opt.Formality,
		}
	}

	err = c.jsonRpcRequest(
		"", methodTranslate,
		newJsonRpcRequest(methodTranslate, &translateParams{
			Priority: -1,
			Lang: &language{
				SourceLangComputed: sourceLang,
				TargetLang:         targetLang,
				UserPreferredLangs: []LangSpec{},
			},
			Jobs:            jobs,
			CommonJobParams: cjp,
		}),
		&jsonRpcResponse{Result: res})

	return
}
