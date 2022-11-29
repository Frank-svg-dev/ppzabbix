package host

type UpdateHostOpts struct {
	HostID         string           `json:"hostid"`
	TemplatesClear []TemplatesClear `json:"templates_clear"`
	Status         int              `json:"status"` //启用主机, 0为启用, 1为禁用
}

type TemplatesClear struct {
	TemplateID string `json:"templateid"`
}

type UpdateHostModel struct {
	Jsonrpc string         `json:"jsonrpc"`
	Method  string         `json:"method"`
	Params  UpdateHostOpts `json:"params"`
	Auth    string         `json:"auth"`
	ID      float64        `json:"id"`
}
