<template>
  <Toolbar :title="$route.params.cluster" location="/clusters"></Toolbar>
</template>

<script>
import Vue from "vue";
import Toolbar from "@/components/Toolbar.vue";

export default Vue.extend({
  name: "ClusterToolbar",

  components: {
    Toolbar
  }
});
</script>

<style scoped>
</style>
