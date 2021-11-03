import { createStore } from 'vuex'

export default createStore({
  state: {
    socket: null,
  },
  mutations: {
    updateConnection(state, socket) {
      state.socket = socket
    }
  },
  actions: {
    connect(ctx, jwt) {
      if (ctx.getters.socket == null) {
        console.log(process.env.VUE_APP_WEBSOCKET + "/" + jwt)
        const socket = new WebSocket(process.env.VUE_APP_WEBSOCKET + "/" + jwt);
        console.log()
        ctx.commit('updateConnection', socket)
      }
    },
    disconnect(ctx){
      ctx.commit('updateConnection', null)
    }
  },
  getters: {
    socket(state) {
      return state.socket
    }
  },
  modules: {
  }
})
