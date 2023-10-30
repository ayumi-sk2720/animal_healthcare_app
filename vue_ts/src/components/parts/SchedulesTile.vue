<script setup lang="ts">
import { defineProps } from "vue";
import ScheduleItem from "@/components/parts/ScheduleItem.vue";
import Schedule from "@/models/Schedule.ts"; // TODO: An import path can only end with a '.ts' extension when 'allowImportingTsExtensions' is enabled.ts(5097)

defineProps<{
  // schedules: Schedule[];
}>();
</script>

<template>
  <div class="p-2 lg:px-6 bg-blue rounded-lg">
    <div
      class="block p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
    >
      <h5
        class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
      >
        次の予定
      </h5>
      <ScheduleItem date="2023.09.30" description="トリミング" />
      <ScheduleItem date="2023.10.01" description="通院" />
    </div>
  </div>
</template>
