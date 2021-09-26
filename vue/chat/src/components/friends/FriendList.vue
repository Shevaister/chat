<template>
  <div>
    <div v-for="friend in friends" :key="friend.id">
      <div class="user-container">
        <img src="@/1.png" alt="img" />
        <p>{{ friend.login }}</p>
        <button
          @click="this.$parent.Unfriend(friend.id)"
          type="button"
          class="btn btn-danger btn-right"
        >
          Unfriend
        </button>
        <button type="button" class="btn btn-primary btn-right">
          Open chat
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["friends"],
};
</script>

<style scoped>
.user-container {
  margin-top: 5px;
  margin-left: 10%;
  width: 80%;
  border-radius: 5px;
  padding-top: 5px;
  padding-left: 5px;
  padding-right: 5px;
  padding-bottom: 40px;
  background: lightgray;
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
}
</style>