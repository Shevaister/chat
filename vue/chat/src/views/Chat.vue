<template>
  <div class="i">
    <Header />
    <ChatPreviewList
      :chatsPreview="chatsPreview"
      :chatId="currentChatId"
      @getChatMessages="getChatMessages"
    />
    <MessageList :chat="chat" :chatWith="chatWith" />
    <SendMessageBox v-if="chatWith.id != null" @send="send" />
  </div>
</template>

<script>
import { jwtCheck } from "@/utils/jwtCheck";
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
      jwt: null,
      socket: null,
      chatWith: {},
      message: {},
      currentChatId: 0,
      pingMsg: JSON.stringify({
        SenderId: 0,
        ReceiverId: 0,
        Text: "ping",
      }),
    };
  },

  created() {
    let jwt = jwtCheck(localStorage.getItem("jwt"));
    if (!jwt) {
      this.$router.push("/");
    } else {
      this.jwt = jwt;
    }
  },

  async beforeMount() {
    var response = await axios({
      method: "get",
      url: process.env.VUE_APP_API_URL + "/a" + "/chatspreview",
      headers: {
        Authorization: "Bearer" + " " + this.jwt,
      },
    }).catch(function (error) {
      if (error.response) {
        alert("error");
      }
    });
    console.log(response)
    for (let i = 0; i < response.data.length; i++) {
      var rawTime = new Date(response.data[i].time);
      var time = rawTime.toLocaleString();
      if (time == "Invalid Date") {
        time = "";
      }
      response.data[i].time = time;
    }

    this.chatsPreview = response.data

    this.$store.dispatch("connect", this.jwt);
    this.socket = this.$store.getters.socket;

    this.socket.onopen = () => {
      const msg = this.pingMsg
      this.socket.send(msg);
    };

    this.socket.onmessage = (event) => {
      this.readWsMessage(JSON.parse(event.data));
    };

    this.socket.onclose = () => {
      alert("whoops looks like you open this site in another tab, to make this page active again reload it");
      //this.$router.go()
    };

    this.socket.onerror = () => {
      alert("error");
      this.$router.go()
    };

    if (this.$route.query.id != undefined) {
      this.getChatMessages(this.$route.query.id);
    }
    this.$router.replace("/chat");
  },

  methods: {
    jwtCheck,

    async getChatMessages(id) {
      var response = await axios({
        method: "get",
        url: process.env.VUE_APP_API_URL + "/a" + "/chat" + "/" + id,
        headers: {
          Authorization: "Bearer" + " " + this.jwt,
        },
      }).catch(function (error) {
        if (error.response) {
          alert("error");
        }
      });
      for (let i = 0; i < response.data.chat.length; i++) {
        var rawTime = new Date(response.data.chat[i].CreatedAt);
        var time = rawTime.toLocaleString();
        response.data.chat[i].CreatedAt = time;
      }
      console.log(response)
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
      this.socket.send(JSON.stringify(obj));
    },

    async ping() {
      const socket = this.socket;
      const pingMsg = this.pingMsg;

      await setTimeout(function () {
        socket.send(pingMsg);
      }, 30000);
    },

    readWsMessage(data) {
      console.log(data)
      if (data.text == "pong" || data.code == 1) {
        this.ping();
      } else if (data.code == 0) {
        alert("you cant send a message to this user");
      } else {
        var id = this.chatsPreview.findIndex((x) => x.chatId === data.chatId);
        var rawTime = new Date(data.time);
        var time = rawTime.toLocaleString();
        if (id == -1) {
          this.chatsPreview.push({
            chatId: data.chatId,
            messageId: data.messageId,
            senderAvatar: data.userAvatar,
            senderLogin: data.userLogin,
            senderId: data.senderId,
            text: data.text,
            time: time,
          });
        } else {
          this.chatsPreview[id].messageId = data.messageId;
          this.chatsPreview[id].text = data.text;
          this.chatsPreview[id].time = time;
          //this.chatsPreview[id].time
        }
        if (data.chatId == this.currentChatId) {
          this.chat.push({
            ChatID: data.chatId,
            CreatedAt: time,
            ID: data.messageId,
            SenderID: data.userId,
            Text: data.text,
          });
        }
      }
    },
  },
};
</script>

