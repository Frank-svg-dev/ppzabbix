package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/templates"
	"testing"
)

var ZbxClient = NewClientProvider("http://192.168.119.100/zabbix/api_jsonrpc.php", "Admin", "zabbix")

/*
只获取template的ID
result: map[id:0 jsonrpc:2.0 result:[map[templateid:10001]]
*/
func TestGetTemplatesList(t *testing.T) {
	res, err := GetTemplatesList(ZbxClient, templates.TemplatesOpts{Params: templates.Params{
		Output: []string{"extends"},
	}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

/*
创建主机时所需要的模板参数按照这个来就可以
result: map[id:0 jsonrpc:2.0 result:[map[available:0 name:Template OS Linux by Zabbix agent templateid:10001]]
*/
func TestGetTemplatesListByCreateHost(t *testing.T) {
	res, err := GetTemplatesList(ZbxClient, templates.TemplatesOpts{Params: templates.Params{
		Output: []string{"extends", "name", "available"},
	}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

/*
创建模板
result: map[id:0 jsonrpc:2.0 result:map[templateids:[10459]]]  回执为interface需要自己断言
*/
func TestCreateNewTemplate(t *testing.T) {
	res, err := CreateNewTemplate(ZbxClient, templates.CreateTemplatesOpts{
		Name:        "GoName",
		Host:        "GoHost",
		Description: "123123123",
		Groups:      []templates.CreateTemplateGroupsOpts{templates.CreateTemplateGroupsOpts{GroupId: 18}},
		Hosts:       []templates.CreateTemplatesHostsOpts{templates.CreateTemplatesHostsOpts{HostId: "10084"}},
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestDeleteTemplates(t *testing.T) {
	res, err := DeleteTemplates(ZbxClient, templates.DeleteTemplateOpts{TemplatesId: []string{"10459"}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestUpdateTemplate(t *testing.T) {
	res, err := UpdateTemplate(ZbxClient, templates.UpdateTemplateOpts{
		TemplateID: "10352",
		Name:       "gogotest",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
