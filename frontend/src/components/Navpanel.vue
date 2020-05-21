<template>
  <transition name="slide">
    <div @click="close" v-if="isActive" class="sidebar">
      <ul>
        <li class="section"><router-link to="/home">home</router-link></li>
        <li class="section-header"><span class="pl-4">compute</span></li>
        <ul>
          <li class="section">
            <router-link to="/inventory">inventory</router-link>
          </li>
          <li class="section">
            <router-link to="/clusters">clusters</router-link>
          </li>
        </ul>
      </ul>
    </div>
  </transition>
</template>

<script>
import Vue from "vue";
import store from "@/store";

export default Vue.extend({
  name: "Sidebar",

  computed: {
    isActive: {
      get() {
        return store.state.showSidebar;
      },

      set(value) {
        store.commit("showSidebar", value);
      }
    }
  },

  methods: {
    close() {
      store.commit("showSidebar", false);
    }
  }
});
</script>

<style scoped>
a {
  @apply pl-4;
  display: inline-block;
  width: 100%;
}

ul {
  @apply list-none;
}

.sidebar {
  @apply fixed top-0 left-0 h-full mt-12 z-40 bg-white shadow-lg;
  width: 300px;
}

.section-header {
  @apply pt-3 pb-3 uppercase w-full;
}

.section {
  @apply pt-2 pb-2 capitalize;
}

.slide-enter-active,
.slide-leave-active {
  transition: transform 100ms ease;
}

.slide-enter,
.slide-leave-to {
  transform: translateX(-100%);
  transition: all 100ms ease-in 0s;
}
</style>
