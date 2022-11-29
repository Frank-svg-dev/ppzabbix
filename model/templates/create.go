package templates

type CreateTemplateModel struct {
	JsonRpc string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  CreateTemplatesOpts `json:"params"`
	Auth    string              `json:"auth"`
	Id      float64             `json:"id"`
}

type CreateTemplatesOpts struct {
	Name        string                     `json:"name"`
	Host        string                     `json:"host"`
	Description string                     `json:"description"`
	Groups      []CreateTemplateGroupsOpts `json:"groups"`
	Hosts       []CreateTemplatesHostsOpts `json:"hosts"`
}

type CreateTemplateGroupsOpts struct {
	GroupId float64 `json:"groupid"`
}

type CreateTemplatesHostsOpts struct {
	HostId string `json:"hostid"`
}
