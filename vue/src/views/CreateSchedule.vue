<template>
    <div class="create-schedule">
        <h1>予定を追加</h1>
    </div>
    <div>
        <label for="pet-select">ペット名</label>
        <select name="pets" id="pet-select">
            <option value="">-- 選択してください --</option>
            <option v-for="pet in pets" :key="pet.id" :value="pet.name">{{ pet.name }}</option>
        </select>
    </div>
    <div>
        <label for="title">タイトル</label>
        <input type="text" id="title" placeholder="例）トリミング">
    </div>
    <div>
        <label for="datetime-local">日時</label>
        <input type="datetime-local" id="datetime-local">
    </div>
    <div>
        <label for="place">場所</label>
        <input type="text" id="place" placeholder="例）ペテモ 立川店">
    </div>

    <button type="submit">作成する</button>
</template>

<script>
export default {
    data () {
        return {
            pets: [
                {
                    id: 1,
                    name: '麦'
                },
                {
                    id: 2,
                    name: 'チャイ'
                }
            ],
        }
    }
}
</script>