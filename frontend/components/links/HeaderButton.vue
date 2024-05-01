<template>
  <NuxtLink class="header-button" :class="{ active: props.active }" :to="to">
    {{ props.text }}
  </NuxtLink>
</template>

<script lang="ts" setup>
interface Props {
  text: string;
  active: boolean;
  to: string;
}

const props = withDefaults(defineProps<Props>(), {
  text: "Button",
  active: false,
  to: "/",
});
</script>

<style lang="scss" scoped>
a {
  text-decoration: none;
}

.header-button {
  cursor: pointer;
  text-align: center;
  font-size: 16px;
  font-weight: 500;
  padding: 10px 20px;
  border: none;
  border-bottom: 2px solid transparent;
  transition: border-color 0.3s ease;
  color: #555;

  &.active {
    color: #000;
  }

  &.active,
  &:hover {
    border-color: #000;
  }
}
</style>
