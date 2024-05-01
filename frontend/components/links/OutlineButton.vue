<template>
  <NuxtLink class="outline-button" :to="to">
    {{ props.text }}
  </NuxtLink>
</template>

<script lang="ts" setup>
interface Props {
  text: string;
  to: string;
}

const props = withDefaults(defineProps<Props>(), {
  text: "Button",
  to: "/",
});
</script>

<style lang="scss" scoped>
a {
  text-decoration: none;
}

.outline-button {
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px #000 solid;
  border-radius: 0.25rem;
  background-color: rgb(239 246 255);
  padding: 0.375rem 1.25rem;
  font-size: 0.875rem;
  text-align: center;
  line-height: 1.25rem;
  color: #000;
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 0.15s;

  &:hover {
    background-color: #000;
    color: #fff;
  }
}
</style>
