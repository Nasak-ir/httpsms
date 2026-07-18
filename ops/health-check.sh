#!/usr/bin/env bash
set -Eeuo pipefail

APP_ROOT=${APP_ROOT:-/opt/nasak-sms}
REPO_DIR=${REPO_DIR:-$APP_ROOT/repository}
COMPOSE_ENV=${COMPOSE_ENV:-$APP_ROOT/shared/compose.env}

curl --fail --silent --show-error --max-time 10 http://127.0.0.1:3800/health >/dev/null
curl --fail --silent --show-error --max-time 10 http://127.0.0.1:3801/ >/dev/null

docker compose --env-file "$COMPOSE_ENV" -f "$REPO_DIR/compose.production.yml" ps --status running --format json >/dev/null
echo "Nasak SMS API and panel are healthy."
