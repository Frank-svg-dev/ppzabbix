package templates

type TemplatesOpts struct {
	Params Params `json:"params"`
}

type GetTemplatesList struct {
	Jsonrpc string  `json:"jsonrpc"`
	Method  string  `json:"method"`
	Params  Params  `json:"params"`
	Auth    string  `json:"auth"`
	ID      float64 `json:"id"`
}
type Filter struct {
	Host []string `json:"host"`
}
type Params struct {
	Output      []string `json:"output"`
	Filter      Filter   `json:"filter"`
	SelectHosts []string `json:"selectHosts"`
}
