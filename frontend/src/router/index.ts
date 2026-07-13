import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/login',
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../pages/login.vue'),
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../pages/register.vue'),
    },
    {
      path: '/main',
      name: 'main',
      component: () => import('../pages/main.vue'),
    }
  ],
})

export default router
