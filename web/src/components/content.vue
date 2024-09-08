<template>
    <h1 style="text-align: center;"><slot>{{ title }}</slot></h1>
    <div class="editor">
      <!-- <textarea class="input" :value="input" @input="update"></textarea> -->
      
      <div class="output" v-html="output"></div>
    </div>
</template>

<script setup>
import { marked } from 'marked'
import { debounce } from 'lodash-es'
import { ref, computed} from 'vue'

// const title = ref('这是标题')
// const title = props.title
const input = ref('#### 欢迎来到longanote')
const props = defineProps({
    title: {
        type: String,
        default: "",
    },
    content: {
      type: String,
      default: "",
    }
})
const output = computed(() => marked(props.content))

const update = debounce((e) => {
  input.value = e.target.value
}, 100)



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
  width: 50%;
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