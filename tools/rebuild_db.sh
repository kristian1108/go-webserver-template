#!/bin/bash

set -e

DB_NAME=""
DB_USER=""
OLD_DB_NAME="$DB_NAME"_old

psql -U "$USER" -d postgres << EOF
SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE pid <> pg_backend_pid() AND datname = '$DB_NAME';
DROP DATABASE IF EXISTS $OLD_DB_NAME;

-- Give people 1 chance to fix a terrible mistake
ALTER DATABASE $DB_NAME RENAME TO $OLD_DB_NAME;

CREATE ROLE $DB_USER SUPERUSER LOGIN PASSWORD 'test';
CREATE DATABASE $DB_NAME;
EOF

dbmate  --url "postgres://$DB_USER:test@localhost:5432/$DB_NAME?sslmode=disable" --migrations-dir src/migrations up