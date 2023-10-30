<!-- Vite + Vue 3 + TypeScript + Tailwind CSS の簡易環境構築 | https://zenn.dev/showy1/articles/c5d1b5d33552be -->
<script setup lang="ts">
import TopSection from "@/components/parts/TopSection.vue";
import CardTile from "@/components/parts/CardTile.vue";
</script>

<template>
  <div>
    <!-- <TopSection msg="Custom Message!!!" /> -->
    <CardTile title="予定" description="ホゲホゲ" date="2020/10/30" />
  </div>
</template>
