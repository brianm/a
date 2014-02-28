#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [ "$1" = "activate" ]
then    
    $SHELL --init-file $DIR/$0 -i
    exit 0
fi

. ~/.bashrc

export PROMPT_DIRTRIM=2
export GOPATH=$DIR/_workspace
export PATH=$DIR/_workspace/bin/:$PATH
plum="\[\033[38;5;54m\]"
reset="\[\033[00m\]"
export PS1="[${plum}$PROJECT${reset}] $PS1"
cd _workspace/src/$PACKAGE
