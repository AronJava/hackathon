<template>
  <div class="article">
    <div class="top">
      <img class="avatar" src=""/>
    </div>
  </div>
</template>

<script>
import { fetchGet,fetchPost } from '@/utils/fetch';

export default {
  created() {
    this.getImg()
  },
  name: "post",
  data() {
      return {
        avator: ''
      }
  },
  methods: {
    click() {
      window.location.href = "/post#maodian";
    },
    getImg(){
      var that = this;
      fetchGet('https://www.easy-mock.com/mock/5d031cab0916f02402ce982b/jike/api/toutiao').then(res=>{
        console.log(res)
      })
    }
  }
};
</script>


