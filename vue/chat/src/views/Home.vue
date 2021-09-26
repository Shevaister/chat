<template>
    <div>
        <Header/>
        <p class = "t"> hi </p>
    </div>
</template>

<script>
import router from '@/router'
import Header from '@/components/Header.vue'
export default {
    components: {
        Header
    },

    beforeMount() {
        if (localStorage.getItem('jwt') == undefined) {
            router.push('/') 
        }
    },

    methods: {
       
    }
}
</script>

<style scoped>
.t {
    font-size: 1.25rem;
}
</style>