<template>
  <section class="flex items-center gap-2 bg-gray-100 px-4 py-4">
    <Input
      v-model="searchQuery"
      type="text"
      placeholder="Search for a task..."
      class="flex-grow shadow-md px-4 rounded-full h-14"
    />
    <Button size="icon" class="flex-shrink-0 rounded-full w-14 h-14" @click="onSearch">
      <Search class="w-6 h-6" />
    </Button>
  </section>

  <section class="space-y-6 bg-white mt-4 px-4">
    <div class="flex flex-col gap-4 overflow-y-auto">
      <div
        v-for="task in filteredTasks"
        :key="task.id"
        class="relative rounded-xl w-full max-w-full h-64 cursor-pointer overflow-clip"
        @click="onTaskClick(task.id)"
      >
        <img
          :src="task.imageUrl"
          :alt="task.title"
          class="bg-gray-100 w-full h-full object-cover"
        />

        <div class="bottom-0 left-0 absolute p-2 w-full">
          <div class="bg-black/40 backdrop-blur-md p-4 rounded-xl text-white">
            <p class="line-clamp-1 font-semibold">{{ task.title }}</p>
            <div class="flex items-center gap-1 mt-1 text-sm">
              <MapPin class="w-4 h-4" />
              <p class="text-xs">{{ task.address }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import { Input } from '@/core/components/ui/input'
import { MapPin, Search } from 'lucide-vue-next'
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()
const route = useRoute()

const searchQuery = ref<string>((route.query.search as string) || '')

// Trigger search
const onSearch = () => {
  router.push({ query: { ...route.query, search: searchQuery.value } })
}

const publicTasks = [
  {
    id: 1,
    imageUrl: 'https://cdn.britannica.com/93/171293-050-D99BEDB2/Graffiti-Berlin-Wall.jpg',
    title: 'Graffiti Removal ',
    description:
      "Remove graffiti from public walls and benches to improve the neighborhood's appearance.",
    address: '123 Main Street, Cityville',
  },
  {
    id: 2,
    title: 'Park Litter Cleanup',
    imageUrl:
      'https://www.nydailynews.com/wp-content/uploads/migration/2020/08/10/D5J74UEKWVEV3C6ON4D7RBBGRA.jpg',
    description:
      'Collect and dispose of litter in the local park to keep it clean and safe for visitors.',
    address: '456 Park Avenue, Cityville',
  },
  {
    id: 3,
    title: 'Community Garden Maintenance',
    imageUrl:
      'https://www.lenexa.com/files/sharedassets/city/v/1/parks-places/parks-outdoors/images/park-electric-community-garden.jpg?dimension=pageimage&w=1140',
    description:
      'Weed, water, and maintain the community garden to support local sustainability efforts.',
    address: '789 Greenway Road, Cityville',
  },
  {
    id: 4,
    title: 'Bus Stop Cleaning',
    imageUrl: 'https://pbs.twimg.com/media/FlkJhezacAAeTMl.jpg:large',
    description:
      'Clean and sanitize the benches and surrounding area of the neighborhood bus stop.',
    address: '101 Bus Depot Lane, Cityville',
  },
  {
    id: 5,
    title: 'Library Bookshelf Organization',
    imageUrl:
      'https://cdn.apartmenttherapy.info/image/upload/f_jpg,q_auto:eco,c_fill,g_auto,w_1500,ar_1:1/at%2Fhouse%20tours%2Farchive%2FGrace%20G.%2F288b7b2d8fd0914f603f025761f1081d1d5aaf25',
    description: 'Organize bookshelves and help with light cleaning in the community library.',
    address: '202 Knowledge Drive, Cityville',
  },
]

const filteredTasks = computed(() => {
  const searchQuery = ((route.query.search as string) || '').toLowerCase()
  if (!searchQuery) return publicTasks
  return publicTasks.filter(
    (task) =>
      task.title.toLowerCase().includes(searchQuery) ||
      task.description.toLowerCase().includes(searchQuery) ||
      task.address.toLowerCase().includes(searchQuery),
  )
})

// Handle task click
const onTaskClick = (taskId: number) => {
  router.push(`/public-tasks/${taskId}`)
}
</script>
