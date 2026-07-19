import { defineStore } from 'pinia'
import type { User as FirebaseUser } from 'firebase/auth'
import { setAuthHeader, setApiKey } from '~/composables/useApi'
import type { EntitiesUser } from '~~/shared/types/api'

export interface AuthUser {
  email: string | null
  displayName: string | null
  id: string
}

interface TokenAuthSession {
  email: string | null
  expiresAt: number
  id: string
  idToken: string
  refreshToken: string
}

interface FirebaseEmailAuthData {
  email: string
  expires_in: string
  id_token: string
  local_id: string
  refresh_token: string
}

interface FirebaseEmailAuthResponse {
  data: FirebaseEmailAuthData
}

const TOKEN_AUTH_SESSION_KEY = 'httpsms_token_auth_session'
const TOKEN_REFRESH_MARGIN_MS = 2 * 60 * 1000

export const useAuthStore = defineStore('auth', () => {
  const authStateChanged = ref(false)
  const authUser = ref<AuthUser | null>(null)
  const user = ref<EntitiesUser | null>(null)
  const { apiFetch } = useApi()

  async function setAuthUserAction(newUser: AuthUser | null | undefined) {
    const userChanged = newUser?.id !== authUser.value?.id
    authUser.value = newUser ?? null
    authStateChanged.value = true

    if (userChanged && newUser !== null) {
      await Promise.all([loadUser(), loadPhones()])
    }
  }

  async function onAuthStateChanged(firebaseUser: FirebaseUser | null) {
    if (firebaseUser == null) {
      if (await loadTokenAuthSession()) {
        return
      }
      authUser.value = null
      user.value = null
      authStateChanged.value = true
      setAuthHeader(null)
      setApiKey('')
      return
    }
    clearTokenAuthSession()
    setAuthHeader(await firebaseUser.getIdToken())
    const { uid, email, displayName } = firebaseUser
    authUser.value = { id: uid, email, displayName }
    authStateChanged.value = true
  }

  async function onIdTokenChanged(firebaseUser: FirebaseUser | null) {
    if (firebaseUser == null) {
      if (await loadTokenAuthSession()) {
        return
      }
      setAuthHeader(null)
      setApiKey('')
      return
    }
    clearTokenAuthSession()
    setAuthHeader(await firebaseUser.getIdToken())
  }

  async function authenticateWithEmailPassword(
    email: string,
    password: string,
    mode: 'sign_in' | 'sign_up',
  ) {
    const response = await apiFetch<FirebaseEmailAuthResponse>(
      '/v1/auth/email',
      {
        method: 'POST',
        body: { email, password, mode },
      },
    )

    setTokenAuthSession(response.data)
    await loadUser()
  }

  async function loadTokenAuthSession(): Promise<boolean> {
    const session = getStoredTokenAuthSession()
    if (session === null) {
      return false
    }

    if (session.expiresAt <= Date.now() + TOKEN_REFRESH_MARGIN_MS) {
      if (!(await refreshTokenAuthSession(session))) {
        clearTokenAuthSession()
        return false
      }
      return loadAuthenticatedUserFromTokenSession()
    }

    setAuthHeader(session.idToken)
    authUser.value = {
      id: session.id,
      email: session.email,
      displayName: null,
    }
    authStateChanged.value = true
    return loadAuthenticatedUserFromTokenSession()
  }

  async function refreshTokenAuthSession(
    session: TokenAuthSession,
  ): Promise<boolean> {
    try {
      const response = await apiFetch<FirebaseEmailAuthResponse>(
        '/v1/auth/email/refresh',
        {
          method: 'POST',
          body: { refresh_token: session.refreshToken },
        },
      )
      setTokenAuthSession(response.data)
      return true
    } catch {
      return false
    }
  }

  function setTokenAuthSession(data: FirebaseEmailAuthData) {
    const expiresInSeconds = Number.parseInt(data.expires_in || '3600', 10)
    const session: TokenAuthSession = {
      id: data.local_id,
      email: data.email || null,
      idToken: data.id_token,
      refreshToken: data.refresh_token,
      expiresAt: Date.now() + Math.max(expiresInSeconds - 30, 60) * 1000,
    }

    setAuthHeader(session.idToken)
    authUser.value = {
      id: session.id,
      email: session.email,
      displayName: null,
    }
    authStateChanged.value = true
    storeTokenAuthSession(session)
  }

  async function loadAuthenticatedUserFromTokenSession(): Promise<boolean> {
    try {
      await Promise.all([loadUser(), loadPhones()])
      return true
    } catch {
      user.value = null
      setAuthHeader(null)
      setApiKey('')
      clearTokenAuthSession()
      return false
    }
  }

  function getStoredTokenAuthSession(): TokenAuthSession | null {
    if (typeof window === 'undefined') {
      return null
    }
    try {
      const raw = window.localStorage.getItem(TOKEN_AUTH_SESSION_KEY)
      if (!raw) {
        return null
      }
      const session = JSON.parse(raw) as Partial<TokenAuthSession>
      if (
        !session.id ||
        !session.idToken ||
        !session.refreshToken ||
        !session.expiresAt
      ) {
        return null
      }
      return session as TokenAuthSession
    } catch {
      return null
    }
  }

  function storeTokenAuthSession(session: TokenAuthSession) {
    if (typeof window === 'undefined') {
      return
    }
    window.localStorage.setItem(TOKEN_AUTH_SESSION_KEY, JSON.stringify(session))
  }

  function clearTokenAuthSession() {
    if (typeof window === 'undefined') {
      return
    }
    window.localStorage.removeItem(TOKEN_AUTH_SESSION_KEY)
  }

  async function loadUser() {
    try {
      const response = await apiFetch<{ data: EntitiesUser }>('/v1/users/me')
      user.value = response.data
      setApiKey(response.data.api_key)
      return response.data
    } catch (error) {
      user.value = null
      setApiKey('')
      throw error
    }
  }

  async function updateUser(payload: { owner?: string; timezone?: string }) {
    const phonesStore = usePhonesStore()
    if (payload.owner) {
      phonesStore.setOwner(payload.owner)
    }

    const activePhone = phonesStore.activePhone
    if (!activePhone) return

    const response = await apiFetch<{ data: EntitiesUser }>('/v1/users/me', {
      method: 'PUT',
      body: {
        active_phone_id: activePhone.id,
        timezone: payload.timezone ?? user.value?.timezone,
      },
    })

    setApiKey(response.data.api_key)
    user.value = response.data
  }

  async function deleteUserAccount(): Promise<string> {
    await apiFetch<{ message: string }>('/v1/users/me', {
      method: 'DELETE',
    })
    return 'Your account has been deleted successfully'
  }

  async function rotateApiKey(userId: string): Promise<EntitiesUser> {
    const response = await apiFetch<{ data: EntitiesUser }>(
      `/v1/users/${userId}/api-keys`,
      {
        method: 'DELETE',
      },
    )
    user.value = response.data
    setApiKey(response.data.api_key)
    return response.data
  }

  function resetState() {
    user.value = null
    authUser.value = null
    authStateChanged.value = true
    setAuthHeader(null)
    setApiKey('')
    clearTokenAuthSession()
  }

  function loadPhones() {
    const phonesStore = usePhonesStore()
    return phonesStore.loadPhones(false)
  }

  return {
    authStateChanged,
    authUser,
    user,
    setAuthUserAction,
    onAuthStateChanged,
    onIdTokenChanged,
    authenticateWithEmailPassword,
    loadTokenAuthSession,
    loadUser,
    updateUser,
    deleteUserAccount,
    rotateApiKey,
    resetState,
  }
})
