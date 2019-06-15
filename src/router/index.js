import Vue from 'vue'
import Router from 'vue-router'
import article from '@/components/article/index'
import post from '@/components/post/index'
import activity from '@/components/activity/index'
import send from '@/components/send/index'
import ing from '@/components/activity/children/ing'
import learn from '@/components/activity/children/learn'
import work from '@/components/activity/children/work'
import wonActivity from '@/components/post/children/wonActivity'
import life from '@/components/post/children/life'
import technology from '@/components/post/children/technology'
import postDetail from '@/components/postDetail'
import error from '@/components/error'
import mine from '@/components/mine/index'
import danmu from '@/components/post/danmu'
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
        default:post
      },
      children:[
        {
          path:"",
          name:"wonActivity",
          component: wonActivity,
          meta:{
            index:0
          }
        },
        {
          path:"wonActivity",
          component: wonActivity,
        },{
          path:"life",
          component: life,
        },{
          path:"technology",
          component: technology,
        },{
          path:"danmu",
          component: danmu,
        },
      ]
    },,{
      path:'/send',
      name:"send",
      components:{
        default:send
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
      path:'/mine/:id?',
      component:mine,
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
    },
    {
      path: '/postDetail',
      name:"postDetail",
      component: postDetail,
      beforeEnter(to,from,next){
        next();
      },
      meta:{
        index:0
      }
    },
  ]
})
