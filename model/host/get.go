package host

type GetHostListOpts struct {
	OutPut                []string      `json:"output"`
	Filter                GetHostFilter `json:"filter"`
	SelectGroups          []string      `json:"selectGroups"`
	SelectParentTemplates []string      `json:"selectParentTemplates"`
}

type GetHostListModel struct {
	Jsonrpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  GetHostListOpts `json:"params"`
	Auth    string          `json:"auth"`
	ID      float64         `json:"id"`
}

type GetHostFilter struct {
	Host []string `json:"host"`
}
