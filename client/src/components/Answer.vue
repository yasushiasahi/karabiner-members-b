<template>
  <div>
    <p v-if="isCollect">正解</p>
    <p v-else>不正解</p>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component
export default class Answer extends Vue {
  @Prop()
  private isCollect!: boolean
}
</script>

<style scoped lang="scss">
div {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.1);
  position: absolute;
  top: 0;
  left: 0;
  p {
    color: red;
    font-size: 300px;
    font-weight: bold;
  }
}
</style>
