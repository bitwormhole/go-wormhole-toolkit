package autoconfig

// Run 函数运行命令“autoconfig”的功能
func Run(args *CommandArgs) (*Context, error) {

	context := &Context{}
	context.Args = args
	finder := &ConfigPathFinder{}

	err := finder.Init(context)
	if err != nil {
		return nil, err
	}

	err = finder.FindGoMod(context)
	if err != nil {
		return nil, err
	}

	err = finder.FindProjectConfig(context)
	if err != nil {
		return nil, err
	}

	err = finder.FindConfigPackageDir(context)
	if err != nil {
		return nil, err
	}

	err = finder.FindSourceFiles(context)
	if err != nil {
		return nil, err
	}

	sourcefiles := context.PackageDirectory.Sources

	for index := range sourcefiles {
		item := sourcefiles[index]
		err = item.LoadAST()
		if err != nil {
			return nil, err
		}
	}

	return context, nil
}
