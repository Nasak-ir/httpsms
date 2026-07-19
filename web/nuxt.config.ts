// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-01-01',

  ssr: false,

  modules: [
    '@nuxt/eslint',
    'vuetify-nuxt-module',
    '@pinia/nuxt',
    '@nuxtjs/seo',
  ],

  css: ['vuetify/styles', '~/assets/main.css'],

  site: {
    url: process.env.APP_URL || 'https://sms.nasak.ir',
    name: process.env.APP_NAME || 'پیامک نسک',
    description:
      'درگاه پیامک اختصاصی نسک؛ ارسال و دریافت پیامک با سیم‌کارت خودتان و بدون اعتبار نرم‌افزاری.',
    defaultLocale: 'fa',
  },

  // Authenticated app routes that should never appear in search engines or the
  // sitemap. The marketing/blog/legal pages remain public and indexable.
  robots: {
    disallow: [
      '/messages',
      '/threads',
      '/settings',
      '/billing',
      '/bulk-messages',
      '/heartbeats',
      '/phone-api-keys',
      '/search-messages',
    ],
  },

  sitemap: {
    exclude: [
      '/messages',
      '/threads',
      '/threads/**',
      '/settings',
      '/billing',
      '/bulk-messages',
      '/heartbeats/**',
      '/phone-api-keys',
      '/search-messages',
      '/login',
    ],
  },

  // The app ships as a client-rendered SPA (ssr: false) served statically, so
  // runtime Satori OG-image generation is not available. Curated static OG
  // images are set per page instead.
  ogImage: {
    enabled: false,
  },

  // Static download assets in /public are valid at runtime but are not Nuxt
  // routes, so exclude them from link validation to avoid false positives.
  linkChecker: {
    excludeLinks: ['/templates/**'],
  },

  build: {
    transpile: ['vuetify', 'chart.js', 'vue-chartjs', 'v-phone-input'],
  },

  vite: {
    define: {
      'process.env.DEBUG': false,
    },
    optimizeDeps: {
      include: [
        '@mdi/js',
        'chartjs-adapter-moment',
        'date-fns',
        'firebase/app',
        'firebase/auth',
        'highlight.js/lib/core',
        'libphonenumber-js',
        'pusher-js',
        'qrcode',
      ],
    },
  },

  vuetify: {
    vuetifyOptions: {
      theme: {
        defaultTheme: 'nasakDark',
        themes: {
          nasakDark: {
            dark: true,
            colors: {
              background: '#0b121a',
              surface: '#101923',
              primary: '#31b77a',
              secondary: '#78a9d4',
              error: '#ef6461',
              warning: '#e3aa4e',
              success: '#31b77a',
              info: '#58a6d8',
            },
          },
          nasakLight: {
            dark: false,
            colors: {
              background: '#fbf8f0',
              surface: '#fffdf7',
              primary: '#167b55',
              secondary: '#355b7d',
              error: '#b42318',
              warning: '#956000',
              success: '#167b55',
              info: '#176b9b',
            },
          },
        },
      },
      icons: {
        defaultSet: 'mdi-svg',
      },
    },
  },

  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.API_BASE_URL || 'http://localhost:8000',
      clientVersion: process.env.GITHUB_SHA || 'dev',
      appUrl: process.env.APP_URL || 'https://sms.nasak.ir',
      appName: process.env.APP_NAME || 'پیامک نسک',
      appGithubUrl:
        process.env.APP_GITHUB_URL || 'https://github.com/Nasak-ir/httpsms',
      appDocumentationUrl:
        process.env.APP_DOCUMENTATION_URL ||
        'https://github.com/Nasak-ir/httpsms/blob/main/docs/NASAK_SMS.md',
      appDownloadUrl:
        process.env.APP_DOWNLOAD_URL ||
        'https://github.com/Nasak-ir/httpsms/releases/latest/download/NasakSms.apk',
      appEnv: process.env.APP_ENV || 'production',
      checkoutUrl: process.env.CHECKOUT_URL || '',
      enterpriseCheckoutUrl: process.env.ENTERPRISE_CHECKOUT_URL || '',
      cloudflareTurnstileSiteKey:
        process.env.CLOUDFLARE_TURNSTILE_SITE_KEY || '',
      pusherKey: process.env.PUSHER_KEY || '',
      pusherCluster: process.env.PUSHER_CLUSTER || '',
      firebaseApiKey: process.env.FIREBASE_API_KEY || '',
      firebaseAuthDomain: process.env.FIREBASE_AUTH_DOMAIN || '',
      firebaseProjectId: process.env.FIREBASE_PROJECT_ID || '',
      firebaseStorageBucket: process.env.FIREBASE_STORAGE_BUCKET || '',
      firebaseMessagingSenderId: process.env.FIREBASE_MESSAGING_SENDER_ID || '',
      firebaseAppId: process.env.FIREBASE_APP_ID || '',
      firebaseMeasurementId: process.env.FIREBASE_MEASUREMENT_ID || '',
    },
  },

  nitro: {
    prerender: {
      routes: [],
      failOnError: false,
    },
  },

  routeRules: {
    '/messages': { robots: false },
    '/threads': { robots: false },
    '/threads/**': { robots: false },
    '/settings': { robots: false },
    '/billing': { robots: false },
    '/bulk-messages': { robots: false },
    '/heartbeats/**': { robots: false },
    '/phone-api-keys': { robots: false },
    '/search-messages': { robots: false },
  },

  app: {
    head: {
      titleTemplate: '%s',
      title: 'پیامک نسک | درگاه پیامک اختصاصی',
      htmlAttrs: { lang: 'fa', dir: 'rtl' },
      script: [],
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        {
          name: 'description',
          content:
            'ارسال و دریافت پیامک با گوشی اندرویدی و سیم‌کارت خودتان، بدون خرید اعتبار نرم‌افزاری.',
        },
        { name: 'format-detection', content: 'telephone=no' },
        { name: 'twitter:card', content: 'summary_large_image' },
        {
          property: 'og:title',
          content: 'پیامک نسک | درگاه پیامک اختصاصی',
        },
        {
          property: 'og:description',
          content:
            'ارسال و دریافت پیامک با گوشی اندرویدی و سیم‌کارت خودتان، بدون اعتبار نرم‌افزاری.',
        },
        {
          property: 'og:image',
          content: 'https://sms.nasak.ir/brand/nasak-logo.png',
        },
      ],
      link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
    },
  },
})
