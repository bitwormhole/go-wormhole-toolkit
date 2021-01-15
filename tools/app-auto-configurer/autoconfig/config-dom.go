package autoconfig

// ComponentElement 表示一个组件的配置项
type ComponentElement struct {
	ID    string `json:"id"`
	Class string `json:"class"`
	Type  string `json:"type"`

	Scope         string            `json:"scope"`
	InitMethod    string            `json:"initMethod"`
	DestroyMethod string            `json:"destroyMethod"`
	Inject        map[string]string `json:"inject"`
}

// ConfigDOM 是配置文件的 “gss.config.json” DOM
type ConfigDOM struct {
	Components []*ComponentElement `json:"components"`
	Package    string              `json:"package"`
}
