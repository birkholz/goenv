# README #

## goenv ##
* goenv is a simple tool that mimics python's virtualenv, for use in go projects.
* Version 0.1
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### Set up ###

Clone repository:

```
#!bash

hg clone https://bitbucket.org/fernandosanchezjr/goenv goenv
```
Enter goenv folder, activate, and install:

```
#!bash

cd goenv
source bin/activate
make_install
```
Optionally, you may deactivate the goenv:

```
#!bash

deactivate
```

### Usage ###

Once goenv is installed, you may use it similarly to virtualenv

```
#!bash

goenv ~/projects/test_go_project
cd ~/projects/test_go_project
source bin/activate
```
Note: you may source the activate script from any location

```
#!bash

source ~/projects/test_go_project/bin/activate
```
#### Built-in shell commands ####
goenv sets a few built-in commands for convenience

```
#!bash

build
build_all
make_clean
make_install
make_uninstall
deactivate
```

* build: builds go files in current folder using **go install**
* build_all: builds go files in all folders using **go install**
* make_clean: cleans all files generated in the active goenv's bin folder during compilation (except for activate script)
* make_install: builds all folders, and installs all files contained in the active goenv's bin folder to /usr/local/bin
* make_uninstall: builds all folders, and uninstalls all files contained in the active goenv's bin folder from /usr/local/bin if they exist
* deactivate: deactivates the goenv

### Dependencies ###

* bash
* find
* install
* A working Go 1.3+ installation

### Contribution guidelines ###

* Have a bug or an issue? Have a pull request? File an issue on the issue tracker!