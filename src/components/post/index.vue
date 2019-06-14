<template>
    <div class="post">
       <div class="post_area">
            <div class="top">
                <img class="avatar" v-bind:src="avator">
                <div class="date">
                    <span class="text1">{{dataYang}}</span>
                    <span class="text2">{{dataYin}}</span>
                </div>
                <input type="text" placeholder="彭世豪采茶纪" class="search">
            </div>
            <div class="post_content">
                <ul class="post_daohang">
                    <router-link :to='{path:"/post/wonActivity"}' tag="li" exact><i class="wonA_icon"></i>精彩活动</router-link>
                    <router-link  :to='{path:"/post/life"}' tag="li"><i class="life_icon"></i>日常生活</router-link>
                    <router-link  to="/post/technology" tag="li"><i class="tech_icon"></i>技术大咖</router-link>
                </ul>
                <router-view class=""></router-view>
            </div>
       </div>
    </div>
</template>
<script>
import { fetchGet,fetchPost } from '@/utils/fetch';
export default {
    name:"post",
    created() {
        this.getImg();
        this.getLun();
    },
    data() {
      return {
        avator: '',
        dataYang:'2019.06.15',
        dataYin:'农历五月十三'
      }
    },
    methods:{
         click() {
      window.location.href = "/post#maodian";
    },
    getImg(){
      var that = this;
      fetchGet('https://www.easy-mock.com/mock/5d031cab0916f02402ce982b/jike/api/toutiao').then(res=>{
        this.avator = res.data.data.avatar;
      })
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

