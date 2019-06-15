<template>
    <div class="postDetail">
       <div class="topBar">
            <i class="return" @click="jsReturn"></i>
            <i class="diandiandian"></i>
       </div>
       <div class="content">
           <div class="detail_disc">
                <img class="avatar" v-bind:src="avator">
                <div class="right">
                    <div class="author">强叔</div>
                    <div class="diTime">发布于20分钟前</div>
                </div>
           </div>
           <div class="title">电竞霸主再燃激情 ROG确认参展2019ChinJoyBTOC</div>
           <div class="text_content">
             <div>
               <p>2019年第十七届ChinaJoy将于8月2日8月5日在上海新国际博览中心举行。中国国际数码互动娱乐展览会（ChinaJoy）2019年第十七届ChinaJoy将于8月2日8月5日在上海新国际博览中心举行。中国国际数码互动娱乐展览会（ChinaJoy）2019年第十七届ChinaJoy将于8月2日8月5日在上海新国际博览中心举行。中国国际数码互动娱乐展览会（ChinaJoy）2019年第十七届将于8月2日8月5日在上海新国际博览中心举行。中国国际数码互动娱乐展览会（ChinaJoy）2019年第十七届将于8月2日8月5日在上海新国际博览中心举行。中国国际数码互动娱乐展览会。</p> 
               <p>作为全球数字娱乐领域最具影响力的盛会，聚焦着海内外近千家展商及全世界游戏玩家的热切目光，无疑是全球顶尖数字娱乐产品和技术的最佳宣传平台。</p>
             </div>
             <img class="newsImg" src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1560574409264&di=b5c4ddc35f1a018e78f86e1342f7c28f&imgtype=0&src=http%3A%2F%2Fupload.newhua.com%2F2012%2F0214%2F1329182738783.jpg" alt="">
           </div>
       </div>
       <div class="tuijian">
         <div class="tujian_bar">
           <span class="text1">相关推荐</span>
           <span class="text2">RELATED POSTS</span>
           <span class="text3">换一批</span>
         </div>

         <div class="tuijian_tiezi">
            <div class="post" @click="tiezi(1)">
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

            <div class="post" @click="tiezi(1)">
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

            <div class="post" @click="tiezi(1)">
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
         </div>
       </div>

       <div class="comment">
         <div class="com_til">
           <span class="com_text">评论</span>
           <span class="com_num">12</span>
         </div>
         <div class="com_area"  >
           <div v-for="item in itemList">
             <div class="com_item" >
             <div class="top">
                <img class="avatar" v-bind:src="avator">
                <div>
                  <span class="author">{{item.author}}</span>
                  <div class="content_text">{{item.content}}</div>
                </div>
                 <div class="com_time">1分钟前</div>
             </div>
           </div>
         </div>
           </div>
       </div>
       <div class="com_wrap">
         <input type="text" class="com_inp" placeholder="评论一下..." v-model="value"  @keyup.enter="jsCom">
         <div class="bg">
            <span class="bg_img bg_huifu"></span>
            <!-- <span>12</span> -->
         </div>
         <div class="bg">
            <span class="bg_img bg_shoucang"></span>
         </div>
         <div class="bg">
            <span class="bg_img bg_share"></span>
         </div>
       </div>
    </div>
</template>
<script>
import { fetchGet,fetchPost } from '@/utils/fetch';
var author = "强叔";
var itemList = [{
  "avatar":"https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=427078936,2574206403&fm=26&gp=0.jpg",
  "author":"强叔",
  "content":"覆盖游戏、动漫、互联网影视与音乐、网络文学、电子竞技、智能娱乐软件与硬件以及新生娱乐业态等数字娱乐多领域，汇集全球顶级数",
  "time":"1分钟前"
},{
  "avatar":"https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=427078936,2574206403&fm=26&gp=0.jpg",
  "author":"龙哥",
  "content":"覆盖游戏、动漫、互联网影视与音乐、网络文学、电子竞技、智能娱乐软件与硬件以及新生娱乐业态等数字娱乐多领域，汇集全球顶级数",
  "time":"1分钟前"
},{
  "avatar":"https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=427078936,2574206403&fm=26&gp=0.jpg",
  "author":"xiangs",
  "content":"覆盖游戏、动漫、互联网影视与音乐、网络文学、电子竞技、智能娱乐软件与硬件以及新生娱乐业态等数字娱乐多领域，汇集全球顶级数",
  "time":"1分钟前"
},{
  "avatar":"https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=427078936,2574206403&fm=26&gp=0.jpg",
  "author":"强叔",
  "content":"覆盖游戏、动漫、互联网影视与音乐、网络文学、电子竞技、智能娱乐软件与硬件以及新生娱乐业态等数字娱乐多领域，汇集全球顶级数",
  "time":"1分钟前"
}]
export default {
    name:"postDetail",
    created() {
        this.getImg();
        this.getLun();
    },
    data() {
      return {
        avator: '',
        dataYang:'2019.06.15',
        dataYin:'农历五月十三',
        itemList:itemList,
        author:author,
        value:""
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
    jsCom(e){
      // console.log(e);
      // console.log(this.value)
      var text = {
        "avatar":"https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=427078936,2574206403&fm=26&gp=0.jpg",
        "author":author,
        "content":this.value,
        "time":"1秒前"
      };
      this.value = "";
      itemList.push(text);
      this.itemList = itemList;
      console.log(this.itemList);
    },
    jsReturn(){
        this.$router.back();
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

