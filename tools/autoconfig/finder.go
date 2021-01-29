package autoconfig

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

// ConfigPathFinder 是一个配置路径查找器
type ConfigPathFinder struct {
}

// Init 方法初始化一个 CommandContext 对象
func (inst *ConfigPathFinder) Init(ctx *Context) error {

	args := ctx.Args

	if args == nil {
		args = &CommandArgs{}
		ctx.Args = args
	}

	pwd := args.PWD
	path, err := os.Getwd()

	if err != nil {
		return err
	}

	if pwd == nil {
		pwd = fs.Default().GetPath(path)
		args.PWD = pwd
	}

	return nil
}

// FindGoMod 方法从当前路径向上查找 go.mod 文件
func (inst *ConfigPathFinder) FindGoMod(ctx *Context) error {

	pwd := ctx.Args.PWD
	path1 := pwd
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
		return errors.New("cannot find go.mod in path:" + pwd.Path())
	}

	ctx.GoMod = &GoModFile{
		Path:    pathGoMod,
		Context: ctx,
	}

	return nil
}

// FindProjectConfig 查找项目的配置
func (inst *ConfigPathFinder) FindProjectConfig(ctx *Context) error {

	filename := ctx.Args.ConfigFileName

	if filename == "" {
		filename = "project.config.json"
	}

	goModFile := ctx.GoMod.Path
	projectDir := goModFile.Parent()
	gssConfigFile := projectDir.GetChild(filename)

	if !gssConfigFile.IsFile() {
		return errors.New("the file is not exists: " + gssConfigFile.Path())
	}

	ctx.ProjectConfig = &ProjectConfigFile{
		Path:    gssConfigFile,
		Context: ctx,
	}

	loader := &GssConfigLoader{}
	err := loader.Load(ctx.ProjectConfig)
	return err
}

// FindConfigPackageDir  查找配置包目录
func (inst *ConfigPathFinder) FindConfigPackageDir(ctx *Context) error {

	cfgFile := ctx.ProjectConfig.Path
	pkgDirHref := ctx.ProjectConfig.DOM.Package

	if pkgDirHref == "" {
		return errors.New("no string value with name 'package' in file " + cfgFile.Path())
	}

	pkgDirPath := cfgFile.GetHref(pkgDirHref)
	if !pkgDirPath.IsDir() {
		return errors.New("the directory is not exists: " + pkgDirPath.Path())
	}

	ctx.PackageDirectory = &PackageDirectory{
		Path:    pkgDirPath,
		Context: ctx,
	}

	return nil
}

// FindSourceFiles 查找配置源文件
func (inst *ConfigPathFinder) FindSourceFiles(ctx *Context) error {

	dir := ctx.PackageDirectory.Path

	if !dir.IsDir() {
		return errors.New("the path is not a dir: " + dir.Path())
	}

	list := dir.GetItemList()
	dest := make([]*SourceFile, 0)

	for index := range list {
		item := list[index]
		filename := item.Name()
		if !strings.HasSuffix(filename, ".go") {
			continue
		}
		fmt.Println("find source " + item.Path())
		source := &SourceFile{
			Path:    item,
			Context: ctx,
		}
		dest = append(dest, source)
	}

	ctx.PackageDirectory.Sources = dest
	return nil
}
