package autoconfig

import (
	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

// CommandArgs  表示一个命令上下文
type CommandArgs struct {
	Context        *Context
	PWD            fs.Path
	ConfigFileName string // default("") = project.config.json
}

// GoModFile  表示一个围绕 go.mod 文件周边的环境
type GoModFile struct {
	Context *Context
	Path    fs.Path
}

// ProjectConfigFile  表示一个围绕 gss.config.json 文件周边的环境
type ProjectConfigFile struct {
	Context *Context
	Path    fs.Path
	DOM     *ConfigDOM
}

// PackageDirectory  表示一个围绕 go package 文件夹周边的环境
type PackageDirectory struct {
	Context *Context
	Path    fs.Path
	Builder *ConfigGoBuilder
	Sources []*SourceFile
}

// SourceFile  表示一个围绕 *.go 源文件周边的环境
type SourceFile struct {
	Context *Context
	Path    fs.Path
}

// Context  自动配置工具的上下文
type Context struct {
	Args             *CommandArgs
	GoMod            *GoModFile
	ProjectConfig    *ProjectConfigFile
	PackageDirectory *PackageDirectory
}
