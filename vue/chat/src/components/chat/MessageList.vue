<template>
  <div>
    <div class="chatbox">
      <div v-for="message in chat" :key="message.messageId">
        <MessageFormFrom v-if="message.SenderID === chatWith.id">
          <template v-slot:image>
            <img class="right" src="@/1.png" alt="" />
          </template>
          <template v-slot:message>
            {{ message.Text }}
          </template>
          <template v-slot:time>
            {{ message.CreatedAt }}
          </template>
        </MessageFormFrom>
        <MessageFormTo v-else>
          <template v-slot:image>
            <img src="@/1.png" alt="" />
          </template>
          <template v-slot:message>
            {{ message.Text }}
          </template>
          <template v-slot:time>
            {{ message.CreatedAt }}
          </template>
        </MessageFormTo>
      </div>
    </div>
  </div>
</template>

<script>
import MessageFormTo from "@/components/chat/MessageFormTo";
import MessageFormFrom from "@/components/chat/MessageFormFrom";
export default {
  props: ["chat", "chatWith"],

  components: {
    MessageFormTo,
    MessageFormFrom,
  },

  watch: {
    chat: {
      deep: true,
      handler() {
        this.$nextTick(() => {
          this.scrollDown();
        });
      },
    },
  },

  methods: {
    scrollDown() {
      var messageBody = document.querySelector(".chatbox");
      messageBody.scrollTop =
        messageBody.scrollHeight - messageBody.clientHeight;
    },
  },
};
</script>


<style scoped>
.chatbox {
  display: block;
  position: relative;
  width: auto;
  height: 75vh;
  
  overflow: auto;
}

img {
  float: left;
  max-width: 60px;
  width: 100%;
  margin-right: 20px;
  border-radius: 50%;
}

img.right {
  float: right;
  margin-left: 20px;
  margin-right: 0;
}

@media screen and (min-width: 0px) and (max-width: 1000px){
  .chatbox {
    width:auto;
    margin-right: 1%;
  }
}
</style>