<template>
  <div class="container">
    <ul id="clusters" class="list-none">
      <li v-for="item in clusters" :key="item.metadata.uid">
        <div>
          <h3>
            {{ item.metadata.name }}
          </h3>
          <h4>
            {{ item.spec.infrastructureRef.kind }}
          </h4>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      clusters: []
    };
  },

  mounted: function() {
    window.wails.Events.On("clusters", clusters => {
      console.log(clusters);

      if (clusters) {
        this.clusters = clusters;
      }
    });
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
