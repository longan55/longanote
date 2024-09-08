<script setup>
// import ref from 'vue'
import Menu from './components/menu.vue';
import Content from './components/content.vue';


// const title = ref('这是默认标题')
// const content = ref('### 这是默认内容')
const handleClick = (val) => {
    console.log('val',val)
}
</script>

<template>
  <div class="catalogue">
  <div class="c1">
    <Menu @click="handleClick"></Menu>
  </div>
  </div>
  <div class="content">
    <Content title="1111111", content="#### 欢迎来到longanote!!!" ></Content>
  </div>
</template>

<style>
.c1{
  margin-left: 0%; 
  margin-top: 80%;
}
.catalogue {
  margin-left: 15%; /* 左边距 200px */
  position: fixed; /* 固定在屏幕上 */
  top: 0;
  left: 0;
  width: 12%; /* 视口宽度的 30% */
  height: 100vh; /* 视口高度的 100% */
  background-color: #fcf4f4; /* 背景颜色 */
  border-right: 1px solid #d2f7ce; /*右边框 */
  padding: 0px; /* 内边距 */
  box-sizing: border-box; /* 包括边框和内边距在内 */
  overflow-y: auto; /* 允许纵向滚动 */
}
.content {
  margin-left: 30%; /* 左边距 200px */
  position: fixed; /* 固定在屏幕上 */
  top: 0;
  left: 0;
  width: 40%; /* 视口宽度的 30% */
  height: 100vh; /* 视口高度的 100% */
  background-color: #ffffff; /* 背景颜色 */
  border-right: 1px solid #ddd; /* 右边框 */
  padding: 20px; /* 内边距 */
  box-sizing: border-box; /* 包括边框和内边距在内 */
  overflow-y: auto; /* 允许纵向滚动 */
}
</style>