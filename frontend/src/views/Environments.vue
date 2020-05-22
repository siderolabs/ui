<template>
  <table class="table-fixed text-left w-full">
    <thead class="text-sm">
      <tr>
        <th class="px-4 py-2">Name</th>
        <th class="px-4 py-2">Kernel</th>
        <th class="px-4 py-2">Initrd</th>
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
          {{ item.spec.kernel.url }}
        </td>
        <td class="border px-4 py-2 text-sm">{{ item.spec.initrd.url }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  name: "Environments",

  data() {
    return {
      items: []
    };
  },

  created: function() {
    window.wails.Events.On("environments", environments => {
      if (environments) {
        this.items = environments;
      }
    });
  },

  mounted: function() {
    window.backend.Environments.Environments().then(environments => {
      this.items = environments;
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
