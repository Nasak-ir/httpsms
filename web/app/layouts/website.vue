<script setup lang="ts">
import { mdiBookOpenOutline, mdiCellphoneLink, mdiLogin } from '@mdi/js'

const authStore = useAuthStore()
const appStore = useAppStore()
</script>

<template>
  <VApp class="nasak-shell">
    <VAppBar flat color="transparent" class="nasak-header">
      <VContainer class="d-flex align-center">
        <NuxtLink to="/" class="brand-link" aria-label="صفحه اصلی پیامک نسک">
          <VImg src="/brand/nasak-mark.png" width="44" height="44" alt="نسک" />
          <span>پیامک نسک</span>
        </NuxtLink>
        <VSpacer />
        <ThemeToggle />
        <VBtn
          :href="appStore.appData.documentationUrl"
          :prepend-icon="mdiBookOpenOutline"
          variant="text"
          class="d-none d-sm-flex"
        >
          مستندات
        </VBtn>
        <VBtn
          v-if="authStore.authUser"
          to="/threads"
          :prepend-icon="mdiCellphoneLink"
          color="primary"
          variant="flat"
        >
          پنل
        </VBtn>
        <VBtn
          v-else
          to="/login"
          :prepend-icon="mdiLogin"
          color="primary"
          variant="flat"
        >
          ورود
        </VBtn>
      </VContainer>
    </VAppBar>

    <VMain>
      <AppToast />
      <slot />
    </VMain>

    <VFooter color="transparent" class="border-t">
      <VContainer class="d-flex flex-wrap align-center ga-3 text-body-small">
        <span>پیامک نسک؛ درگاه اختصاصی ارسال با سیم‌کارت شما</span>
        <VSpacer />
        <NuxtLink to="/privacy-policy">حریم خصوصی</NuxtLink>
        <NuxtLink to="/terms-and-conditions">شرایط استفاده</NuxtLink>
      </VContainer>
    </VFooter>
  </VApp>
</template>

<style scoped>
.nasak-header {
  border-bottom: 1px solid var(--nasak-border);
  backdrop-filter: blur(16px);
}

.brand-link {
  display: flex;
  gap: 10px;
  align-items: center;
  color: inherit;
  font-size: 1.05rem;
  font-weight: 600;
  text-decoration: none;
}

.border-t {
  border-top: 1px solid var(--nasak-border);
}

a {
  color: inherit;
  text-decoration: none;
}
</style>
