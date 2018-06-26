# EZREP 

## Usage of init:

| Command     | Description                                                                                          |
| ----------- | ---------------------------------------------------------------------------------------------------- |
| -c, -config | The alternate full file name to the ezrep.json file to use  during execution. (default "ezrep.yaml") |
| -i, -init   | Use this command to create a basic yaml configuration file.                                          |

## Usage of execute:

| Command      | Description                                                                                         |
| ------------ | --------------------------------------------------------------------------------------------------- |
| -c, -config  | The alternate full file name to the ezrep.json file to use during execution. (default "ezrep.yaml") |
| -p, -preview | Use this flag to run ezrep and display all of the changes that would have been made.                |
| -r, -root    | The path to the directory in which to search and make changes. (default "./")                       |

## Usage of export:

| Command     | Description                                                                                         |
| ----------- | --------------------------------------------------------------------------------------------------- |
| -c, -config | The alternate full file name to the ezrep.json file to use during execution. (default "ezrep.yaml") |
| -e, -export | Use this flag to parse the input arguments and print out the value of the desired variable.         |

## Examples:

Initialize a project with a default configuration file ->

```shell
  ezrep -i
  ezrep -init
  ezrep -i -c myconfig.yml
  ezrep -init -config myconfig.yml
```

Export a variable to stdout ->

```shell
  ezrep -e Version $VAR1 $VAR2 ...
  ezrep -export Version $VAR1 $VAR2 ...
  ezrep -e Version -c myconfig.yaml $VAR1 $VAR2 ...
  ezrep -export Version -config myconfig.yaml $VAR1 $VAR2 ...
```

Execute changes to files ->

```shell
  exrep $VAR1 $VAR2 ...
  exrep -p $VAR1 $VAR2 ...
  exrep -preview $VAR1 $VAR2 ...
  exrep -c myconfig.yml -r ./src $VAR1 $VAR2 ...
  exrep -config myconfig.yml -root ./src $VAR1 $VAR2 ...
```

## A practical application

If you wanted to replace all of the version numbers in your assembly info &| project 
files, you could do the following:

** Create an ezrep.yaml file **

```shell
exrep -init
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
ezrep "Version=1.2.3.4"
```
