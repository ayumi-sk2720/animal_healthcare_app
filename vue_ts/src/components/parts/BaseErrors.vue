<template>
  <div v-for="error of errors" :key="error.$uid">
    <div class="text-red-700 font-bold">{{ error.$message }}</div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "BaseErrors",
  props: {
    errors: {
      type: Object,
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
