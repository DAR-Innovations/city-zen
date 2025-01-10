<template>
  <div class="w-full min-h-screen text-sm">
    <div class="flex items-center gap-4 justify-between p-4 bg-background">
      <div>
        <p class="font-medium text-base">Diar Begisbayev</p>
        <p class="text-gray-500">Public Volunteer</p>
      </div>

      <div>
        <Button class="rounded-full" size="icon" variant="outline">
          <Bell class="h-10 w-10" />
        </Button>
      </div>
    </div>

    <div class="w-full h-full">
      <MapboxMap access-token="...">
        <MapboxMarker position="[0, 0]" />
      </MapboxMap>
    </div>

    <div class="fixed bottom-0 left-0 w-full flex justify-center p-4">
      <div class="flex items-center gap-2">
        <div
          class="rounded-full p-1 bg-gray-100 bg-opacity-80 backdrop-blur-md flex items-center gap-2"
        >
          <Button
            v-for="nav in tabbarNavigation"
            :key="nav.pathName"
            size="icon"
            variant="outline"
            class="rounded-full"
          >
            <component :is="nav.icon" class="text-xl" />
          </Button>
        </div>

        <Button class="rounded-full h-full">
          <Plus class="w-12 h-12" />
        </Button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { MapboxMap, MapboxMarker } from '@studiometa/vue-mapbox-gl'
import 'mapbox-gl/dist/mapbox-gl.css'
import { Button } from '@/core/components/ui/button'
import { Bell, Home, Plus, Search, User, type LucideIcon } from 'lucide-vue-next'

const tabbarNavigation: { pathName: string; icon: LucideIcon }[] = [
  {
    pathName: 'HOME',
    icon: Home,
  },
  {
    pathName: 'SEARCH',
    icon: Search,
  },
  {
    pathName: 'PROFILE',
    icon: User,
  },
]
</script>

<style scoped></style>
