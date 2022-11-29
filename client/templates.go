package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/client"
	"github.com/Frank-svg-dev/ppzabbix/model/templates"
	"github.com/Frank-svg-dev/ppzabbix/utils"
)

func GetTemplatesList(client *client.RESTCLIENT, opts templates.TemplatesOpts) (map[string]interface{}, error) {
	body := templates.GetTemplatesList{
		Jsonrpc: "2.0",
		Method:  "template.get",
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

func CreateNewTemplate(client *client.RESTCLIENT, opts templates.CreateTemplatesOpts) (map[string]interface{}, error) {
	body := templates.CreateTemplateModel{
		JsonRpc: "2.0",
		Method:  "template.create",
		Params:  opts,
		Auth:    client.Result,
		Id:      client.Id,
	}

	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteTemplates(client *client.RESTCLIENT, opts templates.DeleteTemplateOpts) (map[string]interface{}, error) {
	body := templates.DeleteTemplateModel{
		JsonRpc: "2.0",
		Method:  "template.delete",
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

func UpdateTemplate(client *client.RESTCLIENT, opts templates.UpdateTemplateOpts) (map[string]interface{}, error) {
	body := templates.UpdateTemplateModel{
		JsonRpc: "2.0",
		Method:  "template.update",
		Params:  opts,
		Auth:    client.Result,
		Id:      client.Id,
	}

	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err
	}

	return result, nil
}
