<template>
    <div>
        <Header/>
        <div class = "container">
            <ChangeLogin :jwt = "jwt"/>
            <ChangePassword :jwt = "jwt"/>
            <ChangeAvatar :jwt = "jwt"/>
        </div>
    </div>
</template>

<script>
import ChangeLogin from '@/components/settings/ChangeLogin'
import ChangeAvatar from '@/components/settings/ChangeAvatar'
import ChangePassword from '@/components/settings/ChangePassword'
import Header from '@/components/Header'
import router from '@/router'
export default {
    data() {
        return {
            jwt: localStorage.getItem('jwt'),
        }
    },

    beforeMount() {
        if (localStorage.getItem('jwt') == undefined) {
            router.push('/') 
        }
    },

    components: {
        Header,
        ChangeLogin,
        ChangePassword,
        ChangeAvatar,
    },
}
</script>

<style scoped>
.container {
    margin-top: 5%;
    width: 50%;
}
</style>