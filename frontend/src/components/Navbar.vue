<template>
  <div class="bg-white border-solid border-b shadow-md">
    <ul class="flex justify-between items-center h-full mx-2 p-0 list-none">
      <Burger></Burger>
      <li>
        <img
          class="h-12 p-1"
          alt="Company logo"
          src="@/assets/images/logo.png"
        />
      </li>
      <li></li>
    </ul>
  </div>
</template>

<script>
import Vue from "vue";
import Burger from "./Burger.vue";

export default Vue.extend({
  name: "Navbar",
  components: {
    Burger
  }
});
</script>

<style scoped>
</style>
