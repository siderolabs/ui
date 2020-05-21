<template>
  <div class="flex flex-col items-left justify-between">
    <div class="mb-2">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
        Region
      </label>
      <input
        v-model="region"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        type="text"
        name="region"
        id=""
        placeholder="Region"
      />
    </div>
    <div class="mb-2">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
        Instance Type
      </label>
      <input
        v-model="instanceType"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        type="text"
        name="instance-type"
        id=""
        placeholder="Instance Type"
      />
    </div>
    <div class="mb-2">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
        Subnet
      </label>
      <input
        v-model="subnet"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        type="text"
        name="subnet"
        id=""
        placeholder="Subnet ID"
      />
    </div>
    <div class="mb-2">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
        SSH Key
      </label>
      <input
        v-model="sshKeyName"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        type="text"
        name="ssh-key-name"
        id=""
        placeholder="SSH Key Name"
      />
    </div>
    <div class="mb-2">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
        VPC
      </label>
      <input
        v-model="vpc"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        type="text"
        name="vpc"
        id=""
        placeholder="VPC"
      />
    </div>
    <div class="mb-2">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
        Security Group
      </label>
      <input
        v-model="securityGroup"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        type="text"
        name="security-group"
        id=""
        placeholder="Security Group"
      />
    </div>
    <div class="mb-2">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
        AMI
      </label>
      <input
        v-model="ami"
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        type="text"
        name="ami"
        id=""
        placeholder="AMI"
      />
    </div>
    <div class="mt-5">
      <button
        v-on:click="create"
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 border border-blue-700 rounded"
        type="button"
      >
        Create
      </button>
    </div>
  </div>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  name: "NewAWSCluster"
});
</script>

<style scoped>
.form-text {
  @apply shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight;
}
</style>
