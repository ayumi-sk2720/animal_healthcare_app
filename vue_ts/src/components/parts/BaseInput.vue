<script lang="ts">
import { defineComponent, PropType, SetupContext } from "vue";
import BaseErrors from "@/components/parts/BaseErrors.vue";
import TitleLabel from "@/components/parts/TitleLabel.vue";

// fileやcolorはUIが違いすぎるので別枠で作成
// もしくはカラーピッカーなど別で作ったほうが見栄えがいいものは除外(場合によっては、time、dateも分ける？)
type InputAttr =
  | "text"
  | "number"
  | "email"
  | "url"
  | "password"
  | "tel"
  | "date"
  | "time"
  | "datetime-local";

export default defineComponent({
  name: "BaseInput",
  props: {
    id: {
      type: String as PropType<string>,
      required: true,
    },
    name: {
      type: String as PropType<string>,
      required: true,
    },
    type: {
      type: String as PropType<InputAttr>,
      required: true,
    },
    value: {
      type: String as PropType<string>,
      required: true,
    },
    placeholder: {
      type: String as PropType<string>,
      required: true,
    },
    errors: {
      type: Object,
      required: true,
    },
    labelTitle: {
      type: String,
      required: true,
    },
  },
  setup(props, context: SetupContext) {
    const updateValue = (e: Event) => {
      if (e.target instanceof HTMLInputElement) {
        context.emit("update:value", e.target.value);
      }
    };
    return {
      props,
      updateValue,
    };
  },
  components: { BaseErrors, TitleLabel },
});
</script>
<template>
  <div>
    <TitleLabel :label="labelTitle" />
    <input
      :id="id"
      :name="name"
      :type="type"
      :value="value"
      :placeholder="placeholder"
      @input="updateValue"
      class="block w-full p-5 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-md focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
    />
    <BaseErrors :errors="errors" />
  </div>
</template>


<style scoped>
input {
  width: 100%;
}
</style>
