# EZREP 

[![pipeline status](https://gitlab.com/adleatherwood/ezrep/badges/develop/pipeline.svg)](https://gitlab.com/adleatherwood/ezrep/commits/develop)

## Options for init:

| Command | Usage    | Description                                                                                          |
| ------- | -------- | ---------------------------------------------------------------------------------------------------- |
| -c      | optional | The alternate full file name to the ezrep.json file to use  during execution. (default "ezrep.yaml") |

## Options for export:

| Command | Usage    | Description                                                                                         |
| ------- | -------- | --------------------------------------------------------------------------------------------------- |
| -c      | optional | The alternate full file name to the ezrep.json file to use during execution. (default "ezrep.yaml") |
| -v      | required | The name of the variable from the config to searh for and return.                                   |

## Options for process:

| Command | Usage    | Description                                                                                         |
| ------- | -------- | --------------------------------------------------------------------------------------------------- |
| -c      | optional | The alternate full file name to the ezrep.json file to use during execution. (default "ezrep.yaml") |
| -p      | optional | Use this flag to run ezrep and display all of the changes that would have been made.                |
| -r      | optional | The path to the directory in which to search and make changes. (default "./")                       |

## Examples:

Initialize a project with a default configuration file ->

```shell
  ezrep init
  ezrep init -c myconfig.yml
```

Export a variable to stdout ->

```shell
  ezrep export -v Version $VAR1 $VAR2 ...
  ezrep export -v Version -c myconfig.yaml $VAR1 $VAR2 ...
```

Process changes to files ->

```shell
  exrep process $VAR1 $VAR2 ...
  exrep process -p -c myconfig.yml -r ./src $VAR1 $VAR2 ...
```

## A practical application

If you wanted to replace all of the version numbers in your assembly info &| project 
files, you could do the following:

** Create an ezrep.yaml file **

```shell
ezrep init
```

** Modify the contents **

```yaml
variables: 
- name  : Version
  find  : ([vV]ersion=)([\d\.]+)(.*) 
  group : 2         
tasks: 
- variable : Version
  filemask : AssemblyInfo\.cs
  find     : ([vV]ersion\(")([\d\.]+)("\))
  replace  : ${1}%s${3}
- variable : Version
  filemask : \.[cf]sproj
  find     : (<[vV]ersion>)([\d\.]+)(<\/[vV]ersion>)
  replace  : ${1}%s${3}
```

** Execute the application **

```
ezrep process "Version=1.2.3.4"
```

[Icon Source](http://www.iconarchive.com/show/minimal-fruit-icons-by-alex-t.html)
