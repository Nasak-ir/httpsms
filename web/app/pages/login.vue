<script setup lang="ts">
definePageMeta({
  middleware: ['guest'],
  layout: 'blank',
})

useSeoMeta({
  title: 'ورود | پیامک نسک',
  description: 'ورود امن به پنل درگاه پیامک اختصاصی نسک.',
})

const route = useRoute()
const to = computed(() => (route.query.to as string) || '/threads')
</script>

<template>
  <VContainer fluid class="login-shell">
    <div class="theme-control"><ThemeToggle /></div>
    <VRow align="center" justify="center" class="w-100">
      <VCol cols="12" sm="9" md="6" lg="4" xl="3">
        <div class="brand">
          <VImg src="/brand/nasak-mark.png" width="64" height="64" alt="نسک" />
          <div>
            <h1>پیامک نسک</h1>
            <p>درگاه اختصاصی ارسال با سیم‌کارت شما</p>
          </div>
        </div>
        <VCard class="login-card" flat>
          <VCardText>
            <ClientOnly>
              <FirebaseAuth :to="to" />
            </ClientOnly>
          </VCardText>
        </VCard>
        <VBtn to="/" variant="text" block class="mt-3"
          >بازگشت به صفحه اصلی</VBtn
        >
      </VCol>
    </VRow>
  </VContainer>
</template>

<style scoped>
.login-shell {
  display: flex;
  min-height: 100dvh;
  align-items: center;
  background:
    linear-gradient(145deg, rgb(49 183 122 / 8%), transparent 45%),
    rgb(var(--v-theme-background));
}

.theme-control {
  position: fixed;
  z-index: 2;
  top: 12px;
  left: 12px;
}

.brand {
  display: flex;
  gap: 16px;
  align-items: center;
  justify-content: center;
  margin-bottom: 24px;
}

.brand h1 {
  margin: 0;
  font-size: 1.6rem;
}

.brand p {
  margin: 4px 0 0;
  color: #9fb0c1;
}

.login-card {
  padding: 12px;
  border: 1px solid var(--nasak-border);
  border-radius: 8px;
}
</style>
