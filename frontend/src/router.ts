import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from './modules/landing/pages/landing-page.vue'
import { DEFAULT_TITLE, TITLE_TEMPLATE } from './core/constants/seo.constant'

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  scrollBehavior() {
    return { top: 0 }
  },
  routes: [{ path: '/', component: LandingPage }],
})

router.beforeEach(async (to, _from, next) => {
  // Set page title
  document.title = to.meta?.title ? TITLE_TEMPLATE(to.meta.title as string) : DEFAULT_TITLE

  return next()
})
