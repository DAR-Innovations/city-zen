import HomeLayout from '@/core/layouts/home/home-layout.vue'
import { createRouter, createWebHistory } from 'vue-router'
import { DEFAULT_TITLE, TITLE_TEMPLATE } from './core/constants/seo.constant'
import HomePage from './modules/landing/pages/home-page.vue'
import ProfileIssuesPage from './modules/profile/pages/profile-issues-page.vue'
import ProfilePage from './modules/profile/pages/profile-page.vue'
import PublicTaskDetailsPage from './modules/public-tasks/pages/public-task-details-page.vue'
import PublicTasksPage from './modules/public-tasks/pages/public-tasks-page.vue'

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  scrollBehavior() {
    return { top: 0 }
  },
  routes: [
    {
      path: '/',
      component: HomeLayout,
      children: [
        { path: '', component: HomePage },

        { path: 'public-tasks', component: PublicTasksPage },
        { path: 'public-tasks/:id', component: PublicTaskDetailsPage },

        { path: 'profile', component: ProfilePage },
        { path: 'profile/issues', component: ProfileIssuesPage },
      ],
    },
  ],
})

router.beforeEach(async (to, _from, next) => {
  // Set page title
  document.title = to.meta?.title ? TITLE_TEMPLATE(to.meta.title as string) : DEFAULT_TITLE

  return next()
})
