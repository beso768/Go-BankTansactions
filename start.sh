#!/bin/sh

set -e

echo "run db migration"
/app/migrate.linux-amd64 -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"