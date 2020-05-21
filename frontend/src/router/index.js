import Vue from "vue";
import VueRouter from "vue-router";
import Home from "@/views/Home.vue";
import Clusters from "@/views/Clusters.vue";
import Cluster from "@/views/Cluster.vue";
import ClusterSidebar from "@/components/ClusterSidebar.vue";
import ClusterToolbar from "@/components/ClusterToolbar.vue";
import Inventory from "@/views/Inventory.vue";
import InventorySidebar from "@/components/InventorySidebar.vue";
import Servers from "@/views/Servers.vue";
import ServerClasses from "@/views/ServerClasses.vue";
import Toolbar from "@/components/Toolbar.vue";
import Machines from "@/views/Machines.vue";
import Pools from "@/views/Pools.vue";
import NewCluster from "@/views/NewCluster.vue";
import NewClusterToolbar from "@/components/NewClusterToolbar.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
  },
  {
    path: "/inventory",
    name: "inventory",
    component: Inventory,
    components: {
      default: Inventory,
      sidebar: InventorySidebar,
    },
    children: [
      {
        path: "/servers",
        name: "servers",
        component: Servers,
      },
      {
        path: "/serverclasses",
        name: "server classes",
        component: ServerClasses,
      },
    ],
  },
  {
    path: "/clusters",
    components: {
      default: Clusters,
      toolbar: Toolbar,
    },
  },
  {
    path: "/clusters/:cluster",
    name: "cluster",
    components: {
      default: Cluster,
      sidebar: ClusterSidebar,
      toolbar: ClusterToolbar,
    },
    children: [
      {
        path: "/machines",
        name: "machines",
        component: Machines,
      },
      {
        path: "/pools",
        name: "pools",
        component: Pools,
      },
    ],
  },
  {
    path: "/new",
    name: "new cluster",
    components: {
      default: NewCluster,
      toolbar: NewClusterToolbar,
    },
  },
];

const router = new VueRouter({
  mode: "abstract",
  base: process.env.BASE_URL,
  routes,
});

export default router;
