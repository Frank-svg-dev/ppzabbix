package templates

type UpdateTemplateOpts struct {
	TemplateID string `json:"templateid"`
	Name       string `json:"name"`
}

type UpdateTemplateModel struct {
	JsonRpc string             `json:"jsonrpc"`
	Method  string             `json:"method"`
	Params  UpdateTemplateOpts `json:"params"`
	Auth    string             `json:"auth"`
	Id      float64            `json:"id"`
}
