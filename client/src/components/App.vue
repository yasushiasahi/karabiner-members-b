<template>
  <div class="top">
    <template v-if="users.length">
      <div v-for="user in users">
        <div>{{user.id}}</div>
        <div>{{user.name}}</div>
        <div>{{user.pass}}</div>
        <div>{{user.token}}</div>
      </div>
    </template>
    <h1 v-else>データを読み込み中</h1>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import axios from 'axios'

const instance = axios.create({
  baseURL: 'http://127.0.0.1/api/',
})

interface User {
  id: number
  name: string
  pass: string
  token: string
}

@Component
export default class HelloWorld extends Vue {
  @Prop()
  private msg!: string

  private users: User[] = []

  private mounted(): void {
    instance.get('user').then((res) => {
      this.users = res.data
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.top {
  background-color: wheat;
  min-height: 100vh;
}
</style>
