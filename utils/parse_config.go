package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

// ParseConfig принимает path и config. Находит файл конфигурации
// и записывает все данные в указатель config
func ParseConfig(path string, config interface{}) error {
	if path == "" {
		return fmt.Errorf("[ParseConfig] - path for config is empty")
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("[ParseConfig] - could not open file by <%s> path. error: %v", path, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("[ParseConfig] - could not read file by <%s> path. error: %v", path, err)
	}
	if err = jsoniter.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("[ParseConfig] - could not unmarshal file by <%s> path. error: %v", path, err)
	}
	v := validator.New()
	if err := v.Struct(config); err != nil {
		return fmt.Errorf("[ParseConfig] - could not validate config. error: %v", err)
	}
	return nil
}
