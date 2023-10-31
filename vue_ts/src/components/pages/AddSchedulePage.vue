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
// TODO : Vue3を用いた「リアルタイムバリデーション」の実装方法 | https://blog.labol.co.jp/entry/real-time-validation-implementation-method
// TODO : VueのComposition APIでフォームのバリデーション実装 | https://tech.mof-mof.co.jp/blog/validation-vue-composition-api/
// TODO : api request
// TODO : api処理の共通化
/***
 * |
 * | [参考]
 * |
 * 汎用的なフォームコンポーネントを作成したい | https://qiita.com/taka_1156/items/c65700208072b1f30337
 * 使用例 | https://qiita.com/taka_1156/items/c65700208072b1f30337#%E4%BD%BF%E7%94%A8%E6%96%B9%E6%B3%95
 */
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
