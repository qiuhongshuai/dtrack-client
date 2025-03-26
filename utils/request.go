package utils

import (
	"encoding/json"
	"fmt"
	"resty.dev/v3"
)

func ParserResponse(res *resty.Response, keys ...string) (map[string]interface{}, error) {
	if !res.IsSuccess() {
		return nil, fmt.Errorf("请求失败: %s", res.String())
	}
	temp := make(map[string]interface{})
	//fmt.Println(res.String())
	if err := json.Unmarshal(res.Bytes(), &temp); err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	flatten(temp, "", result)
	if len(keys) == 0 {
		return result, nil
	}
	r := make(map[string]interface{})
	for _, key := range keys {
		if v, ok := result[key]; ok {
			r[key] = v
		}
	}

	return r, nil
}
func flatten(data map[string]interface{}, prefix string, result map[string]interface{}) {
	for key, value := range data {
		currentKey := key
		if prefix != "" {
			currentKey = prefix + "." + key
		}
		//fmt.Println(key)
		switch v := value.(type) {
		case map[string]interface{}:
			flatten(v, currentKey, result)
		case []interface{}:
			for i, item := range v {
				if subMap, ok := item.(map[string]interface{}); ok {
					flatten(subMap, fmt.Sprintf("%s[%d]", currentKey, i), result)
				} else {
					result[fmt.Sprintf("%s[%d]", currentKey, i)] = item
				}
			}
		default:
			result[currentKey] = value
		}
	}
}

// Contains 函数用于判断元素是否存在于切片中，适用于任何可比较的类型
func Contains[T comparable](slice []T, element T) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}
