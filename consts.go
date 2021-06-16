package godeepl

const (
	EndpointLogin  = "https://w.deepl.com/account"
	EndpointPublic = "https://www2.deepl.com/jsonrpc"
	EndpointPro    = "https://api.deepl.com/jsonrpc"

	methodLogin         = "login"
	methodTranslate     = "LMT_handle_jobs"
	methodSplitSentence = "LMT_split_into_sentences"

	jsonRPCVersion = "2.0"
)
