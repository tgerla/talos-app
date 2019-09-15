  <template>
  <div class="data-table">
    <el-table :data="value" :height="height">
      <el-table-column prop="filesystem" label="Filesystem" min-width="180"></el-table-column>
      <el-table-column prop="size" label="Size" min-width="180"></el-table-column>
      <el-table-column prop="available" label="Available" min-width="180"></el-table-column>
      <el-table-column prop="mounted_on" label="Mounted On" min-width="180"></el-table-column>
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
