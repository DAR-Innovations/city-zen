import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './modules/landing/pages/home-page.vue'
import { DEFAULT_TITLE, TITLE_TEMPLATE } from './core/constants/seo.constant'

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  scrollBehavior() {
    return { top: 0 }
  },
  routes: [{ path: '/', component: HomePage }],
})

router.beforeEach(async (to, _from, next) => {
  // Set page title
  document.title = to.meta?.title ? TITLE_TEMPLATE(to.meta.title as string) : DEFAULT_TITLE

  return next()
})
