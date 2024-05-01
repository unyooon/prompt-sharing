<template>
  <div class="header">
    <div class="header__logo">Service Name</div>
    <div class="header__buttons">
      <LinksHeaderButton text="Home" to="/" :active="checkActive('/')" />
      <LinksHeaderButton
        text="Features"
        to="/features"
        :active="checkActive('/features')"
      />
      <LinksOutlineButton text="Feedback" to="/" />
      <ButtonsFillButton text="Join Now" :is-show-array="true" />
    </div>
  </div>
</template>

<script lang="ts" setup>
const router = useRouter();

const checkActive = (path: string) => {
  return router.currentRoute.value.path === path;
};
</script>

<style lang="scss" scoped>
.header {
  width: 100%;

  padding: 10px 20px;
  padding-top: 0.75rem;
  margin-top: 0.5rem;
  margin-bottom: 4rem;

  display: flex;
  justify-content: space-between;
  align-items: center;

  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
    Liberation Mono, Courier New, monospace;

  &__logo {
  }

  &__buttons {
    display: flex;
    gap: 1.25rem;
  }
}
</style>
