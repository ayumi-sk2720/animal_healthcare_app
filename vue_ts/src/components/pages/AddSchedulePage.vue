<script setup lang="ts">
import { defineProps } from "vue";
import SubmitButton from "@/components/parts/SubmitButton.vue";
import TitleLabel from "@/components/parts/TitleLabel.vue";
import BaseInput from "@/components/parts/BaseInput.vue";
import HorizontalLine from "@/components/parts/HorizontalLine.vue";

const clickEvent = () => {
  console.log("hello clickEvent");
};

defineProps<{}>();
// 対応すべきTODOや、調査等については、「App.vue」参照
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
        <SubmitButton label="作成する" :click_event="clickEvent" />
      </div>
    </div>
  </div>
</template>
