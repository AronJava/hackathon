<template>
    <div class="post_three post_life">
      <div class="post_area">
        <div class="post" @click="tiezi(2)">
            <div class="dian2"></div>
            <div class="content">
                <div class="text_area">
                    <span class="text_disc">科技中心技术大咖倾囊相授：全角度全方位解读研发效能提升之道</span>
                    <span class="time">2分钟前</span>
                    <div class="huifu">
                        <img src="../../../assets/images/huifu.png" alt="">
                        <span>0</span>
                    </div>
                </div>
                <img class="pic_area" src="http://img3.myhsw.cn/2019-06-13/xf8ac5z9.jpg?60" alt="">
           </div>
        </div>


         <div class="post" >
            <div class="dian2"></div>
            <div class="content">
                <div class="text_area">
                    <span class="text_disc">科技中心技术大咖倾囊相授：全角度全方位解读研发效能提升之道</span>
                    <span class="time">2分钟前</span>
                    <div class="huifu">
                        <img src="../../../assets/images/huifu.png" alt="">
                        <span>0</span>
                    </div>
                </div>
                <img class="pic_area" src="http://img3.myhsw.cn/2019-06-13/xf8ac5z9.jpg?60" alt="">
           </div>
        </div>



         <div class="post" >
            <div class="dian2"></div>
            <div class="content">
                <div class="text_area">
                    <span class="text_disc">科技中心技术大咖倾囊相授：全角度全方位解读研发效能提升之道</span>
                    <span class="time">2分钟前</span>
                    <div class="huifu">
                        <img src="../../../assets/images/huifu.png" alt="">
                        <span>0</span>
                    </div>
                </div>
                <img class="pic_area" src="http://img3.myhsw.cn/2019-06-13/xf8ac5z9.jpg?60" alt="">
           </div>
        </div>

      </div>
    </div>
</template>

<script>
import { fetchGet,fetchPost } from '@/utils/fetch';
import slider from 'vue-concise-slider'

export default {
  created() {
    this.getImg();
  },
  name: "post",
  data() {
      return {
        dataYang:'2019.06.15',
        dataYin:'农历五月十三',
        //图片列表[arr]
         pages:[
     {
     html: '<p class="slider1">slider1</p>',
     style: {
      'background': '#1bbc9b'
      }
     },
     {
      html: '<p class="slider2">slider2</p>',
      style: {
       'background': '#4bbfc3'
      }
     },
     {
      html: '<p class="slider3">slider3</p>',
      style: {
       'background': '#7baabe'
      }
     }
    ],
    //Sliding configuration [obj]
    sliderinit: {
     currentPage: 0,
     thresholdDistance: 500,
     thresholdTime: 100,
     autoplay:1000,
     loop:true,
     direction:'vertical',
     infinite:1,
     slidesToScroll:1,
     timingFunction: 'ease',
     duration: 300
    }
   }
  },
  mounted(){
      // console.log("1111111111");
      // new Swiper('.swiper-container', {
      //   pagination: {
      //     el: '.swiper-pagination'
      //   },
      //   loop : true,
      //   autoplay:true
      // });
  },
  methods: {
    getImg(){
    //   fetchGet('https://www.easy-mock.com/mock/5d031cab0916f02402ce982b/jike/api/toutiao').then(res=>{
    //     this.avator = res.data.data.avatar;
    //   })
    },
    tiao(addre){
      window.location.href=addre;
    },
    tiezi(param){
      window.location.href="/post/danmu?id="+param;
    },
  },
  components: {
      slider
  }
};
</script>


