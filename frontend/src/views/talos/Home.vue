<template>
  <div>
    <sui-card-group :items-per-row="3">
      <sui-card v-for="(item, key) in items" :key="item">
        <sui-card-content>
          <sui-card-header>
            {{ key }}
          </sui-card-header>
        </sui-card-content>
        <sui-card-content>
          <sui-card-meta>Endpoint: {{ item.server }}</sui-card-meta>
        </sui-card-content>
        <sui-card-content extra>
          <sui-button v-on:click="openTalosControlPlane($event, item.contextName)">Talos Control Plane</sui-button>
        </sui-card-content>
      </sui-card>
    </sui-card-group>
  </div>
</template>

<script>
import Vue from "vue";
import router from "@/router";

export default Vue.extend({
  name: "Clusters",

  data() {
    return {
      items: []
    };
  },

  created: function() {
    window.wails.Events.On("clusters", clusters => {
      if (clusters) {
        this.items = clusters;
      }
    });
  },

  mounted: function() {
    window.backend.Clusters.Clusters().then(clusters => {
      this.items = clusters;
    });
  },

  methods: {
    openTalosControlPlane(e, item) {
      router.push(`/cluster/${item}`);
    }
  }
});
</script>
