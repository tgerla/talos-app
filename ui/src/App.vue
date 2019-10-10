<template>
  <div id="app">
    <el-container>
      <div style="width: auto;">
        <div style="text-align: center;">
          <img id="logo" src="./assets/logo.png" />
        </div>

        <el-button
          type="text"
          v-model="collapse"
          @click="toggleCollapse"
          icon="el-icon-s-unfold"
          style="font-size: 24px; padding-left: 20px;"
        ></el-button>
        <el-menu
          default-active="2"
          class="el-menu-vertical"
          @open="handleOpen"
          @close="handleClose"
          :collapse="collapse"
        >
          <el-submenu index="1">
            <template slot="title">
              <i class="el-icon-s-grid"></i>
              <span slot="title">Clusters</span>
            </template>
            <el-tree :data="treeData" />
          </el-submenu>
          <el-menu-item index="2">
            <i class="el-icon-s-tools"></i>
            <span slot="title">Settings</span>
          </el-menu-item>

          <el-divider></el-divider>

          <router-link :to="{ path: '/cluster' }">
            <el-menu-item index="3">
              <i class="el-icon-s-tools"></i>
              <span slot="title">TEST Cluster View</span>
            </el-menu-item>
          </router-link>

          <router-link :to="{path: '/node' }">
            <el-menu-item index="4">
              <i class="el-icon-s-tools"></i>
              <span slot="title">TEST Node View</span>
            </el-menu-item>
          </router-link>

        </el-menu>
      </div>
      <el-container>
        <el-main>
          <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/' }">Local Clusters</el-breadcrumb-item>
            <el-breadcrumb-item :to="{ path: '/' }">cluster-1</el-breadcrumb-item>
            <el-breadcrumb-item>node 1</el-breadcrumb-item>
          </el-breadcrumb>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Node from "./components/Node.vue";

export default {
  name: "app",
  components: {
    Node,
  },

  mounted() {
    this.update("nodes");
  },


  data() {
    return {
      collapse: false,
      drawer: false,
      nodes: [],
      node: "",
      version: "",
      treeData: [],
      options: [
        {
          name: "node-1"
        },
        {
          name: "node-2"
        }
      ]
    };
  },

  methods: {
    toggleCollapse() {
      this.collapse = !this.collapse;
    },

    handleShutdown() {
      this.drawer = false;
      const h = this.$createElement;
      this.$confirm(
        "This will shutdown the node. To power it back on, use the platform's interface. Continue?",
        "Warning",
        {
          confirmButtonText: "Shutdown",
          cancelButtonText: "Cancel",
          type: "warning"
        }
      )
        .then(() => {
          this.$message({
            type: "success",
            message: h("p", null, [h("span", null, "Shutdown requested")])
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: h("p", null, [h("span", null, "Shutdown canceled")])
          });
        });
    },

    handleReset() {
      this.drawer = false;
      const h = this.$createElement;
      this.$confirm("This will reset the node. Continue?", "Warning", {
        confirmButtonText: "Reset",
        cancelButtonText: "Cancel",
        type: "warning"
      })
        .then(() => {
          this.$message({
            type: "success",
            message: h("p", null, [h("span", null, "Reset requested")])
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: h("p", null, [h("span", null, "Reset canceled")])
          });
        });
    },

    handleReboot() {
      this.drawer = false;
      const h = this.$createElement;
      this.$confirm("This will reboot the node. Continue?", "Warning", {
        confirmButtonText: "Reboot",
        cancelButtonText: "Cancel",
        type: "warning"
      })
        .then(() => {
          this.$message({
            type: "success",
            message: h("p", null, [h("span", null, "Reboot requested")])
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: h("p", null, [h("span", null, "Reboot canceled")])
          });
        });
    },

    handleUpgrade() {
      this.drawer = false;
      const h = this.$createElement;
      this.$confirm("This will upgrade the node. Continue?", "Warning", {
        confirmButtonText: "Upgrade",
        cancelButtonText: "Cancel",
        type: "warning"
      })
        .then(() => {
          this.$message({
            type: "success",
            message: h("p", null, [h("span", null, "Upgrade requested")])
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: h("p", null, [h("span", null, "Upgrade canceled")])
          });
        });
    },

    // FIXME: this is duplicate code of the update() method in Node.vue,
    // but I'm sure there's a better way to not repeat ourselves.
    update(type) {
      let ws = new WebSocket(
        "ws://localhost:" + global.backendPort + "/web/app/events"
      );

      ws.onmessage = message => {
        let obj = JSON.parse(message.data);
        if (obj.type == type) {
          this[type] = JSON.parse(obj.event);
          console.log(this[type]);
        }

        // FIXME: there's probably a better place for this "business logic" to construct
        // the menu data structure (treeData) from the raw JSON returned by the k8s API call.
        switch(type) {
          case "nodes":
            // build the sidebar cluster menu based on the "get nodes" result from k8s
            var treeData = [{label: "Local Cluster", children: []}];
            for(var node of this['nodes']) {
              treeData[0]['children'].push({label: node.metadata.name});
            }
            console.log(treeData);
            this['treeData'] = treeData;
        }
        this.loading = false;
      };

      ws.onopen = () => {
        this.loading = true;
        ws.send(
          JSON.stringify({
            event: type
          })
        );
      };
    }
  }
};
</script>

<style>
@import url("https://fonts.googleapis.com/css?family=Lato&display=swap");

#app {
  font-family: "Lato", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

#logo {
  height: 75px;
  width: 75px;
  padding: 10px;
}

.el-menu-vertical:not(.el-menu--collapse) {
  width: 200px;
  min-height: 100%;
}

ul.el-menu a {
  text-decoration: none;
}

</style>
