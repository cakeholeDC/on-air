#!/bin/bash

# cd ~/dev/on-air
cd "$(dirname "$0")";
CWD="$(pwd)"
echo $CWD
echo $PATH
echo $HOME
export PATH="$HOME/.poetry/bin:$PATH"
echo $PATH

poetry run python3 app.py
