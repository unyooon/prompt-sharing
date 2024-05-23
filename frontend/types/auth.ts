import type { JWT } from "next-auth/jwt";

export interface ExtendedJWT extends JWT {
  providerInfo: ProviderInfo;
}

export interface ProviderInfo {
  provider: string | undefined;
  access_token: string | undefined;
  expire_at: number | undefined;
  refresh_token: string | undefined;
  refresh_error: string | undefined;
}
