import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Info.vue'
import About from '../views/About.vue'
import AddSchedule from '../views/AddSchedule.vue'

const routes = [
  {
    path: '/',
    name: 'top',
    component: Home
  },
  {
    path: '/pet/info',
    name: 'info',
    component: Home
  },
  {
    path: '/about',
    name: 'about',
    component: About
  },
  {
    path: '/pet/schedule',
    name: 'schedule',
    component: AddSchedule
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
