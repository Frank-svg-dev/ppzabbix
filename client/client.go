package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/client"
	"github.com/Frank-svg-dev/ppzabbix/utils"
)

// NewClientProvider 初始化客户端
func NewClientProvider(url, username, password string) *client.RESTCLIENT {

	url = url + "/zabbix/api_jsonrpc.php"

	opts := client.ClientOpts{
		JsonRpc: "2.0",
		Method:  "user.login",
		Params: client.UserAuthParam{
			User:     username,
			Password: password,
		},
		Id:   0,
		Auth: nil,
	}

	result, err := utils.PostProcessingJSON(opts, url)

	if err != nil {
		panic(err)
	}

	return &client.RESTCLIENT{
		JsonRpc: result["jsonrpc"].(string),
		Url:     url,
		Result:  result["result"].(string),
		Id:      result["id"].(float64),
	}
}
