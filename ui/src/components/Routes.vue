  <template>
  <div class="data-table">
    <el-table :data="value" :height="height">
      <el-table-column prop="interface" label="Interface" min-width="180"></el-table-column>
      <el-table-column prop="destination" label="Destination" min-width="180"></el-table-column>
      <el-table-column prop="gateway" label="Gateway" min-width="180"></el-table-column>
      <el-table-column prop="source" label="Source" min-width="180"></el-table-column>
      <el-table-column prop="family" label="Family" min-width="180"></el-table-column>
      <el-table-column prop="protocol" label="Protocol" min-width="180"></el-table-column>
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
