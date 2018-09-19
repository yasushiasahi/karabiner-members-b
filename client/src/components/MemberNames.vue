<template>
  <div>
    <span v-for="member in members" v-bind:key="member.name" v-on:click="() => handleClick(member.name)">
      {{ member.name }}
    </span>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Emit, Vue } from 'vue-property-decorator'
import { Member } from './../types'

@Component
export default class MemberNames extends Vue {
  @Prop()
  private members!: Member[]

  @Emit('handleClick')
  handleClick(name: string): void {}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
span {
  display: inline-block;
  padding: 3px 10px;
  margin: 5px;
  font-size: 20px;
  background-color: lime;
  border-radius: 15px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.5);
  cursor: pointer;
}
</style>
