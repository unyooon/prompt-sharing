<template>
  <div class="fill-button">
    {{ props.text }}
    <svg
      v-if="isShowArray"
      stroke="currentColor"
      fill="currentColor"
      stroke-width="0"
      viewBox="0 0 20 20"
      aria-hidden="true"
      class="svg"
      height="1em"
      width="1em"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        fill-rule="evenodd"
        d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z"
        clip-rule="evenodd"
      ></path>
    </svg>
  </div>
</template>

<script lang="ts" setup>
interface Props {
  text: string;
  isShowArray: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  text: "Button",
  isShowArray: false,
});
</script>

<style lang="scss" scoped>
a {
  text-decoration: none;
}

.fill-button {
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid #000;
  border-radius: 0.25rem;
  background-color: #000;
  padding: 0.375rem 1.25rem;
  text-align: center;
  font-size: 0.875rem;
  line-height: 1.25rem;
  color: #fff;
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 0.15s;

  &:hover {
    background-color: #fff;
    color: #000;
  }
}

.svg {
  margin-left: 0.5rem;
}
</style>
