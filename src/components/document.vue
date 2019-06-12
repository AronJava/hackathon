<template>
    <div style="height:2000px;">
        我是document组件
        <div id="maodian" @click="click">锚点</div>
    </div>
</template>
<script>
export default {
    name:"docuemnt",
    methods:{
        click(){
            window.location.href = "/docuemnt#maodian";
        }
    }
}
</script>

