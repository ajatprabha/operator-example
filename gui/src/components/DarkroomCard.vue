<template lang="pug">
  div
    v-toolbar(:color='statusColor("deployed")' dark)
      v-btn.hidden-xs-only(icon)
        v-icon mdi-image-search-outline
      v-toolbar-title {{ darkroom.name }}
      v-spacer
      v-btn(:to="{ name: 'source-detail', params: { name: darkroom.name } }" small outlined)
        | View
        v-icon(right='') mdi-share
    v-card
      v-card-text
        p Domains:
        ul
          li(v-for='(domain, i) in darkroom.status.domains' :key='i') {{ domain }}
      v-card-actions.pa-3
        v-icon(:color='statusColor("deployed")' small) mdi-checkbox-blank-circle
        span.ml-2 {{ "deployed" | capitalize }}
        v-spacer
          | {{ darkroom.type }}
          v-icon.ml-2(small='') {{ iconClass(darkroom.type) }}
</template>

<script>
import statusColorMap from "./utils/statusColorMap";
import iconClassMap from "./utils/iconClassMap";

export default {
  name: "DarkroomCard",
  props: {
    darkroom: { type: Object, required: true }
  },
  methods: {
    statusColor: val => statusColorMap.get(val),
    iconClass: val => iconClassMap.get(val)
  }
};
</script>

<style scoped></style>
