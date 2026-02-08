#!/bin/sh

set -e

# echo "run db migration"
# /app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
# moving migration to main.go, before server starts

echo "start the app"
exec "$@"