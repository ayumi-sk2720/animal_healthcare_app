<template>
    <div class="home mx-auto max-w-3xl items-center justify-between p-6 lg:px-8">
        <div>
            <h1>{{ msg }}</h1>
        </div>
        <div>
            <p v-if="msg.length > 0">
                {{ msg }}
            </p>
            <p v-else>
                no text
            </p>
            <input type="text" v-model="msg">
            <button @click="clear()">clear</button>
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