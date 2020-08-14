<template>
  <sui-message 
    error 
    v-if="err" 
    header="Load failed"
    :content="err"
  />
  <sui-message 
    v-else-if="!err && items.length == 0 && !loading" 
  >
    No entities
  </sui-message>
  <sui-segment v-else>
    <div style="padding: 1em" v-if="loading">
      <sui-loader active/>
    </div>
    <sui-table single-line v-else>
      <sui-table-header>
        <sui-table-row>
          <slot name="header"></slot> 
        </sui-table-row>
      </sui-table-header>
      <sui-table-body>
        <slot v-for="(item, index) in items" :key="index" name="item" :item="item"></slot>
      </sui-table-body>
    </sui-table>
  </sui-segment>
</template>

<script>
import Vue from "vue";
import wailsBackend from "@/mixins/wailsBackend.js";

export default Vue.extend({
  name: "ListView",
  mixins: [wailsBackend],

  props: {
    runtimeObject: {
      type: String,
      required: true
    },
    contextName: {
      type: String,
      required: true
    },
    idField: {
      type: String,
      required: true
    }
  },

  data() {
    return {
      items: [],
      loading: false,
      err: null,
    };
  },

  mounted() {
    this.loading = true;
    this.err = null;
    this.bind(this.contextName, this.runtimeObject, this.updateList)
      .then(() => this.loading = false)
      .catch(err => this.err = err);
  },

  beforeDestroy() {
    this.unbind(this.contextName, this.runtimeObject, this.updateList);
  },

  methods: {
    findIndex(obj) {
      return this.items.findIndex(element => this.getID(element) == this.getID(obj));
    },

    getID(obj) {
      var parts = this.idField.split(".");
      var res = obj;
      for(var i = 0; i < parts.length - 1; i++) {
        res = obj[parts[i]];
        if(!res) {
          return null;
        }
      }

      return res[parts[parts.length-1]];
    },

    updateList(event) {
      var index;
      switch(event.type) {
        case "added":
          if(this.findIndex(event.payload) != -1) {
            return;
          }

          this.items.push(event.payload);
          break;
        case "removed":
          index = this.findIndex(event.payload);
          if(index == -1) {
            return;
          }
          this.splice(index, 1);
          break;
        case "updated":
          index = this.findIndex(event.payload.oldObj);
          if(index == -1) {
            return;
          }
          this.items[index] = event.payload.newObj;
          break;
      }
    }
  }

});
</script>
