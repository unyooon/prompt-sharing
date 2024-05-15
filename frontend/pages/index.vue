<template>
  <div class="home">
    <div class="home__search-container">
      <div class="home__head-text">Search the prompts you're looking for</div>
      <div class="home__search-box">
        <SearchesPromptSearch />
      </div>
    </div>
    <div class="home__prompt-container">
      <!-- <CardsPromptCard />
      <CardsPromptCard />
      <CardsPromptCard />
      <CardsPromptCard />
      <CardsPromptCard />
      <CardsPromptCard />
      <CardsPromptCard /> -->
    </div>
  </div>
</template>

<script lang="ts" setup>
const { data } = useAuth();

console.log(data.value?.providerInfo);
</script>

<style lang="scss" scoped>
.home {
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

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
