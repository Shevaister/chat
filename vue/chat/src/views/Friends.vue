<template>
    <div>
        <Header/>
        <div class = "container">
        <FriendList :friends="friends"/>
        </div>
    </div>
</template>

<script>
import FriendList from '@/components/friends/FriendList'
import axios from 'axios'
import Header from '@/components/Header'
export default {
    data() {
        return{
            friends: [],
            jwt: localStorage.getItem('jwt'),
        }
    },

    components: {
        FriendList,
        Header,
    },
    
    beforeCreate() {
        if (localStorage.getItem('jwt') == undefined) {
            router.push('/') 
        }
    },

    async created() {
        const response = await axios({
            method: 'get',
            url: 'http://localhost:8000/a/friends',
            headers: {
                Authorization: 'Bearer' + ' ' + this.jwt,
            },
            data: {
            }
        })
        .catch(function (error) {
            if (error.response) {
                alert("This user already in your Blocklist")
            }
        })

        this.friends = response.data
        console.log(response)

        
    },
         
    methods: {
         async Unfriend(friendId) {
            const response = await axios({
                method: 'post',
                url: 'http://localhost:8000/a/unfriend',
                headers: {
                    Authorization: 'Bearer' + ' ' + this.jwt,
                },
                data: {
                    user2Id: friendId,
                }
            })
            .catch(function (error) {
                if (error.response) {
                    alert("this user is not your friend")
                }
            })
            if (response != undefined) {
                var id = this.friends.findIndex((x) => x.id === friendId);
                this.friends.splice(id, 1) 
            }
             
        },
    }
    
}
</script>


<style scoped>
.container {
    margin-top: 5%;
}
</style>