<template>
  <div class="wrapper">
    <Header />
    <div class="main">
      <slot />
    </div>
  </div>
</template>

<script lang="ts" setup></script>

<style lang="scss" scoped>
.wrapper {
  min-height: 100vh;

  max-width: 80rem;
  margin-left: auto;
  margin-right: auto;
  padding-left: 4rem;
  padding-right: 4rem;

  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.main {
  flex-grow: 1;
}
</style>
