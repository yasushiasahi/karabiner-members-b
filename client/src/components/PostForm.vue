<template>
  <div>
    <p>SCORE : {{ score }}</p>
    <input v-model="name" placeholder="名前を入力" />
    <span v-on:click="handleClick">結果を送信</span>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { ais } from './../util'

@Component
export default class PostForm extends Vue {
  @Prop()
  private score!: number
  @Prop()
  private init!: () => void

  private name: string = ''

  private async handleClick(): Promise<any> {
    const { name, score } = this
    ais.post('user', { name, score })
    this.init()
  }
}
</script>

<style scoped lang="scss">
div {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: rgba(255, 255, 255, 0.8);
  color: #333;
  position: absolute;
  top: 0;
  left: 0;
  p {
    font-size: 80px;
    font-weight: bold;
  }
  input {
    font-size: 50px;
    text-align: center;
  }
  span {
    display: block;
    font-size: 40px;
    background-color: lime;
    margin: 100px 0;
    padding: 5px 50px;
    border-radius: 20px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.5);
    cursor: pointer;
  }
}
</style>
