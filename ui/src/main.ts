import Vue from 'vue';
import App from './App.vue';
import './plugins/element.js';

// tslint:disable-next-line:no-var-requires
require('./assets/css/main.css');

Vue.config.productionTip = false;

new Vue({
  render: (h) => h(App),
}).$mount('#app');
