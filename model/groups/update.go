package groups

type RenameHostGroupsModel struct {
	Jsonrpc string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  RenameHostGroupOpts `json:"params"`
	Auth    string              `json:"auth"`
	ID      float64             `json:"id"`
}

type RenameHostGroupOpts struct {
	GroupID string `json:"groupid"`
	Name    string `json:"name"`
}
