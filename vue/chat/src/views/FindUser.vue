<template>
    <div>
        <Header/>
        <div class="input-group container">
            <input :value = "login" @input = "login = $event.target.value" type="search" class="form-control rounded" placeholder="Search" aria-label="Search" aria-describedby="search-addon" />
            <button @click = "FindUser" type="button" class="btn btn-outline-primary">search</button>
        </div>
        <div v-if = "completed === 1" class = "user-container">
            <img  src = "@/1.png" alt = "img">
            <p :user = "user"> {{user.login}} </p>
            <button @click = "AddUser" type="button" class="btn btn-success user-container-btn">Add</button>
            <button type="button" class="btn btn-primary user-container-btn">Open chat</button>
            <button @click = "BlockUser" type="button" class="btn btn-danger user-container-btn">Block</button>
        </div>
        <div v-if = "completed === 0" class = "none">
            <span> No users found </span>
        </div>
    </div>
</template>

<script>
import Header from '@/components/Header'
import axios from 'axios';
export default {
    components: {
        Header
    },

    beforeMount() {
        if (localStorage.getItem('jwt') == undefined) {
            router.push('/') 
        }
    },

    data() {
        return {
            completed: -1,
            login: '',
            user: {id : '',login: '',},
            jwt: localStorage.getItem('jwt'),
        }
    },

    methods: {
        async FindUser() {
            const response = await axios({
                method: 'post',
                url: 'http://localhost:8000/a/find',
                headers: {
                    Authorization: 'Bearer' + ' ' + this.jwt,
                },
                data: {
                    login: this.login,
                }
            })
            .catch(function () {})

            if (response != undefined) {
                this.completed = 1
                this.user.login = response.data.login
                this.user.id = response.data.id
            } else {
                this.completed = 0
            }


        },

        async AddUser() {
            const response = await axios({
                method: 'post',
                url: 'http://localhost:8000/a/add',
                headers: {
                    Authorization: 'Bearer' + ' ' + this.jwt,
                },
                data: {
                    user2Id: this.user.id,
                }
            })
            .catch(function (error) {
                if (error.response) {
                    alert("This user already in your friendlist")
                }
            })
            
            if (response != undefined) {
                alert("Success")
            }
        },

        async BlockUser() {
            const response = await axios({
                method: 'post',
                url: 'http://localhost:8000/a/block',
                headers: {
                    Authorization: 'Bearer' + ' ' + this.jwt,
                },
                data: {
                    user2Id: this.user.id,
                }
            })
            .catch(function (error) {
                if (error.response) {
                    alert("This user already in your Blocklist")
                }
            })

            if (response != undefined) {
                alert("Success")
            }
        },


    }
}
</script>

<style scoped>
.none {
    margin-top: 5%;
    text-align: center;
}

.container {
    margin-top: 5%;
    width: 90%;
}

.user-container {
    margin-top: 5%;
    margin-left: 10%;
    width: 80%;
    border-radius: 5px;
    padding-top: 5px;
    padding-left: 5px;
    padding-right: 5px;
    padding-bottom: 40px;
    background: lightgray;
}

.user-container-btn {
    margin-right: 5px;
    float:right;
}

img {
  float: left;
  max-width: 60px;
  width: 100%;
  margin-right: 20px;
  border-radius: 50%;
}
</style>