<template>
  <Dialog v-model:open="isDialogOpen">
    <DialogTrigger as-child>
      <Button class="flex flex-shrink-0 justify-center items-center rounded-full w-16 h-16">
        <Plus class="!w-7 !h-7" :strokeWidth="1.2" />
      </Button>
    </DialogTrigger>
    <DialogContent class="rounded-3xl sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>Create an Issue</DialogTitle>
        <DialogDescription>
          Upload an image, provide a comment, and optionally include your location.
        </DialogDescription>
      </DialogHeader>

      <form id="issueForm" class="space-y-6" @submit="onSubmit">
        <!-- Image Upload Field -->
        <FormField name="image" v-slot="{ componentField }">
          <FormItem>
            <FormLabel>Upload Image</FormLabel>
            <FormControl>
              <Input
                type="file"
                accept="image/*"
                @change="(e: Event) => handleImageUpload(e, componentField)"
                class="block border-gray-300 border rounded-lg w-full text-gray-900 text-sm cursor-pointer focus:outline-none"
              />
            </FormControl>
            <FormMessage />
            <div v-if="previewImage" class="mt-4">
              <img
                :src="previewImage"
                alt="Uploaded Preview"
                class="border rounded-xl w-full h-44 object-cover"
              />
            </div>
          </FormItem>
        </FormField>

        <!-- Comment Field -->
        <FormField name="comment" v-slot="{ componentField }">
          <FormItem>
            <FormLabel>Comment</FormLabel>
            <FormControl>
              <Textarea
                placeholder="Describe the issue..."
                v-bind="componentField"
                class="border focus:border-blue-500 rounded-lg focus:ring w-full h-24 text-sm focus:outline-none"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <!-- Geolocation Info -->
        <div v-if="location" class="text-gray-500 text-sm">
          <p><strong>Location:</strong> {{ location.latitude }}, {{ location.longitude }}</p>
        </div>
        <div v-else class="text-gray-500 text-sm">
          <p>Enable location services to include your location.</p>
        </div>

        <DialogFooter>
          <Button type="submit" form="issueForm">Submit Issue</Button>
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/core/components/ui/dialog'
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'
import { Input } from '@/core/components/ui/input'
import { Textarea } from '@/core/components/ui/textarea'
import { toTypedSchema } from '@vee-validate/zod'
import { Plus } from 'lucide-vue-next'
import { useForm } from 'vee-validate'
import { onMounted, reactive, ref } from 'vue'
import * as z from 'zod'

const isDialogOpen = ref(false)

// State for the uploaded image
const previewImage = ref<string | null>(null)

// Geolocation state
const location = reactive<{ latitude: number | null; longitude: number | null }>({
  latitude: null,
  longitude: null,
})

// Get user's location
const fetchGeolocation = () => {
  if ('geolocation' in navigator) {
    navigator.geolocation.getCurrentPosition(
      (position) => {
        location.latitude = position.coords.latitude
        location.longitude = position.coords.longitude
      },
      (error) => {
        console.error('Error fetching geolocation:', error)
      },
    )
  } else {
    console.warn('Geolocation not supported by the browser.')
  }
}

// Call geolocation fetch on component mount
onMounted(fetchGeolocation)

// Zod schema for form validation
const formSchema = toTypedSchema(
  z.object({
    image: z
      .instanceof(File)
      .optional()
      .refine(
        (file) => file && file.size < 5 * 1024 * 1024, // Ensure file size is less than 5MB
        'File size should be less than 5MB',
      ),
    comment: z
      .string()
      .min(5, 'Comment must be at least 5 characters')
      .max(200, 'Comment must be less than 200 characters'),
  }),
)

// Initialize form using useForm
const { handleSubmit } = useForm({
  validationSchema: formSchema,
})

// Handle image upload
function handleImageUpload(
  event: Event,
  componentField: { onChange: (value: File | null) => void },
) {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    const file = target.files[0]
    previewImage.value = URL.createObjectURL(file)
    componentField.onChange(file) // Pass the file to the form validation
  }
}

// Form submission handler
const onSubmit = handleSubmit((values) => {
  const issueData = {
    ...values,
    location: {
      latitude: location.latitude,
      longitude: location.longitude,
    },
  }

  console.log('Issue Submitted:', issueData)
  previewImage.value = null
  isDialogOpen.value = false
})
</script>

<style scoped></style>
