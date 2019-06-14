<template>
    <div>
        这里是热门讨论区，有任何技术或者生活的小随想都可以提问交流~
        <div id="maodian" @click="click">锚点</div>
    </div>
</template>
<script>
export default {
    name:"post",
    methods:{
        click(){
            window.location.href = "/post#maodian";
        }
    }
}
</script>

