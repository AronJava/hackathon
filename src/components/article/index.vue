<template>
  <div class="article">
    <div class="top">
      <img class="avatar" v-bind:src="avator">
      <div class="date">
        <span class="text1">{{dataYang}}</span>
        <span class="text2">{{dataYin}}</span>
      </div>
      <input type="text" placeholder="彭世豪采茶纪" class="search">
    </div>
    <div class="crousel">
       <!-- Swiper -->
      <div class="swiper-container">
        <div class="swiper-wrapper">
          <div class="swiper-slide"><img src="../../assets/images/banner.png" alt=""></div>
          <div class="swiper-slide"><img src="../../assets/images/banner.png" alt=""></div>
          <div class="swiper-slide"><img src="../../assets/images/banner.png" alt=""></div>
          <div class="swiper-slide"><img src="../../assets/images/banner.png" alt=""></div>
          <div class="swiper-slide"><img src="../../assets/images/banner.png" alt=""></div>
        </div>
        <!-- Add Pagination -->
        <div class="swiper-pagination"></div>
      </div>
    </div>
    <div class="activity">
      <div class="item btn2"><img src="../../assets/images/btn2_bg.png" alt=""><span>精彩活动</span></div>
      <div class="item btn3"><img src="../../assets/images/btn3_bg.png" alt=""><span>日常生活</span></div>
      <div class="item btn4"><img src="../../assets/images/btn4_bg.png" alt=""><span>技术大咖</span></div>
    </div>
    <div class="post_area">

      <div class="post" >
        <div class="dian2"></div>
        <div class="content">
          <div class="text_area">
            <span class="text_disc">科技中心技术大咖倾囊相授：全角度全方位解读研发效能提升之道</span>
            <span class="time">2分钟前</span>
            <div class="huifu">
              <img src="../../assets/images/huifu.png" alt="">
              <span>0</span>
            </div>
          </div>
          <img class="pic_area" src="http://img3.myhsw.cn/2019-06-13/xf8ac5z9.jpg?60" alt="">
        </div>
      </div>

      <div class="post" >
        <div class="dian3"></div>
        <div class="content">
          <div class="text_area">
            <span class="text_disc">科技中心技术大咖倾囊相授：全角度全方位解读研发效能提升之道</span>
            <span class="time">2分钟前</span>
            <div class="huifu">
              <img src="../../assets/images/huifu.png" alt="">
              <span>0</span>
            </div>
          </div>
          <img class="pic_area" src="http://img3.myhsw.cn/2019-06-13/xf8ac5z9.jpg?60" alt="">
        </div>
      </div>

      <div class="post" >
        <div class="dian1"></div>
        <div class="content">
          <div class="text_area">
            <span class="text_disc">科技中心技术大咖倾囊相授：全角度全方位解读研发效能提升之道</span>
            <span class="time">2分钟前</span>
            <div class="huifu">
              <img src="../../assets/images/huifu.png" alt="">
              <span>0</span>
            </div>
          </div>
          <img class="pic_area" src="http://img3.myhsw.cn/2019-06-13/xf8ac5z9.jpg?60" alt="">
        </div>
      </div>
       <div class="post bg_gonggao" >
        <div class="content">
          <div class="text_area">
            <span class="text_disc">科技中心技术大咖倾囊相授：全角度全方位解读研发效能提升之道</span>
            <span class="time time_gonggao">2分钟前</span>
            <div class="huifu">
              <img src="../../assets/images/huifu_gonggao.png" alt="">
              <span>0</span>
            </div>
          </div>
          <img class="pic_area" src="http://img3.myhsw.cn/2019-06-13/xf8ac5z9.jpg?60" alt="">
        </div>
      </div>
      <div class="bottom"></div>
    </div>
  </div>
</template>

<script>
import { fetchGet,fetchPost } from '@/utils/fetch';
import Swiper from '@/utils/swiper.min.js'

export default {
  created() {
    this.getImg();
    this.getLun();
  },
  name: "post",
  data() {
      return {
        avator: '',
        dataYang:'2019.06.15',
        dataYin:'农历五月十三'
      }
  },
  methods: {
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
};
</script>


