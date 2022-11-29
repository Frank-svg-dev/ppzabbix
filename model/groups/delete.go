package groups

type DeleteHostGroupsModel struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Auth    string   `json:"auth"`
	ID      float64  `json:"id"`
}
