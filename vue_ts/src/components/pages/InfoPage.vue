<!-- Vite + Vue 3 + TypeScript + Tailwind CSS の簡易環境構築 | https://zenn.dev/showy1/articles/c5d1b5d33552be -->
<script lang="ts">
import { PetSummary } from "@/apis/petRepository";
import CardTile from "@/components/parts/CardTile.vue";
import ProfileTile from "@/components/parts/ProfileTile.vue";
import SchedulesTile from "@/components/parts/SchedulesTile.vue";
import SpinnerTile from "@/components/parts/Spinner.vue";
import { defineComponent, reactive, ref } from "vue";

// const schedules = [new Schedule("トリミング", "2023.09.30"), new Schedule("通院","2023.10.01")];
// TODO : api request
// TODO : api処理の共通化
// TODO: 【TypeScript】Axiosのリクエストとレスポンスに型をつける | https://zenn.dev/mkt_engr/articles/axios-req-res-typescript
// TODO: Vue 3 + TypeScript ではじめるリポジトリパターン | https://qiita.com/blasg5884/items/394184cbf2f3ea760d4a
// TODO: RepositoryFactoryパターンをVueのAPIリクエストに導入する | https://tech.forstartups.com/entry/2021/07/27/194946

export default defineComponent({
  name: "InfoPage",
  components: {
    CardTile,
    ProfileTile,
    SchedulesTile,
    SpinnerTile,
  },
  setup() {
    const state = reactive<PetSummary>({
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

    return {
      state,
      isLoading,
    };
  },
  mounted() {
    const response = async () => {
      // TODO: ローディングアニメーションの制御も、いちいち使う側でやりたくない | おそらくこの処理をうまくレイヤー化できれば、HTTPリクエスト・レスポンスのテスト化が可能？
      const { data } = await this.$repository.pet.getPetSummary(2);
      await new Promise((resolve) => setTimeout(resolve, 1000));
      this.isLoading = false;
      const { pet, memo, schedules } = data;

      this.state.pet = pet;
      this.state.memo = memo;
      this.state.schedules = schedules;
    };
    response();
  },
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
