<template>
    <div class="create-schedule mx-auto block max-w-3xl items-center justify-around p-6 lg:px-8 bg-emerald-500 lg:rounded-lg text-gray-700">
        <div class="p-4 pt-0 lg:px-6 font-bold">
            <h1>予定を追加</h1>
        </div>
        <div class="p-4 lg:px-6 bg-white rounded-lg grid justify-center">
            <div class="max-w-7xl p-4">
                <label for="title" class="block p-2">タイトル</label>
                <input type="text" id="title" placeholder="例）トリミング" class="block w-full p-2 border border-gray-700 rounded">
            </div>
            <div class="max-w-7xl p-4">
                <label for="datetime-local" class="block p-2">日時</label>
                <input type="datetime-local" id="datetime-local" class="block w-full p-2 border border-gray-700 rounded">
            </div>
            <div class="max-w-7xl p-4">
                <label for="place" class="block p-2">場所</label>
                <input type="text" id="place" placeholder="例）ペテモ 立川店" class="block w-full p-2 border border-gray-700 rounded">
            </div>
            <div class="p-4 text-center">
                <button type="submit" class="w-4/5 p-3 bg-orange-500 text-white rounded">作成する</button>
            </div>
        </div>
    </div>
</template>
