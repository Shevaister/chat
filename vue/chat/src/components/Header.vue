<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
  <div class="container-fluid">
    <a class="navbar-brand" href="/">Chat</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">
          <a class="nav-link" href="/home">Home</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/friends">Friends</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/chat">Chat</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/find">Find user</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/settings">Settings</a>
        </li>
      </ul>
      <a @click = "SignOut" class="nav-link l" href="/">sign out</a>
    </div>
  </div>
</nav>
</template>

<script>
export default {
    methods: {
        SignOut() {
            localStorage.removeItem('jwt')
        }
    }
}
</script>

<style scoped>
.l {
  text-decoration: none;
  padding: 0rem;
  color: #6c757d; 
}
</style>