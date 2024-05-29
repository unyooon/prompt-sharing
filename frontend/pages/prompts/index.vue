<template>
  <div class="prompts">
    <div class="prompts__sidebar">
      <SidebarsPromptsMenu />
    </div>
    <div class="prompts__wrapper">
      <div class="prompts__main">
        <template v-for="prompt in prompts">
          <CardsPromptCard
            :title="prompt.description"
            :prompt="prompt.text"
            :tags="[prompt.llmName]"
            :show-delete="true"
          />
        </template>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { ReadPromptResponse } from "~/types/api";

const prompts = ref<ReadPromptResponse[]>([]);
const isLoading = ref(true);

const fetchPrompts = async () => {
  try {
    const api = useApi();
    const response = await api.prompts.get({ page: 1, size: 30, isMine: true });
    if (response) {
      prompts.value = response.data;
    }
  } catch (err) {
    console.error("Error fetching prompts:", err);
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchPrompts();
});
</script>

<style lang="scss" scoped>
.prompts {
  display: flex;

  &__wrapper {
    display: flex;
    justify-content: center;
    width: 100%;
  }

  &__main {
    width: 80%;
    padding-top: 2rem;
    padding-bottom: 2rem;
    margin-top: 1.25rem;

    gap: 1.5rem;
    column-count: 2;

    @media (max-width: 1280px) {
      display: flex;
      flex-direction: column;
      align-items: center;
      column-count: 1;
    }
  }
}
</style>
