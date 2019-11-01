  <template>
  <div>
    <PackChart :value=value />
  </div>

</template>

<script>
import PackChart from './Chart.vue';

export default {
  components: {
    PackChart,
  },
  props: ['value', 'height'],

  methods: {
    formatTime: function(timestamp) {
      const t = new Date(0);
      t.setUTCSeconds(timestamp.seconds);
      return t.toLocaleString();
    },
  },
};
</script>
