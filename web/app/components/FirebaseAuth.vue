<script setup lang="ts">
import {
  getAuth,
  signInWithPopup,
  GoogleAuthProvider,
  GithubAuthProvider,
  signInWithEmailAndPassword,
  createUserWithEmailAndPassword,
  sendPasswordResetEmail,
} from 'firebase/auth'
import { mdiGoogle, mdiGithub, mdiEmail, mdiEye, mdiEyeOff } from '@mdi/js'
import type { User as FirebaseUser } from 'firebase/auth'
import { ErrorMessages } from '~/utils/errors'

const props = withDefaults(
  defineProps<{
    to?: string
  }>(),
  { to: '/' },
)

const router = useRouter()
const authStore = useAuthStore()
const notificationsStore = useNotificationsStore()
const appStore = useAppStore()

const loading = ref(false)
const showEmailForm = ref(false)
const isSignUp = ref(false)
const showForgotPassword = ref(false)
const resetEmailSent = ref(false)
const showPassword = ref(false)
const email = ref('')
const password = ref('')
const generalError = ref('')
const errorMessages = ref(new ErrorMessages())

type LoginMethod = 'google' | 'github' | 'email'
const LAST_LOGIN_METHOD_KEY = 'httpsms_last_login_method'
const lastUsedMethod = ref<LoginMethod | null>(null)

onMounted(() => {
  try {
    const stored = localStorage.getItem(LAST_LOGIN_METHOD_KEY)
    if (stored === 'google' || stored === 'github' || stored === 'email') {
      lastUsedMethod.value = stored
    }
  } catch (error) {
    console.error(error)
  }
})

function clearErrors() {
  errorMessages.value = new ErrorMessages()
  generalError.value = ''
}

function validateEmail(): boolean {
  clearErrors()
  if (!email.value.trim()) {
    errorMessages.value.add('email', 'ایمیل را وارد کنید.')
    return false
  }
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailRegex.test(email.value.trim())) {
    errorMessages.value.add('email', 'فرمت ایمیل معتبر نیست.')
    return false
  }
  return true
}

function validateLoginForm(): boolean {
  clearErrors()
  let valid = true
  if (!email.value.trim()) {
    errorMessages.value.add('email', 'ایمیل را وارد کنید.')
    valid = false
  } else {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    if (!emailRegex.test(email.value.trim())) {
      errorMessages.value.add('email', 'فرمت ایمیل معتبر نیست.')
      valid = false
    }
  }
  if (!password.value) {
    errorMessages.value.add('password', 'رمز عبور را وارد کنید.')
    valid = false
  }
  return valid
}

async function signInWithGoogle() {
  loading.value = true
  try {
    const auth = getAuth()
    const result = await signInWithPopup(auth, new GoogleAuthProvider())
    await onSuccess(result.user, 'google')
  } catch (error: unknown) {
    handleError(error, true)
  } finally {
    loading.value = false
  }
}

async function signInWithGithub() {
  loading.value = true
  try {
    const auth = getAuth()
    const result = await signInWithPopup(auth, new GithubAuthProvider())
    await onSuccess(result.user, 'github')
  } catch (error: unknown) {
    handleError(error, true)
  } finally {
    loading.value = false
  }
}

async function submitEmail() {
  if (!validateLoginForm()) return
  loading.value = true
  try {
    try {
      await submitEmailViaApi()
      return
    } catch (apiError: unknown) {
      if (!shouldRetryWithFirebaseSdk(apiError)) {
        handleApiAuthError(apiError)
        return
      }
    }

    const auth = getAuth()
    let result
    if (isSignUp.value) {
      result = await createUserWithEmailAndPassword(
        auth,
        email.value.trim(),
        password.value,
      )
    } else {
      result = await signInWithEmailAndPassword(
        auth,
        email.value.trim(),
        password.value,
      )
    }
    await onSuccess(result.user, 'email')
  } catch (error: unknown) {
    handleError(error)
  } finally {
    loading.value = false
  }
}

async function submitEmailViaApi() {
  generalError.value = ''
  await authStore.authenticateWithEmailPassword(
    email.value.trim(),
    password.value,
    isSignUp.value ? 'sign_up' : 'sign_in',
  )
  try {
    localStorage.setItem(LAST_LOGIN_METHOD_KEY, 'email')
  } catch (error) {
    console.error(error)
  }
  notificationsStore.addNotification({
    message: 'ورود با موفقیت انجام شد.',
    type: 'success',
  })
  router.push({ path: props.to })
}

async function submitPasswordReset() {
  if (!validateEmail()) return
  loading.value = true
  try {
    const auth = getAuth()
    await sendPasswordResetEmail(auth, email.value.trim())
    resetEmailSent.value = true
  } catch (error: unknown) {
    handleError(error)
  } finally {
    loading.value = false
  }
}

function showForgotPasswordForm() {
  clearErrors()
  resetEmailSent.value = false
  showForgotPassword.value = true
}

function backToSignIn() {
  clearErrors()
  resetEmailSent.value = false
  showForgotPassword.value = false
}

async function onSuccess(user: FirebaseUser, method: LoginMethod) {
  try {
    localStorage.setItem(LAST_LOGIN_METHOD_KEY, method)
  } catch (error) {
    console.error(error)
  }
  notificationsStore.addNotification({
    message: 'ورود با موفقیت انجام شد.',
    type: 'success',
  })
  await authStore.onAuthStateChanged(user)
  router.push({ path: props.to })
}

function shouldRetryWithFirebaseSdk(error: unknown): boolean {
  const apiError = error as {
    response?: { status?: number }
    status?: number
  }
  const status = apiError.response?.status ?? apiError.status
  return status === 502 || status === 503 || status === 504
}

function handleApiAuthError(error: unknown) {
  clearErrors()
  const apiError = error as {
    data?: {
      data?: Record<string, string[]>
      message?: string
    }
  }
  const validationErrors = apiError.data?.data
  if (validationErrors) {
    Object.entries(validationErrors).forEach(([field, messages]) => {
      messages.forEach((message) => errorMessages.value.add(field, message))
    })
    return
  }
  generalError.value =
    apiError.data?.message ||
    'اتصال به سرویس ورود برقرار نشد. چند لحظه بعد دوباره تلاش کنید.'
}

function handleError(error: unknown, isSocial = false) {
  const firebaseError = error as { code?: string; message?: string }
  const code = firebaseError.code || ''

  if (
    code === 'auth/popup-closed-by-user' ||
    code === 'auth/cancelled-popup-request'
  ) {
    return
  }

  if (isSocial) {
    const message = getGeneralErrorMessage(code, firebaseError.message)
    notificationsStore.addNotification({ message, type: 'error' })
    return
  }

  clearErrors()

  switch (code) {
    case 'auth/wrong-password':
      errorMessages.value.add('password', 'رمز عبور درست نیست.')
      break
    case 'auth/invalid-credential':
      errorMessages.value.add('email', 'ایمیل یا رمز عبور درست نیست.')
      errorMessages.value.add('password', 'ایمیل یا رمز عبور درست نیست.')
      break
    case 'auth/user-not-found':
      errorMessages.value.add('email', 'حسابی با این ایمیل پیدا نشد.')
      break
    case 'auth/invalid-email':
      errorMessages.value.add('email', 'فرمت ایمیل معتبر نیست.')
      break
    case 'auth/email-already-in-use':
      errorMessages.value.add('email', 'حسابی با این ایمیل وجود دارد.')
      break
    case 'auth/weak-password':
      errorMessages.value.add('password', 'رمز عبور باید حداقل ۶ کاراکتر باشد.')
      break
    case 'auth/user-disabled':
      errorMessages.value.add('email', 'این حساب غیرفعال شده است.')
      break
    case 'auth/too-many-requests':
      generalError.value =
        'تلاش‌های ناموفق زیاد بود؛ کمی بعد دوباره امتحان کنید.'
      break
    case 'auth/network-request-failed':
      generalError.value =
        'ارتباط با سرور برقرار نشد؛ اتصال اینترنت را بررسی کنید.'
      break
    case 'auth/missing-email':
      errorMessages.value.add('email', 'ایمیل را وارد کنید.')
      break
    default:
      generalError.value =
        firebaseError.message || 'خطای پیش‌بینی‌نشده‌ای رخ داد.'
  }
}

function getGeneralErrorMessage(
  code: string,
  fallback: string | undefined,
): string {
  switch (code) {
    case 'auth/user-not-found':
      return 'حسابی با این ایمیل پیدا نشد.'
    case 'auth/wrong-password':
    case 'auth/invalid-credential':
      return 'اطلاعات ورود معتبر نیست.'
    case 'auth/user-disabled':
      return 'این حساب غیرفعال شده است.'
    case 'auth/too-many-requests':
      return 'تلاش‌های ناموفق زیاد بود؛ کمی بعد دوباره امتحان کنید.'
    case 'auth/network-request-failed':
      return 'ارتباط با سرور برقرار نشد؛ اتصال اینترنت را بررسی کنید.'
    default:
      return fallback || 'خطای پیش‌بینی‌نشده‌ای رخ داد.'
  }
}
</script>

<template>
  <div>
    <v-btn
      block
      color="white"
      size="large"
      class="mb-3 position-relative"
      :loading="loading"
      :disabled="loading"
      @click="signInWithGoogle"
    >
      <v-chip
        v-if="lastUsedMethod === 'google'"
        size="x-small"
        color="primary"
        label
        variant="flat"
        class="position-absolute last-used-chip"
      >
        آخرین روش
      </v-chip>
      <v-icon color="red" :icon="mdiGoogle" class="mr-2" />
      ورود با گوگل
    </v-btn>

    <v-btn
      block
      size="large"
      variant="flat"
      color="black"
      class="mb-3 position-relative"
      :loading="loading"
      :disabled="loading"
      @click="signInWithGithub"
    >
      <v-chip
        v-if="lastUsedMethod === 'github'"
        label
        size="x-small"
        color="primary"
        variant="flat"
        class="position-absolute last-used-chip"
      >
        آخرین روش
      </v-chip>
      <v-icon :icon="mdiGithub" class="mr-2" />
      ورود با گیت‌هاب
    </v-btn>

    <v-btn
      v-if="!showEmailForm"
      block
      size="large"
      variant="flat"
      color="red"
      class="mb-3 position-relative"
      :disabled="loading"
      @click="showEmailForm = true"
    >
      <v-chip
        v-if="lastUsedMethod === 'email'"
        label
        size="x-small"
        color="primary"
        variant="flat"
        class="position-absolute last-used-chip"
      >
        آخرین روش
      </v-chip>
      <v-icon :icon="mdiEmail" class="mr-2" />
      ورود با ایمیل
    </v-btn>

    <!-- Forgot Password Form -->
    <v-form
      v-if="showEmailForm && showForgotPassword"
      class="mt-4"
      @submit.prevent="submitPasswordReset"
    >
      <template v-if="!resetEmailSent">
        <p class="text-body-medium text-medium-emphasis mb-4">
          ایمیل حساب را برای بازیابی رمز عبور وارد کنید.
        </p>
        <v-text-field
          v-model="email"
          label="ایمیل"
          color="primary"
          type="email"
          variant="outlined"
          density="comfortable"
          class="mb-2"
          :error="errorMessages.has('email')"
          :error-messages="errorMessages.get('email')"
        />
        <v-alert
          v-if="generalError"
          type="error"
          density="compact"
          class="mb-3"
        >
          {{ generalError }}
        </v-alert>
        <v-btn
          block
          size="large"
          color="primary"
          type="submit"
          :loading="loading"
        >
          ارسال لینک بازیابی
        </v-btn>
      </template>
      <template v-else>
        <v-alert type="success" density="compact" class="mb-3">
          راهنمای بازیابی رمز عبور به ایمیل شما ارسال شد.
        </v-alert>
      </template>
      <v-btn
        block
        variant="text"
        size="small"
        color="warning"
        class="mt-2"
        @click="backToSignIn"
      >
        بازگشت به ورود
      </v-btn>
    </v-form>

    <!-- Sign In / Sign Up Form -->
    <v-form
      v-if="showEmailForm && !showForgotPassword"
      class="mt-4"
      @submit.prevent="submitEmail"
    >
      <v-text-field
        v-model="email"
        label="ایمیل"
        color="primary"
        type="email"
        variant="outlined"
        density="comfortable"
        class="mb-2"
        :error="errorMessages.has('email')"
        :error-messages="errorMessages.get('email')"
      />
      <v-text-field
        v-model="password"
        label="رمز عبور"
        :type="showPassword ? 'text' : 'password'"
        color="primary"
        variant="outlined"
        density="comfortable"
        class="mb-2"
        :error="errorMessages.has('password')"
        :error-messages="errorMessages.get('password')"
        :append-inner-icon="showPassword ? mdiEyeOff : mdiEye"
        @click:append-inner="showPassword = !showPassword"
      />
      <v-alert v-if="generalError" type="error" density="compact" class="mb-3">
        {{ generalError }}
      </v-alert>
      <v-btn
        v-if="!isSignUp"
        variant="plain"
        size="small"
        color="primary"
        class="mb-3 px-0 mt-n4"
        @click="showForgotPasswordForm"
      >
        رمز عبور را فراموش کرده‌اید؟
      </v-btn>
      <v-btn
        block
        size="large"
        color="primary"
        type="submit"
        :loading="loading"
      >
        {{ isSignUp ? 'ساخت حساب' : 'ورود' }}
      </v-btn>
      <v-btn
        block
        variant="plain"
        size="small"
        color="primary"
        class="mt-2"
        @click="isSignUp = !isSignUp"
      >
        {{ isSignUp ? 'حساب دارید؟ وارد شوید' : 'حساب ندارید؟ ثبت‌نام کنید' }}
      </v-btn>
    </v-form>

    <p class="text-body-small text-medium-emphasis mt-4">
      با ادامه، شرایط استفاده و حریم خصوصی را می‌پذیرید.
      <a
        :href="appStore.appData.url + '/terms-and-conditions'"
        class="text-decoration-none"
      >
        شرایط استفاده
      </a>
      و
      <a
        :href="appStore.appData.url + '/privacy-policy'"
        class="text-decoration-none"
      >
        حریم خصوصی.</a
      >
    </p>
  </div>
</template>

<style scoped>
.last-used-chip {
  top: -8px;
  left: -8px;
  z-index: 1;
}
</style>
