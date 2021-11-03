<template>
  <div>
    <form @submit.prevent class = "container">
      <div class="form-group">
        
        <label for="exampleInputEmail1">Login</label>
        <input
            v-model="login"
            type="text" 
            class="form-control" 
            id="exampleInputEmail1" 
            aria-describedby="emailHelp" 
            placeholder="Enter login"
        >
      </div>
      <div class="form-group">
        <label for="exampleInputPassword1">Password</label>
        <input
            v-model="password"
            type="password" 
            class="form-control" 
            id="exampleInputPassword1" 
            placeholder="Password"
        >
      </div>
      <div class = "condition"> login and password must contain at least 8 characters</div>
      <br>
      <button @click = "SignUp" type="submit" class="btn btn-primary">sign up</button>
      <button @click = "SignIn" type="submit" class="btn btn-primary btnright">sign in</button>
    </form>
  </div>
</template>

<script>
import {jwtCheck} from '@/utils/jwtCheck'
import axios from 'axios';
export default {
    data() {
        return {
            login: '',
            password: '',
        }
    },

    created() {
        if (jwtCheck(localStorage.getItem('jwt'))) {
            this.$router.push('home')
        }
    },

    methods: {
        jwtCheck,
        async SignUp() {
            if (this.login.length < 8 ||  this.password.length < 8) {
                alert("Not enough characters in password or login")
            } else {
            const response = await axios({
                method: 'post',
                url: process.env.VUE_APP_API_URL + "/signup",
                headers: {},
                data: {
                    login: this.login,
                    password: this.password,
                }
            })
            .catch(function (error) {
                if (error.response) {
                    alert("user with this login already exists")
                }
            })
            if (response != undefined) {
                alert("success")
            }} 
        },
        async SignIn() {
            if (this.login < 8 ||  this.password.length < 8) {
                alert("Not enough characters in password or login")
            } else {
            const response = await axios({
                method: 'post',
                url: process.env.VUE_APP_API_URL + "/signin",
                headers: {
                },
                data: {
                    login: this.login,
                    password: this.password,
                }
            })
            .catch(function (error) {
                if (error.response) {
                    alert("wrong login or password")
                }
            })

            if (response != undefined) {
                const now = new Date()
                const item = {
		            value: response.data.jwt,
		            expiry: now.getTime() + 24*60*60000, //24 hours
	            }
                localStorage.setItem('jwt', JSON.stringify(item))
                localStorage.setItem('avatar', response.data.avatar)
                this.$router.push('home')
            }}
        }
    }
}
</script>

<style scoped>  
.container {
  margin-top: 5%;
  width: 50%;
}

.btnright {
    float:right;
}

.condition{
    text-align: center;
}

</style>