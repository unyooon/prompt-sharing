// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,

  app: {
    head: {
      title: "PromptDeck",
      htmlAttrs: {
        lang: "ja",
      },
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          hid: "description",
          name: "description",
          content: "Prompt",
        },
        {
          name: "google-signin-client_id",
          content:
            "1041190716338-21ap5uscs20nmb4e8ag1me4svjnfcjg2.apps.googleusercontent.com",
        },
      ],
      link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
      script: [],
    },
  },

  modules: ["@sidebase/nuxt-auth"],
  auth: {
    provider: {
      type: "authjs",
    },
    baseURL: process.env.NUXT_AUTH_ORIGIN,
  },

  css: ["~/assets/scss/main.scss"],
  vite: {
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: '@use "~/assets/scss/_variables.scss" as *;',
        },
      },
    },
  },
  runtimeConfig: {
    googleClientId: "",
    googleClientSecret: "",
    authOrigin: "http://localhost:3000",
    authSecret: "",
    public: {
      apiBaseUrl: "http://localhost:3000",
    },
  },
});
