<script setup lang="ts">
import {
  mdiArrowRight,
  mdiChartBoxOutline,
  mdiDownload,
  mdiUpload,
} from '@mdi/js'

definePageMeta({ middleware: ['auth'] })
useHead({ title: 'آمار مصرف | پیامک نسک' })

const billingStore = useBillingStore()
const loading = ref(true)

const sent = computed(() => billingStore.billingUsage?.sent_messages ?? 0)
const received = computed(
  () => billingStore.billingUsage?.received_messages ?? 0,
)
const total = computed(() => sent.value + received.value)

onMounted(async () => {
  try {
    await Promise.all([
      billingStore.loadBillingUsage(),
      billingStore.loadBillingUsageHistory(),
    ])
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <VContainer>
    <VAppBar flat>
      <VBtn icon to="/threads" aria-label="بازگشت">
        <VIcon :icon="mdiArrowRight" />
      </VBtn>
      <VToolbarTitle>آمار مصرف</VToolbarTitle>
      <ThemeToggle />
      <VProgressLinear
        :active="loading"
        indeterminate
        absolute
        location="bottom"
        color="primary"
      />
    </VAppBar>

    <VRow class="mt-6">
      <VCol cols="12">
        <VAlert
          type="success"
          variant="tonal"
          :icon="mdiChartBoxOutline"
          title="بدون اعتبار نرم‌افزاری"
        >
          این آمار فقط برای پایش سرویس است. پیامک نسک محدودیت پلن، credit یا
          هزینه اشتراک ندارد؛ هزینه احتمالی اپراتور سیم‌کارت جداگانه است.
        </VAlert>
      </VCol>
      <VCol cols="12" md="4">
        <VCard class="pa-5" variant="outlined">
          <VIcon :icon="mdiUpload" color="primary" />
          <div class="text-h4 nasak-number mt-4">
            {{ sent.toLocaleString('fa-IR') }}
          </div>
          <div class="nasak-muted">پیام ارسال‌شده</div>
        </VCard>
      </VCol>
      <VCol cols="12" md="4">
        <VCard class="pa-5" variant="outlined">
          <VIcon :icon="mdiDownload" color="info" />
          <div class="text-h4 nasak-number mt-4">
            {{ received.toLocaleString('fa-IR') }}
          </div>
          <div class="nasak-muted">پیام دریافت‌شده</div>
        </VCard>
      </VCol>
      <VCol cols="12" md="4">
        <VCard class="pa-5" variant="outlined">
          <VIcon :icon="mdiChartBoxOutline" color="warning" />
          <div class="text-h4 nasak-number mt-4">
            {{ total.toLocaleString('fa-IR') }}
          </div>
          <div class="nasak-muted">مجموع پردازش‌شده</div>
        </VCard>
      </VCol>
    </VRow>
  </VContainer>
</template>
