<template>
  <div class="chats-container">
    <div
      @click="$parent.getChatMessages(chatPreview.senderId)"
      v-for="chatPreview in chatsPreview"
      :key="chatPreview.messageId"
      :class="[
        'chat-preview-container',
        chatPreview.chatId == chatId ? 'active-chat' : 'standard-chat',
      ]"
    >
      <img src="@/1.png" />
      <p class="nick break">{{ chatPreview.senderLogin }}</p>
      <br>
      <span class="message break">{{ chatPreview.text }}</span>
      <span class="time break">{{ chatPreview.time }}</span>
    </div>
  </div>
</template>

<script>
export default {
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
};
</script>


<style scoped>
.chats-container {
  float: left;
  margin-left: 0;
  width: 20%;
  position: relative;
  height: 800px;
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
    margin-right: 3px;
    margin-left: 3px;
    width: 57px;
    height: 57px;
    display:flex;
  }
  .chats-container {
    width: 63px;
  }
}

@media screen and (min-width: 1001px) and (max-width: 2000px) {
  .chat-preview-container .time {
    display: none;
  }
}
</style>