<template lang="pug">
  div
    v-row
      v-col(cols="12")
        h1.display-1.mb-4.font-weight-medium New Source
        v-divider
      v-col
        v-row
          v-col(cols="12")
            h3.headline General
            v-row
              v-col(cols="12" md="6")
                v-text-field(
                  v-model="name"
                  label="Source Name"
                  outlined
                )
              v-col(cols="12" md="6")
                v-select(
                  v-model="sourceType"
                  label="Source Type"
                  :items="SOURCE_ITEMS"
                  :prepend-icon="iconClass(sourceType)"
                  outlined
                )
            v-divider.mb-6
            h3.headline {{ SOURCE_ITEMS.find(el => el.value === sourceType).text }} Settings
            v-row(v-if="sourceType === 'WebFolder'")
              v-col(cols="12" md="8")
                v-text-field(
                  v-model="baseUrl"
                  label="Base URL"
                  placeholder="https://example.com/path/to/images"
                  hide-details="auto"
                  outlined
                )
            v-row(v-else-if="sourceType === 'S3'")
              v-col(cols="12" sm="12" md="6")
                v-text-field(
                  v-model="accessKey"
                  label="Access Key ID"
                  hide-details="auto"
                  dense
                  outlined
                )
              v-col(cols="12" sm="12" md="6")
                v-text-field(
                  v-model="secretKey"
                  label="Secret Access Key"
                  hide-details="auto"
                  dense
                  outlined
                )
              v-col(cols="12" sm="12" md="6")
                v-text-field(
                  v-model="bucket"
                  label="S3 Bucket"
                  hide-details="auto"
                  dense
                  outlined
                )
              v-col(cols="12" sm="12" md="6")
                v-text-field(
                  v-model="pathPrefix"
                  label="Path Prefix"
                  placeholder="path/to/images"
                  hide-details="auto"
                  dense
                  outlined
                )
            v-divider.my-6
            h3.headline Domains
            v-row
              v-col(cols="12" md="6")
                v-combobox(
                  v-model="subDomains"
                  label="Darkroom SubDomains"
                  multiple
                  chips
                  outlined
                  hint="Add subdomain and press enter"
                  append-icon=""
                )
            v-divider.mb-6
            h3.headline Cache Settings
            v-row
              v-col(cols="12" md="6")
                v-select(
                  v-model="cacheTTL"
                  :items="CACHE_TTL_BEHAVIOR"
                  label="Cache TTL Behavior"
                  outlined
                  hide-details="auto"
                )
              v-col(cols="12" md="6")
                v-text-field(
                  v-model="defaultCache"
                  type="number"
                  label="Default Cache TTL (seconds)"
                  outlined
                  hide-details="auto"
                )
        v-row(justify="end")
          v-col(cols="auto")
            v-btn.ma-2(large color="error" :to="{ name: 'sources' }")
              v-icon(left) mdi-cancel
              | Cancel
            v-btn.ma-2(large color="primary" @click="doCreate")
              v-icon(left) mdi-content-save-outline
              | Save
</template>

<script>
import iconClassMap from "../components/utils/iconClassMap";

export default {
  name: "AddSource",
  title: "Add New Source",
  data: () => ({
    name: "",
    sourceType: "WebFolder",
    baseUrl: "",
    accessKey: "",
    secretKey: "",
    bucket: "",
    pathPrefix: "",
    subDomains: [],
    cacheTTL: "RespectOrigin",
    defaultCache: 0
  }),
  computed: {
    SOURCE_ITEMS: () => [
      { text: "Web Folder", value: "WebFolder" },
      { text: "Amazon S3", value: "S3" }
    ],
    CACHE_TTL_BEHAVIOR: () => [
      { text: "Respect Origin", value: "RespectOrigin" },
      { text: "Override Origin", value: "OverrideOrigin" },
      { text: "Enforce Minimum", value: "EnforceMinimum" }
    ]
  },
  methods: {
    iconClass: val => iconClassMap.get(val),
    doCreate() {
      this.$api.createDarkroom(this.payload()).then(() => {
        this.$router.replace({ name: "sources" });
      });
    },
    payload() {
      return this.sourceType === "WebFolder"
        ? {
            name: this.name,
            source: {
              type: "WebFolder",
              baseUrl: this.baseUrl
            },
            subDomains: this.subDomains
          }
        : {
            name: this.name,
            source: {
              type: "S3",
              accessKey: this.accessKey,
              secretKey: this.secretKey,
              region: this.region,
              pathPrefix: this.pathPrefix
            },
            subDomains: this.subDomains
          };
    }
  }
};
</script>

<style scoped></style>
