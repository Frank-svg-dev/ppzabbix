package client

import (
	"github.com/Frank-svg-dev/ppzabbix/model/groups"
	"testing"
)

func TestGetHostGroups(t *testing.T) {
	res, err := GetHostGroups(ZbxClient, groups.HostGroupOpts{
		Params: groups.Params{
			Output: []string{"extends"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetHostGroupsByCreateHost(t *testing.T) {
	res, err := GetHostGroups(ZbxClient, groups.HostGroupOpts{
		Params: groups.Params{
			Output: []string{"extends", "name"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestCreateNewHostGroups(t *testing.T) {
	res, err := CreateNewHostGroups(ZbxClient, groups.CreateHostGroupOpts{Name: "gohostgrouptest"})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestDeleteHostGroups(t *testing.T) {
	res, err := DeleteHostGroups(ZbxClient, []string{"19"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRenameHostGroups(t *testing.T) {
	res, err := RenameHostGroups(ZbxClient, groups.RenameHostGroupOpts{
		GroupID: "18",
		Name:    "go123test",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
