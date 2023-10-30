<script setup lang="ts">
import { defineProps } from "vue";

defineProps<{
  date?: string;
  description?: string;
}>();
</script>

<template>
  <div>
    <div class="flex justify-between">
      <div class="w-6/12">
        <p class="font-normal text-gray-700 dark:text-gray-400">
          {{ date }}
        </p>
      </div>
      <div class="w-6/12">
        <p class="font-normal text-gray-700 dark:text-gray-400">
          {{ description }}
        </p>
      </div>
    </div>
  </div>
</template>
