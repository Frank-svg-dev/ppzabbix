package client

type (
	ClientOpts struct {
		JsonRpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  UserAuthParam `json:"params"`
		Id      int           `json:"id"`
		Auth    []byte        `json:"auth"`
	}

	UserAuthParam struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}

	RESTCLIENT struct {
		JsonRpc string  `json:"jsonrpc"`
		Url     string  `json:"url"`
		Result  string  `json:"result"`
		Id      float64 `json:"id"`
	}
)
