<template>
  <div class="prompt-card">
    <div class="prompt-card__title">
      {{ title }}
    </div>
    <div class="prompt-card__prompt">
      <p class="prompt-card__prompt-text">
        {{ prompt }}
      </p>
    </div>
    <div v-if="tags.length > 0" class="prompt-card__tags">
      <div v-for="tag in tags" class="prompt-card__tag">
        {{ tag }}
      </div>
    </div>
    <div class="prompt-card__menu">
      <div class="prompt-card__menu-btn" @click="copy">
        <img :src="CopyIcon" alt="" />
      </div>
      <!-- <div v-show="showDelete" class="prompt-card__menu-btn" @click="copy">
        <img :src="DeleteIcon" alt="" />
      </div> -->
    </div>
  </div>
</template>

<script lang="ts" setup>
import CopyIcon from "~/assets/images/copy.svg";
import DeleteIcon from "~/assets/images/delete.svg";

interface Props {
  title: string;
  prompt: string;
  tags: string[];
  showDelete?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  title: "Title",
  prompt: "",
  tags: () => [],
  showDelete: false,
});

const copy = () => {
  navigator.clipboard.writeText(props.prompt);
};
</script>

<style lang="scss" scoped>
.prompt-card {
  height: fit-content;
  width: 360px;
  flex: 1 1 0%;
  break-inside: avoid;
  border: 1px solid #000;
  margin-bottom: 1.5rem;

  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
    Liberation Mono, Courier New, monospace;

  &__title {
    border-bottom: solid 1px #000;
    padding: 1rem;
    font-size: inherit;
    font-weight: bold;
  }

  &__prompt {
    padding: 1rem;
    font-size: 0.875rem;
    border-bottom: solid 1px #000;
    overflow-wrap: break-word;
    line-height: 1.5rem;

    &-text {
      margin-top: 1rem;
      margin-bottom: 1rem;

      max-height: 290px;
      overflow: hidden;
    }
  }

  &__tags {
    padding: 1rem;
    border-bottom: solid 1px #000;
  }

  &__tag {
    font-size: 0.75rem;
    line-height: 1rem;

    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
    padding-left: 0.5rem;
    padding-right: 0.5rem;

    border: solid 1px #000;
    border-radius: 0.25rem;
    background-color: rgb(239 246 255);

    width: fit-content;
  }

  &__menu {
    width: 100%;
    padding: 0.75rem;
    width: 1.75rem;
    height: 1.75rem;

    display: flex;
    align-items: center;
  }

  &__menu-btn {
    cursor: pointer;
    padding: 0.25rem;
    margin-right: 0.5rem;

    img {
      width: 0.75rem;
      height: 0.75rem;
      user-select: none;
    }
  }
}
</style>
