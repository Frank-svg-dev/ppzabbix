package client

import (
	"fmt"
	"github.com/Frank-svg-dev/ppzabbix/model/alerts"
	"github.com/Frank-svg-dev/ppzabbix/model/client"
	"github.com/Frank-svg-dev/ppzabbix/utils"
)

// GetAlerts 获取现存告警
func GetAlerts(client *client.RESTCLIENT, opts alerts.GetAlertsOpts) {
	body := alerts.GetAlertsModel{
		Jsonrpc: "2.0",
		Method:  "alert.get",
		Params:  opts,
		Auth:    client.Result,
		ID:      client.Id,
	}

	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}
