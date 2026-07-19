#!/usr/bin/env bash
set -Eeuo pipefail

APP_ROOT=${APP_ROOT:-/opt/nasak-sms}
REPO_DIR=${REPO_DIR:-$APP_ROOT/repository}
COMPOSE_ENV=${COMPOSE_ENV:-$APP_ROOT/shared/compose.env}

check_http() {
  local url=$1

  curl \
    --fail \
    --silent \
    --show-error \
    --max-time 10 \
    --retry 15 \
    --retry-all-errors \
    --retry-delay 2 \
    "$url" >/dev/null
}

check_http http://127.0.0.1:3800/health
check_http http://127.0.0.1:3801/

docker compose --env-file "$COMPOSE_ENV" -f "$REPO_DIR/compose.production.yml" ps --status running --format json >/dev/null
echo "Nasak SMS API and panel are healthy."
