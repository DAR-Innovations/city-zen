<template>
  <div class="relative flex justify-between items-center py-6 p-4">
    <!-- Back Button -->
    <Button variant="ghost" class="p-0" @click="$router.back()">
      <ChevronLeft class="!w-6 !h-6" />
    </Button>

    <p class="left-1/2 absolute font-semibold text-lg -translate-x-1/2">My Issues</p>

    <div class="w-8"></div>
  </div>

  <section class="flex items-center gap-2 mt-2 px-4">
    <Input
      v-model="searchQuery"
      type="text"
      placeholder="Search for a issues"
      class="flex-grow shadow-md px-4 rounded-full h-14"
    />
    <Button size="icon" class="flex-shrink-0 rounded-full w-14 h-14" @click="onSearch">
      <Search class="w-6 h-6" />
    </Button>
  </section>

  <section class="space-y-6 bg-white mt-6 px-4">
    <div class="flex flex-col gap-4 overflow-y-auto">
      <div
        v-for="issue in issues"
        :key="issue.id"
        class="relative flex-shrink-0 rounded-xl w-full h-64 overflow-clip"
      >
        <img
          :src="issue.imageUrl"
          :alt="issue.title"
          class="bg-gray-100 w-full h-full object-cover"
        />

        <div class="bottom-0 left-0 absolute p-2 w-full">
          <div class="bg-black/40 backdrop-blur-md p-4 rounded-xl text-white">
            <p class="font-semibold">{{ issue.title }}</p>
            <div class="flex items-center gap-1 mt-1 text-sm">
              <MapPin class="w-4 h-4" />
              <p class="text-xs">{{ issue.address }}</p>
            </div>
          </div>
        </div>

        <div class="top-0 left-0 absolute flex justify-between items-center gap-4 p-2 w-full">
          <div
            class="flex items-center gap-1 bg-black/40 backdrop-blur-md p-2 rounded-xl text-gray-300 text-xs"
          >
            <Clock class="w-4 h-4" />
            <p>{{ formatDate(issue.createdAt) }}</p>
          </div>

          <div
            :class="`flex items-center gap-1 bg-black/40 backdrop-blur-md p-2 rounded-xl text-xs ${formattedStatusesColors[issue.status]}`"
          >
            <CheckCheck v-if="issue.status === 'DONE'" class="w-4 h-4" />
            <Check v-else class="w-4 h-4" />

            <p>{{ formattedStatusesLabels[issue.status] }}</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import { Input } from '@/core/components/ui/input'
import { ChevronLeft, MapPin, Search } from 'lucide-vue-next'
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()
const route = useRoute()

const searchQuery = ref<string>((route.query.search as string) || '')

// Trigger search
const onSearch = () => {
  router.push({ query: { ...route.query, search: searchQuery.value } })
}

enum IssueStatus {
  OPEN = 'OPEN',
  DONE = 'DONE',
}

// Issues data with addresses added
const issues = [
  {
    id: 1,
    title: 'Broken Streetlight Repair',
    address: '123 Main Street, Cityville',
    description:
      'A streetlight near the main square is broken and needs to be fixed for pedestrian safety.',
    status: IssueStatus.DONE,
    createdAt: '2025-01-10T09:00:00Z',
    imageUrl:
      'https://media.istockphoto.com/id/598171880/photo/broken-lamp.jpg?s=612x612&w=0&k=20&c=4LjHvpVxg0VUttmbWQtYBF5cS7R_9GVyTuKcXcEJCjw=',
  },
  {
    id: 2,
    title: 'Damaged Park Bench',
    address: '456 Park Avenue, Cityville',
    description: 'A bench in the local park has a broken plank and needs repair or replacement.',
    status: IssueStatus.DONE,
    createdAt: '2025-01-09T15:30:00Z',
    imageUrl:
      'https://media.istockphoto.com/id/1193445391/photo/the-old-bench-that-is-broken-on-the-lawn.jpg?s=612x612&w=0&k=20&c=uLwj2gkw8RDWhIICrTtvYhmuhXj_Fn13c3PIZflpJEY=',
  },
  {
    id: 3,
    title: 'Overflowing Trash Can',
    address: '789 Greenway Road, Cityville',
    description: 'A trash can near the bus stop is overflowing and needs immediate attention.',
    status: IssueStatus.OPEN,
    createdAt: '2025-01-10T08:45:00Z',
    imageUrl:
      'https://blog.marinedebris.noaa.gov/sites/default/files/inline-images/OverflowingTrashCans_NOAA.jpeg',
  },
  {
    id: 4,
    title: 'Leaking Water Fountain',
    address: '101 Bus Depot Lane, Cityville',
    description: 'The water fountain in the central plaza has a leakage issue and requires fixing.',
    status: IssueStatus.DONE,
    createdAt: '2025-01-08T14:00:00Z',
    imageUrl: 'https://i.ytimg.com/vi/jDik5Zzx_t0/maxresdefault.jpg',
  },
  {
    id: 5,
    title: 'Fallen Tree Branch',
    address: '202 Knowledge Drive, Cityville',
    description: 'A large tree branch has fallen on the sidewalk, blocking pedestrian access.',
    status: IssueStatus.OPEN,
    createdAt: '2025-01-10T07:15:00Z',
    imageUrl: 'https://www.organicallygreen.org/wp-content/uploads/2016/02/fallentree.jpg',
  },
]

// Status formatting
const formattedStatusesLabels = {
  OPEN: 'Open',
  DONE: 'Done',
}

const formattedStatusesColors = {
  OPEN: 'text-yellow-300',
  DONE: 'text-green-300',
}

// Helper to format ISO date to a readable format
const formatDate = (dateString: string) => {
  return new Intl.DateTimeFormat('en-US', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
  }).format(new Date(dateString))
}

// Handle task click
const onIssueClick = (issueId: number) => {
  router.push(`/issues/${issueId}`)
}
</script>
