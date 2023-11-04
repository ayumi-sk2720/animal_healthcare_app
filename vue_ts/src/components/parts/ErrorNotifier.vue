<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "ErrorNotifier",
  props: {
    visibility: {
      type: Boolean,
      required: true,
    },
  },
  setup(props) {
    return {
      props,
    };
  },
});
</script>
<template>
  <div
    class="p-4 m-4 text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400"
    role="alert"
    v-if="visibility"
  >
    処理に失敗しました。
  </div>
</template>
