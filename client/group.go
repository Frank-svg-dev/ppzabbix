package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/client"
	"github.com/Frank-svg-dev/ppzabbix/model/groups"
	"github.com/Frank-svg-dev/ppzabbix/utils"
)

func GetHostGroups(client *client.RESTCLIENT, opts groups.HostGroupOpts) (map[string]interface{}, error) {
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

func CreateNewHostGroups(client *client.RESTCLIENT, opts groups.CreateHostGroupOpts) (map[string]interface{}, error) {
	body := groups.CreateNewHostGroupsModel{
		Jsonrpc: "2.0",
		Method:  "hostgroup.create",
		Params:  opts,
		Auth:    client.Result,
		ID:      client.Id,
	}

	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteHostGroups(client *client.RESTCLIENT, opts []string) (map[string]interface{}, error) {
	body := groups.DeleteHostGroupsModel{
		Jsonrpc: "2.0",
		Method:  "hostgroup.delete",
		Params:  opts,
		Auth:    client.Result,
		ID:      client.Id,
	}

	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func RenameHostGroups(client *client.RESTCLIENT, opts groups.RenameHostGroupOpts) (map[string]interface{}, error) {
	body := groups.RenameHostGroupsModel{
		Jsonrpc: "2.0",
		Method:  "hostgroup.update",
		Params:  opts,
		Auth:    client.Result,
		ID:      client.Id,
	}
	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err
	}

	return result, nil
}
