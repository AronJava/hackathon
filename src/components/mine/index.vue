<template>
    <div class="mine">
        <!-- 我是user页面
        <ul>
            <router-link v-for = "item in userList" :to = "'/user/'+item.id" tag = "li" exact = "">{{item.name}}</router-link>
            <router-link :to="{path:'',query:{info:'share'}}"  tag="li" exact="">他的分享</router-link>
            <router-link to="?info=keep"  tag="li" exact="">他的收藏</router-link>
            <div>
                姓名：{{userInfo.name}}
                性别：{{userInfo.sex}}
                爱好：{{userInfo.hobby}}
            </div>
        </ul> -->

        <img src="https://event.ncfwx.com/upload/image/20190614/21-15-mineAva_03.png" class="mine_avator" alt="">
        <div class="disc">
            <div class="name">强叔</div>
            <div class="nicheng">昵称：<span>微甜中年</span></div>
            <div class="disc_msg">承接下水道清理，承重墙加固，清洗油烟机等项目</div>
        </div>
        <div>
            <div class="msg msg1">
                <i class="iconl"></i>
                我的帖子
                <i class="iconr"></i>
            </div>
            <div class="msg msg2">
                <i class="iconl"></i>
                我的评论
               <i class="iconr"></i>
            </div>
            <div class="msg msg3">
                <i class="iconl"></i>
                我的收藏
               <i class="iconr"></i>
            </div>
            <div class="msg msg4">
                <i class="iconl"></i>
                退出登录
               <i class="iconr"></i>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    name:"user",
    data(){
        return {
         
        }
    },
    beforeCreate(){
        console.log("beforeCreate");
    },
    beforeRouteEnter(to,from,next){
        console.log(this);
        next((vm)=>{
            // console.log(vm)
        });
    },
    watch:{
        $route(){
            // this.getData();
        }
    },
    mounted(){
        // this.getData();
    },
    methods:{
        // getData(){
        //     let index = this.$route.params.id;//当前路由的信息，$router是挣整个路由的信息
        //     this.userList.forEach((item)=>{
        //         if(item.id == index){
        //             this.userInfo = this.userList[index-1];
        //         }
        //     })
        // }
    }
}

//或者没有el，就new.Vue().$mount("#app")
//生命周期
</script>