<script setup lang="ts">
import { mdiWeatherNight, mdiWhiteBalanceSunny } from '@mdi/js'
import { useTheme } from 'vuetify'

const theme = useTheme()
const isDark = computed(() => theme.global.current.value.dark)

function toggleTheme() {
  const next = isDark.value ? 'nasakLight' : 'nasakDark'
  theme.global.name.value = next
  if (import.meta.client) {
    localStorage.setItem('nasak-sms-theme', next)
  }
}

onMounted(() => {
  const saved = localStorage.getItem('nasak-sms-theme')
  if (saved === 'nasakLight' || saved === 'nasakDark') {
    theme.global.name.value = saved
  }
})
</script>

<template>
  <VBtn
    :aria-label="isDark ? 'فعال‌کردن تم روشن' : 'فعال‌کردن تم تاریک'"
    :icon="isDark ? mdiWhiteBalanceSunny : mdiWeatherNight"
    variant="text"
    size="large"
    @click="toggleTheme"
  />
</template>
