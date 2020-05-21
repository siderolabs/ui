<template>
  <div id="app">
    <div class="grid">
      <div class="navbar">
        <Navbar />
      </div>
      <div class="navpanel">
        <Navpanel />
      </div>
      <div class="sidebar">
        <router-view name="sidebar"></router-view>
      </div>
      <div class="main">
        <div class="subgrid">
          <div class="toolbar">
            <router-view name="toolbar"></router-view>
          </div>
          <div class="content">
            <div class="w-3/4 m-auto">
              <router-view></router-view>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import Navpanel from "@/components/Navpanel.vue";
import Navbar from "@/components/Navbar.vue";
import "./assets/css/main.css";

export default Vue.extend({
  name: "app",
  components: {
    Navbar,
    Navpanel
  }
});
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

.grid {
  display: grid;
  grid-template-rows: theme("spacing.12") auto;
  grid-template-columns: auto 1fr;
  height: 100vh;
}

.navbar {
  @apply z-50;
  grid-column-start: 1;
  grid-column-end: 3;
}

.navpanel {
  @apply z-40;
  grid-row-start: 2;
  grid-row-end: 3;
}

.sidebar {
  @apply z-30;
  width: 200px;
  grid-row-start: 2;
  grid-row-end: 3;
  grid-column-start: 1;
  grid-column-end: 2;
}

.main {
  grid-column-start: 2;
  grid-row-start: 2;
}

.subgrid {
  display: grid;
  grid-gap: theme("spacing.10");
  grid-template-rows: theme("spacing.12") auto;
  grid-template-columns: 1fr;
  height: 100%;
}

.toolbar {
  grid-row-start: 1;
}

.content {
  grid-row-start: 2;
}
</style>
