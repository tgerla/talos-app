<template>

  <div id="cluster">
    <div id="cluster-ops">
      <el-row>
        <el-col :span="12"><el-divider content-position="left"><h3>Cluster Information</h3></el-divider></el-col>
        <el-col :span="12"><el-divider><h3>Cluster Operations</h3></el-divider></el-col>

      </el-row>
      <el-row>
        <el-col :span="10">
          <p id="node-info-text">The cluster is <span style="color: green; font-weight: bold;">healthy</span>. 10 services are running, 1 is finished. Talos version 0.2,
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
    <p>
      <router-link :to="'/'">Home</router-link>
      <router-link :to="'/node'">Node</router-link>
    </p>
    
    <div>
      <PackChart :tweetData=loadData />
    </div>

  </div>
</template>

<script>

import PackChart from "./components/Chart.vue";
import * as d3 from "d3";

export default {
  name: "Cluster",
  components: {
    PackChart,
  },
    data: function() {
    return {
      loadData: []
    };
  },
  mounted() {
    console.log("App loaded");
    this.fetchData();
  },
  methods: {
    async fetchData() {
      //let data = await d3.json("/   tweets.json");
      this.loadData = [
{"user": "Al", "content": "I really love seafood.", "timestamp": " Mon Dec 23 2013 21:30 GMT-0800 (PST)", "retweets": ["Raj","Pris","Roy"], "favorites": ["Sam"]},
{"user": "Al", "content": "I take that back, this doesn't taste so good.", "timestamp": "Mon Dec 23 2013 21:55 GMT-0800 (PST)", "retweets": ["Roy"], "favorites": []},
{"user": "Al", "content": "From now on, I'm only eating cheese sandwiches.", "timestamp": "Mon Dec 23 2013 22:22 GMT-0800 (PST)", "retweets": [], "favorites": ["Roy","Sam"]},
{"user": "Roy", "content": "Great workout!", "timestamp": " Mon Dec 23 2013 7:20 GMT-0800 (PST)", "retweets": [], "favorites": []},
{"user": "Roy", "content": "Spectacular oatmeal!", "timestamp": " Mon Dec 23 2013 7:23 GMT-0800 (PST)", "retweets": [], "favorites": []},
{"user": "Roy", "content": "Amazing traffic!", "timestamp": " Mon Dec 23 2013 7:47  GMT-0800 (PST)", "retweets": [], "favorites": []},
{"user": "Roy", "content": "Just got a ticket for texting and driving!", "timestamp": " Mon Dec 23 2013 8:05 GMT-0800 (PST)", "retweets": [], "favorites": ["Sam", "Sally", "Pris"]},
{"user": "Pris", "content": "Going to have some boiled eggs.", "timestamp": " Mon Dec 23 2013 18:23 GMT-0800 (PST)", "retweets": [], "favorites": ["Sally"]},
{"user": "Pris", "content": "Maybe practice some gymnastics.", "timestamp": " Mon Dec 23 2013 19:47  GMT-0800 (PST)", "retweets": [], "favorites": ["Sally"]},
{"user": "Sam", "content": "@Roy Let's get lunch", "timestamp": " Mon Dec 23 2013 11:05 GMT-0800 (PST)", "retweets": ["Pris"], "favorites": ["Sally", "Pris"]}
];
    }
  }
};
</script>