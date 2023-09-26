<template>
    <div class="home">
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
            msg: 'This is the top page.'
            }
        },
        methods: {
            clear () {
            this.msg = '';
            }
        }
    }
</script>