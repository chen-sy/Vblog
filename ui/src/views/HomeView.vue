<script setup>
import { ref } from 'vue'
const nowTime = ref(new Date().toLocaleString())

setInterval(() => {
  nowTime.value = new Date().toLocaleString()
}, 1000)

</script>

<template>
  <div>
    <h2>快乐过每一天，现在是：{{ nowTime }}</h2>
  </div>
</template>
