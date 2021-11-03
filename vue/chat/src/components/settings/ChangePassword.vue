<template>
    <form @submit.prevent>
        <div class="form-group">
            <label for="exampleInputPassword1">Change password</label>
            <input
                v-model="password" 
                type="password" 
                class="form-control" 
                id="exampleInputPassword1" 
                placeholder="New paassword"
            >
        </div>
        <button @click = "ChangePassword" type="button" class="btn btn-primary b">Update</button>
    </form>
</template>

<script>
import axios from 'axios'
export default {
    props: [
        'jwt'
    ],

    data() {
        return {
            password: '',
        }
    },

    methods: {
        async ChangePassword() {
            if (this.password.length < 8){
                alert("password must contain at least 8 characters")
            } else {
            const response = await axios({
                method: 'post',
                url: process.env.VUE_APP_API_URL + "/a" + "/changepassword",
                headers: {
                    Authorization: 'Bearer' + ' ' + this.jwt,
                },
                data: {
                    password: this.password,
                }
            })
            .catch(function (error) {
                if (error.response) {
                    alert("error")
                }
            })
            if (response != undefined) {
                this.password = ''
                alert("success")
            }} 
        }
    }
}
</script>

<style scoped>
.b {
    margin-top: 1%;
    margin-bottom: 1%;
}
</style>
