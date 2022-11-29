package host

type ClientOpts struct {
	HostName      string          `json:"host"`
	HostInterface []HostInterface `json:"hostInterfaces"`
	GroupID       string          `json:"groupID"`
	Tags          []Tags          `json:"tags"`
	TemplateID    string          `json:"templateID"`
}

type CreateHostModel struct {
	Jsonrpc string  `json:"jsonrpc"`
	Method  string  `json:"method"`
	Params  Params  `json:"params"`
	Auth    string  `json:"auth"`
	ID      float64 `json:"id"`
}

type Params struct {
	Host        string          `json:"host"`
	VisibleName string          `json:"visiblename"`
	Interfaces  []HostInterface `json:"interfaces"`
	Groups      []Groups        `json:"groups"`
	Tags        []Tags          `json:"tags"`
	Templates   []Templates     `json:"templates"`
}
type HostInterface struct {
	Type  int    `json:"type"`
	Main  int    `json:"main"`
	Useip int    `json:"useip"`
	IP    string `json:"ip"`
	DNS   string `json:"dns"`
	Port  string `json:"port"`
}
type Groups struct {
	Groupid string `json:"groupid"`
}
type Tags struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}
type Templates struct {
	Templateid string `json:"templateid"`
}
