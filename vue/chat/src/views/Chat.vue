<template>
  <div class = "i">
    <Header />
    <ChatPreviewList :chatsPreview="chatsPreview" :chatId="currentChatId" />
    <MessageList :chat="chat" :chatWith="chatWith" />
    <SendMessageBox v-if="chatWith.id != null" />
  </div>
</template>

<script>
import axios from "axios";
import Header from "@/components/Header";
import MessageList from "@/components/chat/MessageList";
import SendMessageBox from "@/components/chat/SendMessageBox";
import ChatPreviewList from "@/components/chat/ChatPreviewList";
export default {
  components: {
    MessageList,
    Header,
    ChatPreviewList,
    SendMessageBox,
  },

  data() {
    return {
      chatsPreview: [],
      chat: [],
      jwt: localStorage.getItem("jwt"),
      socket: null,
      chatWith: {},
      message: {},
      currentChatId: 0,
    };
  },

  beforeCreate() {
    if (localStorage.getItem('jwt') == undefined) {
            router.push('/') 
        }
    console.log(this.$route.query.id)
    this.$router.replace('/chat')
  },

  async created() {
    var response = await axios({

      method: "get",
      url: "http://localhost:8000/a/chatspreview",
      headers: {
        Authorization: "Bearer" + " " + this.jwt,
      },
      data: {},
    }).catch(function (error) {
      if (error.response) {
        alert("error");
      }
    });

    for (let i = 0; i < response.data.length; i++) {
      var rawTime = new Date(response.data[i].time);
      var time = rawTime.toLocaleString();
      if (time == "Invalid Date"){
        time = ""
      }
      response.data[i].time = time;
    }


    this.chatsPreview = response.data;

    this.socket = new WebSocket("ws://localhost:8000/ws/" + this.jwt);

    this.socket.onmessage = (event) => {
      this.readWsMessage(JSON.parse(event.data));
    };
    /*this.socket.onerror = function(event) {
      console.error("WebSocket error observed:", event);
    }*/
  },

  methods: {
    async getChatMessages(id) {
      var response = await axios({
        method: "get",
        url: "http://localhost:8000/a/chat/" + id,
        headers: {
          Authorization: "Bearer" + " " + this.jwt,
        },
        data: {},
      }).catch(function (error) {
        if (error.response) {
          alert("error");
        }
      });
      console.log(this.$route.query.userId)
      for (let i = 0; i < response.data.chat.length; i++) {
        var rawTime = new Date(response.data.chat[i].CreatedAt);
        var time = rawTime.toLocaleString();
        response.data.chat[i].CreatedAt = time;
      }

      this.chat = response.data.chat;
      this.chatWith = response.data.user;
      this.currentChatId = response.data.chatId;
    },

    send(text) {
      var obj = {
        SenderId: 0,
        ReceiverId: this.chatWith.id,
        Text: text,
      };
      var a = this.socket.send(JSON.stringify(obj));
      console.log(a)
    },

    readWsMessage(data) {
      var id = this.chatsPreview.findIndex((x) => x.chatId === data.chatId);
      var rawTime = new Date(data.time);
      var time = rawTime.toLocaleString();
      if (id == -1) {
        this.chatsPreview.push({
          chatId: data.chatId,
          messageId: data.messageId,
          senderAvatar: data.senderAvatar,
          senderLogin: data.senderLogin,
          senderId: data.senderId,
          text: data.text,
          time: time,
        });
      } else {
        this.chatsPreview[id].messageId = data.messageId;
        this.chatsPreview[id].text = data.text;
        this.chatsPreview[id].time = time;
      }
      if (data.chatId == this.currentChatId) {
        this.chat.push({
          ChatID: data.chatId,
          CreatedAt: time,
          ID: data.messageId,
          SenderID: data.senderId,
          Text: data.text,
        });
      }
    },
  },
};
</script>

<style scoped>
</style>