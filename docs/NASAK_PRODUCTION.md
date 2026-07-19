# Nasak SMS production deployment

Nasak SMS is a self-hosted fork of httpSMS. It turns an Android phone into an
SMS gateway and exposes a web panel, HTTP API, webhooks, delivery reports,
rate limiting, message expiration, and optional end-to-end encryption.

## Architecture

- Public panel: `https://sms.nasak.ir`
- Public API: `https://sms.nasak.ir/api/v1/...`
- API and panel bind only to `127.0.0.1` on the server.
- PostgreSQL and Redis have no host ports.
- Nginx is the only public entry point.
- Runtime secrets live in `/opt/nasak-sms/shared` with mode `0600`.
- Database backups are created hourly and verified with `pg_restore --list`.

## Required external configuration

httpSMS uses Firebase Authentication for panel users and Firebase Cloud
Messaging to wake the Android gateway phone. Create one Firebase project and
provide all three matching artifacts:

1. Firebase Web SDK values for `/opt/nasak-sms/shared/compose.env`.
2. A compact service-account JSON value for `FIREBASE_CREDENTIALS` in
   `/opt/nasak-sms/shared/api.env`.
3. `google-services.json` for the Android build.

Enable Email/Password sign-in in Firebase Authentication. Add
`sms.nasak.ir` to Firebase Authentication authorized domains.

The web panel first uses the Firebase Web SDK directly. If a user's browser
cannot reach Firebase, email/password authentication falls back to
`/api/v1/auth/email`, which exchanges credentials for Firebase tokens from the
API server. On networks where the server cannot reach Google/Firebase directly,
set `FIREBASE_AUTH_PROXY_URL` in `/opt/nasak-sms/shared/api.env` to an approved
HTTP or SOCKS5 proxy. If the Firebase API key has HTTP referrer restrictions,
set `FIREBASE_AUTH_REFERER_URL=https://sms.nasak.ir/` so the server-side
fallback sends the same origin metadata as the browser. Do not store proxy
credentials in the repository.

Set `LOCAL_AUTH_SECRET` to a long random value. When Firebase Identity Toolkit is
unreachable from the server, the same endpoint falls back to Nasak-hosted
email/password credentials hashed with bcrypt and returns a short-lived local
bearer token accepted by the API middleware.

Install the downloaded service-account JSON without copying it into the
repository or printing its private key:

```bash
sudo /opt/nasak-sms/repository/ops/install-firebase-credentials.sh \
  /tmp/firebase-service-account.json
```

The installer verifies that the service account belongs to the configured
Firebase project, atomically updates `api.env`, and removes the source file.

Cloudflare Turnstile is optional. When both keys are empty, authenticated
message search stays available without any external CAPTCHA request. The API
still enforces account ownership, phone limits, validation, and authentication.

## Server layout

```text
/opt/nasak-sms/
  repository/       Git checkout
  shared/
    compose.env     Compose and public Firebase config (0600)
    api.env         API secrets and service-account JSON (0600)
  backups/          Verified local PostgreSQL dumps (0700)
```

## First deployment

```bash
sudo /opt/nasak-sms/repository/ops/install-docker-ubuntu.sh
sudo install -d -m 0750 /opt/nasak-sms/{shared,backups}
sudo cp deploy/compose.env.example /opt/nasak-sms/shared/compose.env
sudo cp deploy/api.env.example /opt/nasak-sms/shared/api.env
sudo chmod 600 /opt/nasak-sms/shared/*.env
# Fill the runtime values, then:
sudo /opt/nasak-sms/repository/ops/deploy.sh
```

Install the Nginx file and request the certificate:

```bash
sudo cp ops/nginx-sms.nasak.ir.conf /etc/nginx/sites-available/sms.nasak.ir
sudo ln -sfn /etc/nginx/sites-available/sms.nasak.ir /etc/nginx/sites-enabled/sms.nasak.ir
sudo nginx -t && sudo systemctl reload nginx
sudo certbot --nginx -d sms.nasak.ir --redirect --non-interactive --agree-tos
```

## Android gateway

The Android app is built with the same Firebase project's `google-services.json`.
The file is stored only as the encrypted GitHub Actions secret
`ANDROID_GOOGLE_SERVICES_JSON_BASE64` and is created inside the runner for the
duration of the build. Release signing uses these additional secrets:

- `ANDROID_RELEASE_KEYSTORE_BASE64`
- `ANDROID_RELEASE_KEYSTORE_PASSWORD`
- `ANDROID_RELEASE_KEY_ALIAS`
- `ANDROID_RELEASE_KEY_PASSWORD`

Pushes to `main` produce a signed workflow artifact. Tags matching `sms-v*`
also publish `NasakSms.apk` in a GitHub Release. The signing keystore must be
kept in an offline recovery kit because every future update must use the same
key.

The Nasak build defaults to this server URL:

```text
https://sms.nasak.ir/api
```

Sign in with the same Firebase account as the web panel, grant SMS permissions,
register the SIM phone number, and keep battery optimization disabled for the
gateway application.

The production fork does not include the upstream Axiom or Sentry telemetry.
Application diagnostics stay on the device unless a Nasak-owned monitoring
provider is deliberately configured in a future release.

## Operations

```bash
sudo systemctl status nasak-sms
sudo /opt/nasak-sms/repository/ops/health-check.sh
sudo /opt/nasak-sms/repository/ops/backup.sh
sudo docker compose --env-file /opt/nasak-sms/shared/compose.env \
  -f /opt/nasak-sms/repository/compose.production.yml logs --tail=200 api
```

Never commit production `.env`, Firebase service-account JSON,
`google-services.json`, database dumps, or API keys.

For servers where `proxy.golang.org` is unavailable, set `GOPROXY` in
`compose.env` to an approved regional mirror. This affects build-time module
downloads only and is not used by the running API.
