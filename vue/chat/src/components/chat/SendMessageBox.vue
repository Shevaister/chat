<template>
    <div class="send-container"> 
      <form @submit.prevent>
        <input
          v-model="text"
          @keyup.enter="send"
          class="send-text"
          type="text"
        />
        <input
          @click="send"
          class="send-button"
          type="button"
          value="send"
        />
      </form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            text: ""
        }
    },

    methods: {
        send() {
            this.$emit('send', this.text)
            this.text = ""
        }
    }
}
</script>

<style scoped>
.send-container {
  position: fixed;
  margin-top: 1%;
  width: 100%;
  bottom: 2%;
}

.send-text {
  font-size: 20px;
  width: 80%;
}

.send-button {
  width: 20%;
  background: none;
  border: 1px solid black;
  font-size: 20px;
}

.send-button:hover {
  background: lightgray;
}
</style>