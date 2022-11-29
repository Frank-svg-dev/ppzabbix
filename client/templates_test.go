package client

import (
	"fmt"
	"github.com/Frank-svg-dev/ppzabbix/model/templates"
	"testing"
)

func TestGetTemplatesList(t *testing.T) {
	zbxClient := NewClientProvider("http://192.168.119.100/zabbix/api_jsonrpc.php", "Admin", "zabbix")
	res, err := GetTemplatesList(zbxClient, templates.TemplatesOpts{Params: templates.Params{
		Output: []string{"extends"},
	}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetTemplatesListByCreateHost(t *testing.T) {
	zbxClient := NewClientProvider("http://192.168.119.100/zabbix/api_jsonrpc.php", "Admin", "zabbix")
	res, err := GetTemplatesList(zbxClient, templates.TemplatesOpts{Params: templates.Params{
		Output: []string{"extends", "name", "available"},
	}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	fmt.Println(res)
}
