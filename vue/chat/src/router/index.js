import { createRouter, createWebHistory } from 'vue-router'
import Settings from '@/views/Settings.vue'
import Chat from '@/views/Chat.vue'
import Auth from '@/views/Auth.vue'
import Home from '@/views/Home.vue'
import FindUser from '@/views/FindUser.vue'
import Friends from '@/views/Friends.vue'

const routes = [
  {
    path: '/',
    component: Auth
  },
  {
    path: '/settings',
    component: Settings
  },
  {
    path: '/chat',
    component: Chat
  },
  {
    path: '/home',
    component: Home
  },
  {
    path: '/find',
    component: FindUser
  },
  {
    path: '/friends',
    component: Friends
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
