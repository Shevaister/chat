<template>
    <div>
        <Header/>
        <div class = "container">
            <UserSearch @findUser = "findUser"/>
            <UserList :users = "users" :completed = "completed" @addUser = "addUser" @blockUser = "blockUser" @openChat = "openChat"/>
        </div>
        
    </div>
</template>

<script>
import {jwtCheck} from '@/utils/jwtCheck'
import UserList from '@/components/findUser/UserList'
import UserSearch from '@/components/findUser/UserSearch'
import Header from '@/components/Header'
import axios from 'axios';
export default {
    components: {
        Header,
        UserList,
        UserSearch,
    },

    data() {
        return {
            completed: -1,
            login: '',
            users: [],
            jwt: null,
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
        async findUser(login) {
            const response = await axios({
                method: 'post',
                url: process.env.VUE_APP_API_URL + "/a" + "/find",
                headers: {
                    Authorization: 'Bearer' + ' ' + this.jwt,
                },
                data: {
                    login: login,
                }
            })
            .catch(function () {})

            if (response.data.length != 0) {
                this.completed = 1
                this.users = response.data
            } else {
                this.completed = 0
                this.users = response.data
            }


        },

        async addUser(id) {
            const response = await axios({
                method: 'post',
                url: process.env.VUE_APP_API_URL + "/a" + "/add",
                headers: {
                    Authorization: 'Bearer' + ' ' + this.jwt,
                },
                data: {
                    user2Id: id,
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

        async blockUser(id) {
            const response = await axios({
                method: 'post',
                url: process.env.VUE_APP_API_URL + "/a" + "/block",
                headers: {
                    Authorization: 'Bearer' + ' ' + this.jwt,
                },
                data: {
                    user2Id: id,
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

        openChat(id) {
            this.$router.push("/chat/?id=" + id)
        }


    }
}
</script>

<style scoped>
.container {
  width: 90%;
}
</style>