/**
 * This composable is used to perform custom fetch operations on the client-side only.
 * It uses the `useFetch` composable to make HTTP requests with authentication headers.
 */

interface FetchOptions {
  [key: string]: any;
}

export const useCustomFetch = async <T>(
  path: string,
  params: FetchOptions = {}
) => {
  const auth = useAuth();
  const token = (auth.data.value as any).providerInfo.access_token;
  const config = useRuntimeConfig();
  const url = config.public.apiBaseUrl + path;

  const result = await $fetch<T>(url, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
    ...params,
  });

  return result;
};
