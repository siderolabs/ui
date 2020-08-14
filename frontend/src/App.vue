<template>
  <div id="app">
    <div class="main-menu">
      <sui-menu secondary >
        <sui-menu-item>
          <sui-button basic size="mini" icon="bars" v-on:click="toggleSidebar"/>
        </sui-menu-item>
        <sui-menu-item>
          <breadcrumb/>
        </sui-menu-item>
      </sui-menu>
    </div>
    <div class="content">
      <sui-sidebar-pushable>
        <sui-menu
          is="sui-sidebar"
          :visible="sidebarVisible"
          animation="overlay"
          width="thin"
          inverted
          vertical
          pointing
          style="overflow: hidden"
        >
          <router-link is="sui-menu-item" to="/"> <sui-icon name="home" /> Home </router-link>
          <router-view name="sidebar"/>
        </sui-menu>
        <sui-sidebar-pusher class="router-view">
          <router-view/>
        </sui-sidebar-pusher>
      </sui-sidebar-pushable>
    </div>
  </div>
</template>

<script>
  import 'semantic-ui-css/semantic.min.css';
  import Vue from "vue";
  import Breadcrumb from "@/components/Breadcrumb.vue"

  export default Vue.extend({
    name: "app",

    components: {
      Breadcrumb,
    },

    data() {
      return {
        sidebarVisible: false,
      }
    },

    methods: {
      toggleSidebar() {
        this.sidebarVisible = !this.sidebarVisible;
      },
    },
  });
</script>

<style>
.main-menu {
  background: linear-gradient(0deg, rgba(241,241,241,1) 0%, rgba(254,254,254,1) 100%);
  border-bottom: 1px solid #cdcdcd;
}

.content {
  border-top: 1px solid #ffffff;
  background: rgba(241,241,241,1);
}

.router-view {
  padding: 1em;
  height: 100%;
}

#app {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}

#app > .main-menu {
  flex-grow: 0;
}

#app > .content {
  flex-grow: 1;
  position: relative;
}

#app > .content > div {
  width: 100%;
  height: 100%;
  position: absolute;
}

</style>
