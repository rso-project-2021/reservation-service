#!/bin/sh

set -e

echo "run db migration"
db_source=$(jq -r .db_source config.json)
/app/migrate -path /app/migration -database "$db_source" -verbose up

echo "start the app"
exec "$@"