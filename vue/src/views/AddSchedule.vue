<template>
    <div class="create-schedule mx-auto block max-w-3xl items-center justify-around p-6 lg:px-8 bg-emerald-500 lg:rounded-lg text-gray-700">
        <div class="p-4 pt-0 lg:px-6 font-bold">
            <h1>予定を追加</h1>
        </div>
        <div class="p-4 lg:px-6 bg-white rounded-lg grid justify-center">
            <div class="max-w-7xl p-4">
                <label for="title" class="block p-2">タイトル</label>
                <input type="text" id="title" placeholder="例）トリミング" class="block w-full p-2 border border-gray-700 rounded">
                <input type="text" id="nameText">
            </div>
            <div class="max-w-7xl p-4">
                <label for="datetime-local" class="block p-2">日時</label>
                <input type="datetime-local" id="date" class="block w-full p-2 border border-gray-700 rounded">
            </div>
            <div class="max-w-7xl p-4">
                <label for="location" class="block p-2">場所</label>
                <input type="text" id="location" placeholder="例）ペテモ 立川店" class="block w-full p-2 border border-gray-700 rounded">
            </div>
            <div class="p-4 text-center">
                <button type="button" @click="addSchedule" class="w-4/5 p-3 bg-orange-500 text-white rounded">作成する</button>
            </div>
        </div>
    </div>
</template>

<script>
// TODO: POST送信の記述が正しいか確かめる

// function addSchedule() {
//     const title = document.getElementById('title').value
//     const date = document.getElementById('date').value
//     const location = document.getElementById('location').value
//     console.log(title, date, location)
// }

import axios from 'axios';
// axios.defaults.headers.common = {
//     'X-Requested-With': 'XMLHttpRequest',
//     'X-CSRF-TOKEN' : document.querySelector('meta[name="csrf-token"]').getAttribute('content')
// };
export default {
    data(){
        return {
            date: "",
            title: "",
            location: "",
        };
    },
    methods:{
    addSchedule: function () {
        const date = document.getElementById('date').value
        const title = document.getElementById('title').value
        const location = document.getElementById('location').value
        axios
            .post("http://127.0.0.1:4010/pet/schedule", {
                date: date,
                title: title,
                location: location,
            }, {
                headers: {'Content-Type': 'application/json'}
            })
            .then((res) => {
                console.log(res);
                this.posts = res.data.posts;
            })
            .catch((err) => {
                console.log(date)
                console.log(title)
                console.log(location)
                console.log(err.response);
            });
        },
    },
};
</script>
