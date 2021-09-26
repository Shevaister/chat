<template>
  <div>
    <form @submit.prevent class = "container">
      <div class="form-group">
        <label for="exampleInputEmail1">Login</label>
        <input
            :value = "login"
            @input = "login = $event.target.value"  
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
            :value = "password"
            @input = "password = $event.target.value" 
            type="password" 
            class="form-control" 
            id="exampleInputPassword1" 
            placeholder="Password"
        >
      </div>
      <br>
      <button @click = "SignUp" type="submit" class="btn btn-primary">sign up</button>
      <button @click = "SignIn" type="submit" class="btn btn-primary btnright">sign in</button>
    </form>
  </div>
</template>

<script>
import router from '@/router'
import axios from 'axios';
export default {

    data() {
        return {
            login: '',
            password: '',
            socket: null,
        }
    },
    beforeMount() {
            
            if (localStorage.getItem('jwt') != undefined) {
               router.push('home') 
            }
        },
    methods: {
        async SignUp() {
            const response = await axios({
                method: 'post',
                url: 'http://localhost:8000/signup',
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
            } 
        },
        async SignIn() {
            const response = await axios({
                method: 'post',
                url: 'http://localhost:8000/signin',
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
                localStorage.setItem('jwt', response.data)
                router.push('home')
            }
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

</style>