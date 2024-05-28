<template>
  <div class="profile-menu">
    <template v-for="item in items">
      <SidebarsPromptsItem
        :icon="item.icon"
        :text="item.text"
        :is-active="route.fullPath === item.path"
        @click="() => router.push(item.path)"
      />
    </template>
  </div>
</template>

<script lang="ts" setup>
interface SidebarItem {
  icon: string;
  text: string;
  path: string;
}

const route = useRoute();
const router = useRouter();

const items: SidebarItem[] = [
  { icon: "user", text: "My Prompts", path: "/prompts" },
  { icon: "user", text: "Create", path: "/prompts/create" },
];
</script>

<style lang="scss" scoped></style>
