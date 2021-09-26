<template>
    <form @submit.prevent>
        <div class="form-group">
            <label for="exampleInputEmail1">Change login</label>
            <input
                :value = "login"
                @input = "login = $event.target.value"   
                type="text" 
                class="form-control" 
                id="exampleInputEmail1" 
                aria-describedby="emailHelp" 
                placeholder="New login"
            >
        </div>
        <button @click = "ChangeLogin" type="submit" class="btn btn-primary b">Update</button>
    </form>
</template>

<script>
import axios from 'axios'
export default {
    props: [
        'jwt',
    ],

    data() {
        return {
            login: '',
        }
    },

    methods: {
        async ChangeLogin() {
            const response = await axios({
                method: 'post',
                url: 'http://localhost:8000/a/changelogin',
                headers: {
                    Authorization: 'Bearer' + ' ' + this.jwt,
                },
                data: {
                    login: this.login,
                }
            })
            .catch(function (error) {
                if (error.response) {
                    alert("user with this login already exists")
                }
            })
            if (response != undefined) {
                this.login = ''
                alert("success")
            } 
        },

    }
}
</script>

<style scoped>
.b {
    margin-top: 1%;
    margin-bottom: 1%;
}
</style>