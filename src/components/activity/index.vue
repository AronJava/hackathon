<template>
  <div class="activity">
    <div class="msg system_msg">
      <i class="iconl s_icon"></i>
      系统消息
      <i></i>
    </div>
    <div class="msg comment_msg">
      <i class="iconl c_icon"></i>
      评论消息
      <i></i>
    </div>
  </div>
</template>

<script>
export default {
  name: 'activity',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
