package groups

type HostGroupOpts struct {
	Params Params `json:"params"`
}

type GetHostGroupsList struct {
	Jsonrpc string  `json:"jsonrpc"`
	Method  string  `json:"method"`
	Params  Params  `json:"params"`
	Auth    string  `json:"auth"`
	ID      float64 `json:"id"`
}

type Params struct {
	Output []string `json:"output"`
	Filter Filter   `json:"filter"`
}

type Filter struct {
	Name []string `json:"name"`
}

type HostGroupsResult struct {
	Jsonrpc string   `json:"jsonrpc"`
	Result  []Result `json:"result"`
	ID      int      `json:"id"`
}
type Result struct {
	Groupid  string `json:"groupid"`
	Name     string `json:"name"`
	Internal string `json:"internal"`
}
