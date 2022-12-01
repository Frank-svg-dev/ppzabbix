package client

import (
	"fmt"
	"github.com/Frank-svg-dev/ppzabbix/model/trigger"
	"testing"
)

func TestGetTriggerListByFault(t *testing.T) {
	res, err := GetTriggerList(ZbxClient, trigger.TriggerOpts{
		OutPut:            []string{"triggerid", "description", "priority"},
		Filter:            trigger.TriggerFilter{Value: 1},
		SortField:         "priority",
		SortOrder:         "DESC",
		ExpandDescription: 1,
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	fmt.Println(res)
}

func TestUpdateTriggers(t *testing.T) {
	res, err := UpdateTriggers(ZbxClient, trigger.UpdateTriggerOpts{TriggerId: "16199", Status: 0})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestDeleteTriggers(t *testing.T) {
	res, err := DeleteTriggers(ZbxClient, trigger.DeleteTriggerOpts{TriggersId: []string{"10011"}})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
