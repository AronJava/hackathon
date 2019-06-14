<template>
  <div id="app">
    <div class="nav-box">
     <ul class="nav">
       <router-link :to = '{path:"/article"}' tag="li" event="mouseover" class="ml5"><i class="boottom_btn"></i>头条</router-link>
       <router-link  :to = '{path:"/post"}' active-class="a" tag="li"><i class="boottom_btn"></i>论坛</router-link>
       <router-link  :to = '{path:"/send"}' active-class="a" tag="li"><i class="boottom_btn"></i>发新帖</router-link>
       <router-link  to="/person" tag="li"><i class="boottom_btn"></i>消息</router-link>
       <router-link  to="/mine/index" tag="li"><i class="boottom_btn"></i>我的</router-link>
    </ul>
   </div>
   <!-- <button @click = "forward">前进</button>
   <button @click = "go">go</button>
   <button @click = "back">后退</button>
   <button @click = "push">push</button>
   <button @click = "replace">replace</button> -->
   <!-- <router-view name="slides"></router-view> -->
   <transition :name = "name" mode="out-in">
      <router-view class="center content"></router-view>
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

