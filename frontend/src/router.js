import Vue from "vue";
import VueRouter from "vue-router";
import Home from "@/views/talos/Home.vue";
import HomeSidebar from "@/components/talos/HomeSidebar.vue";
import Cluster from "@/views/talos/Cluster.vue";
import ClusterSidebar from "@/components/talos/ClusterSidebar.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "",
    name: "Home",
    components: {
      default: Home,
      sidebar: HomeSidebar,
    },
  },
  {
    path: "/cluster/:cluster",
    name: ":cluster",
    components: {
      default: Cluster,
      sidebar: ClusterSidebar,
    },
    meta: {
      breadcrumbs: [
        {text: "Home", to: "/"},
      ]
    }
  },
];

const router = new VueRouter({
  mode: "abstract",
  base: process.env.BASE_URL,
  routes,
});

export default router;
