<template>
  <Toolbar location="/clusters"></Toolbar>
</template>

<script>
import Vue from "vue";
import Toolbar from "@/components/Toolbar.vue";

export default Vue.extend({
  name: "NewClusterToolbar",

  components: {
    Toolbar
  }
});
</script>

<style scoped>
</style>
