package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Usage() {
	fmt.Println("Usage: goenv [destination_folder]")
}

func EnsurePathExists(path string) {
	os.MkdirAll(path, 0755)
}

func CheckIfExists(path string) bool {
	if _, err := os.Stat(path) ; err != nil {
		return ! os.IsNotExist(err)
	} else {
		return true
	}
}

func CreateGoEnv(path string) {
	if CheckIfExists(path) {
		fmt.Println("Refreshing existing goenv in", path)
	} else {
		fmt.Println("Creating new goenv in", path)
	}
	EnsurePathExists(path)
	binPath := fmt.Sprintf("%s/bin", path)
	EnsurePathExists(binPath)
	WriteActivateScript(binPath, path)
	EnsurePathExists(fmt.Sprintf("%s/pkg", path))
	EnsurePathExists(fmt.Sprintf("%s/src", path))
}

func WriteActivateScript(binPath string, path string) {

	activateFishScript := `# This file must be used with "source bin/activate.fish" *from fish* (http://fishshell.com)
# you cannot run it directly

set -x PARENT_PATH ` + path + `

if test -z $GO_ENV
    set -gx GO_ENV $PARENT_PATH
    echo "Activating goenv in $GO_ENV"
    set -Ux GOPATH $GO_ENV
    set -gx BIN_PATH "$GOPATH/bin"
    set -gx OLD_PATH $PATH
    set PATH $PATH $BIN_PATH
else if test $GO_ENV != $PARENT_PATH
    set -gx GO_ENV $PARENT_PATH
    echo "Switching to goenv in $GO_ENV"
    set -Ux GOPATH $GO_ENV

    set -gx BIN_PATH "$GOPATH/bin"
    set PATH $OLD_PATH $BIN_PATH
else
    echo "Already activated goenv in $GO_ENV"
end

function get_deps
    go get >/dev/null 2>&1
end

function install_sh
    set -l LS_RESULTS (eval "ls -1 *.sh 2>/dev/null | wc -l")
    if test $LS_RESULTS != "0"
        chmod 755 *.sh
        cp *.sh $BIN_PATH/
    end
end

function build
    get_deps
    set -g BUILD_RESULT (eval "go install 2>&1")
    if test -n $BUILD_RESULT
        echo $BUILD_RESULT
    else
        install_sh
    end
end

function list_installables
    set -l LS_RESULTS (eval "ls -1 $BIN_PATH/* 2>/dev/null | grep -v activate")
    for result in $LS_RESULTS
        echo (basename $result)
    end
end

function install_to_system
    for installable in (eval list_installables)
        sudo install -m 755 $BIN_PATH/$installable /usr/local/bin/$installable
    end
end

function make_install
    set -e BUILD_RESULT
    build_all
    if test -n $BUILD_RESULT
        echo $BUILD_RESULT
    else
        install_to_system
    end
end


function uninstall_from_system
    for installable in (eval list_installables)
        if test -e /usr/local/bin/$installable
            sudo rm /usr/local/bin/$installable
        end
    end
end


function make_uninstall
    build_all
    if test -n $BUILD_RESULT
        echo $BUILD_RESULT
    else
        uninstall_from_system
    end
end


function deactivate
    set PATH $OLD_PATH
    set -e BIN_PATH
    set -e GO_ENV
    set -e GOPATH
    set -e OLD_PATH
    functions -e get_deps install_sh build list_installables install_to_system
    functions -e make_install uninstall_from_system make_uninstall
    functions -e list_buildable_folders build_all make_clean deactivate
end


function list_buildable_folders
    find $GO_ENV -iname '*.go' | xargs -L 1 dirname $1 | sort -u
end


function build_all
    set -l OLD_PWD (eval pwd)
    for folder in (eval list_buildable_folders)
        cd $folder
        build
        if test -n $BUILD_RESULT
            echo $BUILD_RESULT
        end
    end
    cd $OLD_PWD
end

function make_clean
    for installable in (eval list_installables)
        rm $BIN_PATH/$installable
    end
end
`
	ioutil.WriteFile(fmt.Sprintf("%s/activate", binPath), []byte(activateScript), 0755)
	ioutil.WriteFile(fmt.Sprintf("%s/activate.fish", binPath), []byte(activateFishScript), 0755)
}
