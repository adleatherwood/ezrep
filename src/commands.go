package main

import (
	"flag"
)

const (
	previewUsage  = "Use this flag to run ezrep and display all of the changes that would have been made."
	configUsage   = "The alternate full file name to the ezrep.json file to use during execution."
	rootUsage     = "The path to the directory in which to search and make changes."
	variableUsage = "The name of the variable from the config to searh for and return."
)

const (
	initDefault     = false
	previewDefault  = false
	configDefault   = "ezrep.yaml"
	rootDefault     = "./"
	variableDefault = ""
)

type undefinedArgs struct {
	isInit         bool
	isPreview      bool
	configFile     string
	rootPath       string
	variabletValue string
	initFlags      *flag.FlagSet
	exportFlags    *flag.FlagSet
	processFlags   *flag.FlagSet
}

func newUndefinedArgs() *undefinedArgs {
	undefined := undefinedArgs{
		isInit:         initDefault,
		isPreview:      previewDefault,
		configFile:     configDefault,
		rootPath:       rootDefault,
		variabletValue: variableDefault,
		initFlags:      flag.NewFlagSet("init", flag.PanicOnError),
		exportFlags:    flag.NewFlagSet("export", flag.PanicOnError),
		processFlags:   flag.NewFlagSet("execute", flag.PanicOnError),
	}

	undefined.initFlags.StringVar(&undefined.configFile, "c", configDefault, configUsage)
	undefined.exportFlags.StringVar(&undefined.configFile, "c", configDefault, configUsage)
	undefined.exportFlags.StringVar(&undefined.variabletValue, "v", variableDefault, variableUsage)
	undefined.processFlags.StringVar(&undefined.configFile, "c", configDefault, configUsage)
	undefined.processFlags.StringVar(&undefined.rootPath, "r", rootDefault, rootUsage)
	undefined.processFlags.BoolVar(&undefined.isPreview, "p", previewDefault, previewUsage)

	return &undefined
}

type initArgs struct {
	configFile string
}

type exportArgs struct {
	configFile string
	variable   string
	inputs     []string
}

type processArgs struct {
	isPreview  bool
	configFile string
	rootPath   string
	inputs     []string
}

type runMode int

const (
	runHelp    runMode = iota
	runInit    runMode = iota
	runExport  runMode = iota
	runProcess runMode = iota
)

type parameters struct {
	runMode     runMode
	initArgs    initArgs
	exportArgs  exportArgs
	processArgs processArgs
}

func parseParameters(undefined *undefinedArgs, osArgs []string) parameters {
	debug := func(a *flag.Flag) {
		//fmt.Println(">", a.Name, "value=", a.Value)
	}

	switch {
	case len(osArgs) == 1 || osArgs[1] == "-h" || osArgs[1] == "-help":
		return parameters{
			runMode: runHelp,
		}
	case osArgs[1] == "init":
		undefined.initFlags.Parse(osArgs[2:])
		undefined.initFlags.VisitAll(debug)
		return parameters{
			runMode: runInit,
			initArgs: initArgs{
				configFile: undefined.configFile,
			},
		}
	case osArgs[1] == "export":
		undefined.exportFlags.Parse(osArgs[2:])
		undefined.exportFlags.VisitAll(debug)
		return parameters{
			runMode: runExport,
			exportArgs: exportArgs{
				configFile: undefined.configFile,
				variable:   undefined.variabletValue,
				inputs:     undefined.exportFlags.Args(),
			},
		}
	case osArgs[1] == "process":
		undefined.processFlags.Parse(osArgs[2:])
		undefined.processFlags.VisitAll(debug)
		return parameters{
			runMode: runProcess,
			processArgs: processArgs{
				isPreview:  undefined.isPreview,
				configFile: undefined.configFile,
				rootPath:   undefined.rootPath,
				inputs:     undefined.processFlags.Args(),
			},
		}
	default:
		return parameters{
			runMode: runHelp,
		}
	}
}
