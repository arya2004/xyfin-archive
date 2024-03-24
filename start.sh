#!/bin/sh

set -e

echo "run database migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start server"
exec "$@"