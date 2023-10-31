<script setup lang="ts">
import { defineProps } from "vue";

defineProps<{
  label?: string;
  click_event?: object;
}>();
</script>

<template>
  <button
    type="submit"
    class="w-4/5 p-3 bg-orange-500 text-white rounded"
    :click="click_event"
  >
    {{ label }}
  </button>
</template>
