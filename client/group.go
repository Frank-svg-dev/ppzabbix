package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/client"
	"github.com/Frank-svg-dev/ppzabbix/model/groups"
	"github.com/Frank-svg-dev/ppzabbix/utils"
)

func GetHostGroups(client *client.RESTCLIENT, opts groups.HostGroupOpts) (interface{}, error) {
	body := groups.GetHostGroupsList{
		Jsonrpc: "2.0",
		Method:  "hostgroup.get",
		Params:  opts.Params,
		Auth:    client.Result,
		ID:      client.Id,
	}
	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err
	}

	return result, nil
}
