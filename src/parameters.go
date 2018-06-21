package main

import (
	"flag"
)

const (
	initUsage    = "Use this command to create a basic yaml configuration file."
	previewUsage = "Use this flag to run ezrep and display all of the changes that would have been made."
	configUsage  = "The alternate full file name to the ezrep.json file to use during execution."
	rootUsage    = "The path to the directory in which to search and make changes."
	exportUsage  = "Use this flag to parse the input arguments and print out the value of the desired variable."
)

const (
	initDefault    = false
	previewDefault = false
	configDefault  = "ezrep.yaml"
	rootDefault    = "./"
	exportDefault  = ""
)

type undefinedArgs struct {
	isInit       bool
	isPreview    bool
	configFile   string
	rootPath     string
	exportValue  string
	initFlags    *flag.FlagSet
	executeFlags *flag.FlagSet
	exportFlags  *flag.FlagSet
}

func newUndefinedArgs() *undefinedArgs {
	undefined := undefinedArgs{
		isInit:       initDefault,
		isPreview:    previewDefault,
		configFile:   configDefault,
		rootPath:     rootDefault,
		exportValue:  exportDefault,
		initFlags:    flag.NewFlagSet("init", flag.PanicOnError),
		executeFlags: flag.NewFlagSet("execute", flag.PanicOnError),
		exportFlags:  flag.NewFlagSet("export", flag.PanicOnError),
	}

	undefined.initFlags.BoolVar(&undefined.isInit, "init", initDefault, initUsage)
	undefined.initFlags.BoolVar(&undefined.isInit, "i", initDefault, "init (shorthand)")
	undefined.initFlags.StringVar(&undefined.configFile, "config", configDefault, configUsage)
	undefined.initFlags.StringVar(&undefined.configFile, "c", configDefault, "config (shorthand)")

	undefined.executeFlags.BoolVar(&undefined.isPreview, "preview", previewDefault, previewUsage)
	undefined.executeFlags.BoolVar(&undefined.isPreview, "p", previewDefault, "preview (shorthand)")
	undefined.executeFlags.StringVar(&undefined.rootPath, "root", rootDefault, rootUsage)
	undefined.executeFlags.StringVar(&undefined.rootPath, "r", rootDefault, "root (shorthand)")
	undefined.executeFlags.StringVar(&undefined.configFile, "config", configDefault, configUsage)
	undefined.executeFlags.StringVar(&undefined.configFile, "c", configDefault, "config (shorthand)")

	undefined.exportFlags.StringVar(&undefined.exportValue, "export", exportDefault, exportUsage)
	undefined.exportFlags.StringVar(&undefined.exportValue, "e", exportDefault, "export (shorthand)")
	undefined.exportFlags.StringVar(&undefined.configFile, "config", configDefault, configUsage)
	undefined.exportFlags.StringVar(&undefined.configFile, "c", configDefault, "config (shorthand)")

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

type executeArgs struct {
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
	runExecute runMode = iota
)

type parameters struct {
	runMode     runMode
	initArgs    initArgs
	exportArgs  exportArgs
	executeArgs executeArgs
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
	case osArgs[1] == "-i" || osArgs[1] == "-init":
		undefined.initFlags.Parse(osArgs[1:])
		undefined.initFlags.VisitAll(debug)
		return parameters{
			runMode: runInit,
			initArgs: initArgs{
				configFile: undefined.configFile,
			},
		}
	case osArgs[1] == "-e" || osArgs[1] == "-export":
		undefined.exportFlags.Parse(osArgs[1:])
		undefined.exportFlags.VisitAll(debug)
		return parameters{
			runMode: runExport,
			exportArgs: exportArgs{
				configFile: undefined.configFile,
				variable:   undefined.exportValue,
				inputs:     undefined.exportFlags.Args(),
			},
		}
	default:
		undefined.executeFlags.Parse(osArgs[1:])
		undefined.executeFlags.VisitAll(debug)
		return parameters{
			runMode: runExecute,
			executeArgs: executeArgs{
				isPreview:  undefined.isPreview,
				configFile: undefined.configFile,
				rootPath:   undefined.rootPath,
				inputs:     undefined.executeFlags.Args(),
			},
		}
	}
}
