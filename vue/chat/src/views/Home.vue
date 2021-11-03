<template>
    <div>
        <Header/>
        <p class = "text"> Welcome to our chat, here you can search and add your friends or someone else and real-time chat with them, also you can block any user you want, our chat is combination of the latest webserver tecnologies, enjoy </p>
    </div>
</template>

<script>
import {jwtCheck} from '@/utils/jwtCheck'
import Header from '@/components/Header.vue'
export default {
    components: {
        Header
    },

    created() {
        if (!jwtCheck(localStorage.getItem('jwt'))) {
            this.$router.push('/')
        }
    },

    methods: {
       jwtCheck,
    }
}
</script>

<style scoped>
.text {
    font-size: 1.25rem;
    margin-top: 2%;
    margin-left: 5%;
    width:90%
}
</style>