<template>
  <a-menu :style="{ width: '200px', height: '100%' }" :default-open-keys="['0']" :default-selected-keys="['0_2']">
    <a-sub-menu :key=parent.title :title=parent.title v-for="parent in items">
      <!-- <template #icon><icon-bulb></icon-bulb></template> -->
      <div v-if="parent.child != null" v-for="son in parent.child">
        <div v-if="son.child != null">
          <a-sub-menu :key=son.title :title=son.title>
            <a-menu-item v-for="grandson in son.child" :key=grandson.title @click="GetArticle(grandson.path)">{{ grandson.title }}</a-menu-item>
          </a-sub-menu>
        </div>
        <div v-else>
          <a-menu-item :key=son.title>{{ son.title }}</a-menu-item>
        </div>
      </div>
    </a-sub-menu>
  </a-menu>
</template>

<script setup>
import { ref, onMounted, defineEmits} from 'vue';
import axios from 'axios';
const emits = defineEmits(["click"])
var items = ref([])
onMounted(async () => {
  try {
    const response = await fetch('http://127.0.0.1:9000/getMenu');
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    // 解析 JSON 数据
    const data = await response.json();
    items.value = data.Data
  } catch (error) {
    console.error('There was a problem with the fetch operation:', error);
  }
});

async function GetArticle(path) {
  try {
    const response = await axios.post('http://127.0.0.1:9000/getArticle', new URLSearchParams({
      path: path,
    }), {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
    });
    emits('click', response.data);
  } catch (error) {
    console.error('Error sending data:', error);
  }
}
</script>

<style scoped></style>