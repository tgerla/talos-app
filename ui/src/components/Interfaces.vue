  <template>
  <div class="data-table">
    <el-table :data="value" :height="height">
      <el-table-column prop="index" label="Index" min-width="180"></el-table-column>
      <el-table-column prop="mtu" label="MTU" min-width="180"></el-table-column>
      <el-table-column prop="name" label="Name" min-width="180"></el-table-column>
      <el-table-column prop="hardwareaddr" label="MAC" min-width="180"></el-table-column>
      <!-- <el-table-column prop="flags" label="Flags" min-width="180"></el-table-column> -->
      <el-table-column prop="ipaddress" label="Address" min-width="180"></el-table-column>
    </el-table>
  </div>
</template>
  </div>
</template>

<script>
export default {
  props: ["value", "height"]
};
</script>
