package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/client"
	"github.com/Frank-svg-dev/ppzabbix/model/host"
	"github.com/Frank-svg-dev/ppzabbix/utils"
)

// CreateHosts  添加主机
func CreateHosts(client *client.RESTCLIENT, opts host.ClientOpts) (interface{}, error) {
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

	return result["result"], err
}
