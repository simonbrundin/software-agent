<script setup lang="ts">
import { format, isToday } from 'date-fns'
import type { Conversation, Message } from '~/types'

const props = defineProps<{
  conversation: Conversation
}>()

const emits = defineEmits(['close', 'send-message'])

const messagesContainer = ref<HTMLElement | null>(null)
const newMessage = ref('')
const loading = ref(false)

const toast = useToast()

onMounted(() => {
  scrollToBottom()
})

watch(() => props.conversation.messages, () => {
  nextTick(() => {
    scrollToBottom()
  })
}, { deep: true })

function scrollToBottom() {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

function formatMessageTime(date: string) {
  const d = new Date(date)
  if (isToday(d)) {
    return format(d, 'HH:mm')
  }
  return format(d, 'dd MMM HH:mm')
}

function onSubmit() {
  if (!newMessage.value.trim()) return

  loading.value = true

  setTimeout(() => {
    const newMsg: Message = {
      id: props.conversation.messages.length + 1,
      body: newMessage.value,
      date: new Date().toISOString(),
      isOwn: true
    }

    emits('send-message', newMsg)
    newMessage.value = ''

    toast.add({
      title: 'Message sent',
      description: 'Your message has been sent successfully',
      icon: 'i-lucide-check-circle',
      color: 'success'
    })

    loading.value = false
  }, 500)
}
</script>

<template>
  <UDashboardPanel id="chat">
    <UDashboardNavbar :title="conversation.participant.name" :toggle="false">
      <template #leading>
        <UButton
          icon="i-lucide-x"
          color="neutral"
          variant="ghost"
          class="-ms-1.5"
          @click="emits('close')"
        />
      </template>

      <template #right>
        <UTooltip text="Voice call">
          <UButton
            icon="i-lucide-phone"
            color="neutral"
            variant="ghost"
          />
        </UTooltip>

        <UTooltip text="Video call">
          <UButton icon="i-lucide-video" color="neutral" variant="ghost" />
        </UTooltip>

        <UTooltip text="More options">
          <UButton
            icon="i-lucide-ellipsis-vertical"
            color="neutral"
            variant="ghost"
          />
        </UTooltip>
      </template>
    </UDashboardNavbar>

    <div class="flex items-center gap-4 p-4 sm:px-6 border-b border-default">
      <UAvatar
        v-bind="conversation.participant.avatar"
        :alt="conversation.participant.name"
        size="lg"
      />

      <div class="min-w-0">
        <p class="font-semibold text-highlighted">
          {{ conversation.participant.name }}
        </p>
        <p class="text-muted text-sm">
          {{ conversation.participant.email }}
        </p>
      </div>
    </div>

    <div
      ref="messagesContainer"
      class="flex-1 p-4 sm:p-6 overflow-y-auto space-y-4"
    >
      <div
        v-for="message in conversation.messages"
        :key="message.id"
        class="flex"
        :class="[message.isOwn ? 'justify-end' : 'justify-start']"
      >
        <div
          class="max-w-[75%] rounded-2xl px-4 py-2.5"
          :class="[
            message.isOwn
              ? 'bg-primary text-primary-fg rounded-br-sm'
              : 'bg-muted text-muted-fg rounded-bl-sm'
          ]"
        >
          <p class="whitespace-pre-wrap">
            {{ message.body }}
          </p>
          <p
            class="text-xs mt-1.5 opacity-70"
            :class="[message.isOwn ? 'text-right' : 'text-left']"
          >
            {{ formatMessageTime(message.date) }}
          </p>
        </div>
      </div>
    </div>

    <div class="pb-4 px-4 sm:px-6 shrink-0">
      <UCard variant="subtle" class="mt-auto">
        <form @submit.prevent="onSubmit">
          <div class="flex items-end gap-2">
            <UTextarea
              v-model="newMessage"
              color="neutral"
              variant="none"
              required
              autoresize
              placeholder="Type a message..."
              :rows="1"
              :disabled="loading"
              class="flex-1"
              :ui="{ base: 'resize-none' }"
              @keydown.enter.exact.prevent="onSubmit"
            />

            <UButton
              type="submit"
              color="primary"
              :loading="loading"
              icon="i-lucide-send"
            />
          </div>
        </form>
      </UCard>
    </div>
  </UDashboardPanel>
</template>
