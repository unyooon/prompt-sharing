<template>
  <div class="home">
    <div class="home__search-container">
      <div class="home__head-text">Search the prompts you're looking for</div>
      <div class="home__create-prompt-button-container"></div>
      <div class="home__search-box">
        <SearchesPromptSearch
          placeholder="Search for tag, prompt, or username"
          :value="query"
          @change="(v) => (query = v)"
        />
      </div>
    </div>
    <div class="home__prompt-container">
      <template v-for="prompt in prompts">
        <CardsPromptCard
          :title="prompt.description"
          :prompt="prompt.text"
          :tags="[prompt.llmName]"
        />
      </template>
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { ReadPromptResponse } from "~/types/api";

const prompts = ref<ReadPromptResponse[]>([]);
const isLoading = ref(true);
const query = ref("");

const fetchPrompts = async () => {
  try {
    const api = useApi();
    const response = await api.prompts.get(1, 30);
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
.home {
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  &__create-prompt-button-container {
    width: 100%;
    display: flex;
    justify-content: center;
    margin-top: 2rem;
    margin-bottom: 2rem;
  }

  &__create-prompt-button {
    max-width: 240px;
  }

  &___search-container {
    width: 100%;
    max-width: 80rem;
    padding-left: 4rem;
    padding-right: 4rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }

  &__head-text {
    font-size: 2rem;
    font-weight: 700;
    margin-top: 2rem;
    margin-bottom: 2rem;
    line-height: 1;

    font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
      Liberation Mono, Courier New, monospace;
  }

  &__search-box {
    width: 100%;
    min-width: 420px;
    display: flex;
  }

  &__prompt-container {
    padding-top: 2rem;
    padding-bottom: 2rem;
    margin-top: 1.25rem;

    gap: 1.5rem;
    column-count: 2;
  }
}
</style>
