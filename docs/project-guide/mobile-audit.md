# ممیزی موبایل Android و شواهد اجرایی

این یادداشت روی commit `3b44528b16b06e8edfb44da824f6cb7997a571de` (شاخهٔ `codex/documentation-audit`) تهیه شده است. پروژهٔ `httpsms` یک درگاه SMS self-hosted با کلاینت Android در `android/` است. این گزارش فقط ممیزی مستندات و کنترل‌های کم‌خطر است؛ ارسال پیام واقعی، اتصال به Firebase واقعی و انتشار APK انجام نشد.

## شواهد قابل بازتولید

| کنترل | دستور/مسیر | نتیجه |
| --- | --- | --- |
| وضعیت مرجع | `git rev-parse HEAD` | `3b44528b16b06e8edfb44da824f6cb7997a571de` |
| راهنمای گراف | `graphify query "Android application build, security-sensitive configuration, and test entry points"` | `NOT_RUN`: executable `graphify` در PATH این محیط موجود نیست |
| Java | `java -version` | `PASS`: Temurin/Oracle Java `27.0.1` مشاهده شد |
| Gradle/Android build | `android/gradlew.bat --version` | `NOT_RUN`: wrapper تلاش کرد `gradle-9.4.1-bin.zip` را از `services.gradle.org` دریافت کند و در مهلت محیط کامل نشد؛ assemble/test موفق ادعا نمی‌شود |
| API tests | `go version; go test ./...` در `api/` | `NOT_RUN`: `go` در PATH موجود نیست |
| Web tests | `npm --version` / اسکریپت `web/package.json` | `NOT_RUN`: `npm` در PATH موجود نیست؛ اسکریپت `test` نیز فعلاً فقط پیام `No tests configured yet` چاپ می‌کند |
| secrets scan سطح مسیر | `rg` روی `.gitignore` و workflow | `PASS`: `firebase-credentials.json`، `google-services.json`، keystoreها و فایل‌های محیطی از commit مستثنا شده‌اند |

## کنترل‌های امنیتی مشاهده‌شده در کد

- [AndroidManifest.xml](../../android/app/src/main/AndroidManifest.xml) پشتیبان‌گیری خودکار را با `android:allowBackup="false"` خاموش می‌کند و `FileProvider` را non-exported نگه می‌دارد.
- مجوزهای حساس `READ_SMS`، `SEND_SMS`، `RECEIVE_SMS`، `READ_CALL_LOG` و `REQUEST_IGNORE_BATTERY_OPTIMIZATIONS` در manifest هستند؛ این‌ها برای درگاه SMS قابل انتظارند اما باید در release و سیاست حریم خصوصی به‌صورت صریح توجیه و در دستگاه واقعی آزموده شوند.
- receiverهای SMS، تماس و boot با `exported="true"` اعلام شده‌اند و برای SMS/تماس permission حفاظتی دارند. تست نصب/دریافت broadcast در این محیط انجام نشد؛ بررسی owner و حداقل‌سازی سطح export قبل از release لازم است.
- `android/app/build.gradle.kts` از SDK هدف/compile برابر 36 و حداقل SDK برابر 28 استفاده می‌کند. release با `isMinifyEnabled = false` ساخته می‌شود؛ این مورد ریسک افشای کد و حجم را بالا می‌برد و باید تصمیم release آن ثبت شود.
- امضای release فقط وقتی فعال می‌شود که چهار متغیر محیطی keystore حاضر باشند. workflow نیز secretهای signing را در pull request مطالبه نمی‌کند و APK را با `apksigner verify --verbose` بررسی می‌کند.
- workflow اندروید پیش از build، endpointهای upstream و telemetry ناخواسته را grep می‌کند و برای Firebase، project id و package name را اعتبارسنجی می‌کند. این کنترل متنی است و جایگزین تست runtime/TLS نیست.
- کد Android در `MainActivity.kt` و `FirebaseMessagingService.kt` در صورت فعال بودن تنظیم debug، `Timber.DebugTree` نصب می‌کند. چون پیامک، شماره و توکن‌ها داده حساس‌اند، release باید تنظیم debug logging را غیرفعال یا redaction آن را اثبات کند.

## کارهای باقیمانده و معیار پذیرش

1. در runner دارای Android SDK، Java پشتیبانی‌شده و شبکهٔ artifact، `./gradlew --no-daemon :app:testDebugUnitTest :app:assembleRelease` اجرا و log/exit code ذخیره شود؛ سپس `apksigner verify --verbose` و `apkanalyzer` روی APK انجام شود.
2. روی دستگاه آزمایشی بدون سیم‌کارت تولیدی، اعطای حداقل مجوزها، دریافت SMS/MMS، reboot receiver، foreground service و notification permission آزموده شود؛ متن پیام و شمارهٔ واقعی وارد artifact نشود.
3. بررسی استاتیک وابستگی‌ها (`gradle dependencies`/نسخه‌های Firebase و OkHttp) و اسکن secret در CI اجرا شود. در صورت انتشار، `isMinifyEnabled`/R8 و سیاست logging release تصمیم‌گیری و مستند شود.
4. endpoint پیش‌فرض `https://sms.nasak.ir/api` و گواهی TLS در محیط staging با تست certificate/hostname تأیید شود؛ هیچ تستی به production SMS ارسال نکند.

نتیجهٔ این سند «شاهد کد/CI» است و معادل تأیید runtime، امنیت عملیاتی یا انتشار APK نیست.

