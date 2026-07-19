# پیامک نسک

پیامک نسک یک درگاه self-hosted برای ارسال و دریافت SMS با سیم‌کارت و گوشی
اندرویدی خود مجموعه است. API، صف ارسال، گزارش تحویل، پیام‌های دریافتی، webhook
و پنل عملیاتی روی زیرساخت نسک اجرا می‌شوند.

## هزینه و محدودیت

- سامانه credit، اشتراک ماهانه یا سقف نرم‌افزاری ندارد.
- آمار ارسال و دریافت فقط برای پایش عملیات ذخیره می‌شود و مانع ارسال نیست.
- Firebase Cloud Messaging فقط گوشی را برای دریافت کار از صف بیدار می‌کند.
- تعرفه احتمالی اپراتور سیم‌کارت همچنان اعمال می‌شود و توسط این نرم‌افزار قابل
  حذف نیست.

## معماری

```text
Nasak web apps
      |
      | X-API-Key / HTTPS
      v
sms.nasak.ir/api
      |
      +-- PostgreSQL: users, phones, messages, delivery events, usage
      +-- Redis: cache and lightweight queue coordination
      +-- FCM: wake-up notification
      |
      v
Nasak SMS Android app
      |
      +-- Android SmsManager -> SIM operator
      +-- delivery event -> Nasak API
      +-- incoming SMS -> Nasak API -> webhook
```

هیچ پیامکی از یک سرویس فروش credit عبور نمی‌کند. شماره، کلید API و متن پیام فقط
بین سرویس نسک و گوشی متصل مبادله می‌شوند.

## راه‌اندازی گوشی

1. آخرین APK را از صفحه Release خصوصی پروژه نصب کنید.
2. مجوزهای SMS، اعلان و اجرای پس‌زمینه را فعال کنید.
3. در پنل `sms.nasak.ir` وارد شوید و از «کلید اتصال گوشی» یک کلید محدود بسازید.
4. در اپ اندروید، Server URL را `https://sms.nasak.ir/api` قرار دهید.
5. کلید گوشی و شماره سیم‌کارت را وارد کنید.
6. پس از اتصال، heartbeat گوشی باید در پنل سبز شود.

کلید گوشی با کلید اصلی حساب فرق دارد و فقط به شماره‌های انتخاب‌شده دسترسی دارد.

## ارسال پیام

```bash
curl --request POST 'https://sms.nasak.ir/api/v1/messages/send' \
  --header 'Content-Type: application/json' \
  --header 'X-API-Key: YOUR_ACCOUNT_API_KEY' \
  --data '{
    "from": "+989120000000",
    "to": "+989121111111",
    "content": "پیام آزمایشی نسک",
    "request_id": "your-idempotency-reference"
  }'
```

`from` باید یکی از شماره‌های متصل به همان حساب باشد. پاسخ موفق به معنی ورود پیام
به صف است؛ وضعیت نهایی با eventهای `message.phone.sent`,
`message.phone.delivered` یا `message.send.failed` مشخص می‌شود.

## اتصال به سایت اصلی نسک

```dotenv
SMS_PROVIDER="nasak-gateway"
NASAK_SMS_BASE_URL="https://sms.nasak.ir/api"
NASAK_SMS_API_KEY="..."
NASAK_SMS_FROM="+989..."
NASAK_SMS_SEND_PATH="/v1/messages/send"
```

این متغیرها فقط در env رمزگذاری‌شده production نگهداری می‌شوند. هیچ کلیدی در
repository، Docker image یا مرورگر قرار نمی‌گیرد.

## متغیرهای self-hosted

```dotenv
ENTITLEMENT_ENABLED=false
BILLING_LIMITS_ENABLED=false
PAID_FEATURES_ENABLED=false
DATABASE_DEBUG_LOGS=false
AXIOM_TOKEN=
```

- `ENTITLEMENT_ENABLED=false`: تعداد گوشی‌ها، کلیدها و زمان‌بندی‌ها محدود نیست.
- `BILLING_LIMITS_ENABLED=false`: تعداد پیام‌ها هرگز باعث توقف ارسال نمی‌شود.
- `PAID_FEATURES_ENABLED=false`: routeهای پرداخت و Lemon Squeezy ثبت نمی‌شوند.
- `DATABASE_DEBUG_LOGS=false`: SQL و پارامترهای حساس در لاگ production چاپ
  نمی‌شوند.
- Axiom خالی: telemetry خارجی غیرفعال و logger محلی استفاده می‌شود.

## امنیت

- API فقط روی HTTPS ارائه می‌شود.
- کلید حساب و کلید گوشی باید جدا باشند و دوره‌ای rotate شوند.
- کلیدها در header ارسال شوند؛ قراردادن کلید در query یا body توصیه نمی‌شود.
- متن پیام، کلیدها و payloadهای پیام در application log ثبت نمی‌شوند.
- PostgreSQL و Redis فقط داخل شبکه Docker در دسترس‌اند.
- کانتینر API read-only، بدون Linux capability و با `no-new-privileges` اجرا
  می‌شود.
- برای integration سایت از یک حساب سرویس و کلید اختصاصی استفاده شود.

## Backup و بازیابی

داده پایدار سرویس در volume دیتابیس PostgreSQL است. بکاپ production باید شامل
موارد زیر باشد:

- `pg_dump -Fc` از دیتابیس SMS
- envهای رمزگذاری‌شده `api.env` و `compose.env`
- inventory گوشی‌ها و checksum تنظیمات
- نسخه APK و commit فعال

Redis داده قابل بازسازی است و backup دائمی لازم ندارد. secret خام یا private
Firebase service account نباید وارد Git شود.

## سلامت سرویس

- پنل: `GET https://sms.nasak.ir/healthz`
- API: `GET https://sms.nasak.ir/api/health`
- heartbeat گوشی: از پنل و endpointهای heartbeat
- صف: پیام جدید باید از `pending` به `sent/delivered` یا `failed` برسد.

## توسعه

```bash
docker compose up --build
cd api && go test ./...
cd web && corepack pnpm install && corepack pnpm lint && corepack pnpm build
```

کد بر پایه پروژه متن‌باز `NdoleStudio/httpsms` توسعه یافته و مجوز AGPL آن در
فایل `LICENSE` حفظ شده است. branding، deployment، auth fallback و سیاست
unlimited نسخه نسک در همین repository نگهداری می‌شوند.
