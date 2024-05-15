export default defineNuxtRouteMiddleware(async (to, from) => {
  const { data } = useAuth();

  if (!data.value) {
    return;
  }

  const accessToken = window.localStorage.getItem("accessToken");
  const refreshToken = window.localStorage.getItem("refreshToken");
  const expiresAt = window.localStorage.getItem("expiresAt");

  if (!accessToken || !refreshToken || !expiresAt) {
    window.localStorage.setItem(
      "accessToken",
      (data.value as any).providerInfo.access_token
    );
    window.localStorage.setItem(
      "refreshToken",
      (data.value as any).providerInfo.refresh_token
    );
    window.localStorage.setItem(
      "expiresAt",
      (data.value as any).providerInfo.expire_at
    );
  }
});
