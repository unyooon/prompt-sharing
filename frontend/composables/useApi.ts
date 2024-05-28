import { promptsApi } from "~/composables/api/prompts";
import { llmMastersApi } from "~/composables/api/llmMasters";

export const useApi = () => {
  const prompts = promptsApi();
  const llmMasters = llmMastersApi();

  return {
    prompts,
    llmMasters,
  };
};
