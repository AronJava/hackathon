<template>
    <div class="user">
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
    </div>
</template>
<script>
let userList = [{
    id:1,
    name:"jyy",
    sex:"female",
    hobby:"male"
},{
    id:2,
    name:"wpp",
    sex:"male",
    hobby:"female"
},{
    id:3,
    name:"wxs",
    sex:"male",
    hobby:"fireHair"
},{
    id:4,
    name:"wjk",
    sex:"male",
    hobby:"fireHair"
}] 
export default {
    name:"user",
    data(){
        return {
            userList:userList,
            userInfo:{}
        }
    },
    beforeCreate(){
        console.log("beforeCreate");
    },
    beforeRouteEnter(to,from,next){
        console.log(this);
        next((vm)=>{
            console.log(vm)
        });
    },
    watch:{
        $route(){
            this.getData();
        }
    },
    mounted(){
        this.getData();
    },
    methods:{
        getData(){
            let index = this.$route.params.id;//当前路由的信息，$router是挣整个路由的信息
            this.userList.forEach((item)=>{
                if(item.id == index){
                    this.userInfo = this.userList[index-1];
                }
            })
        }
    }
}

//或者没有el，就new.Vue().$mount("#app")
//生命周期
</script>