<template>
  <table class="table-fixed text-left w-full">
    <thead class="text-sm">
      <tr>
        <th class="px-4 py-2">UUID</th>
        <th class="px-4 py-2">Manufacturer</th>
        <th class="px-4 py-2">CPU</th>
        <th class="px-4 py-2">Memory</th>
        <th class="px-4 py-2">Storage</th>
      </tr>
    </thead>
    <tbody class="text-sm">
      <tr
        v-for="item in items"
        :key="item.metadata.name"
        class="cursor-pointer"
      >
        <td class="border px-4 py-2 text-sm">{{ item.metadata.name }}</td>
        <td class="border px-4 py-2 text-sm">
          {{ item.spec.system.manufacturer }}
        </td>
        <td class="border px-4 py-2 text-sm">{{ item.spec.cpu.version }}</td>
        <td class="border px-4 py-2 text-sm"></td>
        <td class="border px-4 py-2 text-sm"></td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  name: "Servers",

  data() {
    return {
      items: []
    };
  },

  created: function() {
    window.wails.Events.On("servers", servers => {
      if (servers) {
        this.items = servers;
      }
    });
  },

  mounted: function() {
    window.backend.Servers.Servers().then(servers => {
      this.items = servers;
    });
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
