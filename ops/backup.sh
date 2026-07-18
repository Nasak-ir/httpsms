#!/usr/bin/env bash
set -Eeuo pipefail

APP_ROOT=${APP_ROOT:-/opt/nasak-sms}
REPO_DIR=${REPO_DIR:-$APP_ROOT/repository}
SHARED_DIR=${SHARED_DIR:-$APP_ROOT/shared}
BACKUP_DIR=${BACKUP_DIR:-$APP_ROOT/backups}
COMPOSE_ENV=${COMPOSE_ENV:-$SHARED_DIR/compose.env}
STAMP=$(date -u +%Y%m%dT%H%M%SZ)
TARGET="$BACKUP_DIR/httpsms-$STAMP.dump"

install -d -m 0700 "$BACKUP_DIR"
read_env_value() {
  local key=$1 file=$2
  sed -n "s/^${key}=//p" "$file" | tail -n 1 | tr -d '\r'
}

POSTGRES_USER=$(read_env_value POSTGRES_USER "$COMPOSE_ENV")
POSTGRES_DB=$(read_env_value POSTGRES_DB "$COMPOSE_ENV")
POSTGRES_USER=${POSTGRES_USER:-httpsms}
POSTGRES_DB=${POSTGRES_DB:-httpsms}

docker compose --env-file "$COMPOSE_ENV" -f "$REPO_DIR/compose.production.yml" exec -T postgres \
  pg_dump --format=custom --compress=9 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" >"$TARGET"

chmod 600 "$TARGET"
sha256sum "$TARGET" >"$TARGET.sha256"
chmod 600 "$TARGET.sha256"

docker compose --env-file "$COMPOSE_ENV" -f "$REPO_DIR/compose.production.yml" exec -T postgres \
  pg_restore --list <"$TARGET" >/dev/null

find "$BACKUP_DIR" -maxdepth 1 -type f -name 'httpsms-*.dump*' -mtime +7 -delete
echo "Created and verified $TARGET"
