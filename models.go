package godeepl

import "sync/atomic"

// LangSpec is the specification for
// a specific language selector.
type LangSpec string

var rollingId int64

const (
	LangAuto       LangSpec = "auto"
	LangBulgarian  LangSpec = "BG"
	LangCzech      LangSpec = "CS"
	LangDanish     LangSpec = "DA"
	LangGerman     LangSpec = "DE"
	LangGreek      LangSpec = "EL"
	LangEnglish    LangSpec = "EN"
	LangSpanish    LangSpec = "ES"
	LangEstonian   LangSpec = "ET"
	LangFinnish    LangSpec = "FI"
	LangFrench     LangSpec = "FR"
	LangHungarian  LangSpec = "HU"
	LangItalian    LangSpec = "IT"
	LangJapanese   LangSpec = "JA"
	LangLithuanian LangSpec = "LT"
	LangLatvian    LangSpec = "LV"
	LangDutch      LangSpec = "NL"
	LangPolish     LangSpec = "PL"
	LangPortuguese LangSpec = "PT"
	LangRomanian   LangSpec = "RO"
	LangRussian    LangSpec = "RU"
	LangSlovak     LangSpec = "SK"
	LangSlovenian  LangSpec = "SL"
	LangSwedish    LangSpec = "SV"
	LangChinese    LangSpec = "ZH"
)

// LoginResult wraps the response of the login endpoint.
type LoginResult struct {
	Id              int    `json:"id"`
	IsAdministrator bool   `json:"isAdministrator"`
	Email           string `json:"email"`
	Token           string `json:"token"`
}

// TranslationResult wraps a translation response.
type TranslationResult struct {
	Translations          []*Translation `json:"translations"`
	TargetLang            LangSpec       `json:"target_lang"`
	SourceLang            LangSpec       `json:"source_lang"`
	SourceLangIsConfident bool           `json:"source_lang_is_confident"`
	DetectedLanguages     interface{}    `json:"detectedLanguages"` // No clue what type this might be because only {} is always returned
}

// Translation contains an array of beams for the
// translation as well as a quality identificator.
type Translation struct {
	Beams   []*Beam `json:"beams"`
	Quality string  `json:"quality"`
}

// Beam contains one translation alternative for
// a translated text.
type Beam struct {
	ProcessedSentence string `json:"postprocessed_sentence"`
	NumSymbols        int    `json:"num_symbols"`
}

// SplitSentenceResult wraps the response of
// the spit sentence endpoint.
type SplitSentenceResult struct {
	SplittedTexts     [][]string  `json:"splitted_texts"`
	Lang              LangSpec    `json:"lang"`
	LangIsConfident   bool        `json:"lang_is_confident"`
	DetectedLanguages interface{} `json:"detectedLanguages"` // No clue what type this might be because only {} is always returned
}

type jsonRpcRequest struct {
	Version string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

func newJsonRpcRequest(method string, params interface{}) *jsonRpcRequest {
	return &jsonRpcRequest{
		Version: jsonRPCVersion,
		Id:      int(atomic.AddInt64(&rollingId, 1)),
		Method:  method,
		Params:  params,
	}
}

type translateParams struct {
	Priority int `json:"priority"`

	Jobs            []interface{}    `json:"jobs"`
	Lang            *language        `json:"lang"`
	CommonJobParams *commonJobParams `json:"commonJobParams"`
}

type loginParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	KeepLogin bool   `json:"keepLogin"`
}

type translateJob struct {
	Kind               string   `json:"kind"`
	RawEnSentence      string   `json:"raw_en_sentence"`
	RawEnContextBefore []string `json:"raw_en_context_before"`
	RawEnContextAfter  []string `json:"raw_en_context_after"`
	PreferredNumBeams  int      `json:"preferred_num_beams"`
	Quality            string   `json:"quality"`
}

type language struct {
	UserPreferredLangs []LangSpec `json:"user_preferred_langs"`
	LangUserSelected   LangSpec   `json:"lang_user_selected"`
	SourceLangComputed LangSpec   `json:"source_lang_computed"`
	TargetLang         LangSpec   `json:"target_lang"`
}

type commonJobParams struct {
	Formality string `json:"formality"`
}

type jsonRpcResponse struct {
	Version string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Result  interface{} `json:"result"`
}

type splitSentenceParams struct {
	Texts []string  `json:"texts"`
	Lang  *language `json:"lang"`
}
