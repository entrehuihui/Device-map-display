import Vue from 'vue'
import Router from 'vue-router'
import index from '@/pages/index'
import openMap from '@/pages/openMap'
import user from '@/pages/user'
import jump from '@/pages/jump'
import forget from '@/pages/forget'
import register from '@/pages/register'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: index
    },
    {
      path: '/openMap',
      name: 'openMap',
      component: openMap
    },
    {
      path: '/user',
      name: 'user',
      component: user
    },
    {
      path: '/jump',
      name: 'jump',
      component: jump
    },
    {
      path: '/login/forget',
      name: 'forget',
      component: forget
    },
    {
      path: '/login/register',
      name: 'register',
      component: register
    }
  ]
})
