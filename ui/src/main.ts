import Vue from 'vue';
import App from './App.vue';
import Home from './Home.vue';
import Cluster from './Cluster.vue';
import Node from './components/Node.vue';

import './plugins/element.js';

import VueRouter from 'vue-router';
// tell vue to use the router
Vue.use(VueRouter);

// tslint:disable-next-line:no-var-requires
require('./assets/css/main.css');
Vue.config.productionTip = false;

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', name: 'Home', component: Home },
    { path: '/node', name: 'Node', component: Node },
    { path: '/cluster', name: 'Cluster', component: Cluster},
  ]
});

new Vue({
  // define the selector for the root component
  el: '#app',
  render: (h) => h(App),

  mounted() {
    // Prevent blank screen in Electron builds
    this.$router.push('/');
  },

  // pass in the router to the Vue instance
  router,
}).$mount('#app'); // mount the router on the app
