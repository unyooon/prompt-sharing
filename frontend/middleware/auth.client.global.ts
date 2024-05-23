// file: ~/middleware/checkAuth.client.global.ts
export default defineNuxtRouteMiddleware(async (to, from) => {
  const { getSession, signOut } = useAuth();

  if (import.meta.server) return;
  if (to.path === "/signin") return;

  const session = await getSession();

  if (!session.providerInfo) return;

  const providerInfo = session.providerInfo;

  if (providerInfo && providerInfo.refresh_error) {
    signOut();

    // Optionally redirect to login or an error page
    return navigateTo("/signin");
  }
});
