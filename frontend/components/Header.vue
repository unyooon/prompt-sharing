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
      <LinksFillButton
        v-if="isSignIn"
        text="Crate Prompt"
        :is-show-arrow="true"
        to="/prompts/create"
      />
      <LinksFillButton
        v-if="!isSignIn"
        text="Join Now"
        :is-show-arrow="true"
        to="/signin"
      />
      <div
        v-else
        class="header__account-icon"
        ref="accountIcon"
        @click="isShowAccountModal = !isShowAccountModal"
      >
        <img :src="userImg" alt="" />
        <div
          v-show="isShowAccountModal"
          ref="accountModal"
          class="header__account-modal"
        >
          <!-- <div class="header__account-modal-item">Profile</div> -->
          <div class="header__account-modal-item" @click="emits('signOut')">
            Sign Out
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
const router = useRouter();

interface Props {
  isSignIn: boolean;
  userImg: string;
}

interface Emits {
  (e: "signOut"): void;
}

const props = withDefaults(defineProps<Props>(), {
  isSignIn: false,
  userImg: "",
});

const emits = defineEmits<Emits>();

const isShowAccountModal = ref(false);
const accountIcon = ref<HTMLElement | null>(null);
const accountModal = ref<HTMLElement | null>(null);

const checkActive = (path: string) => {
  return router.currentRoute.value.path === path;
};

const handleOutsideClick = (event: MouseEvent) => {
  if (
    accountModal.value &&
    !accountModal.value?.contains(event.target as Node) &&
    !accountIcon.value?.contains(event.target as Node)
  ) {
    isShowAccountModal.value = false;
  }
};

onMounted(() => {
  document.addEventListener("click", handleOutsideClick);
});

onUnmounted(() => {
  document.removeEventListener("click", handleOutsideClick);
});
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

  &__account-icon {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;

    img {
      width: 2.5rem;
      height: 2.5rem;
      border-radius: 50%;
      overflow: hidden;
    }
  }

  &__account-modal {
    position: absolute;
    top: 3rem;
    left: 0;

    flex-direction: column;
    padding-top: 0.75rem;
    padding-bottom: 0.75rem;
    border: 1px solid #ccc;
    border-radius: 0.25rem;
    background-color: #fff;

    &-item {
      cursor: pointer;
      padding-top: 0.5rem;
      padding-bottom: 0.5rem;
      padding-right: 1.25rem;
      padding-left: 1.25rem;
      white-space: nowrap;
    }

    &-item:hover {
      background-color: #f5f5f5;
    }
  }
}
</style>
