package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/groups"
	"testing"
)

func TestGetHostGroups(t *testing.T) {
	zbxClient := NewClientProvider("http://192.168.119.100/zabbix/api_jsonrpc.php", "Admin", "zabbix")
	res, err := GetHostGroups(zbxClient, groups.HostGroupOpts{
		Params: groups.Params{
			Output: "extends",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
