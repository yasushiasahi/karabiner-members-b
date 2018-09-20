<template>
  <div id="app">
    <div v-if="users.length === 0"><p>なうろーでぃんぐ</p></div>
    <Top v-if="users.length !== 0 && !isPlaying"  v-bind:users="users" v-bind:togglePlay="togglePlay"/>
    <div class="play" v-if="isPlaying">
      <p>{{idx}}/{{members.length}}</p>
      <MemberImg v-bind:img="members[idx].img"/>
      <div v-if="users.length !== 0">SCORE : {{ score }}</div>
      <MemberNames v-bind:names="names" v-bind:handleClick="handleClick"/>
      <Answer v-if="isAnswerShown" v-bind:isCollect="isCollect" v-bind:init="init"/>
    <PostForm v-if="isPostFormShown" v-bind:score="score" v-bind:init="init" />
    </div>
  </div>
</template>


<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator'
import { waitSecond, ais, initNames, initMembers } from './util'
import { Member, User } from './types'
import MemberImg from './components/MemberImg.vue'
import MemberNames from './components/MemberNames.vue'
import Answer from './components/Answer.vue'
import PostForm from './components/PostForm.vue'
import Top from './components/Top.vue'

@Component({
  components: {
    MemberImg,
    MemberNames,
    Answer,
    PostForm,
    Top,
  },
})
export default class App extends Vue {
  private members: Member[] = []
  private names: string[] = []
  private users: User[] = []
  private idx: number = 0
  private isCollect: boolean = false
  private isAnswerShown: boolean = false
  private isPostFormShown: boolean = false
  private isPlaying: boolean = false
  private score: number = 0

  private created(): void {
    this.init()
  }

  private async init(): Promise<any> {
    const ms = await ais.get('members')
    if (ms !== null) {
      this.names = initNames(ms)
      this.members = initMembers(ms)
    }

    const us = await ais.get('users')
    if (us !== null) {
      this.users = us
    }

    this.isCollect = false
    this.isAnswerShown = false
    this.isPostFormShown = false
    this.isPlaying = false
    this.score = 0
    this.idx = 0
  }

  private togglePlay(): void {
    this.isPlaying = !this.isPlaying
  }

  private async handleClick(name: string): Promise<any> {
    const { members, idx } = this

    this.isCollect = members[idx].name === name
    if (this.isCollect) {
      const i = this.names.findIndex((n) => n === name)
      this.names.splice(i, 1)
      this.score++
    }

    this.isAnswerShown = true

    const timeoutId = await waitSecond(0.5)
    window.clearTimeout(timeoutId)

    this.isAnswerShown = false

    if (idx + 1 === members.length) {
      this.isPostFormShown = true
      return
    }

    this.idx++
  }
}
</script>


<style lang="scss">
body {
  width: 100%;
  height: 100%;
  margin: 0;
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
}
.play {
  position: relative;
}
</style>
