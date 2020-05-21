import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    showSidebar: false,
  },

  mutations: {
    showSidebar(state, b) {
      state.showSidebar = b;
    },

    toggleSidebar(state) {
      state.showSidebar = !state.showSidebar;
    },
  },

  actions: {},

  modules: {},
});
