  <template>
  <div class="data-table">
    <el-table :data="value" :height="height">
      <el-table-column prop="pid" label="PID" min-width="180"></el-table-column>
      <el-table-column prop="ppid" label="Parent PID" min-width="180"></el-table-column>
      <el-table-column prop="state" label="State" min-width="180"></el-table-column>
      <el-table-column prop="threads" label="Threads" min-width="180"></el-table-column>
      <el-table-column prop="cpu_time" label="CPU Time" min-width="180"></el-table-column>
      <el-table-column prop="virtual_memory" label="Virtual Memory" min-width="180"></el-table-column>
      <el-table-column prop="resident_memory" label="Resident Memory" min-width="180"></el-table-column>
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
