<template>
  <div>
    <router-link to="/new">
      <div>
        <button
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 border border-blue-700 rounded"
          type="button"
        >
          New
        </button>
      </div>
    </router-link>
  </div>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  name: "ClustersToolbar"
});
</script>

<style scoped>
</style>
