package autoconfig

import (
	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

// CommandContext 表示一个命令上下文
type CommandContext struct {
	PWD fs.Path
}

// GoModFileContext 表示一个围绕 go.mod 文件周边的环境
type GoModFileContext struct {
	parent *CommandContext
	Path   fs.Path
}

// GssConfigFileContext 表示一个围绕 gss.config.js 文件周边的环境
type GssConfigFileContext struct {
	parent *GoModFileContext
	Path   fs.Path
	DOM    *ConfigDOM
}

// PackageDirectoryContext 表示一个围绕 go package 文件夹周边的环境
type PackageDirectoryContext struct {
	Path    fs.Path
	parent  *GssConfigFileContext
	Builder *ConfigGoBuilder
	Sources []*SourceFileContext
}

// SourceFileContext 表示一个围绕 *.go 源文件周边的环境
type SourceFileContext struct {
	parent *PackageDirectoryContext
	Path   fs.Path
}
