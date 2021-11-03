<template>
  <div class = "user-list">
    <div v-for="user in users" :key="user.login">
      <div v-if="completed === 1" class="user-container">
        <div v-if="user.avatar != ''">
          <img :src= "photosURL + user.avatar" />
        </div>
        <div v-else>
          <img src= "@/assets/default.png" />
        </div>
        <p :user="user">{{ user.login }}</p>
        <button
          @click="addUser(user.id)"
          type="button"
          class="btn btn-success user-container-btn"
        >
          Add
        </button>
        <button
          @click="openChat(user.id)"
          type="button"
          class="btn btn-primary user-container-btn"
        >
          Open chat
        </button>
        <button
          @click="blockUser(user.id)"
          type="button"
          class="btn btn-danger user-container-btn"
        >
          Block
        </button>
      </div>
    </div>
    <div v-if="completed === 0" class="none">
        <span> No users found </span>
      </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      photosURL: process.env.VUE_APP_IMAGES_URL
    }
  },

  props: ["users", "completed"],

  methods: {
    addUser(id) {
      this.$emit('addUser', id)
    },

    blockUser(id) {
      this.$emit('blockUser', id)
    },

    openChat(id) {
      this.$emit('openChat', id)
    }
  }
};
</script>

<style  scoped>
.none {
    margin-top: 5%;
    text-align: center;
}

.user-list {
  display: block;
  position: relative;
  width: auto;
  height: 80vh;
  overflow: auto;
}

.user-container {
    margin-top: 1vh;
    
    width: auto;
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


