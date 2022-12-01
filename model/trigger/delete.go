package trigger

type DeleteTriggerOpts struct {
	TriggersId []string `json:"TriggersId"`
}

type DeleteTriggerModel struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Auth    string   `json:"auth"`
	Id      float64  `json:"id"`
}
