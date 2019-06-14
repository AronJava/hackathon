<template>
  <div id="app">
    <div class="top" id="appTop">
      <img class="avatar" v-bind:src="avator">
      <div class="date">
        <span class="text1">{{dataYang}}</span>
        <span class="text2">{{dataYin}}</span>
      </div>
      <input type="text" placeholder="彭世豪采茶纪" class="search">
    </div>
    <div class="nav-box">
     <ul class="nav">
       <router-link :to = '{path:"/article"}' tag="li" event="mouseover" class="ml5"><div @click="hasTop"><i class="boottom_btn"></i>头条</div></router-link>
       <router-link  :to = '{path:"/post"}' active-class="a" tag="li"><div @click="hasTop"><i class="boottom_btn"></i>论坛</div></router-link>
       <router-link  :to = '{path:"/send"}' active-class="a" tag="li"><div @click="noTop"><i class="boottom_btn"></i>发新帖</div></router-link>
       <router-link  to="/activity" tag="li"><div @click="hasTop"><i class="boottom_btn"></i>消息</div></router-link>
       <router-link  to="/mine/index" tag="li"><div @click="noTop"><i class="boottom_btn"></i>我的</div></router-link>
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
import { fetchGet,fetchPost } from '@/utils/fetch';
export default {
  name:"app",
  created() {
    this.getImg();
    this.getLun();
  },
  data(){
    return {
      name:'left',
      avator: '',
      dataYang:'2019.06.15',
      dataYin:'农历五月十三'
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
    getImg(){
      var that = this;
      fetchGet('https://www.easy-mock.com/mock/5d031cab0916f02402ce982b/jike/api/toutiao').then(res=>{
        this.avator = res.data.data.avatar;
      })
    },
    noTop(){
      var top = document.getElementById("appTop");
      top.style.display = "none";
    },
    hasTop(){
      var top = document.getElementById("appTop");
      top.style.display = "block";
    },
    getLun(){
      // new Swiper('.swiper-container', {
      //   pagination: {
      //     el: '.swiper-pagination'
      //   },
      //   loop : true,
      //   autoplay:true
      // });
    }
  }
}
</script>

