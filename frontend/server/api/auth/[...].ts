// file: ~/server/api/auth/[...].ts
import GoogleProvider from "next-auth/providers/google";
import { NuxtAuthHandler } from "#auth";
import { ExtendedJWT, ProviderInfo } from "~/types/auth";

const config = useRuntimeConfig();

async function refreshAccessToken(token: ExtendedJWT) {
  try {
    const url =
      "https://oauth2.googleapis.com/token?" +
      new URLSearchParams({
        client_id: config.googleClientId,
        client_secret: config.googleClientSecret,
        grant_type: "refresh_token",
        refresh_token: token.providerInfo.refresh_token ?? "",
      });

    const response = await fetch(url, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      method: "POST",
    });

    const refreshedTokens = await response.json();

    if (!response.ok) {
      throw refreshedTokens;
    }

    token.providerInfo = {
      ...token.providerInfo,
      access_token: refreshedTokens.access_token,
      expire_at: Date.now() + refreshedTokens.expires_in * 1000,
      refresh_token: refreshedTokens.refresh_token,
    };

    return {
      token,
    };
  } catch (error) {
    console.error(error);
    token.providerInfo.refresh_error = "Failed to refresh token";

    return {
      token,
    };
  }
}

export default NuxtAuthHandler({
  providers: [
    // @ts-expect-error You need to use .default here for it to work during SSR. May be fixed via Vite at some point
    GoogleProvider.default({
      clientId: config.googleClientId,
      clientSecret: config.googleClientSecret,
      authorization: {
        params: {
          access_type: "offline",
          prompt: "consent",
        },
      },
    }),
  ],
  session: {
    strategy: "jwt",
  },
  callbacks: {
    jwt: async ({ token, user, account, profile }) => {
      if (account && user) {
        const providerInfo: ProviderInfo = {
          provider: account.provider,
          access_token: account.access_token,
          expire_at: account.expires_at,
          refresh_token: account.refresh_token,
          refresh_error: undefined,
          ...user,
        };
        token.providerInfo = providerInfo;
      }
      if (token.providerInfo) {
        // 現在のUnix時間を取得
        const unixTime = Math.floor(Date.now() / 1000);
        if (
          token.providerInfo.expire_at &&
          unixTime < token.providerInfo.expire_at
        ) {
          return token;
        }
      }

      const result = await refreshAccessToken(token as ExtendedJWT);
      return result.token;
    },
    session: async ({ session, token, user }) => {
      session.providerInfo = token?.providerInfo;

      return session;
    },
  },
});
