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

Cloudflare Turnstile is required only for the public message-search feature.
The rest of the panel can run while its keys are empty.

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

Build the Android application with the same Firebase project's
`google-services.json`. On first launch, set the server URL to:

```text
https://sms.nasak.ir/api
```

Sign in with the same Firebase account as the web panel, grant SMS permissions,
register the SIM phone number, and keep battery optimization disabled for the
gateway application.

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
