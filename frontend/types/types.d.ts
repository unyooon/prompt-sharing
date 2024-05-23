import { DefaultSession, Session } from "next-auth";

interface UserWithId extends DefaultSession["user"] {
  id?: string;
}

declare module "next-auth/jwt" {
  /** Returned by the `jwt` callback and `getToken`, when using JWT sessions */
  interface JWT {
    user: UserWithId;
    providerInfo: ProviderInfo;
  }
}

declare module "next-auth" {
  interface Session extends DefaultSession {
    providerInfo?: ProviderInfo;
  }
}
