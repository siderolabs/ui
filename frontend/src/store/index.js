// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with This
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    breadcrumbs: [],
  },

  mutations: {
    set(state, breadcrumbs) {
      state.breadcrumbs = breadcrumbs;
    },
    push(state, breadcrumb) {
      state.breadcrumbs.push(breadcrumb)
    },
    pop(state) {
      state.breadcrumbs.pop();
    },
    replace(state, payload) {
      const index = state.breadcrumbs.findIndex((breadcrumb) => {
        return breadcrumb.text === payload.find;
      });

      if (index) {
        state.breadcrumbs.splice(index, 1, payload.replace);
      }
    },
    empty(state) {
      state.breadcrumbs = [];
    }
  },

  actions: {},

  modules: {},
});
