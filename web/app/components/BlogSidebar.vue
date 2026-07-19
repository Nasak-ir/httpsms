<script setup lang="ts">
import { mdiBookOpenVariant } from '@mdi/js'

const appStore = useAppStore()
</script>

<template>
  <div>
    <NuxtLink to="/" class="text-decoration-none d-flex align-center">
      <VAvatar color="surface" class="pa-1" size="40" rounded="0">
        <VImg height="36" width="36" :src="'/brand/nasak-mark.png'" />
      </VAvatar>
      <h3 class="text-headline-large ml-1 mb-3 mt-3">پیامک نسک</h3>
    </NuxtLink>
    <p class="mt-0">
      درگاه اختصاصی نسک، گوشی اندرویدی شما را به یک درگاه امن ارسال و دریافت
      پیامک تبدیل می‌کند.
    </p>
    <VBtn variant="tonal" :href="appStore.appData.documentationUrl">
      <VIcon start :icon="mdiBookOpenVariant" />
      مستندات
    </VBtn>
  </div>
</template>
