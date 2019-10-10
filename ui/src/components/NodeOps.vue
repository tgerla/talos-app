<template>
  <div id="node-ops">
    <el-row>
      <el-col :span="12"><el-divider content-position="left"><h3>Node Information</h3></el-divider></el-col>
      <el-col :span="12"><el-divider><h3>Node Operations</h3></el-divider></el-col>

    </el-row>
    <el-row>
      <el-col :span="10">
        <p id="node-info-text">The node is <span style="color: green; font-weight: bold;">healthy</span>. 10 services are running, 1 is finished. Talos version 0.2,
        Kubernetes 1.16.</p>
      </el-col>
      <el-col :span="12" :offset="2">
        <el-button-group>
          <el-button type="info" size="small" @click="handleShutdown" icon="el-icon-switch-button">Shut Down</el-button>
          <el-button type="info" size="small" @click="handleReboot" icon="el-icon-refresh-right">Reboot</el-button>
        </el-button-group>
        <el-divider direction="vertical"></el-divider>
        <el-button-group>
          <el-button type="info" size="small" @click="handleUpgrade" icon="el-icon-top">Upgrade</el-button>
          <el-button type="info" size="small" @click="handleReset" icon="el-icon-refresh">Reset</el-button>
        </el-button-group>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  props: ['value', 'height'],
};
</script>

<style scoped>
#node-info-text {
  margin-top: 0;
}

#node-ops {
  padding-bottom: 30px;
}
</style>