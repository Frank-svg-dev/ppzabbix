package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/client"
	"github.com/Frank-svg-dev/ppzabbix/model/trigger"
	"github.com/Frank-svg-dev/ppzabbix/utils"
)

func GetTriggerList(client *client.RESTCLIENT, opts trigger.TriggerOpts) (interface{}, error) {
	body := trigger.GetTriggerModel{
		Jsonrpc: "2.0",
		Method:  "trigger.get",
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

func UpdateTriggers(client *client.RESTCLIENT, opts trigger.UpdateTriggerOpts) (interface{}, error) {
	body := trigger.UpdateTriggerModel{
		Jsonrpc: "2.0",
		Method:  "trigger.update",
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

func DeleteTriggers(client *client.RESTCLIENT, opts trigger.DeleteTriggerOpts) (interface{}, error) {
	body := trigger.DeleteTriggerModel{
		JsonRpc: "2.0",
		Method:  "trigger.update",
		Params:  opts.TemplatesId,
		Auth:    client.Result,
		Id:      client.Id,
	}

	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err
	}

	return result, nil
}
