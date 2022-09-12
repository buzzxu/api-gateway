package configs

import (
	"fmt"
	"testing"
)

func TestYamlLoader_Load(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("发生错误", r)
		}
	}()
	loader := new(YamlLoader)
	loader.Load()

}
