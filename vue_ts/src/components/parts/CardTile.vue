<script setup lang="ts">
import { defineProps } from "vue";

const props = defineProps<{
  title?: string;
  description?: string;
  date?: string;
}>();
</script>

<template>
  <div class="p-4 my-4 lg:px-6 bg-blue rounded-lg">
    <div
      class="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
    >
      <h5
        class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
      >
        {{ title }}
      </h5>
      <p class="font-normal text-gray-700 dark:text-gray-400">
        {{ description }}
      </p>
    </div>
  </div>
</template>
