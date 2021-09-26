<template>
    <form @submit.prevent>
        <div class="form-group">
            <label for="exampleInputPassword1">Change password</label>
            <input
                :value = "password"
                @input = "password = $event.target.value"  
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
            const response = await axios({
                method: 'post',
                url: 'http://localhost:8000/a/changepassword',
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
            console.log(response)
            if (response != undefined) {
                this.password = ''
                alert("success")
            } 
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
