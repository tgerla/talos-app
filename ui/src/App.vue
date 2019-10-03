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
          </router>          

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
    Node
  },

  data() {
    return {
      collapse: false,
      drawer: false,
      node: "",
      version: "",
      treeData: [{
          label: 'Local Clusters',
          children: [{
            label: 'cluster-1',
            children: [
              {label: 'node 1'},
              {label: 'node 2'},
            ]
          }]
        },{
          label: 'AWS',
          children: [{
            label: 'cluster-2',
            children: [
              {label: 'node 1'},
              {label: 'node 2'},
            ]
          }]
        }, {
          label: 'Azure',
          children: [{
            label: 'cluster-b',
            children: [
              {label: 'node 1'},
              {label: 'node 2'},
            ]
          }, {
            label: 'cluster-c',
            children: [{
              label: 'node 1'
            }]
          }]
        }, {
          label: 'Packet Hosting',
          children: [{
            label: 'cluster-4',
            children: [{
              label: 'node 1'
            }]
          }, {
            label: 'cluster-5',
            children: [{
              label: 'node 1'
            }]
          }]
        }],
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
