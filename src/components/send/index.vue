<template>
  <div>
    <div id="addEvent">
      <!--<select v-model.trim="type">-->
        <!--&lt;!&ndash;<option value="1">技术大咖</option>&ndash;&gt;-->
        <!--&lt;!&ndash;<option value="2">精彩活动</option>&ndash;&gt;-->
        <!--&lt;!&ndash;<option value="3">日常生活</option>&ndash;&gt;-->
        <!--<option v-for="item in options" value="item.value">{{item.text}}</option>-->
      <!--</select>-->
      <!--<span>Selected: {{ selected }}</span>-->

      <select name="type" v-model="options">
      <!-- <option :value="op.value" v-for="op in options" >{{op.text}}</option> -->
      </select>
      <p>
        <input type="text" style="width:370px; height:50px;" v-model.trim="title" placeholder="请输入标题（必填）"/>
      </p >
      <h1>
        <input type="text" style="width:370px; height:500px;" v-model.trim="content" placeholder="请输入内容（必填）"/>
      </h1>
      <p>
        <!--<input id="file" type="file" class="fileupload" v-model.trim="pic" accept="image/*"-->
               <!--capture="camera" v-on:change="viewimg"/>-->
        <!-- <input type="file" name="file1" accept="image/*" v-model.trim="pic" @change="viewimg" ref="uploadPic"> -->
        <!--<input class="fileupload" type="file" v-model.trim="pic" accept="image/*" onchange="viewimg(this)"/>-->
      </p >
      <p>
        <input type="submit" value="发布" @click="addEvent"/>
      </p >
    </div>
  </div>
</template>
<script>
import {fetchGet, fetchPost} from '@/utils/fetch';

export default {
  name: "post",
  created() {
    this.getImg();
    this.getLun();
  },
  data() {
    return {
      avator: '',
      dataYang: '2019.06.15',
      dataYin: '农历五月十三',
      type: 0,
      title: '',
      content: '',
      pic: '',
      options:[
        {value:'null', text:'选择分类'},
        {value:'1', text:'技术大咖'},
        {value:'2', text:'精彩活动'},
        {value:'3', text:'日常生活'}
      ],
      selected: 0,
    }
  },
  // created(){
  //
  // },
  methods: {
    click() {
      window.location.href = "/post#maodian";
    },
    getImg() {
      var that = this;
      fetchGet('https://www.easy-mock.com/mock/5d031cab0916f02402ce982b/jike/api/toutiao').then(res => {
        this.avator = res.data.data.avatar;
      })
    },
    getLun() {
      // new Swiper('.swiper-container', {
      //   pagination: {
      //     el: '.swiper-pagination'
      //   },
      //   loop : true,
      //   autoplay:true
      // });
    },
    addEvent() {
      let _self = this;
      let params = {type:_self.type, title:_self.title, content:_self.content, pic:_self.pic};
      if (_self.title && _self.content){
        fetchPost('http://10.145.0.05/event/addEvent', {inputText: params});
      }

    },
    setImagePreview(docObj, imgObjPreview) {
      if (docObj.files && docObj.files[0]) {
        imgObjPreview.style.display = 'block';
        imgObjPreview.src = window.URL.createObjectURL(docObj.files[0]);
      }
    },
    viewimg($event) {
      //获取当前的input标签
      var currentObj = event.currentTarget;
      //找到要预览的图片img标签，亦可动态生成
      var img = currentObj.parentNode.children[0];
      this.setImagePreview(currentObj, img);
    },
    // 单张图片上传
    uploadPic() {
      var inputs = document.getElementsByClassName("fileupload");
      for (var i = 0; i < inputs.length; i++) {
        //图片转base64上传
        var file = inputs[i].files;
        if (file[0]) {
          var reader = new FileReader();
          reader.readAsDataURL(file[0]);
          reader.onload = function (e) {
            var event = this;
            console.log(event.result);
            fetchPost('http://10.145.0.05/event/addPic', {base64: event.result});
          };
        }
      }
    }
  }
}

</script>
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