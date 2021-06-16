package godeepl

import "sync/atomic"

type LangSpec string

var RollingId int64

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

type JsonRpcRequest struct {
	Version string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

func NewJsonRpcRequest(method string, params interface{}) *JsonRpcRequest {
	return &JsonRpcRequest{
		Version: jsonRPCVersion,
		Id:      int(atomic.AddInt64(&RollingId, 1)),
		Method:  method,
		Params:  params,
	}
}

type TranslateParams struct {
	Priority int `json:"priority"`

	Jobs            []interface{}    `json:"jobs"`
	Lang            *Lang            `json:"lang"`
	CommonJobParams *CommonJobParams `json:"commonJobParams"`
}

type LoginParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	KeepLogin bool   `json:"keepLogin"`
}

type TranslateJob struct {
	Kind               string   `json:"kind"`
	RawEnSentence      string   `json:"raw_en_sentence"`
	RawEnContextBefore []string `json:"raw_en_context_before"`
	RawEnContextAfter  []string `json:"raw_en_context_after"`
	PreferredNumBeams  int      `json:"preferred_num_beams"`
	Quality            string   `json:"quality"`
}

type Lang struct {
	UserPreferredLangs []LangSpec `json:"user_preferred_langs"`
	LangUserSelected   LangSpec   `json:"lang_user_selected"`
	SourceLangComputed LangSpec   `json:"source_lang_computed"`
	TargetLang         LangSpec   `json:"target_lang"`
}

type CommonJobParams struct {
	Formality string `json:"formality"`
}

type JsonRpcResponse struct {
	Version string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Result  interface{} `json:"result"`
}

type LoginResult struct {
	Id              int    `json:"id"`
	IsAdministrator bool   `json:"isAdministrator"`
	Email           string `json:"email"`
	Token           string `json:"token"`
}

type TranslationResult struct {
	Translations          []*Translation `json:"translations"`
	TargetLang            LangSpec       `json:"target_lang"`
	SourceLang            LangSpec       `json:"source_lang"`
	SourceLangIsConfident bool           `json:"source_lang_is_confident"`
	DetectedLanguages     interface{}    `json:"detectedLanguages"` // no clue what type this might be because only {} is always returned
}

type Translation struct {
	Beams   []*Beam `json:"beams"`
	Quality string  `json:"quality"`
}

type Beam struct {
	ProcessedSentence string `json:"postprocessed_sentence"`
	NumSymbols        int    `json:"num_symbols"`
}

type SplitSentenceParams struct {
	Texts []string `json:"texts"`
	Lang  *Lang    `json:"lang"`
}

type SplitSentenceResult struct {
	SplittedTexts     [][]string  `json:"splitted_texts"`
	Lang              LangSpec    `json:"lang"`
	LangIsConfident   bool        `json:"lang_is_confident"`
	DetectedLanguages interface{} `json:"detectedLanguages"` // no clue what type this might be because only {} is always returned
}
