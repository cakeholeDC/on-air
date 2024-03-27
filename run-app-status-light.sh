#!/bin/bash
cd "$(dirname "$0")" || exit;

PATH="$PATH:/usr/bin:$HOME/.local/bin"

poetry run python app.py
