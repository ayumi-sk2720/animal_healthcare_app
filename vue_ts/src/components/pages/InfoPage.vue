<!-- Vite + Vue 3 + TypeScript + Tailwind CSS の簡易環境構築 | https://zenn.dev/showy1/articles/c5d1b5d33552be -->
<script setup lang="ts">
import { PetSummary } from "@/apis/petRepository";
import repositoryFactory, { Repositories } from "@/apis/repositoryFactory";
import CardTile from "@/components/parts/CardTile.vue";
import ProfileTile from "@/components/parts/ProfileTile.vue";
import SchedulesTile from "@/components/parts/SchedulesTile.vue";
import SpinnerTile from "@/components/parts/Spinner.vue";
import { key } from "@/main";
import { inject, onMounted, ref } from "vue";

// injectで撮るからには、provideで登録する必要がある
// 文字列だと一致するチェックが緩い（＝取れないかもしれない) *** is possibily undefinedの理由
// const repository = inject<Repositories>("repository");
const repository = inject(key);

// const schedules = [new Schedule("トリミング", "2023.09.30"), new Schedule("通院","2023.10.01")];
// TODO : api request
// TODO : api処理の共通化
// TODO: 【TypeScript】Axiosのリクエストとレスポンスに型をつける | https://zenn.dev/mkt_engr/articles/axios-req-res-typescript
// TODO: Vue 3 + TypeScript ではじめるリポジトリパターン | https://qiita.com/blasg5884/items/394184cbf2f3ea760d4a
// TODO: RepositoryFactoryパターンをVueのAPIリクエストに導入する | https://tech.forstartups.com/entry/2021/07/27/194946
const state = ref<PetSummary>({
  pet: {
    name: "",
    age: "",
    sex: "",
    nowWeight: "",
    targetWeight: "",
    birthDay: "",
  },
  memo: {
    title: "",
    date: "",
  },
  schedules: [],
});
const isLoading = ref(true);
// const clickHoge = () => {
//   const repository = inject<Repositories>("repository");
// }
onMounted(() => {
  const response = async () => {
    // TODO: ローディングアニメーションの制御も、いちいち使う側でやりたくない | おそらくこの処理をうまくレイヤー化できれば、HTTPリクエスト・レスポンスのテスト化が可能？
    // const { data } = await this.$repository.pet.getPetSummary(2); // リクエストできなくなっちゃった、、、

    // Injectをする場所がsetupの箇所です（何かの関数の中ではダメです）
    // 大前提となる処理（inject, stateの初期化、イベントの定義の順かな）は、importの直下、がよくやります
    console.log(repository);
    if (!repository) {
      return;
    }
    const { data } = repository.pet.getPetSummary(2);
    console.log(data); // undefinedになってしまう
    await new Promise((resolve) => setTimeout(resolve, 1000));
    isLoading.value = false;
    const { pet, memo, schedules } = data;

    state.value.pet = pet;
    state.value.memo = memo;
    state.value.schedules = schedules;
  };
  response();
});
</script>
<template>
  <div>
    <SpinnerTile :is-loading="isLoading" />
    <div
      class="p-3 bg-emerald-500 h-screen"
      :class="isLoading ? 'opacity-20' : ''"
    >
      <ProfileTile
        :name="state.pet.name"
        :age="state.pet.age"
        :sex="state.pet.sex"
        :birthday="state.pet.birthDay"
      />
      <div class="flex justify-between">
        <div class="w-6/12">
          <CardTile
            title="体重"
            :description="state.pet.nowWeight"
            size="max-w-lg"
          />
        </div>
        <div class="w-6/12">
          <CardTile
            title="目標体重"
            :description="state.pet.targetWeight"
            size="max-w-lg"
          />
        </div>
      </div>
      <CardTile
        title="気になるメモ"
        :description="state.memo.title"
        :date="state.memo.date"
      />
      <SchedulesTile :schedules="state.schedules" />
    </div>
  </div>
</template>
