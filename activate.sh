#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo ". ~/.bashrc" > /tmp/go_activate.$$
make env >> /tmp/go_activate.$$
if [ -f .bash_local ]
then
    echo ". $DIR/.bash_local" >> /tmp/go_activate.$$
fi

echo "Entering workspace, exit shell to escape"
bash --init-file /tmp/go_activate.$$ -i

echo "Exiting workspace"
rm /tmp/go_activate.$$
