package trigger

type UpdateTriggerOpts struct {
	TriggerId string `json:"triggerid"`
	Status    int    `json:"status"`
}

type UpdateTriggerModel struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  UpdateTriggerOpts `json:"params"`
	Auth    string            `json:"auth"`
	ID      float64           `json:"id"`
}
