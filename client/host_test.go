package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/host"
	"testing"
)

func TestCreateInterFaceHost(t *testing.T) {
	zbxClient := NewClientProvider("http://192.168.119.100/zabbix/api_jsonrpc.php", "Admin", "zabbix")
	result, err := CreateHosts(zbxClient, host.ClientOpts{
		HostName: "gotest",
		HostInterface: []host.HostInterface{host.HostInterface{
			Type:  1,
			Main:  1,
			Useip: 1,
			IP:    "192.168.111.111",
			DNS:   "",
			Port:  "10050",
		}},
		Tags:       []host.Tags{},
		GroupID:    "18",
		TemplateID: "10047",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetHostsList(t *testing.T) {
	res, err := GetHostsList(ZbxClient, host.GetHostListOpts{OutPut: []string{"hostid", "name"}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetHostListByFilter(t *testing.T) {
	res, err := GetHostsList(ZbxClient, host.GetHostListOpts{OutPut: []string{"hostid", "name", "available"}, Filter: host.GetHostFilter{Host: []string{"Zabbix server"}}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetHostListByGetHostGroup(t *testing.T) {
	res, err := GetHostsList(ZbxClient, host.GetHostListOpts{OutPut: []string{"hostid", "name", "available"},
		Filter: host.GetHostFilter{Host: []string{"Zabbix server"}}, SelectGroups: []string{"extend"}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetHostListByTemplates(t *testing.T) {
	res, err := GetHostsList(ZbxClient, host.GetHostListOpts{OutPut: []string{"hostid", "name", "available"},
		Filter: host.GetHostFilter{Host: []string{"Zabbix server"}}, SelectGroups: []string{"extend"}, SelectParentTemplates: []string{"templateid", "name"}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestDeleteHosts(t *testing.T) {
	res, err := DeleteHosts(ZbxClient, host.DeleteHostOpts{HostIds: []string{"10458"}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestUpdateHostByEnableMonitor(t *testing.T) {
	res, err := UpdateHost(ZbxClient, host.UpdateHostOpts{
		HostID: "10461",
		Status: 0,
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestUpdateHostByDeleteTemplates(t *testing.T) {
	res, err := UpdateHost(ZbxClient, host.UpdateHostOpts{
		HostID:         "10461",
		TemplatesClear: []host.TemplatesClear{host.TemplatesClear{TemplateID: "10352"}},
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
