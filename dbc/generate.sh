#!/bin/bash

DBS=("bootnode" "kernel")
DBNAME=$1

SHELL_FOLDER=$(cd "$(dirname "$0")" || exit;pwd)
echo "Current wd:${SHELL_FOLDER}"

function fileRemove() {
    rm -vf "$SHELL_FOLDER/$1"/*.go
    find "$SHELL_FOLDER/$1"/* -type d ! -name 'schema' -print0 | xargs -0 rm -rvf
    return
}

function schemaGenerate() {
    go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert --feature sql/lock \
--template "$SHELL_FOLDER"/template \
--template glob="$SHELL_FOLDER/template/*.tmpl" "$SHELL_FOLDER/$1"/schema
    return
}

for i in "${!DBS[@]}" ; do
    dbname=${DBS[$i]}
    if [ ! -d "${dbname}" ]; then
    continue
    fi

    echo "Process $dbname module"

    echo "Delete old schema file"
    fileRemove "$dbname"

    echo "Generate new schema"
    schemaGenerate "$dbname"
done