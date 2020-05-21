<template>
  <Sidebar
    header="inventory"
    :sections="[
      { name: 'servers', to: `/servers` },
      {
        name: 'server classes',
        to: `/serverclasses`
      }
    ]"
  ></Sidebar>
</template>

<script>
import Vue from "vue";
import Sidebar from "@/components/Sidebar.vue";

export default Vue.extend({
  name: "InventorySidebar",

  components: {
    Sidebar
  }
});
</script>

<style scoped>
</style>
