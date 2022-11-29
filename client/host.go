package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/client"
	"github.com/Frank-svg-dev/ppzabbix/model/host"
	"github.com/Frank-svg-dev/ppzabbix/utils"
)

// CreateHosts  添加主机
func CreateHosts(client *client.RESTCLIENT, opts host.ClientOpts) (map[string]interface{}, error) {
	body := host.CreateHostModel{
		Jsonrpc: "2.0",
		Method:  "host.create",
		Params: host.Params{
			Host:        opts.HostName,
			VisibleName: opts.HostName,
			Interfaces:  opts.HostInterface,
			Tags:        opts.Tags,
			Groups:      []host.Groups{host.Groups{Groupid: opts.GroupID}},
			Templates:   []host.Templates{host.Templates{Templateid: opts.TemplateID}},
		},
		Auth: client.Result,
		ID:   client.Id,
	}

	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err

	}

	return result, err
}

func GetHostsList(client *client.RESTCLIENT, opts host.GetHostListOpts) (map[string]interface{}, error) {
	body := host.GetHostListModel{
		Jsonrpc: "2.0",
		Method:  "host.get",
		Params:  opts,
		Auth:    client.Result,
		ID:      client.Id,
	}
	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err

	}

	return result, err
}

func DeleteHosts(client *client.RESTCLIENT, opts host.DeleteHostOpts) (map[string]interface{}, error) {
	body := host.DeleteHostModel{
		Jsonrpc: "2.0",
		Method:  "host.delete",
		Params:  opts.HostIds,
		Auth:    client.Result,
		ID:      client.Id,
	}

	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err

	}

	return result, err
}

func UpdateHost(client *client.RESTCLIENT, opts host.UpdateHostOpts) (map[string]interface{}, error) {
	body := host.UpdateHostModel{
		Jsonrpc: "2.0",
		Method:  "host.update",
		Params:  opts,
		Auth:    client.Result,
		ID:      client.Id,
	}

	result, err := utils.PostProcessingJSON(body, client.Url)

	if err != nil {
		return nil, err

	}

	return result, err
}
