package autoconfig

import (
	"errors"
	"fmt"
	"os"

	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

// ConfigPathFinder 是一个配置路径查找器
type ConfigPathFinder struct {
}

// Init 方法初始化一个 CommandContext 对象
func (inst *ConfigPathFinder) Init() (*CommandContext, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	pwd := fs.Default().GetPath(path)
	ctx2 := &CommandContext{
		PWD: pwd,
	}
	return ctx2, nil
}

// FindGoMod 方法从当前路径向上查找 go.mod 文件
func (inst *ConfigPathFinder) FindGoMod(ctx *CommandContext) (*GoModFileContext, error) {

	path1 := ctx.PWD
	var pathGoMod fs.Path = nil

	for {
		file := path1.GetChild("go.mod")

		if file.Exists() && file.IsFile() {
			pathGoMod = file
			break
		}

		path2 := path1.Parent()
		if path2 == nil {
			break
		} else {
			path1 = path2
		}
	}

	if pathGoMod == nil {
		return nil, errors.New("cannot find go.mod in path:" + ctx.PWD.Path())
	}

	ctx2 := &GoModFileContext{
		Path:   pathGoMod,
		parent: ctx,
	}

	return ctx2, nil
}

// FindGssConfig 查找项目的配置
func (inst *ConfigPathFinder) FindGssConfig(ctx *GoModFileContext) (*GssConfigFileContext, error) {

	goModFile := ctx.Path
	projectDir := goModFile.Parent()
	gssConfigFile := projectDir.GetChild("/app.config.json")

	if !gssConfigFile.IsFile() {
		return nil, errors.New("the file is not exists: " + gssConfigFile.Path())
	}

	result := &GssConfigFileContext{
		Path:   gssConfigFile,
		parent: ctx,
	}

	return result, nil
}

// FindConfigPackageDir  查找配置包目录
func (inst *ConfigPathFinder) FindConfigPackageDir(ctx *GssConfigFileContext) (*PackageDirectoryContext, error) {

	cfgFile := ctx.Path
	pkgDirHref := ctx.DOM.Package

	if pkgDirHref == "" {
		return nil, errors.New("no string value with name 'package' in file " + cfgFile.Path())
	}

	pkgDirPath := cfgFile.GetHref(pkgDirHref)
	if !pkgDirPath.IsDir() {
		return nil, errors.New("the directory is not exists: " + pkgDirPath.Path())
	}
	pdContext := &PackageDirectoryContext{
		Path:   pkgDirPath,
		parent: ctx,
	}
	return pdContext, nil
}

// FindConfigFiles 查找配置文件
func (inst *ConfigPathFinder) FindConfigFiles(ctx *PackageDirectoryContext) ([]*SourceFileContext, error) {
	return nil, nil
}

// FindSourceFiles 查找配置源文件
func (inst *ConfigPathFinder) FindSourceFiles(ctx *PackageDirectoryContext) ([]*SourceFileContext, error) {

	if !ctx.Path.IsDir() {
		return nil, errors.New("the path is not a dir: " + ctx.Path.Path())
	}

	list := ctx.Path.GetItemList()
	dest := make([]*SourceFileContext, 0)

	for index := range list {
		item := list[index]
		fmt.Println("find source " + item.Path())
		source := &SourceFileContext{
			Path:   item,
			parent: ctx,
		}
		dest = append(dest, source)
	}

	ctx.Sources = dest
	return dest, nil
}
