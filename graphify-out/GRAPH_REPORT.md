# Graph Report - httpsms  (2026-07-19)

## Corpus Check
- 383 files · ~311,042 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 3361 nodes · 6800 edges · 245 communities (202 shown, 43 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 312 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `3cf9c339`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Settings
- Container
- helpers_test.go
- PhoneNotificationService
- HttpSmsApiService
- MessageThreadService
- api.ts
- index.vue
- WebhookService
- handler
- MessageService
- SendSmsWorker
- index.vue
- MainActivity
- Logger
- Tracer
- index.vue
- SentReceiver
- BillingUsage
- index.vue
- SettingsViewModel
- DiscordService
- Cache
- Integration3CXService
- BillingService
- Values
- User
- Context
- MessageHandlerValidator
- IndexParams
- Nuxt 4 + Vuetify 4 Migration Implementation Plan
- FirebaseAuth.vue
- index.vue
- UserService
- MessageThreadHandler
- PhoneService
- MessageSendScheduleHandlerValidator
- Client
- Migration Order (Tasks)
- Response
- LemonsqueezyService
- googlePushQueue
- filters.ts
- UserID
- index.vue
- SIM
- EventDispatcher
- dependencies
- HeartbeatService
- formatEventPayload
- Message
- PhoneAPIKeyHandler
- MessageSendScheduleService
- hermesNotificationEmailFactory
- MessageSendSchedule
- request
- zerologLogger
- .getCallLog
- container.go
- Heartbeat
- DiscordStore
- message_service.go
- [id].vue
- StickyNotificationService
- PhoneHeartbeatMissedPayload
- UserHandler
- MessageListener
- gormPhoneAPIKeyRepository
- WebhookStore
- response
- File Structure
- devDependencies
- .MessageService
- NewMongoDB
- user_service.go
- PhoneAPIKeyService
- Design
- Build, Test, and Lint Commands
- MessageThreadHeader.vue
- index.vue
- errors.ts
- NotificationEmailFactory
- gormPhoneRepository
- mongoHeartbeatMonitorRepository
- otelTracer
- Affiliates Landing Page — Design
- scripts
- New
- HermesGeneratorConfig
- gormHeartbeatMonitorRepository
- .ValidateStore
- Contributor Covenant Code of Conduct
- Login "Last Used" Badge — Design
- index.vue
- Email
- HeartbeatIndex
- UserHandlerValidator
- EmailNotificationService
- noopLogger
- Webhook Email Payload Formatting
- httpSMS
- auth.ts
- compilerOptions
- UserEmailFactory
- attachment_repository.go
- UserListener
- gormHeartbeatRepository
- MessageThreadIndex
- PhoneAPIKeyIndex
- response.go
- gormLogger
- resetErrors
- messages.ts
- default.vue
- Phone
- PhoneAPIKey
- MessagePhoneSendingPayload
- HeartbeatListener
- Getting Started page — Design
- playwright
- Integration Tests
- Self Host Setup - Docker
- billing_usage_test.go
- bulk_message_handler.go
- testClient
- EmailNotificationListener
- MemoryAttachmentRepository
- Nasak SMS production deployment
- Running Locally
- BootReceiver
- PhoneNumberValidator.kt
- NewSMTPEmailService
- message_send_schedule_test.go
- user_test.go
- MessageAPISentPayload
- BillingListener
- PhoneAPIKeyListener
- BillingUsageHistory
- BulkMessage
- MessageIndex
- File Structure
- Message Thread Archive UI Design
- website.vue
- parseErrors
- package.json
- main
- memoryCache
- redisCache
- HeartbeatMonitor
- Context
- Global Constraints
- CopyButton.vue
- LoadingButton.vue
- countries.ts
- HeartbeatWorker
- gradlew
- factory
- MessageCallMissedPayload
- MessageNotificationSentPayload
- MessagePhoneDeliveredPayload
- MessagePhoneReceivedPayload
- MessagePhoneSentPayload
- MessageSendRetryPayload
- MessageCallMissed
- emulator_fcm_client.go
- .Check
- service
- validateAttachmentURL
- Features
- BackButton.vue
- end-to-end-encryption-to-sms-messages.vue
- index.vue
- index.vue
- Nuxt Minimal Starter
- MessageNotificationFailedPayload
- MessageSendScheduleDeletedPayload
- PhoneHeartbeatOnlinePayload
- .WebhookSendFailed
- .handleRequest
- .OnMessagePhoneReceived
- .onUserAccountCreated
- .onMessageSendScheduleDeleted
- MessageResponse
- PhoneResponse
- UserResponse
- Security Policy
- Troubleshooting
- Architecture
- AppToast.vue
- saveSchedule
- scheduleDayEnabled
- ExampleInstrumentedTest
- ExampleUnitTest
- BulkMessage
- .onUserAccountDeleted
- .Send
- .Send
- TestShouldCheckUnarchive
- .ValidateEvent
- backup.sh
- deploy.sh
- install-docker-ubuntu.sh
- login.vue
- app.ts
- Constants.kt
- health-check.sh
- generate-firebase-credentials.sh
- BlogInfo.vue
- BlogSidebar.vue
- FixedHeader.vue
- LoadingDashboard.vue
- forward-incoming-sms-from-phone-to-webhook.vue
- grant-send-and-read-sms-permissions-on-android.vue
- how-to-send-sms-messages-from-excel.vue
- send-bulk-sms-from-csv-file-with-no-code.vue
- send-sms-from-android-phone-with-python.vue
- send-sms-when-new-row-is-added-to-google-sheets-using-zapier.vue
- qrcode
- getWeekday
- saveEmailNotifications
- github.com/NdoleStudio/httpsms
- github.com/NdoleStudio/httpsms/tests

## God Nodes (most connected - your core abstractions)
1. `Tracer` - 260 edges
2. `UserID` - 230 edges
3. `Logger` - 197 edges
4. `Container` - 145 edges
5. `New()` - 76 edges
6. `MessageService` - 66 edges
7. `Client` - 54 edges
8. `request` - 42 edges
9. `UserService` - 41 edges
10. `User` - 40 edges

## Surprising Connections (you probably didn't know these)
- `newAPIClient()` --calls--> `New()`  [INFERRED]
  tests/helpers_test.go → api/pkg/discord/client.go
- `newPhoneClient()` --calls--> `New()`  [INFERRED]
  tests/helpers_test.go → api/pkg/discord/client.go
- `randomEncryptionKey()` --calls--> `New()`  [INFERRED]
  tests/helpers_test.go → api/pkg/discord/client.go
- `setupPhone()` --calls--> `New()`  [INFERRED]
  tests/helpers_test.go → api/pkg/discord/client.go
- `setupWebhook()` --calls--> `New()`  [INFERRED]
  tests/helpers_test.go → api/pkg/discord/client.go

## Import Cycles
- None detected.

## Communities (245 total, 43 thin omitted)

### Community 0 - "Settings"
Cohesion: 0.07
Nodes (17): AppCompatActivity, Bundle, Context, String, LoginActivity, Boolean, Context, Long (+9 more)

### Community 1 - "Container"
Cohesion: 0.09
Nodes (8): App, Database, Float64Histogram, Handler, NewContainer(), Container, Resource, RoundTripper

### Community 2 - "helpers_test.go"
Cohesion: 0.10
Nodes (60): defaultClientConfig(), T, TestWithBaseURL(), TestWithHTTPClient(), WithApplicationID(), WithBaseURL(), WithBotToken(), WithHTTPClient() (+52 more)

### Community 3 - "PhoneNotificationService"
Cohesion: 0.06
Nodes (30): Time, UUID, SIM, Time, UUID, Time, UUID, Context (+22 more)

### Community 4 - "HttpSmsApiService"
Cohesion: 0.06
Nodes (37): Encrypter, String, create(), HttpSmsApiService, Array, Boolean, Context, Int (+29 more)

### Community 5 - "MessageThreadService"
Cohesion: 0.06
Nodes (23): Time, UUID, SIM, Time, UUID, Context, Event, Context (+15 more)

### Community 6 - "api.ts"
Cohesion: 0.04
Nodes (57): useBillingStore, EntitiesBillingUsage, EntitiesDiscord, EntitiesMessageSendSchedule, EntitiesMessageSendScheduleWindow, EntitiesPhoneAPIKey, EntitiesSIM, EntitiesSubscriptionName (+49 more)

### Community 7 - "index.vue"
Cohesion: 0.03
Nodes (48): activeDiscord, activePhone, activeSchedule, activeWebhook, apiKey, apiKeyShow, authStore, avatarUrl (+40 more)

### Community 8 - "WebhookService"
Cohesion: 0.08
Nodes (20): StringArray, Time, UUID, Context, Event, Context, DB, UUID (+12 more)

### Community 9 - "handler"
Cohesion: 0.06
Nodes (11): App, Ctx, Router, Values, Ctx, Router, Ctx, Router (+3 more)

### Community 10 - "MessageService"
Cohesion: 0.11
Nodes (7): BulkMessage, Context, Event, Message, service, HandleMessageParams, MessageService

### Community 11 - "SendSmsWorker"
Cohesion: 0.11
Nodes (22): android, ArrayList, Context, Message, PendingIntent, Result, String, Worker (+14 more)

### Community 12 - "index.vue"
Cohesion: 0.05
Nodes (41): authStore, canResendSelected, config, deleteMessages(), errorMessages, errorTitle, exportMessages(), fetchMessages() (+33 more)

### Community 13 - "MainActivity"
Cohesion: 0.08
Nodes (22): Int, String, LogEntry, LogzTree, AppCompatActivity, Boolean, Bundle, Context (+14 more)

### Community 14 - "Logger"
Cohesion: 0.09
Nodes (29): NewBillingListener(), NewDiscordListener(), NewEmailNotificationListener(), NewHeartbeatListener(), NewIntegration3CXListener(), NewMarketingListener(), NewMessageListener(), NewMessageSendScheduleListener() (+21 more)

### Community 15 - "Tracer"
Cohesion: 0.09
Nodes (20): Ctx, Router, NewHeartbeatHandler(), App, NewLemonsqueezyHandler(), Authenticated(), Handler, NewGormHeartbeatMonitorRepository() (+12 more)

### Community 16 - "index.vue"
Cohesion: 0.06
Nodes (30): authStore, canResend(), config, contact, contactIsPhoneNumber, deleteMessage(), form, { formatPhoneNumber } (+22 more)

### Community 17 - "SentReceiver"
Cohesion: 0.08
Nodes (22): DeliveredMessageWorker, DeliveredReceiver, FailedMessageWorker, BroadcastReceiver, Context, Intent, Result, String (+14 more)

### Community 18 - "BillingUsage"
Cohesion: 0.10
Nodes (16): Time, UUID, T, TestComputeBillingCycle(), TestDaysInMonth(), computeBillingCycle(), daysInMonth(), Context (+8 more)

### Community 19 - "index.vue"
Cohesion: 0.05
Nodes (33): authStore, billingStore, checkoutURL, config, dialog, enterpriseCheckoutURL, errorMessages, { formatDecimal, formatTimestamp } (+25 more)

### Community 20 - "SettingsViewModel"
Cohesion: 0.10
Nodes (16): AppCompatActivity, Boolean, Bundle, SettingsActivity, Boolean, String, SettingsScreen(), SwitchSetting() (+8 more)

### Community 21 - "DiscordService"
Cohesion: 0.11
Nodes (15): Time, UUID, Context, DB, UUID, Context, Event, Map (+7 more)

### Community 22 - "Cache"
Cohesion: 0.09
Nodes (17): NewMemoryCache(), NewRedisCache(), Phone, Handler, PhoneAPIKeyAuth(), NewGormBillingUsageRepository(), NewGormMessageThreadRepository(), NewGormPhoneAPIKeyRepository() (+9 more)

### Community 23 - "Integration3CXService"
Cohesion: 0.11
Nodes (14): Time, UUID, Context, Event, Context, DB, Context, Event (+6 more)

### Community 24 - "BillingService"
Cohesion: 0.09
Nodes (20): App, Ctx, NewIntegration3CxHandler(), APIKeyAuth(), getAPIKeyFromRequest(), Ctx, Handler, BearerAPIKeyAuth() (+12 more)

### Community 25 - "Values"
Cohesion: 0.08
Nodes (13): Time, SIM, Time, Time, Context, Values, MessageAttachment, MessageBulkSend (+5 more)

### Community 26 - "User"
Cohesion: 0.11
Nodes (10): UUID, Location, Time, UUID, Context, DB, Time, AuthContext (+2 more)

### Community 27 - "Context"
Cohesion: 0.11
Nodes (17): SIM, Time, UUID, SIM, Time, UUID, Context, Duration (+9 more)

### Community 28 - "MessageHandlerValidator"
Cohesion: 0.11
Nodes (14): Context, Ctx, NewDiscordHandler(), NewMessageHandler(), NewDiscordHandlerValidator(), NewMessageHandlerValidator(), Context, Time (+6 more)

### Community 29 - "IndexParams"
Cohesion: 0.13
Nodes (13): Time, UUID, BulkMessage, Context, DB, Message, UUID, MessageStatus (+5 more)

### Community 30 - "Nuxt 4 + Vuetify 4 Migration Implementation Plan"
Cohesion: 0.07
Nodes (28): File Structure, Nuxt 4 + Vuetify 4 Migration Implementation Plan, Task 10: Create Pinia Store — Billing, Task 11: Create Middleware, Task 12: Port Layouts, Task 13: Port Components — Toast, LoadingDashboard, LoadingButton, BackButton, Task 14: Port Components — CopyButton, FixedHeader, BlogAuthorBio, BlogInfo, NuxtLogo, Task 15: Port Components — FirebaseAuth, MessageThread, MessageThreadHeader (+20 more)

### Community 31 - "FirebaseAuth.vue"
Cohesion: 0.11
Nodes (28): appStore, authStore, backToSignIn(), clearErrors(), email, errorMessages, generalError, getGeneralErrorMessage() (+20 more)

### Community 32 - "index.vue"
Cohesion: 0.08
Nodes (26): activePhoneApiKey, activePhoneNumber, appStore, authStore, config, createPhoneApiKey(), deleteApiKey(), deleteApiKeyDialog (+18 more)

### Community 33 - "UserService"
Cohesion: 0.14
Nodes (9): Time, Context, service, ApiResponseData, APIResponseRelationshipsSubscriptionInvoice, UserAPIKeyRotatedPayload, Reader, UserService (+1 more)

### Community 34 - "MessageThreadHandler"
Cohesion: 0.11
Nodes (13): Ctx, Router, NewBillingHandler(), Ctx, Router, NewMessageThreadHandler(), NewBillingHandlerValidator(), NewMessageThreadHandlerValidator() (+5 more)

### Community 35 - "PhoneService"
Cohesion: 0.12
Nodes (14): Router, NewBulkMessageHandler(), NewPhoneHandler(), Ctx, Router, NewWebhookHandler(), service, NewBulkMessageHandlerValidator() (+6 more)

### Community 36 - "MessageSendScheduleHandlerValidator"
Cohesion: 0.13
Nodes (10): Ctx, Router, Context, MessageSendScheduleWindow, Values, NewMessageSendScheduleHandlerValidator(), MessageSendScheduleHandler, MessageSendScheduleStore (+2 more)

### Community 37 - "Client"
Cohesion: 0.10
Nodes (16): BearerAuth(), Handler, Context, NewGoogleCloudStorageAttachmentRepository(), NewGormDiscordRepository(), NewGormWebhookRepository(), NewDiscordService(), NewEmulatorFCMClient() (+8 more)

### Community 38 - "Migration Order (Tasks)"
Cohesion: 0.07
Nodes (27): Architecture, Components: Class-based → `<script setup>`, Decisions, Directory Structure (Nuxt 4), Dynamic Routes: `_id` → `[id]`, Firebase: `this.$fire.auth` → VueFire composables, httpSMS Frontend Migration: Nuxt 2 + Vuetify 2 → Nuxt 4 + Vuetify 4, Key Migration Patterns (+19 more)

### Community 39 - "Response"
Cohesion: 0.09
Nodes (13): Context, Context, Context, Request, Context, ApplicationService, ChannelService, CommandCreateRequest (+5 more)

### Community 40 - "LemonsqueezyService"
Cohesion: 0.12
Nodes (13): Time, Time, Time, Time, Context, service, SubscriptionName, UserSubscriptionCancelledPayload (+5 more)

### Community 41 - "googlePushQueue"
Cohesion: 0.14
Nodes (12): Context, Duration, Context, Duration, Event, Span, Context, Duration (+4 more)

### Community 42 - "filters.ts"
Cohesion: 0.14
Nodes (19): parts, props, appStore, notificationsStore, phonesStore, threadsStore, useFilters(), capitalize() (+11 more)

### Community 43 - "UserID"
Cohesion: 0.13
Nodes (11): Time, Time, Context, MessageSendScheduleWindow, UUID, sanitizeWindows(), UserID, UserAccountCreatedPayload (+3 more)

### Community 44 - "index.vue"
Cohesion: 0.10
Nodes (21): authStore, bulkOrders, errorMessages, errorTitle, fetchBulkOrders(), { formatTimestamp }, formFile, loading (+13 more)

### Community 45 - "SIM"
Cohesion: 0.10
Nodes (8): UUID, Context, Values, SIM, PhoneDelete, PhoneFCMToken, PhoneIndex, PhoneUpsert

### Community 46 - "EventDispatcher"
Cohesion: 0.17
Nodes (11): Ctx, Router, NewEventsHandler(), EmulatorPushQueue(), Float64Histogram, NewEventDispatcher(), NewGooglePushQueue(), EventsHandler (+3 more)

### Community 47 - "dependencies"
Cohesion: 0.09
Nodes (22): dependencies, chart.js, chartjs-adapter-moment, date-fns, firebase, flag-icons, highlight.js, libphonenumber-js (+14 more)

### Community 48 - "HeartbeatService"
Cohesion: 0.25
Nodes (8): Context, service, Time, UUID, HeartbeatMonitorParams, HeartbeatMonitorStoreParams, HeartbeatService, HeartbeatStoreParams

### Community 49 - "formatEventPayload"
Cohesion: 0.17
Nodes (17): eventPayloadJSONStringEnd(), eventPayloadNextNonSpace(), formatEventPayload(), Builder, highlightEventPayloadJSON(), indentEventPayloadJSON(), isEventPayloadDigit(), isEventPayloadNumberCharacter() (+9 more)

### Community 50 - "Message"
Cohesion: 0.19
Nodes (4): StringArray, Time, UUID, Message

### Community 51 - "PhoneAPIKeyHandler"
Cohesion: 0.18
Nodes (9): NewMessageSendScheduleHandler(), App, Ctx, NewPhoneAPIKeyHandler(), service, NewPhoneAPIKeyHandlerValidator(), PhoneAPIKeyHandler, EntitlementService (+1 more)

### Community 52 - "MessageSendScheduleService"
Cohesion: 0.18
Nodes (11): NewGormMessageSendScheduleRepository(), NewGormPhoneNotificationRepository(), service, NewMessageSendScheduleService(), NewNotificationService(), NewPhoneHandlerValidator(), MessageSendScheduleRepository, PhoneNotificationRepository (+3 more)

### Community 53 - "hermesNotificationEmailFactory"
Cohesion: 0.11
Nodes (12): Hermes, UUID, SIM, Time, UUID, SIM, Time, UUID (+4 more)

### Community 54 - "MessageSendSchedule"
Cohesion: 0.20
Nodes (8): Time, UUID, Context, DB, UUID, MessageSendSchedule, MessageSendScheduleWindow, gormMessageSendScheduleRepository

### Community 56 - "zerologLogger"
Cohesion: 0.21
Nodes (5): Event, SpanContext, NewZerologLogger(), Level, zerologLogger

### Community 57 - ".getCallLog"
Cohesion: 0.19
Nodes (11): BroadcastReceiver, Context, Int, Intent, Pair, Result, String, Worker (+3 more)

### Community 58 - "container.go"
Cohesion: 0.18
Nodes (12): axiomLogger(), consoleLogger(), getGCEInstanceID(), DB, Interface, instanceID(), isLocal(), logDriver() (+4 more)

### Community 59 - "Heartbeat"
Cohesion: 0.17
Nodes (10): Time, UUID, Collection, Context, Database, NewMongoHeartbeatRepository(), Heartbeat, mongoHeartbeatRepository (+2 more)

### Community 60 - "DiscordStore"
Cohesion: 0.14
Nodes (5): Context, Values, DiscordIndex, DiscordStore, DiscordUpdate

### Community 61 - "message_service.go"
Cohesion: 0.21
Nodes (13): PhoneNumber, SIM, Time, UUID, MessageEventName, HandleMessageFailedParams, MessageCheckExpired, MessageGetOutstandingParams (+5 more)

### Community 62 - "[id].vue"
Cohesion: 0.12
Nodes (13): authStore, chartData, chartOptions, dataTableHeaders, dataTableItems, { formatPhoneNumber, formatTimestamp }, heartbeats, HeartbeatTableItem (+5 more)

### Community 63 - "StickyNotificationService"
Cohesion: 0.15
Nodes (9): Int, Intent, Service, StickyNotificationService, IBinder, Notification, NotificationRequest, NotificationType (+1 more)

### Community 64 - "PhoneHeartbeatMissedPayload"
Cohesion: 0.12
Nodes (10): Time, UUID, Time, UUID, Time, UUID, Event, PhoneHeartbeatCheckPayload (+2 more)

### Community 65 - "UserHandler"
Cohesion: 0.21
Nodes (4): Ctx, Router, NewUserHandler(), UserHandler

### Community 66 - "MessageListener"
Cohesion: 0.31
Nodes (3): Context, Event, MessageListener

### Community 67 - "gormPhoneAPIKeyRepository"
Cohesion: 0.23
Nodes (5): Context, DB, Phone, UUID, gormPhoneAPIKeyRepository

### Community 68 - "WebhookStore"
Cohesion: 0.15
Nodes (5): Context, Values, WebhookIndex, WebhookStore, WebhookUpdate

### Community 69 - "response"
Cohesion: 0.16
Nodes (11): BulkMessage, BillingUsageResponse, BillingUsagesResponse, BulkMessagesResponse, DiscordResponse, DiscordsResponse, MessageSendScheduleResponse, MessageSendSchedulesResponse (+3 more)

### Community 70 - "File Structure"
Cohesion: 0.12
Nodes (15): File Structure, Global Constraints, Self-Review, Task 10: Full validation, Task 11: End-to-end integration test, Task 1: Add `UnarchiveThread` field to the Phone entity, Task 2: Add `UnarchiveThread` to the received-message event payload, Task 3: Populate `UnarchiveThread` on the payload in ReceiveMessage (+7 more)

### Community 71 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, @commitlint/cli, @commitlint/config-conventional, eslint, eslint-config-prettier, husky, lint-staged, @nuxt/eslint (+8 more)

### Community 72 - ".MessageService"
Cohesion: 0.20
Nodes (7): Router, NewAttachmentHandler(), NewGormMessageRepository(), NewMessageService(), AttachmentHandler, AttachmentRepository, MessageRepository

### Community 73 - "NewMongoDB"
Cohesion: 0.18
Nodes (14): createMongoIndexes(), Context, Database, NewMongoDB(), newMongoRegistry(), parseMongoDBName(), uuidDecodeValue(), uuidEncodeValue() (+6 more)

### Community 74 - "user_service.go"
Cohesion: 0.14
Nodes (9): Location, Time, UUID, UserNotificationUpdate, UserPaymentInvoice, UserInvoiceGenerateParams, UserNotificationUpdateParams, UserSendPhoneDeadEmailParams (+1 more)

### Community 75 - "PhoneAPIKeyService"
Cohesion: 0.25
Nodes (4): Context, service, UUID, PhoneAPIKeyService

### Community 76 - "Design"
Cohesion: 0.13
Nodes (14): 1. Data model — `api/pkg/entities/phone.go`, 2. Event payload — `api/pkg/events/message_phone_received_event.go`, 3. Populate the flag — `api/pkg/services/message_service.go`, 4. Thread service trigger — `api/pkg/services/message_thread_service.go`, 5. Listener — `api/pkg/listeners/message_thread_listener.go`, 6. Request/params — phone upsert, 7. Web frontend — `web/`, 8. Swagger (+6 more)

### Community 77 - "Build, Test, and Lint Commands"
Cohesion: 0.13
Nodes (14): Android (Kotlin), Android (Kotlin), Android — Task-Oriented, Event-Driven, API (Go), API (Go), API — Layered Architecture with Event-Driven Processing, Architecture, Build, Test, and Lint Commands (+6 more)

### Community 78 - "MessageThreadHeader.vue"
Cohesion: 0.13
Nodes (11): appStore, authStore, { mdAndDown, lgAndUp }, notificationsStore, owners, phonesStore, route, router (+3 more)

### Community 79 - "index.vue"
Cohesion: 0.14
Nodes (14): errors, { formatPhoneNumber }, formAttachments, formContent, formPhoneNumber, getRecipientNumber(), { mdAndDown, mdAndUp }, notificationsStore (+6 more)

### Community 80 - "errors.ts"
Cohesion: 0.20
Nodes (5): Bag, AxiosLikeError, ErrorMessages, getErrorMessages(), sanitize()

### Community 81 - "NotificationEmailFactory"
Cohesion: 0.25
Nodes (10): generateWebhookSendFailedPlainText(), NewHermesNotificationEmailFactory(), replaceWebhookSendFailedEventPayloadPlaceholder(), T, TestGenerateWebhookSendFailedPlainTextFallsBackToFormattedPayload(), testNotificationEmailFactory(), TestReplaceWebhookSendFailedEventPayloadPlaceholder(), TestWebhookSendFailedFormatsOnlyEventPayload() (+2 more)

### Community 82 - "gormPhoneRepository"
Cohesion: 0.32
Nodes (5): Context, DB, Phone, UUID, gormPhoneRepository

### Community 83 - "mongoHeartbeatMonitorRepository"
Cohesion: 0.24
Nodes (6): Collection, Context, Database, UUID, NewMongoHeartbeatMonitorRepository(), mongoHeartbeatMonitorRepository

### Community 84 - "otelTracer"
Cohesion: 0.37
Nodes (6): functionName(), getName(), Context, Ctx, Span, otelTracer

### Community 85 - "Affiliates Landing Page — Design"
Cohesion: 0.14
Nodes (13): 1. Hero, 2. Why promote httpSMS (benefit cards), 3. How it works (3 numbered steps), 4. FAQ (static two-column grid — no expansion panels), 5. Closing CTA banner, Affiliates Landing Page — Design, Conventions to follow, Files (+5 more)

### Community 86 - "scripts"
Cohesion: 0.14
Nodes (14): scripts, api:models, build, dev, generate, lint, lint:js, lint:prettier (+6 more)

### Community 87 - "New"
Cohesion: 0.28
Nodes (10): New(), T, TestGetSendDelay_BulkIndex_RateBasedDelay(), TestGetSendDelay_BulkIndex_ZeroRate_ReturnsZero(), TestGetSendDelay_IndexZero_ReturnsZero(), TestGetSendDelay_NoSendAtNoIndex_ReturnsZero(), TestGetSendDelay_WithSendAt_ReturnsTimeUntil(), TestGetSendDelay_WithSendAtInPast_ReturnsZero() (+2 more)

### Community 88 - "HermesGeneratorConfig"
Cohesion: 0.15
Nodes (6): Hermes, newHermesTheme(), HermesGeneratorConfig, hermesTheme, StylesDefinition, Theme

### Community 89 - "gormHeartbeatMonitorRepository"
Cohesion: 0.28
Nodes (4): Context, DB, UUID, gormHeartbeatMonitorRepository

### Community 90 - ".ValidateStore"
Cohesion: 0.41
Nodes (5): BulkMessage, Context, Location, Values, FileHeader

### Community 91 - "Contributor Covenant Code of Conduct"
Cohesion: 0.15
Nodes (12): 1. Correction, 2. Warning, 3. Temporary Ban, 4. Permanent Ban, Attribution, Contributor Covenant Code of Conduct, Enforcement, Enforcement Guidelines (+4 more)

### Community 92 - "Login "Last Used" Badge — Design"
Cohesion: 0.15
Nodes (12): Approach, Display, Edge cases, Goal, Login "Last Used" Badge — Design, Out of scope, Problem, Reading the method (+4 more)

### Community 93 - "index.vue"
Cohesion: 0.15
Nodes (12): config, faqPanel, { lgAndUp, mdAndUp, mdAndDown, md, smAndDown, xl }, planMessages, planMonthlyPrice, planYearlyMonthlyPrice, planYearlyPrice, pricing (+4 more)

### Community 94 - "Email"
Cohesion: 0.29
Nodes (6): formatBillingDate(), Hermes, Location, Time, Email, hermesUserEmailFactory

### Community 95 - "HeartbeatIndex"
Cohesion: 0.18
Nodes (4): Context, Values, HeartbeatIndex, HeartbeatStore

### Community 96 - "UserHandlerValidator"
Cohesion: 0.23
Nodes (5): Context, Values, NewUserHandlerValidator(), UserUpdate, UserHandlerValidator

### Community 97 - "EmailNotificationService"
Cohesion: 0.39
Nodes (4): Context, Duration, service, EmailNotificationService

### Community 99 - "Webhook Email Payload Formatting"
Cohesion: 0.17
Nodes (11): 1. Payload formatter, 2. Hermes dictionary rendering, 3. Webhook email factory, Data Flow, Decisions, Design, Error Handling and Security, Out of Scope (+3 more)

### Community 100 - "httpSMS"
Cohesion: 0.17
Nodes (12): Android App, API, API Clients, Chat/forum, Flows, httpSMS, Integration Testing, License (+4 more)

### Community 101 - "auth.ts"
Cohesion: 0.29
Nodes (8): createApiFetch(), setApiKey(), setAuthHeader(), useApi(), useApiComposable(), AuthUser, useAuthStore, EntitiesUser

### Community 102 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, module, moduleResolution, paths, strict, target, files, include (+3 more)

### Community 103 - "UserEmailFactory"
Cohesion: 0.29
Nodes (7): NewHermesUserEmailFactory(), T, TestFormatBillingDate_RendersInProvidedTimezone(), TestUsageLimitAlert_IncludesPercentBreakdownAndLimit(), TestUsageLimitExceeded_IncludesBreakdownAndBillingPeriod(), testUserEmailFactory(), UserEmailFactory

### Community 104 - "attachment_repository.go"
Cohesion: 0.22
Nodes (8): Ctx, AllowedContentTypes(), ContentTypeFromExtension(), ExtensionFromContentType(), SanitizeFilename(), T, TestExtensionFromContentType(), TestSanitizeFilename()

### Community 105 - "UserListener"
Cohesion: 0.40
Nodes (3): Context, Event, UserListener

### Community 106 - "gormHeartbeatRepository"
Cohesion: 0.25
Nodes (5): Context, DB, NewGormHeartbeatRepository(), gormHeartbeatRepository, HeartbeatRepository

### Community 107 - "MessageThreadIndex"
Cohesion: 0.20
Nodes (4): Context, Values, MessageThreadIndex, MessageThreadUpdate

### Community 108 - "PhoneAPIKeyIndex"
Cohesion: 0.20
Nodes (4): Context, Values, PhoneAPIKeyIndex, PhoneAPIKeyStoreRequest

### Community 109 - "response.go"
Cohesion: 0.18
Nodes (10): T, BadRequest, InternalServerError, NoContent, NotFound, Ok, OkString, PaymentRequired (+2 more)

### Community 110 - "gormLogger"
Cohesion: 0.24
Nodes (5): Context, Interface, Time, LogLevel, gormLogger

### Community 111 - "resetErrors"
Cohesion: 0.18
Nodes (11): defaultTimezone(), minuteToClock(), onDiscordCreate(), onDiscordEdit(), onWebhookCreate(), onWebhookEdit(), openCreateSchedule(), openEditSchedule() (+3 more)

### Community 112 - "messages.ts"
Cohesion: 0.20
Nodes (8): SendMessageRequest, SIM, useMessagesStore, useThreadsStore, EntitiesBulkMessage, EntitiesMessage, EntitiesMessageThread, SearchMessagesRequest

### Community 113 - "default.vue"
Cohesion: 0.20
Nodes (9): appStore, authStore, config, hasDrawer, { lgAndUp }, phonesStore, route, startPoller() (+1 more)

### Community 114 - "Phone"
Cohesion: 0.25
Nodes (5): Duration, SIM, Time, UUID, Phone

### Community 115 - "PhoneAPIKey"
Cohesion: 0.25
Nodes (6): StringArray, Time, UUID, PhoneAPIKey, PhoneAPIKeyResponse, PhoneAPIKeysResponse

### Community 116 - "MessagePhoneSendingPayload"
Cohesion: 0.19
Nodes (7): SIM, Time, UUID, Time, UUID, MessagePhoneSendingPayload, MessageSendExpiredCheckPayload

### Community 117 - "HeartbeatListener"
Cohesion: 0.44
Nodes (3): Context, Event, HeartbeatListener

### Community 118 - "Getting Started page — Design"
Cohesion: 0.22
Nodes (8): Conventions to follow, Discoverability, Getting Started page — Design, Goal, Out of scope, Page structure (top → bottom), Placeholders, Route

### Community 119 - "playwright"
Cohesion: 0.28
Nodes (8): BROWSER, npx, axiom, context7, playwright, mcp-remote, @modelcontextprotocol/server-playwright, @upstash/context7-mcp

### Community 120 - "Integration Tests"
Cohesion: 0.22
Nodes (7): Adding New Tests, CI/CD, Integration Tests, Prerequisites, Project Structure, Test Coverage, Test Data

### Community 121 - "Self Host Setup - Docker"
Cohesion: 0.22
Nodes (9): 1. Setup Firebase, 2. Setup SMTP Email service, 3. Setup Cloudflare Turnstile, 4. Download the code, 5. Setup the environment variables, 6. Build and Run, 7. Create the System User, 8. Build the Android App. (+1 more)

### Community 122 - "billing_usage_test.go"
Cohesion: 0.43
Nodes (7): T, TestBillingUsage_IsEntitled_BelowLimit(), TestBillingUsage_IsEntitled_BulkCountExceeding(), TestBillingUsage_IsEntitled_BulkCountFittingExactly(), TestBillingUsage_IsEntitled_ExceedingLimitIsNotEntitled(), TestBillingUsage_IsEntitled_ReachingExactlyLimitIsEntitled(), TestBillingUsage_TotalMessages()

### Community 123 - "bulk_message_handler.go"
Cohesion: 0.32
Nodes (4): encodeBase62(), Ctx, sanitizeFilename(), truncateFilename()

### Community 124 - "testClient"
Cohesion: 0.39
Nodes (6): Builder, testClient(), T, TestPhoneAPIKeyHandler_delete(), TestPhoneAPIKeyHandler_index(), TestPhoneAPIKeyHandler_store()

### Community 125 - "EmailNotificationListener"
Cohesion: 0.46
Nodes (3): Context, Event, EmailNotificationListener

### Community 126 - "MemoryAttachmentRepository"
Cohesion: 0.36
Nodes (4): Context, Map, NewMemoryAttachmentRepository(), MemoryAttachmentRepository

### Community 127 - "Nasak SMS production deployment"
Cohesion: 0.25
Nodes (7): Android gateway, Architecture, First deployment, Nasak SMS production deployment, Operations, Required external configuration, Server layout

### Community 128 - "Running Locally"
Cohesion: 0.25
Nodes (8): 1. Generate Firebase Credentials, 2. Set Environment Variable, 3. Start the Stack, 4. Wait for Seeding, 5. Run Tests, 6. Tear Down, One-Liner, Running Locally

### Community 129 - "BootReceiver"
Cohesion: 0.43
Nodes (4): BootReceiver, BroadcastReceiver, Context, Intent

### Community 130 - "PhoneNumberValidator.kt"
Cohesion: 0.48
Nodes (6): fixNumber(), formatE164(), isValidPhoneNumber(), Boolean, String, PhoneNumberValidator

### Community 131 - "NewSMTPEmailService"
Cohesion: 0.33
Nodes (5): Context, NewSMTPEmailService(), Auth, SMTPConfig, smtpMailer

### Community 132 - "message_send_schedule_test.go"
Cohesion: 0.48
Nodes (6): T, TestResolveScheduledAt_BeforeWindow_ReturnsWindowStart(), TestResolveScheduledAt_InactiveSchedule_ReturnsCurrentUTC(), TestResolveScheduledAt_NilSchedule_ReturnsCurrentUTC(), TestResolveScheduledAt_NoWindows_ReturnsCurrentUTC(), TestResolveScheduledAt_WithinWindow_ReturnsCurrentUTC()

### Community 133 - "user_test.go"
Cohesion: 0.48
Nodes (6): T, TestUser_GetBillingAnchorDay_EmptySubscription(), TestUser_GetBillingAnchorDay_FreeUser(), TestUser_GetBillingAnchorDay_PaidUser(), TestUser_GetBillingAnchorDay_PaidUserDay31(), TestUser_GetBillingAnchorDay_PaidUserNilRenewsAt()

### Community 134 - "MessageAPISentPayload"
Cohesion: 0.29
Nodes (5): SIM, Time, UUID, Duration, MessageAPISentPayload

### Community 135 - "BillingListener"
Cohesion: 0.48
Nodes (3): Context, Event, BillingListener

### Community 136 - "PhoneAPIKeyListener"
Cohesion: 0.48
Nodes (3): Context, Event, PhoneAPIKeyListener

### Community 137 - "BillingUsageHistory"
Cohesion: 0.29
Nodes (3): Context, Values, BillingUsageHistory

### Community 138 - "BulkMessage"
Cohesion: 0.38
Nodes (3): Location, Time, BulkMessage

### Community 140 - "File Structure"
Cohesion: 0.29
Nodes (6): File Structure, Global Constraints, Task 1: Build the event payload formatter, Task 2: Add opt-in rich dictionary rendering to the Hermes theme, Task 3: Wire formatting into the webhook failure email, Webhook Email Payload Formatting Implementation Plan

### Community 141 - "Message Thread Archive UI Design"
Cohesion: 0.29
Nodes (6): Active Thread Styling, Archive and Unarchive Behavior, Error Handling, Goal, Message Thread Archive UI Design, Validation

### Community 142 - "website.vue"
Cohesion: 0.29
Nodes (5): appStore, authStore, { lgAndUp, mdAndUp }, route, router

### Community 143 - "parseErrors"
Cohesion: 0.29
Nodes (7): deleteDiscord(), deleteWebhook(), loadDiscordIntegrations(), loadWebhooks(), parseErrors(), saveDiscord(), saveWebhook()

### Community 144 - "package.json"
Cohesion: 0.29
Nodes (6): lint-staged, *.{css,scss,sass,vue}, *.{js,ts,vue}, name, private, type

### Community 145 - "main"
Cohesion: 0.40
Nodes (4): main(), getEnvWithDefault(), LoadEnv(), splitCommaEnv()

### Community 146 - "memoryCache"
Cohesion: 0.40
Nodes (3): Context, Duration, memoryCache

### Community 147 - "redisCache"
Cohesion: 0.40
Nodes (3): Context, Duration, redisCache

### Community 148 - "HeartbeatMonitor"
Cohesion: 0.33
Nodes (3): Time, UUID, HeartbeatMonitor

### Community 150 - "Global Constraints"
Cohesion: 0.33
Nodes (5): Global Constraints, Message Thread Archive UI Implementation Plan, Task 1: Active Thread Primary Color, Task 2: Archive Without Switching Filters, Task 3: Production Validation

### Community 151 - "CopyButton.vue"
Cohesion: 0.33
Nodes (4): disabled, notificationsStore, props, { smAndDown }

### Community 152 - "LoadingButton.vue"
Cohesion: 0.40
Nodes (5): emit, isClicked, onClick(), props, size

### Community 153 - "countries.ts"
Cohesion: 0.33
Nodes (5): canadaProvinces, countries, getStateOptions(), SelectOption, usStates

### Community 154 - "HeartbeatWorker"
Cohesion: 0.40
Nodes (3): HeartbeatWorker, Result, Worker

### Community 155 - "gradlew"
Cohesion: 0.60
Nodes (3): gradlew script, die(), warn()

### Community 157 - "MessageCallMissedPayload"
Cohesion: 0.40
Nodes (4): SIM, Time, UUID, MessageCallMissedPayload

### Community 158 - "MessageNotificationSentPayload"
Cohesion: 0.40
Nodes (4): Duration, Time, UUID, MessageNotificationSentPayload

### Community 159 - "MessagePhoneDeliveredPayload"
Cohesion: 0.40
Nodes (4): SIM, Time, UUID, MessagePhoneDeliveredPayload

### Community 160 - "MessagePhoneReceivedPayload"
Cohesion: 0.40
Nodes (4): SIM, Time, UUID, MessagePhoneReceivedPayload

### Community 161 - "MessagePhoneSentPayload"
Cohesion: 0.40
Nodes (4): SIM, Time, UUID, MessagePhoneSentPayload

### Community 162 - "MessageSendRetryPayload"
Cohesion: 0.40
Nodes (4): SIM, Time, UUID, MessageSendRetryPayload

### Community 164 - "emulator_fcm_client.go"
Cohesion: 0.60
Nodes (4): emulatorAndroid, emulatorFCMMessage, emulatorFCMRequest, emulatorFCMResponse

### Community 165 - ".Check"
Cohesion: 0.50
Nodes (3): formatEntityName(), Context, EntitlementCheckResult

### Community 167 - "validateAttachmentURL"
Cohesion: 0.60
Nodes (3): Context, saveToCache(), validateAttachmentURL()

### Community 168 - "Features"
Cohesion: 0.40
Nodes (5): Back Pressure, End-to-end Encryption, Features, Message Expiration, Webhook

### Community 169 - "BackButton.vue"
Cohesion: 0.40
Nodes (3): props, router, { smAndDown }

### Community 170 - "end-to-end-encryption-to-sms-messages.vue"
Cohesion: 0.40
Nodes (4): encryptTab, { mdAndUp }, receiveTab, sendTab

### Community 171 - "index.vue"
Cohesion: 0.40
Nodes (4): Article, articles, { mdAndUp, smAndDown }, sortedArticles

### Community 172 - "index.vue"
Cohesion: 0.40
Nodes (4): authStore, { lgAndUp }, phonesStore, threadsStore

### Community 173 - "Nuxt Minimal Starter"
Cohesion: 0.40
Nodes (4): Development Server, Nuxt Minimal Starter, Production, Setup

### Community 174 - "MessageNotificationFailedPayload"
Cohesion: 0.50
Nodes (3): Time, UUID, MessageNotificationFailedPayload

### Community 175 - "MessageSendScheduleDeletedPayload"
Cohesion: 0.50
Nodes (3): Time, UUID, MessageSendScheduleDeletedPayload

### Community 176 - "PhoneHeartbeatOnlinePayload"
Cohesion: 0.50
Nodes (3): Time, UUID, PhoneHeartbeatOnlinePayload

### Community 182 - "MessageResponse"
Cohesion: 0.67
Nodes (3): Message, MessageResponse, MessagesResponse

### Community 183 - "PhoneResponse"
Cohesion: 0.67
Nodes (3): Phone, PhoneResponse, PhonesResponse

### Community 184 - "UserResponse"
Cohesion: 0.50
Nodes (3): Time, UserResponse, UserSubscriptionPaymentsResponse

### Community 185 - "Security Policy"
Cohesion: 0.50
Nodes (3): Reporting a Vulnerability, Security Policy, Supported Versions

### Community 186 - "Troubleshooting"
Cohesion: 0.50
Nodes (4): API fails to start, Seed container fails, Tests timeout waiting for `delivered` status, Troubleshooting

### Community 187 - "Architecture"
Cohesion: 0.50
Nodes (4): Architecture, Components, FCM Redirect, How It Works

### Community 188 - "AppToast.vue"
Cohesion: 0.50
Nodes (3): { lgAndUp }, notificationActive, notificationsStore

### Community 189 - "saveSchedule"
Cohesion: 0.50
Nodes (4): clockToMinute(), deleteSchedule(), loadSendSchedules(), saveSchedule()

### Community 190 - "scheduleDayEnabled"
Cohesion: 0.50
Nodes (4): scheduleAddWindow(), scheduleDayEnabled(), scheduleToggleDay(), scheduleWindowsForDay()

## Knowledge Gaps
- **633 isolated node(s):** `@modelcontextprotocol/server-playwright`, `BROWSER`, `@upstash/context7-mcp`, `mcp-remote`, `Constants` (+628 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **43 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `UserID` connect `UserID` to `PhoneNotificationService`, `MessageThreadService`, `MessageAPISentPayload`, `WebhookService`, `handler`, `BulkMessage`, `MessageIndex`, `MessageService`, `Logger`, `BillingUsage`, `HeartbeatMonitor`, `DiscordService`, `Integration3CXService`, `Values`, `User`, `Context`, `MessageCallMissedPayload`, `MessageNotificationSentPayload`, `MessagePhoneDeliveredPayload`, `MessagePhoneReceivedPayload`, `MessagePhoneSentPayload`, `MessageSendRetryPayload`, `IndexParams`, `UserService`, `Client`, `.Check`, `LemonsqueezyService`, `SIM`, `EventDispatcher`, `MessageNotificationFailedPayload`, `MessageSendScheduleDeletedPayload`, `PhoneHeartbeatOnlinePayload`, `Message`, `.WebhookSendFailed`, `HeartbeatService`, `hermesNotificationEmailFactory`, `MessageSendSchedule`, `Heartbeat`, `message_service.go`, `PhoneHeartbeatMissedPayload`, `gormPhoneAPIKeyRepository`, `WebhookStore`, `user_service.go`, `PhoneAPIKeyService`, `gormPhoneRepository`, `mongoHeartbeatMonitorRepository`, `gormHeartbeatMonitorRepository`, `.ValidateStore`, `UserHandlerValidator`, `gormHeartbeatRepository`, `MessageThreadIndex`, `Phone`, `PhoneAPIKey`, `MessagePhoneSendingPayload`?**
  _High betweenness centrality (0.110) - this node is a cross-community bridge._
- **Why does `Logger` connect `Logger` to `Container`, `PhoneNotificationService`, `MessageThreadService`, `MessageAPISentPayload`, `BillingListener`, `PhoneAPIKeyListener`, `handler`, `WebhookService`, `MessageService`, `Tracer`, `BillingUsage`, `DiscordService`, `Cache`, `Integration3CXService`, `BillingService`, `User`, `MessageHandlerValidator`, `IndexParams`, `UserService`, `MessageThreadHandler`, `PhoneService`, `MessageSendScheduleHandlerValidator`, `Client`, `service`, `LemonsqueezyService`, `googlePushQueue`, `EventDispatcher`, `HeartbeatService`, `PhoneAPIKeyHandler`, `MessageSendScheduleService`, `MessageSendSchedule`, `zerologLogger`, `container.go`, `Heartbeat`, `UserHandler`, `MessageListener`, `gormPhoneAPIKeyRepository`, `.MessageService`, `PhoneAPIKeyService`, `gormPhoneRepository`, `mongoHeartbeatMonitorRepository`, `otelTracer`, `gormHeartbeatMonitorRepository`, `.ValidateStore`, `UserHandlerValidator`, `EmailNotificationService`, `noopLogger`, `UserListener`, `gormHeartbeatRepository`, `gormLogger`, `HeartbeatListener`, `EmailNotificationListener`, `MemoryAttachmentRepository`?**
  _High betweenness centrality (0.088) - this node is a cross-community bridge._
- **Why does `Tracer` connect `Tracer` to `Container`, `NewSMTPEmailService`, `PhoneNotificationService`, `MessageThreadService`, `BillingListener`, `PhoneAPIKeyListener`, `handler`, `WebhookService`, `MessageService`, `Logger`, `memoryCache`, `redisCache`, `BillingUsage`, `DiscordService`, `Cache`, `Integration3CXService`, `BillingService`, `User`, `MessageHandlerValidator`, `IndexParams`, `UserService`, `MessageThreadHandler`, `PhoneService`, `MessageSendScheduleHandlerValidator`, `Client`, `LemonsqueezyService`, `googlePushQueue`, `EventDispatcher`, `HeartbeatService`, `PhoneAPIKeyHandler`, `MessageSendScheduleService`, `MessageSendSchedule`, `container.go`, `Heartbeat`, `UserHandler`, `MessageListener`, `gormPhoneAPIKeyRepository`, `.MessageService`, `PhoneAPIKeyService`, `gormPhoneRepository`, `mongoHeartbeatMonitorRepository`, `otelTracer`, `gormHeartbeatMonitorRepository`, `UserHandlerValidator`, `EmailNotificationService`, `UserListener`, `gormHeartbeatRepository`, `gormLogger`, `HeartbeatListener`, `EmailNotificationListener`, `MemoryAttachmentRepository`?**
  _High betweenness centrality (0.085) - this node is a cross-community bridge._
- **Are the 2 inferred relationships involving `UserID` (e.g. with `BearerAuth()` and `.EventsQueueConfiguration()`) actually correct?**
  _`UserID` has 2 INFERRED edges - model-reasoned connections that need verification._
- **Are the 74 inferred relationships involving `New()` (e.g. with `axiomLogger()` and `consoleLogger()`) actually correct?**
  _`New()` has 74 INFERRED edges - model-reasoned connections that need verification._
- **What connects `@modelcontextprotocol/server-playwright`, `BROWSER`, `@upstash/context7-mcp` to the rest of the system?**
  _633 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Settings` be split into smaller, more focused modules?**
  _Cohesion score 0.06506849315068493 - nodes in this community are weakly interconnected._