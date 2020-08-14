<template>
  <sui-breadcrumb>
    <template v-for="crumb in crumbs">
      <sui-breadcrumb-section link :to="crumb.to" is="router-link" :key="crumb.text">{{ processText(crumb.text) }}</sui-breadcrumb-section>
      <sui-breadcrumb-divider icon="right angle" :key="crumb.text"/>
    </template>
    <template v-if="active">
      <sui-breadcrumb-section active>
        {{ processText(active.text) }}
      </sui-breadcrumb-section>
    </template>
  </sui-breadcrumb>
</template>

<script>

import Vue from "vue";
import { mapState } from 'vuex';

export default Vue.extend({
  name: "Breadcrumb",

  methods: {
    processText(text) {
      if(text[0] == ':') {
        return this.$route.params[text.slice(1, text.length)];
      }

      return text;
    }
  },

  computed: {
    crumbs() {
      return this.breadcrumbs.slice(0, this.breadcrumbs.length - 1);
    },
    active() {
      return this.breadcrumbs[this.breadcrumbs.length - 1];
    },
    ...mapState({
      breadcrumbs: state => state.breadcrumbs
    })
  }
});
</script>
