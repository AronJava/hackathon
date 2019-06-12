<template>
  <div class="person">
    <ul>
       <router-link :to='{path:"/person/ing"}' tag="li" exact><i class="fa fa-home"></i>活动ing</router-link>
       <router-link  :to='{path:"/person/learn"}' tag="li">活动ed</router-link>
       <!-- <router-link  to="/person/work" tag="li">work</router-link> -->
     </ul>
      <router-view class="center"></router-view>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
