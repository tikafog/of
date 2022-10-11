#!/bin/bash

DBS=("bootnode" "kernel" "media" "upgrade")
DBNAME=$1

SHELL_FOLDER=$(cd "$(dirname "$0")" || exit;pwd)
echo "Current wd:${SHELL_FOLDER}"

function removeFiles() {
    rm -vf "$SHELL_FOLDER/$1"/*.go
    find "$SHELL_FOLDER/$1"/* -type d ! -name 'schema' -print0 | xargs -0 rm -rvf
    return
}

function generateSchema() {
    go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert --feature sql/lock \
--template "$SHELL_FOLDER"/template \
--template glob="$SHELL_FOLDER/template/*.tmpl" "$SHELL_FOLDER/$1"/schema
    return
}

for i in "${!DBS[@]}" ; do
    dbname=${DBS[$i]}
    if [ ! -d "${dbname}" ]; then
      echo "Directory not found: ${dbname}, skipping ${dbname} generation!"
    continue
    fi

    echo "Process $dbname module"

    echo " ->delete old $dbname schema file:"
    removeFiles "$dbname"

    echo " ->generate new $dbname schema:"
    generateSchema "$dbname"
done