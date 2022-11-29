package groups

type CreateHostGroupOpts struct {
	Name string `json:"name"`
}

type CreateNewHostGroupsModel struct {
	Jsonrpc string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  CreateHostGroupOpts `json:"params"`
	Auth    string              `json:"auth"`
	ID      float64             `json:"id"`
}
