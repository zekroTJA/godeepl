package godeepl

const (
	// Endpoint used for login.
	EndpointLogin = "https://w.deepl.com/account"
	// Endpoint for public, unauthorized requests.
	// Attention: This endpoint is highly rate-limited!
	EndpointPublic = "https://www2.deepl.com/jsonrpc"
	// EndpointPro for authenticated pro-plan requests.
	EndpointPro = "https://api.deepl.com/jsonrpc"

	methodLogin         = "login"
	methodTranslate     = "LMT_handle_jobs"
	methodSplitSentence = "LMT_split_into_sentences"

	jsonRPCVersion = "2.0"
)
