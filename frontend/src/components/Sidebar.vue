<template>
  <div class="sidebar">
    <ul>
      <li class="section-header">
        <span class="pl-4">{{ header }}</span>
      </li>
      <li v-for="section in sections" :key="section.to" class="section">
        <router-link :to="section.to">{{ section.name }}</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  name: "Sidebar",

  props: {
    header: {
      required: true,
      type: String
    },

    sections: {
      required: true,
      type: Array
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
  @apply h-full w-full z-30 bg-gray-100 border-r;
}

.section-header {
  @apply h-12 pt-3 pb-3 uppercase border-t border-b w-full;
}

.section {
  @apply pt-2 pb-2 capitalize;
}
</style>
