<template>
  <div id="node">
    <div id="nodeops">
      <NodeOps :value="version" />
    </div>
    <div>
      <el-tabs  type="border-card" v-model="active" v-loading="loading" @tab-click="handleClick" stretch="true">
        <el-tab-pane label="Services" name="services">
          <Services v-model="services" height="100%" />
        </el-tab-pane>
        <el-tab-pane label="Containers" name="containers">
          <Containers v-model="containers" height="100%" />
        </el-tab-pane>
        <el-tab-pane label="Processes" name="processes">
          <Processes v-model="processes" height="100%" />
        </el-tab-pane>
        <el-tab-pane label="Routes" name="routes">
          <Routes v-model="routes" height="100%" />
        </el-tab-pane>
        <el-tab-pane label="Interfaces" name="interfaces">
          <Interfaces v-model="interfaces" height="100%" />
        </el-tab-pane>
        <el-tab-pane label="Mounts" name="mounts">
          <Mounts v-model="mounts" height="100%" />
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script>
import Services from "./Services.vue";
import Containers from "./Containers.vue";
import Processes from "./Processes.vue";
import Routes from "./Routes.vue";
import Interfaces from "./Interfaces.vue";
import Mounts from "./Mounts.vue";
import NodeOps from "./NodeOps.vue";

export default {
  components: {
    Services,
    Containers,
    Processes,
    Routes,
    Interfaces,
    Mounts,
    NodeOps
  },

  data() {
    return {
      active: "services",
      loading: false,
      version: {},
      services: [],
      containers: [],
      processes: [],
      routes: [],
      interfaces: [],
      mounts: [],
    };
  },

  mounted() {
    this.update(this.active);
    this.update("version");
  },

  methods: {
    handleClick(tab, event) {
      this.update(tab.name);
    },

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
#node {
  padding-top: 40px;
  padding-bottom: 40px;
  padding-left: 50px;
  padding-right: 50px;
  text-align: left;
}
</style>
