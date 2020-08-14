<template>
  <div>
    <h2>Nodes</h2>
    <list-view 
      :contextName="this.$route.params.cluster"
      idField="metadata.name"
      runtimeObject="v1.Node"
      >
      <template v-slot:header>
        <sui-table-header-cell>Name</sui-table-header-cell>
        <sui-table-header-cell>OS Image</sui-table-header-cell>
        <sui-table-header-cell>Kubelet Version</sui-table-header-cell>
      </template>
      <template v-slot:item="props">
        <sui-table-row>
          <sui-table-cell>{{props.item.metadata.name}}</sui-table-cell>
          <sui-table-cell>{{props.item.status.nodeInfo.osImage}}</sui-table-cell>
          <sui-table-cell>{{props.item.status.nodeInfo.kubeletVersion}}</sui-table-cell>
        </sui-table-row>
      </template>
    </list-view>
  </div>
</template>

<script>
import Vue from "vue";
import ListView from "@/components/ListView.vue";

export default Vue.extend({
  name: "Cluster",

  components: {
    ListView
  },
});
</script>
