package client

import (
	"testing"
)

//func TestCreateInterFaceHost(t *testing.T) {
//	zbxClient := NewClientProvider("http://192.168.119.100/zabbix/api_jsonrpc.php", "Admin", "zabbix")
//	result, err := CreateHosts(zbxClient, host.ClientOpts{
//		HostName: "gotest",
//		HostInterface: []host.HostInterface{host.HostInterface{
//			Type:  1,
//			Main:  1,
//			Useip: 1,
//			IP:    "192.168.111.111",
//			DNS:   "",
//			Port:  "10050",
//		}},
//		Tags:       []host.Tags{},
//		GroupID:    "18",
//		TemplateID: "10047",
//	})
//
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(result)
//	fmt.Println(result)
//}

func TestGetHostList(t *testing.T) {

}
