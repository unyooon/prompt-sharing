<template>
  <div class="signin">
    <div class="signin__container">
      <div class="signin__title">Sign In</div>
      <!-- 認証用のボタン配置 -->
      <div class="signin__signin-buttons">
        <div class="signin__google-signin-button" @click="signIn('google')">
          <img :src="GoogleSignInIcon" alt="" />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import GoogleSignInIcon from "~/assets/images/googleSignInIcon--light.svg";

const { signIn, data } = useAuth();
const router = useRouter();

onMounted(() => {
  console.log(data.value?.user);
  if (data.value?.user) {
    router.push("/");
  }
});
</script>

<style lang="scss" scoped>
.signin {
  margin-top: 4rem;
  display: flex;
  justify-content: center;
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
    Liberation Mono, Courier New, monospace;

  &__container {
    width: 100%;
    max-width: 400px;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }

  &__title {
    font-size: 1.5rem;
    font-weight: 500;
    margin-bottom: 1rem;
  }

  &__signin-buttons {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 1rem;
    margin-top: 1.25rem;
    margin-bottom: 1rem;
  }

  &__google-signin-button {
    width: 100%;
    display: flex;
    justify-content: center;

    & img {
      width: 100%;
      max-width: 240px;
      cursor: pointer;
    }
  }
}
</style>
