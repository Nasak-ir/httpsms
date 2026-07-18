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
set -a
# shellcheck disable=SC1090
source "$COMPOSE_ENV"
set +a

docker compose --env-file "$COMPOSE_ENV" -f "$REPO_DIR/compose.production.yml" exec -T postgres \
  pg_dump --format=custom --compress=9 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" >"$TARGET"

chmod 600 "$TARGET"
sha256sum "$TARGET" >"$TARGET.sha256"
chmod 600 "$TARGET.sha256"

docker compose --env-file "$COMPOSE_ENV" -f "$REPO_DIR/compose.production.yml" exec -T postgres \
  pg_restore --list <"$TARGET" >/dev/null

find "$BACKUP_DIR" -maxdepth 1 -type f -name 'httpsms-*.dump*' -mtime +7 -delete
echo "Created and verified $TARGET"
