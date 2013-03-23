#!/bin/bash
echo ". ~/.bashrc" > /tmp/go_env.$$
make env >> /tmp/go_env.$$

echo "Entering WORKSPACE, exit shell to escape"
bash --init-file /tmp/go_env.$$ -i

echo "Exiting WORKSPACE"
rm /tmp/go_env.$$
