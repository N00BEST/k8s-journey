#!/bin/bash

set -euxo pipefail

readonly SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && printf "%q\n" "$(pwd)" )"

pushd "$SCRIPT_DIR"
    cd .. 
    readonly PARENT_DIR="$( printf "%q\n" "$(pwd)" )"
popd

readonly COMPILE_SCRIPT="$SCRIPT_DIR/compile.sh"

declare -a DATABASE_FILES
DATABASE_FILES=(secret configmap volume persistentvolumeclaim deployment service)

declare -a API_FILES
API_FILES=(deployment service)

declare -a filesToMerge 

echo "SCRIPT DIR: $SCRIPT_DIR"
echo "PARENT DIR: $PARENT_DIR"
exit 0

$(
    pushd "$PARENT_DIR/source"
        readonly FILES_SOURCE_DIR="$( printf "%q\n" "$(pwd)" )"
    popd
)

i=0

# Functions 

addFile() {
    echo "$1/$2.yaml"
}

# main

for file in "${DATABASE_FILES[@]}"; do
    filesToMerge[i]=$(addFile $FILES_SOURCE_DIR $file)
    ((i+=1))
done

for file in "${API_FILES[@]}"; do
    filesToMerge[i]=$(addFile $FILES_SOURCE_DIR $file)
    ((i+=1))
done

for file in "${filesToMerge[@]}"; do
    echo $file
done

exit 0
$( $COMPILE_SCRIPT ${filesToMerge[@]} )