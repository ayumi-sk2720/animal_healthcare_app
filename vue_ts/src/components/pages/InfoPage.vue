<!-- Vite + Vue 3 + TypeScript + Tailwind CSS の簡易環境構築 | https://zenn.dev/showy1/articles/c5d1b5d33552be -->
<script setup lang="ts">
import CardTile from "@/components/parts/CardTile.vue";
import ProfileTile from "@/components/parts/ProfileTile.vue";
import SchedulesTile from "@/components/parts/SchedulesTile.vue";
import Schedule from "@/models/Schedule.ts";

// const schedules = [new Schedule("トリミング", "2023.09.30"), new Schedule("通院","2023.10.01")];
// TODO : api request
// TODO : api処理の共通化
// TODO: 【TypeScript】Axiosのリクエストとレスポンスに型をつける | https://zenn.dev/mkt_engr/articles/axios-req-res-typescript
// TODO: Vue 3 + TypeScript ではじめるリポジトリパターン | https://qiita.com/blasg5884/items/394184cbf2f3ea760d4a

</script>

<template>
  <div>
    <!-- <TopSection msg="Custom Message!!!" /> -->
    <ProfileTile />
    <div class="flex justify-between">
      <div class="w-6/12">
        <CardTile title="体重" description="4.8kg" size="max-w-lg" />
      </div>
      <div class="w-6/12">
        <CardTile title="目標体重" description="5.8kg" size="max-w-lg" />
      </div>
    </div>
    <CardTile
      title="気になるメモ"
      description="今月に入って飲む水の量が増えた気がする次に通院した時に先生に相談する"
      date="2020/10/30"
    />
    <SchedulesTile />
  </div>
</template>
