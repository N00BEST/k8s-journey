#!/bin/bash

set -euo pipefail

readonly SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && printf "%q\n" "$(pwd)" )"

pushd "$SCRIPT_DIR" &> /dev/null
    cd .. 
    readonly PARENT_DIR="$( printf "%q\n" "$(pwd)" )"
popd &> /dev/null

readonly COMPILE_SCRIPT="$SCRIPT_DIR/compile.sh"

declare -a SYSTEM_FILES
SYSTEM_FILES=(namespace)

declare -a DATABASE_FILES
DATABASE_FILES=(secret configmap volume persistentvolumeclaim deployment service)

declare -a API_FILES
API_FILES=(deployment service)

declare -a filesToMerge 

readonly FILES_SOURCE_DIR="$PARENT_DIR/source"


i=0

# Functions 

addFile() {
    echo "$1/$2-$3.yaml"
}

# main
for file in "${SYSTEM_FILES[@]}"; do
    filesToMerge[i]=$(addFile $FILES_SOURCE_DIR "system" $file)
    ((i+=1))
done

for file in "${DATABASE_FILES[@]}"; do
    filesToMerge[i]=$(addFile $FILES_SOURCE_DIR "database" $file)
    ((i+=1))
done

for file in "${API_FILES[@]}"; do
    filesToMerge[i]=$(addFile $FILES_SOURCE_DIR "api" $file)
    ((i+=1))
done

$( $COMPILE_SCRIPT ${filesToMerge[@]} )

echo "Deployment finished properly"
exit 0