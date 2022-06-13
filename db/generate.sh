#!/bin/bash

DBS=["bootnode"]
DBNAME=$1

SHELL_FOLDER=$(cd "$(dirname "$0")" || exit;pwd)

echo "${SHELL_FOLDER}"

if [ -z "$DBNAME" ]; then
  echo "Usage: ./generate.sh <dbname>"
  exit 0
fi


echo "delete old schema file"
rm -vf "$SHELL_FOLDER"/"$DBNAME"/*.go
find "$SHELL_FOLDER"/"$DBNAME"/* -type d ! -name 'schema' -print0 | xargs -0 rm -rvf
echo "generate new schema"
#ent generate "$SHELL_FOLDER"/"$DBNAME"/ent/schema
echo -e "ent generate --template $SHELL_FOLDER/internal/template \
--template glob=\"$SHELL_FOLDER/internal/template/*.tmpl\" $SHELL_FOLDER/$DBNAME/schema"
go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert --feature sql/lock \
--template "$SHELL_FOLDER"/internal/template \
--template glob="$SHELL_FOLDER/internal/template/*.tmpl" "$SHELL_FOLDER"/"$DBNAME"/schema
