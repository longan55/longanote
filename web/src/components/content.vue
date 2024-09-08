<template>
    <h1 style="text-align: center;"><slot>{{props.title}}</slot></h1>
    <div class="editor">
      <!-- <textarea class="input" :value="input" @input="update"></textarea> -->
      <div class="output" v-html="output"></div>
    </div>
</template>

<script setup>
import { marked } from 'marked'
import {computed} from 'vue'
// import { debounce } from 'lodash-es'

const props = defineProps(['title','content'])
const output = computed(() => marked(props.content))
// const update = debounce((e) => {
//   input.value = e.target.value
// }, 100)
</script>

<style>
body {
  margin: 0;
}

.editor {
  height: 100vh;
  display: flex;
}

.input,
.output {
  overflow: auto;
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  padding: 0 20px;
}

.input {
  border: none;
  border-right: 1px solid #ccc;
  resize: none;
  outline: none;
  background-color: #f6f6f6;
  font-size: 14px;
  font-family: 'Monaco', courier, monospace;
  padding: 20px;
}

code {
  color: #f66;
}
</style>