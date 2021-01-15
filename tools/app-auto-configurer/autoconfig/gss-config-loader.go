package autoconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// GssConfigLoader 是配置加载器
type GssConfigLoader struct {
}

// Load 加载配置
func (inst *GssConfigLoader) Load(ctx *GssConfigFileContext) error {

	jsonFile, err := os.Open(ctx.Path.Path())
	if err == nil {
		defer jsonFile.Close()
	} else {
		// fmt.Println(err)
		return err
	}

	dom := &ConfigDOM{}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, dom)

	ctx.DOM = dom
	return nil
}
