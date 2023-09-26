<template>
  <Header />
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link> |
    <router-link to="/create-schedule">予定新規作成</router-link>
  </nav>
  <router-view/>
  <Footer />
</template>

<script>
  import Header from './components/Header'
  import Footer from './components/Footer'
  export default {
    components: {
      Header,
      Footer
    },
  }
</script>