<template>
  <div class="p-4">
    <p class="font-medium text-xl">Diar Begisbayev</p>
    <p class="mt-1 text-gray-500">Public Volunteer</p>
  </div>

  <div class="flex flex-col gap-4 mt-4 px-4">
    <Button
      @click="onIssuesClick"
      variant="ghost"
      class="flex justify-between items-center gap-4 bg-transparent p-0 w-full"
    >
      <div class="flex items-center gap-4">
        <MapPinPlus class="!w-5 !h-5" />
        <p class="text-base">My Issues</p>
      </div>

      <ChevronRight class="!w-5 !h-5" />
    </Button>

    <Separator />

    <Button
      @click="onLogoutClick"
      variant="ghost"
      class="flex justify-between items-center gap-4 bg-transparent p-0 w-full"
    >
      <div class="flex items-center gap-4">
        <LogOut class="!w-5 !h-5" />
        <p class="text-base">Log out</p>
      </div>

      <ChevronRight class="!w-5 !h-5" />
    </Button>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import { Separator } from '@/core/components/ui/separator'
import { toast } from '@/core/components/ui/toast'
import { authService } from '@/modules/auth/services/auth.service'
import { useMutation } from '@tanstack/vue-query'
import { ChevronRight, LogOut, MapPinPlus } from 'lucide-vue-next'
import { useRouter } from 'vue-router'

const router = useRouter()

const { mutate: logoutMutation } = useMutation({
  mutationFn: () => authService.logoutUser(),
  onSuccess: () => {
    toast({ title: 'Successfully logged out' })
    router.push('/login')
  },
  onError: () => {
    toast({ title: 'Failed to logout' })
  },
})

const onLogoutClick = () => {
  logoutMutation()
}

const onIssuesClick = () => {
  router.push(`profile/issues`)
}
</script>

<style scoped></style>
