<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <h2>Navigation</h2>
  <router-link :to="{ name: 'home' }">
    Home
  </router-link> |
  <router-link :to="{ name: 'admin' }">
    Admin
  </router-link>

  <h2>Content</h2>
  <RouterView />
</template>

<style scoped>

</style>
