import { useAuthStore } from '../stores/auth'
import { getAuth, signOut } from 'firebase/auth'

export default defineNuxtRouteMiddleware(async (to: { path: string }) => {
  const authStore = useAuthStore()

  if (!authStore.authStateChanged) {
    await new Promise<void>((resolve) => {
      const stop = watch(
        () => authStore.authStateChanged,
        (changed) => {
          if (changed) {
            stop()
            resolve()
          }
        },
        { immediate: true },
      )
    })
  }

  if (authStore.authUser === null) {
    return navigateTo({ path: '/login', query: { to: to.path } })
  }

  if (authStore.user === null) {
    try {
      await authStore.loadUser()
    } catch {
      await signOut(getAuth()).catch(() => undefined)
      authStore.resetState()
      return navigateTo({ path: '/login', query: { to: to.path } })
    }
  }
})
