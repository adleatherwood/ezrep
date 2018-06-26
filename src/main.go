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
		mainInit(parameters.initArgs)
	case runExport:
		mainExport(parameters.exportArgs)
	case runProcess:
		mainProcess(parameters.processArgs)
	case runHelp:
		mainHelp(undefined)
	}
}

func mainInit(args initArgs) {
	file := defaultFileIo()
	yaml := basicYAML()
	config := []byte(yaml)

	// bug: do not overwrite existing file
	fmt.Println(args.configFile)
	file.writeBytes(args.configFile, config)
}

func mainExport(args exportArgs) {
	console := defaultConsoleIo()
	file := defaultFileIo()
	config := loadConfig(file, args.configFile)
	variableMap := config.Variables.execute(args.inputs)
	variable := variableMap[args.variable].value

	if len(variableMap) == 0 {
		console.writeLn("Variable '%s' not found in input arguments.", args.variable)
	}

	console.write(variable)
}

func mainProcess(args processArgs) {
	var system systemIo

	if args.isPreview {
		system = previewSystemIo()
		system.console.writeLn("Preview mode enabled.  No files will be modified.")
	} else {
		system = defaultSystemIo()
	}

	config := loadConfig(system.file, args.configFile)

	variables, updates := doProcess(system.directory, system.file, config, args.rootPath, args.inputs)
	variables.print(system.console)

	if len(variables) == 0 {
		system.console.writeLn("No variables found to process.  Exiting ezrep...")
	} else {
		updates.print(system.console)
	}
}

func doProcess(directory directoryIo, file fileIo, config config, root string, inputs []string) (variableMap, updates) {
	filenames := directory.listFiles(root)
	variables := config.Variables.execute(inputs)

	var updates []update

	if len(variables) > 0 {
		updates = config.Tasks.execute(file, variables, filenames)
	}

	return variables, updates
}

func mainHelp(args *undefinedArgs) {
	console := defaultConsoleIo()

	console.writeLn("")
	console.writeLn("EZREP HELP")
	console.writeLn("------------------------------------------------------------")
	console.writeLn("")
	args.initFlags.Usage()
	console.writeLn("")
	args.exportFlags.Usage()
	console.writeLn("")
	args.processFlags.Usage()
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
	console.writeLn("  ezrep export Version $VAR1 $VAR2 ...")
	console.writeLn("  ezrep export Version -c myconfig.yaml $VAR1 $VAR2 ...")
	console.writeLn("  ezrep export Version -config myconfig.yaml $VAR1 $VAR2 ...")
	console.writeLn("")
	console.writeLn("Process changes to files ->")
	console.writeLn("  exrep process $VAR1 $VAR2 ...")
	console.writeLn("  exrep process -p $VAR1 $VAR2 ...")
	console.writeLn("  exrep process -preview $VAR1 $VAR2 ...")
	console.writeLn("  exrep process -c myconfig.yml -r ./src $VAR1 $VAR2 ...")
	console.writeLn("  exrep process -config myconfig.yml -root ./src $VAR1 $VAR2 ...")
	console.writeLn("")
}
