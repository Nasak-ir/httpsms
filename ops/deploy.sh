#!/usr/bin/env bash
set -Eeuo pipefail

APP_ROOT=${APP_ROOT:-/opt/nasak-sms}
REPO_DIR=${REPO_DIR:-$APP_ROOT/repository}
SHARED_DIR=${SHARED_DIR:-$APP_ROOT/shared}
COMPOSE_ENV=${COMPOSE_ENV:-$SHARED_DIR/compose.env}
COMPOSE_FILE=${COMPOSE_FILE:-$REPO_DIR/compose.production.yml}

required_files=("$COMPOSE_ENV" "$SHARED_DIR/api.env")
for file in "${required_files[@]}"; do
  if [[ ! -s "$file" ]]; then
    echo "Required runtime file is missing or empty: $file" >&2
    exit 1
  fi
done

chmod 600 "$COMPOSE_ENV" "$SHARED_DIR/api.env"
cd "$REPO_DIR"

export GIT_COMMIT
GIT_COMMIT=$(git rev-parse --short=12 HEAD)

compose=(docker compose --env-file "$COMPOSE_ENV" -f "$COMPOSE_FILE")
"${compose[@]}" config --quiet
"${compose[@]}" build --pull
"${compose[@]}" up -d --remove-orphans

for _ in $(seq 1 60); do
  if curl -fsS http://127.0.0.1:3800/health >/dev/null && curl -fsS http://127.0.0.1:3801/ >/dev/null; then
    break
  fi
  sleep 2
done

curl -fsS http://127.0.0.1:3800/health >/dev/null
curl -fsS http://127.0.0.1:3801/ >/dev/null

read_env_value() {
  local key=$1 file=$2
  sed -n "s/^${key}=//p" "$file" | tail -n 1 | tr -d '\r'
}

EVENTS_QUEUE_USER_ID=$(read_env_value EVENTS_QUEUE_USER_ID "$SHARED_DIR/api.env")
EVENTS_QUEUE_USER_API_KEY=$(read_env_value EVENTS_QUEUE_USER_API_KEY "$SHARED_DIR/api.env")
POSTGRES_USER=$(read_env_value POSTGRES_USER "$COMPOSE_ENV")
POSTGRES_DB=$(read_env_value POSTGRES_DB "$COMPOSE_ENV")

POSTGRES_USER=${POSTGRES_USER:-httpsms}
POSTGRES_DB=${POSTGRES_DB:-httpsms}

if [[ -z ${EVENTS_QUEUE_USER_ID:-} || -z ${EVENTS_QUEUE_USER_API_KEY:-} ]]; then
  echo "System queue user is not configured." >&2
  exit 1
fi

"${compose[@]}" exec -T postgres psql \
  --username "$POSTGRES_USER" \
  --dbname "$POSTGRES_DB" \
  --set=system_user_id="$EVENTS_QUEUE_USER_ID" \
  --set=system_api_key="$EVENTS_QUEUE_USER_API_KEY" <<'SQL'
INSERT INTO users (id, api_key, email, timezone, subscription_name, created_at, updated_at)
VALUES (:'system_user_id', :'system_api_key', 'system@sms.nasak.ir', 'Asia/Tehran', 'free', NOW(), NOW())
ON CONFLICT (id) DO UPDATE SET api_key = EXCLUDED.api_key, updated_at = NOW();
SQL

"${compose[@]}" restart api
"$REPO_DIR/ops/health-check.sh"

echo "Nasak SMS deployed from commit $GIT_COMMIT"
