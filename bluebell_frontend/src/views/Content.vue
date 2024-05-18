<template>
  <div class="content">
    <div class="left">
      <div class="container">
        <!-- <div class="post">
          <a class="vote">
            <span class="iconfont icon-up"></span>
          </a>
          <span class="text">50.2k</span>
          <a class="vote">
            <span class="iconfont icon-down"></span>
          </a>
        </div> -->
        <div class="l-container">
          <h4 class="con-title">{{ post.title }}</h4>
          <div class="con-info">{{ post.content }}</div>
        </div>
      </div>

      <div class="comment">
        <div class="c-left">
          <div class="line"></div>
          <div class="c-arrow">
            <a class="up down"></a>
          </div>
        </div>

        <div class="c-right">
          <div class="c-user-info">
<!--            根据用户id和文章判断是否以点赞，收藏，有则高亮-->
            <span class="name mr"><i class="iconfont icon-comment"></i> {{ post.commentNumber }}</span>
            <span class="num mr"  @click="isCol"><i class="iconfont icon-collect" :class="{isCollect:isCollect}"></i> {{ post.collectNumber }}</span>
            <span class="num mr"  @click="isStar"><i class="iconfont icon-star"   :class="{starAlready:starAlready}"></i> {{ post.starNumber }}</span>
            <div class="user-btn" @click="comment()">
              <span class="btn-item">
                <i class="iconfont icon-pinglun"></i>comment
              </span>
            </div>
          </div>
          <div
            class="c-content"
            v-for="item of commentList"
            :key="item.comment_id"
          >
            <div  class="time">
              用户名
              <span>{{ getFormattedDate(item.create_time) }}</span>
            </div>
            <span>{{ item.content }}</span>
        </div>
        </div>
      </div>
    </div>

<!--写死的不用管下面的代码-->
    <div class="right">
      <div class="topic-info">
        <h5 class="t-header"></h5>
        <div class="t-info">
          <a class="avatar"></a>
          <span class="topic-name">b/{{ post.community_name }}</span>
        </div>
        <p class="t-desc">森林狼总冠军</p>
        <ul class="t-num">
          <li class="t-num-item">
            <p class="number">5.2m</p>
            <span class="unit">Members</span>
          </li>
          <li class="t-num-item">
            <p class="number">5.2m</p>
            <span class="unit">Members</span>
          </li>
        </ul>
        <div class="date">Created Apr 10, 2008</div>
        <button class="topic-btn">JOIN</button>
      </div>
    </div>
  </div>
</template>

<script>
//使用import语句导入了一个名为formatDate的函数，这个函数用于将日期格式化为更易读的格式，具体是将评论的时间进行处理，因为后端传回来的时间数据格式有问题。
import formatDate from "../components/formatFDate";
// 使用export default导出了Content组件。具体什么作用不是很清楚
export default {
  name: "Content",//声明组件名
  data() { //data函数是什么意思？实现的文件是自己写的，还是内置？
    return {
      post: {},
      commentList: [],
      isCollect:false,
      starAlready:false,
    };
  },
  methods: {
    getFormattedDate(timeStr) {//关于时间的格式修改
      return formatDate(timeStr);
    },
    //下面这个方法是怎么调用的呢？打开页面就调用？
    getPostDetail() {//params.id是哪里来的，打开这个组件网页时传来的吗？
      this.$axios({
        method: "get",
        url: "/post/" + this.$route.params.id,
      })
        .then((response) => {
          console.log(1, response.data);
          if (response.code == 1000) {
            this.post = response.data;//this.post是早先定义的响应式数据属性，在这里为其赋值，
          } else {
            console.log(response.msg);
          }
          this.getCommentList(response.data.post_id);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    //上面方法运行完就会自动运行下面的方法。
    getCommentList(id) {
      this.$axios({
        method: "get",
        url: "/comment",
        params: { postId: id },
      })
        .then((response) => {
          console.log(1, response.data);
          this.commentList = response.data;
        })
        .catch((error) => {
          console.log(error);
        });
    },
    // 收藏
    deleteCollect() {
      this.$axios({
        method: "delete",
        url: "/collect",
        params: {postId: this.$route.params.id}
      })
          .then((response) => {
            console.log(1, response.data);
            if (response.code == 1000) {
              this.isCollect = false//this.post是早先定义的响应式数据属性，在这里为其赋值，
              this.getPostDetail()

            } else {
              console.log(response.msg);
            }
          })
          .catch((error) => {
            console.log(error);
          });
    },
    // 是否收藏
    getPostCollect(){
      this.$axios({
        method: "get",
        url: "/collect",
        params:{postId:this.$route.params.id}
      })
          .then((response) => {
            console.log(1, response.data);
            if (response.code == 1000) {
              this.isCollect = response.data;//this.post是早先定义的响应式数据属性，在这里为其赋值，
              this.getPostDetail()
            } else {
              console.log(response.msg);
            }
          })
          .catch((error) => {
            console.log(error);
          });
    },
    postCollect() {
      this.$axios({
        method: "post",
        url: "/collect",
        data:{parentId:this.post.post_id}
      })
          .then((response) => {
            console.log(1, response.data);
            if (response.code == 1000) {
              this.isCollect = true;//this.post是早先定义的响应式数据属性，在这里为其赋值，
              this.getPostDetail()
            } else {
              console.log(response.msg);
            }
          })
          .catch((error) => {
            console.log(error);
          });
    },
    comment() {
      console.log(11);
    },
    isCol(){
      // this.isCollect=!this.isCollect
      if (this.isCollect) {
        this.deleteCollect()
      } else {
        this.postCollect()
      }
    },

    // 点赞
    isStar(){
      if (this.starAlready) {
        this.deleteStar()
      } else {
        this.createStar()
      }
    },
    deleteStar() {
      this.$axios({
        method: "delete",
        url: "/star",
        params:{postId:this.post.post_id}
      })
          .then((response) => {
            console.log(1, response.data);
            if (response.code == 1000) {
              this.starAlready = false;//this.post是早先定义的响应式数据属性，在这里为其赋值，
              this.getPostDetail()
            } else {
              console.log(response.msg);
            }
          })
          .catch((error) => {
            console.log(error);
          });
    },
    createStar() {
      this.$axios({
        method: "post",
        url: "/star",
        data:{parentId:this.post.post_id,targetId:this.post.post_id}
      })
          .then((response) => {
            console.log(1, response.data);
            if (response.code == 1000) {
              this.starAlready = true;//this.post是早先定义的响应式数据属性，在这里为其赋值，
              this.getPostDetail()
            } else {
              console.log(response.msg);
            }
          })
          .catch((error) => {
            console.log(error);
          });
    },
    getStarAlready(){
      // console.log("111111")
      this.$axios({
        method: "get",
        url: "/star",
        params:{postId:this.$route.params.id}
      })
          .then((response) => {
            console.log(1, response.data);
            if (response.code == 1000) {
              this.starAlready = response.data;//this.post是早先定义的响应式数据属性，在这里为其赋值，
            } else {
              console.log(response.msg);
            }
          })
          .catch((error) => {
            console.log(error);
          });
    },
  },
  mounted: function () {
    this.getPostDetail();
    this.getPostCollect();
    this.getStarAlready()
  },
};
</script>

<style lang="less" scoped>
.mr {
  margin-right: 10px;
}
.content {
  max-width: 100%;
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin: 0 auto;
  padding: 20px 24px;
  margin-top: 48px;
  .left {
    flex-grow: 1;
    max-width: 740px;
    border-radius: 4px;
    word-break: break-word;
    background: #ffffff;
    border: #edeff1;
    flex: 1;
    margin: 32px;
    margin-right: 12px;
    padding-bottom: 30px;
    position: relative;
    .container {
      width: 100%;
      height: auto;
      position: relative;
      .post {
        align-items: center;
        box-sizing: border-box;
        display: -ms-flexbox;
        display: flex;
        -ms-flex-direction: column;
        flex-direction: column;
        height: 100%;
        left: 0;
        padding: 8px 4px 8px 0;
        position: absolute;
        top: 0;
        width: 40px;
        border-left: 4px solid transparent;
        // background: #f8f9fa;
        .text {
          color: #1a1a1b;
          font-size: 12px;
          font-weight: 700;
          line-height: 16px;
          pointer-events: none;
          word-break: normal;
        }
      }
      .l-container {
        padding: 15px;
        // margin-left: 40px;
        .con-title {
          color: #000000;
          font-size: 18px;
          font-weight: 500;
          line-height: 22px;
          text-decoration: none;
          word-break: break-word;
        }
        .con-info {
          margin: 20px 0;
          padding: 15px 0;
          border-bottom: 1px solid grey;
        }
        .con-cover {
          height: 512px;
          width: 100%;
          background: url("https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1585999647247&di=7e9061211c23e3ed9f0c4375bb3822dc&imgtype=0&src=http%3A%2F%2Fi1.hdslb.com%2Fbfs%2Farchive%2F04d8cda08e170f4a58c18c45a93c539375c22162.jpg")
            no-repeat;
          background-size: cover;
          margin-top: 10px;
          margin-bottom: 10px;
        }
      }
    }
    .comment {
      width: 100%;
      height: auto;
      position: relative;
      .c-left {
        .line {
          border-right: 2px solid #edeff1;
          // width: 20px;
          height: 100%;
          position: absolute;
          left: 20px;
        }
        .c-arrow {
          display: flex;
          display: -webkit-flex;
          position: absolute;
          z-index: 2;
          flex-direction: column;
          left: 12px;
          background: #ffffff;
          padding-bottom: 5px;
        }
      }
      .c-right {
        margin-left: 40px;
        padding-right: 10px;
        position: relative;
        .c-user-info {
          margin-bottom: 10px;
        .isCollect{
          background-color:yellow;
        }
        .starAlready{
          background-color: red;
        }
          .user-btn {
            position: absolute;
            top: 0px;
            right: 10px;
            font-size: 12px;
            display: flex;
            display: -webkit-flex;
            .btn-item {
              display: flex;
              display: -webkit-flex;
              align-items: center;
              margin-right: 10px;
              .iconfont {
                margin-right: 4px;
              }
            }
          }
          .name {
            color: #1c1c1c;
            font-size: 12px;
            font-weight: 400;
            line-height: 16px;
          }
          .num {
            padding-left: 4px;
            font-size: 12px;
            font-weight: 400;
            line-height: 16px;
            color: #7c7c7c;
          }
        }
        .c-content {
          font-family: Noto Sans, Arial, sans-serif;
          font-size: 14px;
          font-weight: 400;
          line-height: 21px;
          word-break: break-word;
          color: rgb(26, 26, 27);
          margin-bottom: 15px;
          .time{
            color: #7c7c7c;
          }
        }
      }
    }
  }
  .right {
    flex-grow: 0;
    width: 312px;
    margin-top: 32px;
    .topic-info {
      width: 100%;
      // padding: 12px;
      cursor: pointer;
      background-color: #ffffff;
      color: #1a1a1b;
      border: 1px solid #cccccc;
      border-radius: 4px;
      overflow: visible;
      word-wrap: break-word;
      padding-bottom: 30px;
      .t-header {
        width: 100%;
        height: 34px;
        background: #0079d3;
      }
      .t-info {
        padding: 0 12px;
        display: flex;
        display: -webkit-flex;
        width: 100%;
        height: 54px;
        align-items: center;
        .avatar {
          width: 54px;
          height: 54px;
          background: url("../assets/images/avatar.png") no-repeat;
          background-size: cover;
          margin-right: 10px;
        }
        .topic-name {
          height: 100%;
          line-height: 54px;
          font-size: 16px;
          font-weight: 500;
        }
      }
      .t-desc {
        font-family: Noto Sans, Arial, sans-serif;
        font-size: 14px;
        line-height: 21px;
        font-weight: 400;
        word-wrap: break-word;
        margin-bottom: 8px;
        padding: 0 12px;
      }
      .t-num {
        padding: 0 12px 20px 12px;
        display: flex;
        display: -webkit-flex;
        align-items: center;
        border-bottom: 1px solid #edeff1;
        .t-num-item {
          list-style: none;
          display: flex;
          display: -webkit-flex;
          flex-direction: column;
          width: 50%;
          .number {
            font-size: 16px;
            font-weight: 500;
            line-height: 20px;
          }
          .unit {
            font-size: 12px;
            font-weight: 500;
            line-height: 16px;
            word-break: break-word;
          }
        }
      }
      .date {
        font-family: Noto Sans, Arial, sans-serif;
        font-size: 14px;
        font-weight: 400;
        line-height: 18px;
        margin-top: 20px;
        padding: 0 12px;
      }
      .topic-btn {
        width: 286px;
        height: 34px;
        line-height: 34px;
        color: #ffffff;
        margin: 12px auto 0 auto;
        background: #003f6d;
        border-radius: 4px;
        box-sizing: border-box;
        margin-left: 13px;
      }
    }
  }
}
</style>