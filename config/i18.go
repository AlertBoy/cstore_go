package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

var dic *map[interface{}]interface{}

func loadLocal(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(file), m)
	if err != nil {
		return err
	}
	dic = &m
	return nil
}

func T(key string) string {
	dict := *dic
	keys := strings.Split(key, ".")
	for index, path := range keys {
		if len(keys) == (index + 1) {
			for k, v := range dict {
				if k, ok := k.(string); ok {
					if k == path {
						if value, ok := v.(string); ok {
							return value
						}
					}
				}
			}
			return path
		}
		// 如果还有下一层，继续寻找
		for k, v := range dict {
			if ks, ok := k.(string); ok {
				if ks == path {
					if dict, ok = v.(map[interface{}]interface{}); ok == false {
						return path
					}
				}
			} else {
				return ""
			}
		}
	}
	return ""
}
