<template lang="pug">
  v-row
    v-col(cols="12")
      v-row.px-6(justify="space-between")
        h4.display-1 All Sources
        v-btn.white--text(color="secondary" :to="{ name: 'add-source' }")
          v-icon(left dark) mdi-plus
          | New Source
    v-row.px-6
      v-col.pa-2(
        v-for="(darkroom, i) in darkrooms"
        :key="`source-${i}`"
        xs="12"
        sm="6"
        xl="4"
      )
        DarkroomCard(:darkroom="darkroom")
      v-col.text-sm-center(v-if="!darkrooms.length")
        h5.headline.grey--text No sources yet!
</template>

<script>
import DarkroomCard from "../components/DarkroomCard";

export default {
  name: "SourceList",
  title() {
    return this.$options.name;
  },
  components: { DarkroomCard },
  data: () => ({
    darkrooms: []
  }),
  created() {
    this.$api
      .getAllDarkrooms()
      .then(list => (this.darkrooms = list.items))
      .catch(e => console.log(e));
  }
};
</script>

<style scoped></style>
