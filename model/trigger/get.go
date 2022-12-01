package trigger

type TriggerOpts struct {
	OutPut            []string      `json:"output"`
	Filter            TriggerFilter `json:"filter"`
	SortField         string        `json:"sortfield"`
	SortOrder         string        `json:"sortorder"`
	ExpandDescription int           `json:"expandDescription"`
}

type TriggerFilter struct {
	Value int `json:"value"`
}

type GetTriggerModel struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  TriggerOpts `json:"params"`
	Auth    string      `json:"auth"`
	ID      float64     `json:"id"`
}
