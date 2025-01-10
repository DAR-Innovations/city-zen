<template>
  <Dialog>
    <DialogTrigger as-child>
      <Button
        class="flex items-center gap-2 bg-green-600 shadow-md px-4 py-4 rounded-2xl w-full text-white"
      >
        <Check class="!w-6 !h-6" />
        <p>Report Done</p>
      </Button>
    </DialogTrigger>
    <DialogContent class="rounded-3xl sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>Report an Issue</DialogTitle>
        <DialogDescription>
          Upload an image and provide a comment describing the issue.
        </DialogDescription>
      </DialogHeader>

      <form id="reportForm" class="space-y-6" @submit="onSubmit">
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

        <DialogFooter>
          <Button type="submit" form="reportForm"> Submit Report </Button>
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
import { Check } from 'lucide-vue-next'
import { useForm } from 'vee-validate'
import { ref } from 'vue'
import * as z from 'zod'

// State for the uploaded image
const previewImage = ref<string | null>(null)

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
  console.log('Form Submitted:', values)
  previewImage.value = null
})
</script>

<style scoped></style>
