package autoconfig

import (
	"encoding/json"
)

// GssConfigLoader 是配置加载器
type GssConfigLoader struct {
}

// Load 加载配置
func (inst *GssConfigLoader) Load(ctx *ProjectConfigFile) error {

	byteValue, err := ctx.Path.GetIO().ReadBinary()

	if err != nil {
		return err
	}

	dom := &ConfigDOM{}
	json.Unmarshal(byteValue, dom)
	ctx.DOM = dom
	return nil
}
