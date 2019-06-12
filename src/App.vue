<template>
  <div id="app">
    <img class="logo" src="./assets/images/logo.png"/>
    <div class="nav-box">
     <ul class="nav">
       <router-link :to = '{path:"/article"}' tag="li" event="mouseover">文章</router-link>
       <router-link  :to = '{path:"/post"}' active-class="a" tag="li">帖子</router-link>
       <router-link  to="/person" tag="li">活动</router-link>
       <router-link  to="/user" tag="li">我的</router-link>
    </ul>
   </div>
   <!-- <button @click = "forward">前进</button>
   <button @click = "go">go</button>
   <button @click = "back">后退</button>
   <button @click = "push">push</button>
   <button @click = "replace">replace</button> -->
   <router-view name="slides"></router-view>
   <transition :name = "name" mode="out-in">
      <router-view class="center"></router-view>
   </transition>
  </div>
</template>
<script>
export default {
  
  name:"app",
  data(){
    return {
      name:'left'
    }
  },
  watch:{
    $route(to,from){
      var indexTo = to.meta.index;
      var indexFrom = from.meta.index;
      if(indexTo < indexFrom){
        this.name="right";
      }else{
        this.name="left";
      }
    }
  },
  methods:{
    //编程式导航：用js代码控制路由跳转
    //routelink实质上也是通过back之类实现的
    // back(){
    //   this.$router.back();
    // },
    // forward(){
    //   this.$router.forward();
    // },
    // go(){
    //   this.$router.go(-2);
    // },
    // push(){
    //   this.$router.push("/user");
    // },
    // replace(){
    //   this.$router.replace("/user");
    // }
  }
}
</script>

