<template>
  <div class="pb-24 w-full min-h-screen text-sm">
    <router-view />

    <div class="bottom-0 left-0 fixed flex justify-center p-4 w-full">
      <div class="flex items-center gap-1">
        <div class="flex gap-1 bg-gray-300 bg-opacity-50 backdrop-blur-md p-1 rounded-full">
          <Button
            v-for="nav in tabbarNavigation"
            :key="nav.pathName"
            size="icon"
            variant="ghost"
            :class="{
              'bg-white text-gray-900 border-border': isActive(nav.pathName),
              'text-gray-900 border-gray-300': !isActive(nav.pathName),
            }"
            class="flex justify-center items-center border rounded-full w-16 h-16"
            @click="onNavClick(nav.pathName)"
          >
            <component :is="nav.icon" class="!w-7 !h-7" :strokeWidth="1.2" />
          </Button>
        </div>

        <IssueCreateDialog />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import IssueCreateDialog from '@/modules/issues/components/issue-create-dialog.vue'
import { Home, Search, User, type LucideIcon } from 'lucide-vue-next'
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()

const onNavClick = (path: string) => {
  router.push(path)
}

// Navigation items
const tabbarNavigation: { pathName: string; icon: LucideIcon }[] = [
  {
    pathName: '/',
    icon: Home,
  },
  {
    pathName: '/public-tasks',
    icon: Search,
  },
  {
    pathName: '/profile',
    icon: User,
  },
]

// Get the current route
const route = useRoute()

const isActive = (path: string) => {
  if (path === '/') {
    return route.path === path
  }
  return route.path.startsWith(path)
}
</script>

<style scoped></style>
