import { promptsApi } from "~/composables/api/prompts";

export const useApi = () => {
  const prompts = promptsApi();

  return {
    prompts,
  };
};
