import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import DetailView from '../views/DetailView.vue'
import AdminView from '../views/AdminView.vue'
import AdminEditView from '../views/AdminEditView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/machine/:id',
      name: 'detail',
      component: DetailView
    },
    {
      path: '/admin',
      name: 'admin',
      component: AdminView
    },
    {
      path: '/admin/edit/:id',
      name: 'admin-edit',
      component: AdminEditView
    }
  ]
})

export default router
