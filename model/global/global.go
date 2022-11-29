package global

type PPZabbixModelReq struct {
	JsonRpc string  `json:"jsonrpc"`
	Method  string  `json:"method"`
	Auth    string  `json:"auth"`
	Id      float64 `json:"id"`
}
