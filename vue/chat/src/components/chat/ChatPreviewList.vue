<template>
  <div class="chats-container">
    <div
      @click="getChatMessages(chatPreview.senderId)"
      v-for="chatPreview in chatsPreview"
      :key="chatPreview.messageId"
      :class="[
        'chat-preview-container',
        chatPreview.chatId == chatId ? 'active-chat' : 'standard-chat',
      ]"
    >
        <div v-if="chatPreview.senderAvatar != ''">
          <img :src= "photosURL + chatPreview.senderAvatar" />
        </div>
        <div v-else>
          <img src= "@/assets/default.png" />
        </div>
        <span v-if="chatPreview.senderLogin.length<14">
          <p class="nick break">{{ chatPreview.senderLogin }}</p>
        </span>
        <span v-else>
          <p class="nick break">{{ chatPreview.senderLogin.substring(0,14)+".." }}</p>
        </span>
        <br>
        <span v-if="chatPreview.text.length<8">
          <span class="message break">{{ chatPreview.text }}</span>
        </span>
        <span v-else>
          <span class="message break">{{ chatPreview.text.substring(0,8)+".." }}</span>
        </span>
        <span class="time break">{{ chatPreview.time }}</span>
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

  props: ["chatsPreview", "chatId"],

  computed: {
    chatsPreview: function () {
      return this.chatsPreview.sort((p1, p2) => {
        let modifier = -1;
        if (p1.messageId < p2.messageId) return -1 * modifier;
        if (p1.messageId > p2.messageId) return 1 * modifier;
        return 0;
      });
    },
  },

  methods: {
    getChatMessages(senderId) {
      this.$emit('getChatMessages', senderId)
    }
  }
};
</script>


<style scoped>
.chats-container {
  float: left;
  margin-left: 0;
  width: 20%;
  position: relative;
  height: 85vh;
  overflow: auto;
}

img {
  float: left;
  max-width: 55px;
  margin-right: 20px;
  border-radius: 50%;
  margin-top: auto;
  margin-bottom: auto;
}

.chat-preview-container {
  height: 65px;
  display: block;
  float: right;
  margin-top: 5px;
  margin-right: 1%;
  margin-left: 1%;
  width: 98%;
  border-radius: 5px;
}

.chat-preview-container .nick {
  display:inline-block;
}

.chat-preview-container .message {
}

.chat-preview-container .time {
  float: right;
}

.standard-chat {
  background-color: #f1f1f1;
}

.active-chat {
  background-color: DodgerBlue;
  color: white;
}

.standard-chat:hover {
  background-color: lightgray;
}

@media screen and (min-width: 0px) and (max-width: 1000px) {
  .chat-preview-container .nick {
    display: none;
  }
  .chat-preview-container .message {
    display: none;
  }
  .chat-preview-container .time {
    display: none;
  }
  .chat-preview-container {
    float: left;
    margin-right: 3px;
    margin-left: 3px;
    width: 57px;
    height: 57px;
    display:flex;
  }
  .chats-container {
    width: 80px;
  }
}

@media screen and (min-width: 1001px) and (max-width: 2000px) {
  .chat-preview-container .time {
    display: none;
  }
}
</style>