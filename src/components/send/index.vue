<template>
    <div class="postDetail send">
       <div class="topBar">
            <i class="return" @click="jsReturn"></i>
            <span class="fabu" @click="fabu">发布</span>
       </div>
       <div class="content">
          <input type="text" v-model="inpTitle" class="tieTitle" placeholder="请输入标题（必填）">
          <textarea name="tieCon"  v-model="inpCom" id="tieCon" cols="300" rows="50" placeholder="请输入内容（必填）"></textarea>
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
    },
    data() {
      return {
        dataYang:'2019.06.15',
        dataYin:'农历五月十三',
        itemList:itemList,
        author:author,
        value:"日常生活",
        inpTitle:"",
        inpCom:""
      }
    },
    methods:{
    click() {
         window.location.href = "/post#maodian";
    },
    getImg(){
      var that = this;
      // fetchGet('https://www.easy-mock.com/mock/5d031cab0916f02402ce982b/jike/api/toutiao').then(res=>{
      //   this.avator = res.data.data.avatar;
      // })
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
    },
    jsReturn(){
        this.$router.back();
    },
    fabu(){
      console.log(this.inpTitle,this.inpCom);
       fetchGet('/api/paper/add',{"title":this.inpTitle,"content":this.inpCom,"type":1,}).then(res=>{
    
      })
    }
    }
}
</script>

