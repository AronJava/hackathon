import Vue from 'vue'
import Router from 'vue-router'
import home from '@/components/home'
import document from '@/components/document'
import person from '@/components/person'
import study from '@/components/person/children/study'
import learn from '@/components/person/children/learn'
import work from '@/components/person/children/work'
import error from '@/components/error'
import user from '@/components/user'
import slides from '@/components/slides'
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
      redirect: "/home"
    },{
      path: '/home',
      name:"home",
      component: home,
      beforeEnter(to,from,next){
        next();
      },
      meta:{
        index:0
      }
    },{
      path:'/document',
      name:"document",
      components:{
        default:document,
        slides:slides
      },
      meta:{
        index:1
      }
    },{
      path: '/person',
      component: person,
      children:[
        {
          path:"",
          name:"person",
          component: study,
          meta:{
            index:2
          }
        },
        {
          path:"study",
          component: study,
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
