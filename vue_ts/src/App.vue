<script setup lang="ts">
import FloatingButton from "@/components/parts/FloatingButton.vue";
</script>
<template>
  <div class="bg-emerald-500 h-screen">
    <!-- <nav>
      <router-link to="/info">Info</router-link>
      <router-link to="/schedule">AddSchedule</router-link>
    </nav> -->
    <router-view />
    <FloatingButton />
  </div>
</template>
