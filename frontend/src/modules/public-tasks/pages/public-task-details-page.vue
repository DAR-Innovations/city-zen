<template>
  <div class="w-full min-h-screen">
    <!-- Top Image Section -->
    <div class="relative">
      <img :src="task.imageUrl" :alt="task.title" class="bg-gray-100 w-full h-80 object-cover" />

      <!-- Exit and Share Buttons -->
      <div class="top-4 left-4 absolute flex items-center gap-2">
        <Button
          @click="onBackClick"
          size="icon"
          class="bg-white bg-opacity-60 shadow-md backdrop-blur-md border rounded-full w-12 h-12"
        >
          <ChevronLeft class="!w-7 !h-7 text-black" />
        </Button>
      </div>
      <div class="top-4 right-4 absolute flex items-center gap-2">
        <Button
          @click="onShareClick"
          size="icon"
          class="bg-white bg-opacity-60 shadow-md backdrop-blur-md border rounded-full w-12 h-12"
        >
          <Share class="!w-5 !h-5 text-black" />
        </Button>
      </div>
    </div>

    <!-- Task Details Section -->
    <section class="p-4">
      <div>
        <h1 class="font-semibold text-xl">{{ task.title }}</h1>

        <div class="flex items-center gap-2 mt-2">
          <MapPin class="w-4 h-4 text-gray-500" />
          <p class="text-gray-500 text-sm">{{ task.address }}</p>
        </div>
      </div>

      <div class="mt-4">
        <PublicTasksReportDialog />
      </div>

      <!-- Description -->
      <div class="mt-5">
        <p class="leading-relaxed">
          {{ task.description }}
        </p>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import { toast } from '@/core/components/ui/toast'
import PublicTasksReportDialog from '@/modules/public-tasks/components/details/public-tasks-report-dialog.vue'
import { ChevronLeft, MapPin, Share } from 'lucide-vue-next'
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()
const route = useRoute()

const onBackClick = () => {
  router.back()
}

const onShareClick = async () => {
  try {
    const currentURL = `${window.location.origin}${route.fullPath}`

    await navigator.clipboard.writeText(currentURL)

    toast({
      title: 'Link Copied',
      description: 'The link has been copied to your clipboard.',
    })
  } catch (error) {
    toast({
      title: 'Error',
      description: 'Failed to copy the link. Please try again.',
    })
    console.error('Failed to copy URL:', error)
  }
}
const task = {
  id: 1,
  imageUrl: 'https://cdn.britannica.com/93/171293-050-D99BEDB2/Graffiti-Berlin-Wall.jpg',
  title: 'Graffiti Removal ',
  description:
    "Remove graffiti from public walls and benches to improve the neighborhood's appearance.",
  address: '123 Main Street, Cityville',
}
</script>

<style scoped>
/* Optional styles if required */
</style>
