<template>
  <div class="container">
    <div class="wrapper">
    <h1>カラビナメンバーズ</h1>
    </div>
    <img alt="カラビナロゴ" src="https://www.karabiner.tech/asset/images/index/logo_main.svg">
    <ul>
    <h2>ランキング</h2>
      <li v-for=" u in users"><span>{{ u.name }}</span><span>SCORE: {{ u.score }}</span></li>
    </ul>
    <div class="wrapper">
    <span v-on:click="togglePlay">スタート</span>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { User } from './../types'

@Component
export default class Top extends Vue {
  @Prop()
  private users!: User[]
  @Prop()
  private togglePlay!: () => void
}
</script>

<style scoped lang="scss">
.container {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.1);
}

.wrapper {
  width: 100%;
  span {
    display: inline-block;
    font-size: 40px;
    background-color: lime;
    margin: 100px 0;
    padding: 5px 50px;
    border-radius: 20px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.5);
    cursor: pointer;
  }
}

ul {
  width: 280px;
  li {
    list-style: none;
    display: flex;
    justify-content: space-between;
    span {
      font-size: 25px;
    }
  }
}
</style>
