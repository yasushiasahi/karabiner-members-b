<template>
  <div>
    <img v-bind:src="img" alt="" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component
export default class MemberImg extends Vue {
  @Prop()
  private img!: string
}
</script>

<style scoped lang="scss">
</style>
