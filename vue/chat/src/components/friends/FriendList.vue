<template>
  <div class = "friend-list">
    <div v-for="friend in friends" :key="friend.id">
      <div class="user-container">
        <div v-if="friend.avatar != ''">
          <img :src= "photosURL + friend.avatar" />
        </div>
        <div v-else>
          <img src= "@/assets/default.png" />
        </div>
        <p>{{ friend.login }}</p>
        <button
          @click="unfriend(friend.id)"
          type="button"
          class="btn btn-danger btn-right"
        >
          Unfriend
        </button>
        <button @click="openChat(friend.id)" type="button" class="btn btn-primary btn-right">
          Open chat
        </button>
      </div>
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

  props: ["friends"],

  methods: {
    unfriend(id) {
      this.$emit('unfriend', id)
    },

    openChat(id) {
      this.$emit('openChat', id)
    }
  }
};
</script>

<style scoped>
.user-container {
  margin-top: 5px;
  width: auto;
  border-radius: 5px;
  padding-top: 5px;
  padding-left: 5px;
  padding-right: 5px;
  padding-bottom: 40px;
  background: lightgray;
}

.friend-list {
  display: block;
  position: relative;
  width: auto;
  height: 80vh;
  overflow: auto;
}

img {
  float: left;
  max-width: 60px;
  width: 100%;
  margin-right: 20px;
  border-radius: 50%;
}

.btn-right {
  float: right;
  margin-right: 5px;
}
</style>