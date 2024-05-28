<template>
  <div class="profile-item">
    <div class="profile-item__bar" :class="{ active: isActive }"></div>
    <div
      class="profile-item__label"
      :class="{ active: isActive }"
      @click="emit('click')"
    >
      <div class="profile-item__label-icon"></div>
      <div class="profile-item__label-text">
        {{ text }}
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
interface Props {
  icon: string;
  text: string;
  isActive: boolean;
}

interface Emits {
  (e: "click"): void;
}

const props = withDefaults(defineProps<Props>(), {
  icon: "",
  text: "",
  isActive: false,
});

const emit = defineEmits<Emits>();
</script>

<style lang="scss" scoped>
.profile-item {
  display: flex;
  height: 32px;

  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
    Liberation Mono, Courier New, monospace;

  &__bar {
    height: 100%;
    width: 4px;
    border-radius: 1rem;
    margin-right: 0.25rem;

    &.active {
      background-color: #0366d6;
    }
  }

  &__label {
    cursor: pointer;
    display: flex;
    width: 100%;
    align-items: center;
    padding-right: 0.75rem;
    padding-left: 0.75rem;
    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
    border-radius: 0.25rem;

    &-icon {
      width: 20px;
      height: 20px;
      background-size: cover;
      margin-right: 10px;
    }

    &-text {
      font-size: 14px;
      color: #24292e;
    }

    &.active {
      background-color: #f6f8fa;
      font-weight: bold;
    }
  }
}
</style>
