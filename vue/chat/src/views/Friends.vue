<template>
  <div>
    <Header />
    <div class="container">
      <FriendSearch @search="search"/>
      <FriendList :friends="searchedFriends" @unfriend="unfriend" @openChat="openChat"/>
    </div>
  </div>
</template>

<script>
import {jwtCheck} from '@/utils/jwtCheck'
import FriendList from "@/components/friends/FriendList";
import FriendSearch from "@/components/friends/FriendSearch";
import axios from "axios";
import Header from "@/components/Header";
export default {

  data() {
    return {
      friends: [],
      searchedFriends: [],
      jwt: null,
    };
  },

  components: {
    FriendSearch,
    FriendList,
    Header,
  },

  created() {
    let jwt = jwtCheck(localStorage.getItem('jwt'))
    if (!jwt) {
          this.$router.push('/')
    } else {
        this.jwt = jwt
    }
  },

  async beforeMount() {
    const response = await axios({
      method: "get",
      url: process.env.VUE_APP_API_URL + "/a" + "/friends",
      headers: {
        Authorization: "Bearer" + " " + this.jwt,
      },
    }).catch(function (error) {
      if (error.response) {
        alert("error");
      }
    });
    this.friends = response.data;
    this.searchedFriends = response.data
  },

  methods: {
    jwtCheck,
    async unfriend(friendId) {
      const response = await axios({
        method: "post",
        url: process.env.VUE_APP_API_URL + "/a" + "/unfriend",
        headers: {
          Authorization: "Bearer" + " " + this.jwt,
        },
        data: {
          user2Id: friendId,
        },
      }).catch(function (error) {
        if (error.response) {
          alert("this user is not your friend");
        }
      });
      if (response != undefined) {
        var id = this.friends.findIndex((x) => x.id === friendId);
        this.friends.splice(id, 1);
      }
    },

    openChat(id) {
        this.$router.push("/chat/?id=" + id)
    },

    search(word) {
      this.searchedFriends = this.friends.filter(friend => friend.login.includes(word))
    }
  },
};
</script>


<style scoped>
.container {
  width: 90%;
}
</style>