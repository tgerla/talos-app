  <template>
  <div class="data-table">
    <el-table :data="value" :height="height">
      <el-table-column prop="namespace" label="Namespace" min-width="180"></el-table-column>
      <el-table-column prop="id" label="ID" min-width="180"></el-table-column>
      <el-table-column prop="image" label="Image" min-width="180"></el-table-column>
      <el-table-column prop="pid" label="PID" min-width="180"></el-table-column>
      <el-table-column prop="status" label="Status" min-width="180"></el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  props: ["value", "height"]
};
</script>
