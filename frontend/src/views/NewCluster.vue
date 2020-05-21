<template>
  <div class="w-1/3 text-sm">
    <h1 class=" mb-4 font-bold">Cluster Configuration</h1>
    <div class="mb-2">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
        Name
      </label>
      <input
        v-model="name"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        type="text"
        name="name"
        id=""
      />
    </div>
    <div class="mb-2">
      <div class="mb-2">
        <label
          class="block text-gray-700 text-sm font-bold mb-2"
          for="username"
        >
          Pod Subnet
        </label>
        <input
          v-model="podSubnet"
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          type="text"
          name="podsubnet"
          id=""
        />
      </div>
      <div class="mb-2">
        <label
          class="block text-gray-700 text-sm font-bold mb-2"
          for="username"
        >
          Provider
        </label>
        <select v-model="provider" name="provider" id="">
          <option value="" disabled selected>Choose a provider</option>
          <option value="NewAWSCluster">AWS</option>
          <!-- <option value="NewAzureCluster">Azure</option> -->
          <!-- <option value="NewGCPCluster">GCP</option> -->
          <!-- <option value="NewMetalCluster">Metal</option> -->
        </select>
      </div>
    </div>
    <div>
      <h1 v-if="provider" class="my-4 mb-4 font-bold">
        Provider Configuration
      </h1>
      <component
        v-bind:is="provider"
        v-bind:name="name"
        v-bind:podSubnet="podSubnet"
      ></component>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import NewAWSCluster from "@/components/NewAWSCluster.vue";

export default Vue.extend({
  name: "NewCluster",

  components: {
    NewAWSCluster
  },

  data() {
    return {
      provider: null,
      name: "",
      podSubnet: "192.168.0.0/16"
    };
  }
});
</script>

<style scoped>
</style>
