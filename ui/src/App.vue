<template>
  <div id="app">
    <el-container>
      <div style="width: auto;">
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
            <el-menu-item-group>
              <span slot="title">AWS</span>
              <el-menu-item index="1-1">prod</el-menu-item>
              <el-menu-item index="1-2">stage</el-menu-item>
              <el-menu-item index="1-3">test</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group>
              <span slot="title">GCP</span>
              <el-menu-item index="2-1">prod</el-menu-item>
              <el-menu-item index="2-2">stage</el-menu-item>
              <el-menu-item index="2-3">test</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group>
              <span slot="title">Packet</span>
              <el-menu-item index="3-1">infra</el-menu-item>
            </el-menu-item-group>
          </el-submenu>
          <el-menu-item index="2">
            <i class="el-icon-s-tools"></i>
            <span slot="title">Settings</span>
          </el-menu-item>
        </el-menu>
      </div>
      <el-container>
        <el-header style="height: 40px; text-align: right; font-size: 12px">
          <el-button icon="el-icon-s-operation" @click="drawer = true" circle></el-button>
        </el-header>

        <el-main>
          <div style="text-align: center;">
            <div>
              <img id="logo" src="./assets/logo.png" />
            </div>
            <div>
              <el-select v-model="node" placeholder="Node" height="20px" width="400px">
                <el-option
                  v-for="item in options"
                  :key="item.name"
                  :label="item.name"
                  :value="item.name"
                ></el-option>
              </el-select>
            </div>
          </div>
          <Node />
          <el-drawer :visible.sync="drawer">
            <el-container>
              <el-main>
                <el-container direction="vertical">
                  <el-button-group>
                    <el-button type="primary" @click="handleShutdown" icon="el-icon-switch-button"></el-button>
                    <el-button type="primary" @click="handleReboot" icon="el-icon-refresh-right"></el-button>
                  </el-button-group>
                  <el-button-group>
                    <el-button type="primary" @click="handleUpgrade" icon="el-icon-top"></el-button>
                    <el-button type="primary" @click="handleReset" icon="el-icon-refresh"></el-button>
                  </el-button-group>
                </el-container>
              </el-main>
            </el-container>
          </el-drawer>
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
      collapse: true,
      drawer: false,
      node: "",
      version: "",
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
  height: 150px;
  width: 150px;
}

.el-menu-vertical:not(.el-menu--collapse) {
  width: 200px;
  min-height: 100%;
}
</style>
