<script setup lang="ts">
import { UseClipboard } from '@vueuse/components'
const slotContainer = ref<HTMLDivElement | null>(null)
const source = ref('');
onMounted(() => {
  if (slotContainer.value) {
    const cleanUrl = window.location.href.slice(0, window.location.href.indexOf("#"))
    source.value = cleanUrl+"#"+slotContainer.value.innerText
  }
})

</script>

<template>
  <div class="flex items-center gap-1">
    <UseClipboard v-slot="{ copy, copied }">
      <Icon
          name="dashicons:admin-links"
          @click="copy(source)"
          :class="[
        'transition-colors duration-300 text-xl cursor-pointer',
        copied ? 'text-green-500' : 'text-gray-400/20 hover:text-gray-400/60'
      ]"
      />
      <div ref="slotContainer"><slot></slot></div>
    </UseClipboard>
  </div>
</template>
