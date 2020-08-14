import "core-js/stable";
import "regenerator-runtime/runtime";
import Vue from "vue";
import App from "./App.vue";
import router from "./router.js";
import store from "./store";
import SuiVue from 'semantic-ui-vue';

Vue.use(SuiVue);

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from "@wailsapp/runtime";

router.beforeEach((to, from, next) => {
  if(to.meta.breadcrumbs) {
    store.commit('set', 
      [...to.meta.breadcrumbs,
       {text: to.name}]
    );
  } else {
    store.commit('set', 
       [{text: to.name}]
    );
  }

  next();
});

Wails.Init(() => {
  new Vue({
    router,
    store,
    render: (h) => h(App),
    mounted() {
      this.$router.replace('/');
    },
  }).$mount("#app");
});
