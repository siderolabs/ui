<template>
  <table class="table-fixed text-left w-full">
    <thead class="text-sm">
      <tr>
        <th class="px-4 py-2">Name</th>
        <th class="px-4 py-2">Channel</th>
        <th class="px-4 py-2">Version</th>
        <th class="px-4 py-2">Size</th>
        <th class="px-4 py-2">Concurrency</th>
      </tr>
    </thead>
    <tbody class="text-sm">
      <tr
        v-for="item in items"
        :key="item.metadata.name"
        class="cursor-pointer"
      >
        <td class="border px-4 py-2 text-sm">{{ item.metadata.name }}</td>
        <td class="border px-4 py-2 text-sm">{{ item.spec.channel }}</td>
        <td class="border px-4 py-2 text-sm">{{ item.status.version }}</td>
        <td class="border px-4 py-2 text-sm">{{ item.status.size }}</td>
        <td class="border px-4 py-2 text-sm">{{ item.spec.concurrency }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  name: "Pools",

  data() {
    return {
      items: []
    };
  }
});
</script>

<style scoped>
.sidebar {
  @apply h-full w-1/5 mt-0 mb-0;
}

.section-header {
  @apply pt-3 pb-3 uppercase;
}

.section {
  @apply ml-2 pt-2 pb-2 capitalize;
}

table,
th,
td {
  border-left: none;
  border-right: none;
}
</style>
