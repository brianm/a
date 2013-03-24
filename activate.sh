#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo ". ~/.bashrc" > /tmp/go_env.$$
make env >> /tmp/go_env.$$
if [ -f .bash_local ]
then
    echo ". $DIR/.bash_local" >> /tmp/go_env.$$
fi

echo "Entering WORKSPACE, exit shell to escape"
bash --init-file /tmp/go_env.$$ -i

echo "Exiting WORKSPACE"
rm /tmp/go_env.$$
