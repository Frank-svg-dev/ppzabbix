package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/client"
	"github.com/Frank-svg-dev/ppzabbix/model/templates"
	"github.com/Frank-svg-dev/ppzabbix/utils"
)

func GetTemplatesList(client *client.RESTCLIENT, opts templates.TemplatesOpts) (interface{}, error) {
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
