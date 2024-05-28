<template>
  <div class="prompts">
    <div class="prompts__sidebar">
      <SidebarsPromptsMenu />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.prompts {
  display: flex;

  &__sidebar {
    width: 20%;
    padding: 1rem;
    border-right: 1px solid #ccc;
  }
}
</style>
