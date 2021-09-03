package utils

import (
	"encoding/json"
	"os"
)

func ReadJson(path string, v interface{}) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, v); err != nil {
		return err
	}
	return nil
}

func WriteJson(path string,data interface{}) error {
	value,err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, value, 0777); err != nil {
		return err
	}
	return nil
}