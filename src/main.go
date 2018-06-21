package main

import (
	"fmt"
	"os"
)

var undefined = newUndefinedArgs()

func main() {
	parameters := parseParameters(undefined, os.Args)

	switch parameters.runMode {
	case runInit:
		doInit(parameters.initArgs)
	case runExport:
		doExport(parameters.exportArgs)
	case runExecute:
		doExecute(parameters.executeArgs)
	case runHelp:
		doHelp(undefined)
	}
}

func doInit(args initArgs) {
	file := defaultFileIo()
	yaml := basicYAML()
	config := []byte(yaml)

	// bug: do not overwrite existing file
	fmt.Println(args.configFile)
	file.writeBytes(args.configFile, config)
}

func doExport(args exportArgs) {
	console := defaultConsoleIo()
	file := defaultFileIo()
	config := loadConfig(file, args.configFile)
	variables := config.Definitions.execute(args.inputs)

	console.write(variables[args.variable].value)
}

func doExecute(args executeArgs) {
	var system systemIo

	if args.isPreview {
		system = previewSystemIo()
	} else {
		system = defaultSystemIo()
	}

	config := loadConfig(system.file, args.configFile)
	filenames := system.directory.listFiles(args.rootPath)

	variables := config.Definitions.execute(args.inputs)
	variables.print(system.console)

	if len(variables) == 0 {
		system.console.writeLn("No variables found to process.  Exiting ezrep...")
		return
	}

	updates := config.Applications.execute(system.file, variables, filenames)
	updates.print(system.console)
}

func doHelp(args *undefinedArgs) {
	console := defaultConsoleIo()

	console.writeLn("")
	console.writeLn("EZREP HELP")
	console.writeLn("------------------------------------------------------------")
	console.writeLn("")
	args.initFlags.Usage()
	console.writeLn("")
	args.executeFlags.Usage()
	console.writeLn("")
	args.exportFlags.Usage()
	console.writeLn("")
	console.writeLn("Examples:")
	console.writeLn("")
	console.writeLn("Initialize a project with a default configuration file ->")
	console.writeLn("  ezrep -i")
	console.writeLn("  ezrep -init")
	console.writeLn("  ezrep -i -c myconfig.yml")
	console.writeLn("  ezrep -init -config myconfig.yml")
	console.writeLn("")
	console.writeLn("Export a variable to stdout ->")
	console.writeLn("  ezrep -e Version $VAR1 $VAR2 ...")
	console.writeLn("  ezrep -export Version $VAR1 $VAR2 ...")
	console.writeLn("  ezrep -e Version -c myconfig.yaml $VAR1 $VAR2 ...")
	console.writeLn("  ezrep -export Version -config myconfig.yaml $VAR1 $VAR2 ...")
	console.writeLn("")
	console.writeLn("Execute changes to files ->")
	console.writeLn("  exrep $VAR1 $VAR2 ...")
	console.writeLn("  exrep -p $VAR1 $VAR2 ...")
	console.writeLn("  exrep -preview $VAR1 $VAR2 ...")
	console.writeLn("  exrep -c myconfig.yml -r ./src $VAR1 $VAR2 ...")
	console.writeLn("  exrep -config myconfig.yml -root ./src $VAR1 $VAR2 ...")
	console.writeLn("")
}
