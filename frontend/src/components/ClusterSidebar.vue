<template>
  <Sidebar
    header="cluster"
    :sections="[
      { name: 'machines', to: `/clusters/${$route.params.cluster}/machines` },
      { name: 'pools', to: `/clusters/${$route.params.cluster}/pools` }
    ]"
  ></Sidebar>
</template>

<script>
import Vue from "vue";
import Sidebar from "@/components/Sidebar.vue";

export default Vue.extend({
  name: "ClusterSidebar",

  components: {
    Sidebar
  }
});
</script>

<style scoped>
</style>
