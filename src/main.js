import Vue from 'vue'
import router from './router'
import App from './App'
import './utils/flexible'
import './utils/swiper.min.js'
import './assets/css/app.css';
import axios from 'axios'


var vm = new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
Vue.prototype.$http = axios
vm.$router.beforeEach(function(to,from,next){
  next();
})
// vm.$router.afterEach(function(to,from,next){
//   console.log(to,from,next);
// })