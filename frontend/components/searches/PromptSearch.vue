<template>
  <div class="prompt-search">
    <input :placeholder="placeholder" type="text" />
  </div>
</template>

<script lang="ts" setup>
interface Props {
  placeholder: string;
}

const props = withDefaults(defineProps<Props>(), {
  placeholder: "Search for tag, prompt, or username",
});
</script>

<style lang="scss" scoped>
.prompt-search {
  width: 100%;
  display: flex;
  justify-content: center;

  input {
    display: block;
    width: 100%;
    border-radius: 0.375rem;
    border: 1px solid #000;
    background-color: #fff;
    padding: 0.625rem 3rem 0.625rem 1.25rem;
    font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
      Liberation Mono, Courier New, monospace;
    font-size: 0.875rem;
    line-height: 1.25rem;
    font-weight: 500;
    --tw-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1),
      0 4px 6px -4px rgba(0, 0, 0, 0.1);
    --tw-shadow-colored: 0 10px 15px -3px var(--tw-shadow-color),
      0 4px 6px -4px var(--tw-shadow-color);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000),
      var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
    max-width: 520px;
  }

  input:focus {
    outline: none;
  }
}
</style>
