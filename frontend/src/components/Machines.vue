<template>
  <div class="container">
    <ul id="machines">
      <li v-for="item in machines" :key="item.metadata.uid">
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
      machines: []
    };
  },

  mounted: function() {
    window.wails.Events.On("machines", machines => {
      console.log(machines);

      if (clusters) {
        this.machines = [machines];
      }
    });
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
