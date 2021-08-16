#!/bin/bash

set -euo pipefail

readonly SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
readonly DESTINATION_FILE="$SCRIPT_DIR/../deploy.yaml"

# functions 

timestamp() {
    date +"%Y-%m-%d %H:%M:%S"
}


# main

echo "# Generated: $(timestamp)" > "$DESTINATION_FILE"

for file in "$@"; do
    (cat $file; echo) >> "$DESTINATION_FILE"
done
