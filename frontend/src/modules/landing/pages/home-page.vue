<template>
  <div class="flex justify-between items-center gap-4 bg-background p-4">
    <div>
      <p class="font-medium text-xl">Diar Begisbayev</p>
      <p class="mt-1 text-gray-500">Public Volunteer</p>
    </div>

    <div>
      <Button class="p-0 rounded-full w-14 h-14" variant="outline">
        <Bell class="w-20 h-20" />
      </Button>
    </div>
  </div>

  <div class="mt-2 px-4">
    <section class="flex items-center gap-2">
      <Input type="text" placeholder="Найти задачу" class="flex-grow px-4 rounded-full h-14" />
      <Button size="icon" class="flex-shrink-0 rounded-full w-14 h-14">
        <Search class="w-6 h-6" />
      </Button>
    </section>

    <HomePublicTasks class="mt-6" />

    <HomeUserIssues class="mt-6" />
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import { Input } from '@/core/components/ui/input'
import HomePublicTasks from '@/modules/landing/components/home-public-tasks.vue'
import HomeUserIssues from '@/modules/landing/components/home-user-issues.vue'
import { Bell, Search } from 'lucide-vue-next' /* PartiallyEnd: #3632/scriptSetup.vue */
</script>

<style scoped></style>
