<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
  <div class="container-fluid">
    <button class="navbar-brand button" @click="Push('home')">Chat</button>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">
          <button @click="Push('home')" class="nav-link button" >Home</button>
        </li>
        <li class="nav-item">
          <button @click="Push('friends')" class="nav-link button">Friends</button>
        </li>
        <li class="nav-item">
          <button @click="Push('chat')" class="nav-link button">Chat</button>
        </li>
        <li class="nav-item">
          <button @click="Push('find')" class="nav-link button">Find user</button>
        </li>
        <li class="nav-item">
          <button @click="Push('settings')" class="nav-link button">Settings</button>
        </li>
      </ul>
      <button @click = "SignOut" class="nav-link sign-out button" >sign out</button>
    </div>
  </div>
</nav>
</template>

<script>
export default {
    methods: {
      Push(destination){
        this.$router.push(destination)
      },

      SignOut() {
        localStorage.removeItem('jwt')
        const socket = this.$store.getters.socket
        if (socket != null) {
          socket.onclose = () => {}
          socket.close()
          this.$store.dispatch("disconnect")
        }
        this.$router.push('/')
      }
    }
}
</script>

<style scoped>
.sign-out {
  text-decoration: none;
  padding: 0rem;
  color: #6c757d; 
}
.button {
  border: none;
  background: none;
}
</style>