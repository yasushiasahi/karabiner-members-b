<template>
  <div id="app">
    <img class="logo" alt="カラビナロゴ" src="https://www.karabiner.tech/asset/images/index/logo_main.svg">
    <div v-if="members.length !== 0">
      <MemberImg v-bind:img="members[idx].img"/>
      <MemberNames v-bind:members="members" v-bind:handleClick="handleClick"/>
    </div>
    <div v-else>
      <p>なうろーでぃんぐ</p>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { Member } from './types'
import MemberImg from './components/MemberImg.vue'
import MemberNames from './components/MemberNames.vue'
import axios, { AxiosResponse, AxiosError } from 'axios'

const instance = axios.create({
  baseURL: 'http://localhost:3000/api/',
})

@Component({
  components: {
    MemberImg,
    MemberNames,
  },
})
export default class App extends Vue {
  private members: Member[] = []
  private idx: number = 0
  private isCollect: boolean = false
  private score: number = 0

  private mounted(): void {
    instance
      .get('members')
      .then((res: AxiosResponse) => {
        this.members = res.data
      })
      .catch((err: AxiosError) => {
        this.members = []
      })
  }

  private handleClick(name: string): void {
    const { members, idx } = this
    if (members[idx].name === name) {
      this.isCollect = true
      this.score++
    } else {
      this.isCollect = false
    }
    this.idx++
  }
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.logo {
  width: 30%;
}
</style>
