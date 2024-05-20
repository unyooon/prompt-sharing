// file: ~/server/api/auth/[...].ts
import GoogleProvider from "next-auth/providers/google";
import { NuxtAuthHandler } from "#auth";

const config = useRuntimeConfig();

export default NuxtAuthHandler({
  providers: [
    // @ts-expect-error You need to use .default here for it to work during SSR. May be fixed via Vite at some point
    GoogleProvider.default({
      clientId: config.googleClientId,
      clientSecret: config.googleClientSecret,
    }),
  ],
  session: {
    strategy: "jwt",
  },
  callbacks: {
    jwt: async ({ token, user, account, profile }) => {
      if (user?.email) {
        token.providerInfo = {
          provider: account?.provider,
          access_token: account?.access_token,
          expire_at: account?.expires_at,
          refresh_token: account?.refresh_token,
          ...user,
        };
      }

      return Promise.resolve(token);
    },
    session: async ({ session, token, user }) => {
      (session as any).providerInfo = token?.providerInfo;

      return Promise.resolve({ ...session });
    },
  },
});
