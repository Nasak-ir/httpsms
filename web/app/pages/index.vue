<script setup lang="ts">
import {
  mdiApi,
  mdiCellphoneCheck,
  mdiCheckDecagramOutline,
  mdiCodeJson,
  mdiLockOutline,
  mdiSendOutline,
} from '@mdi/js'

definePageMeta({ layout: 'website' })

useSeoMeta({
  title: 'پیامک نسک | درگاه پیامک اختصاصی و بدون اعتبار نرم‌افزاری',
  description:
    'گوشی اندرویدی خود را به درگاه امن پیامک تبدیل کنید؛ ارسال API، صف، گزارش تحویل و وب‌هوک بدون خرید اعتبار نرم‌افزاری.',
})

const features = [
  {
    icon: mdiSendOutline,
    title: 'ارسال مستقیم',
    text: 'پیام از API به گوشی شما می‌رسد و با سیم‌کارت انتخاب‌شده ارسال می‌شود.',
  },
  {
    icon: mdiCellphoneCheck,
    title: 'گزارش واقعی',
    text: 'وضعیت‌های در صف، ارسال‌شده، تحویل‌شده و ناموفق ثبت و قابل پیگیری‌اند.',
  },
  {
    icon: mdiApi,
    title: 'API تمیز',
    text: 'برای OTP، اعلان سفارش، پیام گروهی و اتصال سیستم‌های دیگر آماده است.',
  },
  {
    icon: mdiLockOutline,
    title: 'زیرساخت خصوصی',
    text: 'API و دیتابیس روی سرور نسک هستند و URLهای داخلی در اختیار سرویس ثالث نیستند.',
  },
]
</script>

<template>
  <VContainer class="landing">
    <section class="hero">
      <div class="hero-copy">
        <div class="eyebrow">
          <VIcon :icon="mdiCheckDecagramOutline" />
          بدون credit و اشتراک نرم‌افزاری
        </div>
        <h1>پیامک نسک</h1>
        <p>
          درگاه پیامک اختصاصی برای سایت‌ها و فروشگاه‌های نسک. گوشی اندرویدی شما
          پیام‌ها را با سیم‌کارت خودتان ارسال می‌کند؛ ساده، قابل کنترل و بدون
          هزینه سرویس واسط.
        </p>
        <div class="hero-actions">
          <VBtn
            to="/login"
            color="primary"
            size="large"
            :prepend-icon="mdiCodeJson"
          >
            ورود به پنل
          </VBtn>
          <VBtn
            href="/phone-api-keys"
            size="large"
            variant="outlined"
            :prepend-icon="mdiCellphoneCheck"
          >
            اتصال گوشی
          </VBtn>
        </div>
        <p class="carrier-note">
          هزینه احتمالی اپراتور سیم‌کارت مطابق بسته یا تعرفه خط شما محاسبه
          می‌شود.
        </p>
      </div>
      <div class="hero-visual" aria-hidden="true">
        <VImg src="/brand/nasak-logo.png" alt="" width="320" max-width="80vw" />
        <div class="status-line">
          <span class="status-dot" />
          درگاه آماده ارسال
        </div>
      </div>
    </section>

    <section class="feature-grid" aria-label="قابلیت‌های پیامک نسک">
      <article v-for="feature in features" :key="feature.title">
        <VIcon :icon="feature.icon" color="primary" size="28" />
        <h2>{{ feature.title }}</h2>
        <p>{{ feature.text }}</p>
      </article>
    </section>
  </VContainer>
</template>

<style scoped>
.landing {
  padding-top: clamp(32px, 8vw, 96px);
  padding-bottom: 72px;
}

.hero {
  display: grid;
  grid-template-columns: minmax(0, 1.15fr) minmax(300px, 0.85fr);
  gap: clamp(32px, 7vw, 96px);
  align-items: center;
  min-height: min(680px, calc(100dvh - 150px));
}

.hero-copy {
  max-width: 680px;
}

.eyebrow {
  display: inline-flex;
  gap: 8px;
  align-items: center;
  color: rgb(var(--v-theme-primary));
  font-weight: 600;
}

h1 {
  margin: 14px 0 16px;
  font-size: clamp(2.8rem, 8vw, 5.5rem);
  line-height: 1.08;
}

.hero-copy > p {
  max-width: 62ch;
  color: #b8c5d1;
  font-size: clamp(1rem, 2.4vw, 1.18rem);
  line-height: 1.9;
}

.hero-actions {
  display: flex;
  gap: 12px;
  margin-top: 28px;
}

.carrier-note {
  margin-top: 18px;
  font-size: 0.85rem !important;
}

.hero-visual {
  display: grid;
  min-height: 420px;
  place-items: center;
  padding: 36px;
  border: 1px solid var(--nasak-border);
  border-radius: 8px;
  background: #fdfdfd;
}

.status-line {
  display: flex;
  gap: 8px;
  align-items: center;
  color: #173a2c;
  font-weight: 600;
}

.status-dot {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  background: #1f9d66;
  box-shadow: 0 0 0 5px rgb(31 157 102 / 15%);
}

.feature-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 12px;
}

.feature-grid article {
  padding: 24px;
  border: 1px solid var(--nasak-border);
  border-radius: 8px;
  background: rgb(var(--v-theme-surface));
}

.feature-grid h2 {
  margin: 16px 0 8px;
  font-size: 1.05rem;
}

.feature-grid p {
  margin: 0;
  color: rgb(var(--v-theme-on-surface), 0.72);
  line-height: 1.8;
}

@media (width <= 959px) {
  .hero {
    grid-template-columns: 1fr;
  }

  .hero-visual {
    min-height: 320px;
  }

  .feature-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (width <= 599px) {
  .hero-actions {
    align-items: stretch;
    flex-direction: column;
  }

  .feature-grid {
    grid-template-columns: 1fr;
  }
}
</style>
