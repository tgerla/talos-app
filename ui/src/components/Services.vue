  <template>
  <div class="data-table">
    <el-table :data="value" :height="height">
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-timeline reverse="true">
            <el-timeline-item
              v-for="(event, index) in props.row.events.events"
              :key="index"
              :timestamp="formatTime(event.ts)"
            >{{event.msg}}</el-timeline-item>
          </el-timeline>
          <!-- <p>Health: {{ props.row.health }}</p> -->
        </template>
      </el-table-column>
      <el-table-column prop="id" label="ID" min-width="180"></el-table-column>
      <el-table-column prop="state" label="State" min-width="180"></el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  props: ["value", "height"],

  methods: {
    formatTime: function(timestamp) {
      const t = new Date(0);
      t.setUTCSeconds(timestamp.seconds);
      return t.toLocaleString();
    }
  }
};
</script>
