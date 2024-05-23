<template>
  <div class="prompt-select">
    <div class="custom-select">
      <div class="selected-option" @click="toggleDropdown">
        {{ selectedOption.name }}<span class="dropdown-arrow">▼</span>
      </div>
      <div v-show="dropdownOpen" class="options">
        <div
          v-for="option in options"
          :key="option.id"
          :class="['option', { selected: option === selectedOption }]"
          @click="selectOption(option)"
        >
          {{ option.name }}
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
interface Options {
  id: number;
  name: string;
}
interface Props {
  options: Options[];
  selectedOptionId: number;
}

interface Emits {
  (e: "selectOption", optionId: number): void;
}

const props = withDefaults(defineProps<Props>(), {
  options: () => [
    {
      id: 1,
      name: "Option 1",
    },
    {
      id: 2,
      name: "Option 2",
    },
    {
      id: 3,
      name: "Option 3",
    },
  ],
  selectedOptionId: 2,
});

const emits = defineEmits<Emits>();

const selectedOption = ref<Options>({
  id: props.selectedOptionId,
  name:
    props.options.find((option) => option.id === props.selectedOptionId)
      ?.name ?? "",
});
const dropdownOpen = ref(false);

const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value;
};

const selectOption = (option: Options) => {
  selectedOption.value = option;
  dropdownOpen.value = false;
  emits("selectOption", option.id);
};
</script>

<style lang="scss" scoped>
.prompt-select {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;

  .custom-select {
    position: relative;
    width: 100%;
    max-width: 520px;
    border: 1px solid #000;
    border-radius: 0.375rem;
    background-color: #fff;
    cursor: pointer;
    font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
      Liberation Mono, Courier New, monospace;
    font-size: 0.875rem;
    line-height: 1.25rem;
    font-weight: 500;
    padding: 0.625rem 1.25rem;
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000),
      var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);

    .selected-option {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .dropdown-arrow {
      margin-left: auto;
    }

    .options {
      position: absolute;
      top: 100%;
      left: 0;
      right: 0;
      background-color: #fff;
      border: 1px solid #000;
      border-radius: 0.375rem;
      box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000),
        var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
      z-index: 10;

      .option {
        padding: 0.625rem 1.25rem;
        cursor: pointer;

        &.selected {
          background-color: #f0f0f0;
        }

        &:hover {
          background-color: #e0e0e0;
        }
      }
    }
  }
}
</style>
