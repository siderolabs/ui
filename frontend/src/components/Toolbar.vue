<template>
  <div class="navbar">
    <div>
      <button
        v-on:click="back"
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 border border-blue-700 rounded"
        type="button"
      >
        Back
      </button>
    </div>
    <div>
      <slot></slot>
    </div>
    <div>
      <span>{{ title }}</span>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import router from "@/router";

export default Vue.extend({
  name: "Toolbar",

  props: {
    title: {
      type: String,
      required: false
    },

    location: {
      type: String,
      required: true
    }
  },

  methods: {
    back() {
      router.push(this.location);
    }
  }
});
</script>

<style scoped>
.navbar {
  @apply flex justify-between items-center content-between h-12 px-20 border-solid border-b;
}
</style>
