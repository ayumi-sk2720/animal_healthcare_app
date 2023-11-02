<script setup lang="ts">
// 対応すべきTODOや、調査等については、「App.vue」参照
import { defineProps, reactive } from "vue";
import { useVuelidate } from "@vuelidate/core";
import { required } from "@vuelidate/validators";

import SubmitButton from "@/components/parts/SubmitButton.vue";
import TitleLabel from "@/components/parts/TitleLabel.vue";
import BaseInput from "@/components/parts/BaseInput.vue";
import HorizontalLine from "@/components/parts/HorizontalLine.vue";

// Vue.js 3のComposition APIでVuelidate 2を利用するための基礎 | https://reffect.co.jp/vue/vulidate-2/
// Vue3+Vuelidateでexternal validationsを試す | https://zenn.dev/kakkoyakakko/articles/ddac0fb3c4c642
const formData = reactive({
  title: "",
  date: "",
  location: "",
});

const rules = {
  title: { required },
};

const v$ = useVuelidate(rules, formData);

const clickEvent = () => {
  console.log("submit", formData);
};

defineProps<{}>();
</script>
<template>
  <div
    class="create-schedule mx-auto block items-center justify-around p-8 lg:px-8 lg:rounded-lg text-gray-700"
  >
    <div class="p-4 pt-0 lg:px-6 font-bold">
      <h1>予定を追加</h1>
    </div>
    <div class="p-4 bg-white rounded-lg grid shadow-xl">
      <TitleLabel label="タイトル" />
      <BaseInput
        id="title"
        name="title"
        type="text"
        placeholder="例）トリミング"
        value=""
      />
      <HorizontalLine />
      <TitleLabel label="日時" />
      <BaseInput
        id="date"
        name="date"
        type="datetime-local"
        placeholder=""
        value=""
      />
      <HorizontalLine />
      <TitleLabel label="場所" />
      <BaseInput
        id="location"
        name="location"
        type="text"
        placeholder="例）ぺテモ立川店"
        value=""
      />
      <HorizontalLine />
      <div class="p-8 text-center">
        <SubmitButton label="作成する" @click="clickEvent" />
      </div>
    </div>
  </div>
</template>
