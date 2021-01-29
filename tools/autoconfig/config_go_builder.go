package autoconfig

import (
	"strings"
)

// ComponentConfig 组件配置
type ComponentConfig struct {
	componentId            string
	componentClasses       string
	componentInitMethod    string
	componentDestroyMethod string

	holderTypeName    string
	holderSimpleName  string
	targetTypeName    string
	targetSimpleName  string
	targetPackageName string
}

// ConfigGoBuilder 配置生成器
type ConfigGoBuilder struct {
	components map[string]*ComponentConfig
	buffer     strings.Builder
}

func (inst *ConfigGoBuilder) Create() (string, error) {

	table := inst.components

	if table == nil {
		return inst.buffer.String(), nil
	}

	for key := range table {
		item := table[key]
		err := inst.writeComponent(item)
		if err != nil {
			return "", err
		}
	}

	return inst.buffer.String(), nil
}

func (inst *ConfigGoBuilder) writeComponent(comp *ComponentConfig) error {

	// createInstance
	err := inst.writeComponentCreateInstance(comp)
	if err != nil {
		return err
	}

	// inject
	err = inst.writeComponentInject(comp)
	if err != nil {
		return err
	}

	// initMethod
	err = inst.writeComponentInitMethod(comp)
	if err != nil {
		return err
	}

	// destroyMethod
	err = inst.writeComponentDestroyMethod(comp)
	if err != nil {
		return err
	}

	return nil
}

func (inst *ConfigGoBuilder) writeComponentCreateInstance(comp *ComponentConfig) error {

	// TODO
	return nil
}

func (inst *ConfigGoBuilder) writeComponentInitMethod(comp *ComponentConfig) error {

	// func (inst * _holder_type_) Init() error {
	//		return inst.target._init_method_()
	// }

	buf := &inst.buffer
	_holderType := comp.holderSimpleName
	_destroyMethod := comp.componentDestroyMethod

	buf.WriteString("func (inst * ")
	buf.WriteString(_holderType)
	buf.WriteString(") Destroy() error {\n")

	if _destroyMethod == "" {
		buf.WriteString("    return nil\n")
	} else {
		buf.WriteString("    return inst.target.")
		buf.WriteString(_destroyMethod)
		buf.WriteString("()\n")
	}

	buf.WriteString("}\n\n")

	return nil
}

func (inst *ConfigGoBuilder) writeComponentDestroyMethod(comp *ComponentConfig) error {

	// func (inst * _holder_type_) Destroy() error {
	//		return inst.target._destroy_method_()
	// }

	buf := &inst.buffer
	_holderType := comp.holderSimpleName
	_destroyMethod := comp.componentDestroyMethod

	buf.WriteString("func (inst * ")
	buf.WriteString(_holderType)
	buf.WriteString(") Destroy() error {\n")

	if _destroyMethod == "" {
		buf.WriteString("    return nil\n")
	} else {
		buf.WriteString("    return inst.target.")
		buf.WriteString(_destroyMethod)
		buf.WriteString("()\n")
	}

	buf.WriteString("}\n\n")

	return nil
}

func (inst *ConfigGoBuilder) writeComponentInject(comp *ComponentConfig) error {

	// TODO
	return nil
}
