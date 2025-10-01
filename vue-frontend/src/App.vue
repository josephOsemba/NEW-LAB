<template>
  <div class="flex flex-col min-h-screen">
    <div class="sticky top-0 z-50">
      <TopNavBar />
    </div>

    <main class="flex-1 mt-0">
      <router-view />
    </main>

    <SiteFooter />
  </div>
</template>

<script setup>
import TopNavBar from "@/components/TopNavBar.vue"
import SiteFooter from "@/components/SiteFooter.vue"
</script>


