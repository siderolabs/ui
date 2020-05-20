<template>
  <div id="app">
    <Clusters />
    <Machines />
  </div>
</template>

<script>
import Clusters from "./components/Clusters.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    Clusters
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

html {
  height: 100%;
  overflow: hidden;
  background-color: white;
  background-size: 20px 20px;
}
</style>
