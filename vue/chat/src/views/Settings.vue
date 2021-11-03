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
import {jwtCheck} from '@/utils/jwtCheck'
import ChangeLogin from '@/components/settings/ChangeLogin'
import ChangeAvatar from '@/components/settings/ChangeAvatar'
import ChangePassword from '@/components/settings/ChangePassword'
import Header from '@/components/Header'
export default {
    components: {
        Header,
        ChangeLogin,
        ChangePassword,
        ChangeAvatar,
    },

    data() {
        return {
            jwt: null
        }
    },

    created() {
        let jwt = jwtCheck(localStorage.getItem('jwt'))
        if (!jwt) {
            this.$router.push('/')
        } else {
            this.jwt = jwt
        }
    },

    methods: {
        jwtCheck,
    }
}
</script>

<style scoped>
.container {
    margin-top: 5%;
    width: 50%;
}
</style>