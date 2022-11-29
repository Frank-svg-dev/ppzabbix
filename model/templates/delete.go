package templates

type DeleteTemplateOpts struct {
	TemplatesId []string `json:"TemplatesId"`
}

type DeleteTemplateModel struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Auth    string   `json:"auth"`
	Id      float64  `json:"id"`
}
