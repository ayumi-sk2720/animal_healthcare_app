<template>
  <commonHeader></commonHeader>
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </nav>
  <router-view/>
</template>

<script>
import commonHeader from './components/commonHeader'
  export default {
    components: {
      commonHeader
    },
  }
</script>