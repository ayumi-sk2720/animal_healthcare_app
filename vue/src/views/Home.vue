<template>
    <div class="create-schedule mx-auto block max-w-3xl items-center justify-around p-6 lg:px-8 bg-emerald-500 rounded-lg text-gray-700">
        <div class="p-4 my-4 lg:px-6 bg-white rounded-lg flex justify-around">
            <div class="w-5/12">
                <img src="../assets/chai.png" class="w-10/12 h-full rounded-full">
            </div>
            <div class="w-7/12">
                <div class="p-2 flex justify-between">
                    <p class="text-2xl">チャイ</p>
                    <p class="text-2xl">3歳8ヶ月</p>
                </div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">性別</p>
                    <p class="text-lg">メス</p>
                </div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">誕生日</p>
                    <p class="text-lg">2020年1月27日</p>
                </div>
            </div>
        </div>
        <div class="flex justify-between">
            <div class="p-4 lg:px-6 bg-white rounded-lg w-6/12 me-4">
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">体重</p>
                    <p class="text-lg">4.8kg</p>
                </div>
                <div class="p-2 flex items-center">
                    <p class="text-xs py-2 text-gray-500">最終更新</p>
                    <p class="text-xs py-2 ml-4">2023.09.28 10:00</p>
                </div>
            </div>
            <div class="p-4 lg:px-6 bg-white rounded-lg w-6/12">
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">目標体重</p>
                    <p class="text-lg">5.2kg</p>
                </div>
            </div>
        </div>
        <div class="p-4 my-4 lg:px-6 bg-white rounded-lg">
            <div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">今日のお薬</p>
                    <p class="text-lg">なし</p>
                </div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">次の投薬</p>
                    <div class="flex items-center">
                        <p class="text-xs">2023.09.30</p>
                        <p class="text-lg ml-4">ネクスガード</p>
                    </div>
                </div>
            </div>
        </div>
        <div class="p-4 my-4 lg:px-6 bg-white rounded-lg">
            <div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">気になるメモ</p>
                    <p class="text-lg">今月に入って水を飲む量が増えてきた。今度通院した時に先生に相談したい。改行入る？？？？？</p>
                </div>
                <div class="p-2 flex items-center">
                    <p class="text-xs py-2 text-gray-500">最終更新</p>
                    <p class="text-xs py-2 ml-4">2023.09.28 10:00</p>
                </div>
            </div>
        </div>
        <div class="p-4 my-4 lg:px-6 bg-white rounded-lg">
            <div>
                <div class="p-2">
                    <p class="text-sm py-2 text-gray-500">次の予定</p>
                    <div class="flex items-center py-2">
                        <p class="text-xs">2023.09.30</p>
                        <p class="text-lg ml-4">トリミング</p>
                    </div>
                    <div class="flex items-center py-2">
                        <p class="text-xs">2023.09.30</p>
                        <p class="text-lg ml-4">通院</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data () {
        return {
            msg: 'Hello World!'
        }
    },
    methods: {
        clear () {
            this.msg = ''
        }
    }
}
</script>