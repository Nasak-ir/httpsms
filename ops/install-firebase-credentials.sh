#!/usr/bin/env bash
set -Eeuo pipefail

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 /path/to/firebase-service-account.json" >&2
  exit 2
fi

APP_ROOT=${APP_ROOT:-/opt/nasak-sms}
SHARED_DIR=${SHARED_DIR:-$APP_ROOT/shared}
COMPOSE_ENV=${COMPOSE_ENV:-$SHARED_DIR/compose.env}
API_ENV=${API_ENV:-$SHARED_DIR/api.env}
CREDENTIALS_FILE=$1

for command in jq mktemp; do
  if ! command -v "$command" >/dev/null 2>&1; then
    echo "Required command is unavailable: $command" >&2
    exit 1
  fi
done

if [[ ! -s "$CREDENTIALS_FILE" ]]; then
  echo "Firebase credential file is missing or empty." >&2
  exit 1
fi

if [[ ! -s "$COMPOSE_ENV" || ! -s "$API_ENV" ]]; then
  echo "Nasak SMS runtime environment is incomplete." >&2
  exit 1
fi

expected_project_id=$(sed -n 's/^FIREBASE_PROJECT_ID=//p' "$COMPOSE_ENV" | tail -n 1 | tr -d '\r')
if [[ -z "$expected_project_id" ]]; then
  echo "FIREBASE_PROJECT_ID is not configured." >&2
  exit 1
fi

jq --exit-status --arg project "$expected_project_id" '
  .type == "service_account" and
  .project_id == $project and
  (.private_key | type == "string" and length > 0) and
  (.client_email | type == "string" and length > 0)
' "$CREDENTIALS_FILE" >/dev/null

tmp_file=$(mktemp "${API_ENV}.tmp.XXXXXX")
cleanup() {
  rm -f -- "$tmp_file"
}
trap cleanup EXIT

{
  grep -v '^FIREBASE_CREDENTIALS=' "$API_ENV"
  printf 'FIREBASE_CREDENTIALS='
  jq --compact-output . "$CREDENTIALS_FILE"
  printf '\n'
} >"$tmp_file"

chown root:ubuntu "$tmp_file"
chmod 0600 "$tmp_file"
mv -f -- "$tmp_file" "$API_ENV"
trap - EXIT

if command -v shred >/dev/null 2>&1; then
  shred --remove "$CREDENTIALS_FILE"
else
  rm -f -- "$CREDENTIALS_FILE"
fi

echo "Firebase service account installed for project $expected_project_id."
