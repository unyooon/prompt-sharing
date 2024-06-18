<template>
  <div class="features">
    <template v-for="feature in features">
      <div class="features__card">
        <div>
          <div class="features__card-title">{{ feature.title }}</div>
          <div class="features__card-description">
            {{ feature.description }}
          </div>
        </div>
        <div
          v-if="!feature.release"
          class="features__card-tips features__card-tips--upcomming"
        >
          upcoming...
        </div>
        <div v-else class="features__card-tips features__card-tips--release">
          release
        </div>
      </div>
    </template>
  </div>
</template>

<script lang="ts" setup>
interface Features {
  title: string;
  description: string;
  release: boolean;
}

const features: Features[] = [
  {
    title: "Create Prompt",
    description: "Allows you to create prompts",
    release: true,
  },
  {
    title: "Delete Prompt",
    description: "Allows you to delete prompts you have created",
    release: false,
  },
];
</script>

<style lang="scss" scoped>
.features {
  display: grid;
  gap: 1.5rem;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
    Liberation Mono, Courier New, monospace;

  &__card {
    min-width: 200px;
    margin: 0 auto;
    padding: 2rem;
    border: 1px solid #ccc;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    justify-content: space-between;
  }

  &__card-title {
    font-size: 1.125rem;
    font-weight: bold;
    margin-bottom: 1rem;
  }

  &__card-description {
    line-height: 1.5rem;
  }

  &__card-tips {
    width: fit-content;

    margin-top: 0.75rem;
    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
    padding-right: 0.5rem;
    padding-left: 0.5rem;
    border-radius: 0.5rem;

    font-size: 0.75rem;
    font-weight: normal;

    &--upcomming {
      background-color: gray;
      color: #fff;
    }

    &--release {
      background-color: green;
      color: #fff;
    }
  }
}
</style>
