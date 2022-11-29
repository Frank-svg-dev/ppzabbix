package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 请求操作+byte转map
func PostProcessingJSON(ProcessStruct interface{}, url string) (ResultMap map[string]interface{}, err error) {
	jsonBytes, _ := json.Marshal(ProcessStruct)
	body := string(jsonBytes)
	response, err := http.Post(url, "application/json; charset=utf-8", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(content, &result)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}
