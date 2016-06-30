#!/usr/bin/env sh

echo "Waiting for Postgres"
while true; do
    nc -q 1 postgres 5432 2>/dev/null && break
done

echo "Starting service-store"
/go/bin/service-store
