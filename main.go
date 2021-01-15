package main

import (
	"fmt"

	"github.com/bitwormhole/go-wormhole-toolkit/tools/app-auto-configurer/autoconfig"
)

func main() {
	fmt.Println("hello, gss-auto-configurer")

	finder := &autoconfig.ConfigPathFinder{}

	cmdContext, err := finder.Init()
	if err == nil {
		fmt.Println("PWD=" + cmdContext.PWD.Path())
	} else {
		fmt.Println(err)
		return
	}

	goMod, err := finder.FindGoMod(cmdContext)
	if err == nil {
		fmt.Println("go.mod=" + goMod.Path.Path())
	} else {
		fmt.Println(err)
		return
	}

	gssConfig, err := finder.FindGssConfig(goMod)
	if err == nil {
		fmt.Println("gss.config.json=" + gssConfig.Path.Path())
		loader := &autoconfig.GssConfigLoader{}
		err := loader.Load(gssConfig)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println(err)
		return
	}

	pkg, err := finder.FindConfigPackageDir(gssConfig)
	if err == nil {
		fmt.Println("package_dir=" + pkg.Path.Path())
	} else {
		fmt.Println(err)
		return
	}

	finder.FindSourceFiles(pkg)

	fmt.Println("done.")
}
