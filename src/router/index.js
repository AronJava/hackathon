import Vue from 'vue'
import Router from 'vue-router'
import article from '@/components/article/index'
import post from '@/components/post/index'
import activity from '@/components/activity/index'
import ing from '@/components/activity/children/ing'
import learn from '@/components/activity/children/learn'
import work from '@/components/activity/children/work'
import error from '@/components/error'
import user from '@/components/user'
import slides from '@/components/post/slides'
Vue.use(Router)

export default new Router({
  //history需要和后台配合
  mode:"history",
  linkActiveClass:"active",
  scrollBehavior(to,from,savedPosition){
    // if(savedPosition){
    //   return savedPosition;
    // }
    if(to.hash){
      return to.hash;
    }
  },
  routes: [
    {
      path: '/',
      redirect: "/article"
    },{
      path: '/article',
      name:"article",
      component: article,
      beforeEnter(to,from,next){
        next();
      },
      meta:{
        index:0
      }
    },{
      path:'/post',
      name:"post",
      components:{
        default:post,
        slides:slides
      },
      meta:{
        index:1
      }
    },{
      path: '/activity',
      component: activity,
      children:[
        {
          path:"",
          name:"activity",
          component: activity,
          meta:{
            index:2
          }
        },
        {
          path:"ing",
          component: ing,
        },{
          path:"learn",
          component: learn,
        },{
          path:"work",
          component: work,
        }
      ]
    },{
      path:'/user/:id?',
      component:user,
      meta:{
        index:3
      }
    // },{
    //   path:'*',
    //   // component:error,
    //   // redirect: "/home"
    //   redirect(to){
    //     if(to.path == "/123"){
    //       return "/person";
    //     }else if(to.path == "/abc"){
    //       return "/document";
    //     }else{
    //       return "/home";
    //     }
    //   }
    }
  ]
})
