#!/bin/bash
make workspace
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo ". ~/.bashrc" > /tmp/go_activate.$$
make env >> /tmp/go_activate.$$
if [ -f .bash_local ]
then
    echo ". $DIR/.bash_local" >> /tmp/go_activate.$$
else
    echo "No .bash_local, using default from bash_local.example"
    cat bash_local.example >> /tmp/go_activate.$$
fi

$(make env)

ws="$(git remote | grep workspace)"
if [ ! -n "$ws" ]
then
    echo "adding git remote workspace"
    git remote add workspace $project
fi

echo "Entering workspace, exit shell to escape"
bash --init-file /tmp/go_activate.$$ -i

rm /tmp/go_activate.$$

if [ -n "$(cd $project && git status -s)" ]
then
    echo "Files in workspace exist which are not checked in:"
    (cd $project && git status -s)
else
    echo "Exiting workspace"
fi
echo
this="$(git log -n 1 | grep commit | awk '{print $2}')"
if [ -n "$(cd $project && git log $this..HEAD)" ]
then
    echo "changes checked into workspace not present here:"
    echo 
    (cd $project && git log $this..HEAD)
fi

