import Vue from 'vue'
import router from './router'
import App from './App'
import './assets/css/app.css';

var vm = new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
vm.$router.beforeEach(function(to,from,next){
  next();
})
// vm.$router.afterEach(function(to,from,next){
//   console.log(to,from,next);
// })