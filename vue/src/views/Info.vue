<template>
    <div class="create-schedule mx-auto block max-w-3xl items-center justify-around p-6 lg:px-8 bg-emerald-500 lg:rounded-lg text-gray-700">
        <div class="p-4 my-4 lg:px-6 bg-white rounded-lg flex justify-around">
            <div class="w-5/12">
                <img src="../assets/chai.png" class="w-10/12 lg:h-full rounded-full">
            </div>
            <div class="w-7/12">
                <div class="p-2 flex justify-between">
                    <p class="text-2xl">{{ data.name }}</p>
                    <p class="text-2xl">{{ data.age }}</p>
                </div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">性別</p>
                    <p class="text-lg">{{ data.gender }}</p>
                </div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">誕生日</p>
                    <p class="text-lg">{{ data.birthday }}</p>
                </div>
            </div>
        </div>
        <div class="flex justify-between">
            <div class="p-4 lg:px-6 bg-white rounded-lg w-6/12 me-4">
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">体重</p>
                    <p class="text-lg">{{ data.weight }}</p>
                </div>
                <div class="p-2 flex items-center">
                    <p class="text-xs py-2 text-gray-500">最終更新</p>
                    <p class="text-xs py-2 ml-4">{{ data.weightLastUpdate }}</p>
                </div>
            </div>
            <div class="p-4 lg:px-6 bg-white rounded-lg w-6/12">
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">目標体重</p>
                    <p class="text-lg">{{ data.targetWeight}}</p>
                </div>
            </div>
        </div>
        <div class="p-4 my-4 lg:px-6 bg-white rounded-lg">
            <div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">今日のお薬</p>
                    <p class="text-lg">{{ data.todayMedication }}</p>
                </div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">次の投薬</p>
                    <div class="flex items-center">
                        <p class="text-xs">{{ data.nextMedicationDate }}</p>
                        <p class="text-lg ml-4">{{ data.nextMedication }}</p>
                    </div>
                </div>
            </div>
        </div>
        <div class="p-4 my-4 lg:px-6 bg-white rounded-lg">
            <div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">気になるメモ</p>
                    <p class="text-lg">{{ data.latestMemo }}</p>
                </div>
                <div class="p-2 flex items-center">
                    <p class="text-xs py-2 text-gray-500">最終更新</p>
                    <p class="text-xs py-2 ml-4">{{ data.memoLastUpdate }}</p>
                </div>
            </div>
        </div>
        <div class="p-4 my-4 lg:px-6 bg-white rounded-lg">
            <div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">次の予定</p>
                    <div class="flex items-center py-2">
                        <p class="text-xs">{{ data.nextScheduleDate1 }}</p>
                        <p class="text-lg ml-4">{{ data.nextSchedule1 }}</p>
                    </div>
                    <div class="flex items-center py-2">
                        <p class="text-xs">{{ data.nextScheduleDate2 }}</p>
                        <p class="text-lg ml-4">{{ data.nextSchedule2 }}</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { reactive } from 'vue'
const data = reactive({
    name: 'チャイ',
    age: '3歳8ヶ月',
    gender: 'メス',
    birthday: '2020年1月27日',
    weight: '4.8kg',
    weightLastUpdate: '2023.09.28 10:00',
    targetWeight: '5.2kg',
    todayMedication: 'なし',
    nextMedicationDate: '2023.10.05',
    nextMedication: 'ネクスガード',
    latestMemo: '今月に入って水を飲む量が増えてきた。今度通院した時に先生に相談したい。改行入る？？？？？',
    memoLastUpdate: '2023.09.30 10:00',
    nextScheduleDate1: '2023.09.30',
    nextSchedule1: 'トリミング',
    nextScheduleDate2: '2023.10.05',
    nextSchedule2: '通院',
})
</script>
