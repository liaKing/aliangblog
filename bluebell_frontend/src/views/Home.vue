<template>
  <div class="content">
    <div class="left">
      <!-- <h4 class="c-l-title">热门帖子</h4> -->
      <div class="c-l-header">
        <!-- :class判断展示最新帖子还是最热帖子-->
        <div
          class="new btn-iconfont"
          :class="{ active: timeOrder }"
          @click="selectOrder('time')"
        >
          <i class="iconfont icon-find"></i>New
        </div>
        <div
          class="top btn-iconfont"
          :class="{ active: scoreOrder }"
          @click="selectOrder('score')"
        >
          <i class="iconfont icon-rank"></i>Score
        </div>
        <div
          class="new btn-iconfont"
          v-if="role === 1"
          :class="{ active: blogVerify }"
          @click="selectVerify('blogVerify')"
        >
          <i class="iconfont icon-vertify"></i>
          博客审核
        </div>
        <div
          class="top btn-iconfont"
          :class="{ active: uInfoManager }"
          @click="selectManager('uInfoManager')"
          v-if="role === 1"
        >
          <i class="iconfont icon-usermanage"></i>
          <div>用户信息管理</div>
        </div>
        <button class="btn-publish" @click="goPublish">发表</button>
      </div>

      <el-tabs
        v-show="tabsShow"
        v-model="activeName"
        type="border-card"
        @tab-click="handleClick"
      >
        <el-tab-pane label="全部" name="all">
          <ul class="c-l-list">
            <li
              class="c-l-item"
              v-for="(post, index) in postList"
              :key="index + post.id"
            >
              <div class="l-container" @click="goDetail(post.post_id)">
                <h4 class="con-title">
                  {{ post.title }}
                  <div class="con-button" v-if="type == 'blogVerify'">
                    <el-button
                      type="success"
                      size="mini"
                      @click.stop="audit(true, post.post_id)"
                      >审核通过</el-button
                    >
                    <el-button
                      type="warning"
                      size="mini"
                      @click.stop="audit(false, post.post_id)"
                      >审核未通过</el-button
                    >
                  </div>
                </h4>
                <div class="con-memo">
                  <p>{{ post.content }}</p>
                </div>
              </div>
            </li>
          </ul></el-tab-pane
        >
        <el-tab-pane label="控球后卫" name="first">
          <ul class="c-l-list">
            <li
              class="c-l-item"
              v-for="(post, index) in postList"
              :key="index + post.id"
            >
              <div class="l-container" @click="goDetail(post.post_id)">
                <h4 class="con-title">
                  {{ post.title }}
                  <div class="con-button" v-if="type == 'blogVerify'">
                    <el-button
                      type="success"
                      size="mini"
                      @click.stop="audit(true, post.post_id)"
                      >审核通过</el-button
                    >
                    <el-button
                      type="warning"
                      size="mini"
                      @click.stop="audit(false, post.post_id)"
                      >审核未通过</el-button
                    >
                  </div>
                </h4>
                <div class="con-memo">
                  <p>{{ post.content }}</p>
                </div>
              </div>
            </li>
          </ul></el-tab-pane
        >
        <el-tab-pane label="得分后卫" name="second">
          <ul class="c-l-list">
            <li
              class="c-l-item"
              v-for="(post, index) in postList"
              :key="index + post.id"
            >
              <div class="l-container" @click="goDetail(post.post_id)">
                <h4 class="con-title">
                  {{ post.title }}
                  <div class="con-button" v-if="type == 'blogVerify'">
                    <el-button
                      type="success"
                      size="mini"
                      @click.stop="audit(true, post.post_id)"
                      >审核通过</el-button
                    >
                    <el-button
                      type="warning"
                      size="mini"
                      @click.stop="audit(false, post.post_id)"
                      >审核未通过</el-button
                    >
                  </div>
                </h4>
                <div class="con-memo">
                  <p>{{ post.content }}</p>
                </div>
              </div>
            </li>
          </ul></el-tab-pane
        >
        <el-tab-pane label="小前锋" name="third">
          <ul class="c-l-list">
            <li
              class="c-l-item"
              v-for="(post, index) in postList"
              :key="index + post.id"
            >
              <div class="l-container" @click="goDetail(post.post_id)">
                <h4 class="con-title">
                  {{ post.title }}
                  <div class="con-button" v-if="type == 'blogVerify'">
                    <el-button
                      type="success"
                      size="mini"
                      @click.stop="audit(true, post.post_id)"
                      >审核通过</el-button
                    >
                    <el-button
                      type="warning"
                      size="mini"
                      @click.stop="audit(false, post.post_id)"
                      >审核未通过</el-button
                    >
                  </div>
                </h4>
                <div class="con-memo">
                  <p>{{ post.content }}</p>
                </div>
              </div>
            </li>
          </ul></el-tab-pane
        >
        <el-tab-pane label="大前锋" name="fourth">
          <ul class="c-l-list">
            <li
              class="c-l-item"
              v-for="(post, index) in postList"
              :key="index + post.id"
            >
              <div class="l-container" @click="goDetail(post.post_id)">
                <h4 class="con-title">
                  {{ post.title }}
                  <div class="con-button" v-if="type == 'blogVerify'">
                    <el-button
                      type="success"
                      size="mini"
                      @click.stop="audit(true, post.post_id)"
                      >审核通过</el-button
                    >
                    <el-button
                      type="warning"
                      size="mini"
                      @click.stop="audit(false, post.post_id)"
                      >审核未通过</el-button
                    >
                  </div>
                </h4>
                <div class="con-memo">
                  <p>{{ post.content }}</p>
                </div>
              </div>
            </li>
          </ul></el-tab-pane
        >
        <el-tab-pane label="中锋" name="fifth">
          <ul class="c-l-list">
            <li
              class="c-l-item"
              v-for="(post, index) in postList"
              :key="index + post.id"
            >
              <div class="l-container" @click="goDetail(post.post_id)">
                <h4 class="con-title">
                  {{ post.title }}
                  <div class="con-button" v-if="type == 'blogVerify'">
                    <el-button
                      type="success"
                      size="mini"
                      @click.stop="audit(true, post.post_id)"
                      >审核通过</el-button
                    >
                    <el-button
                      type="warning"
                      size="mini"
                      @click.stop="audit(false, post.post_id)"
                      >审核未通过</el-button
                    >
                  </div>
                </h4>
                <div class="con-memo">
                  <p>{{ post.content }}</p>
                </div>
              </div>
            </li>
          </ul></el-tab-pane
        >
      </el-tabs>

      <!-- 用户信息管理 -->
      <ul v-if="isUser" class="c-l-list">
        <li class="c-l-item" v-for="(info, index) in infoList" :key="index">
          <div class="flex-container">
            <div class="flex">
              <h4 class="flex-title">{{ info.username }}</h4>
              <span class="con-memo">{{
                getFormattedDate(info.createTime)
              }}</span>
            </div>
            <span
              class="btn"
              v-if="info.delFlg"
              @click="recoverDel(info.username)"
              >撤销删除</span
            >
            <span class="btn" v-else @click="deleteUser(info.user_id)"
              >删除用户</span
            >
          </div>
        </li>
      </ul>
      <!-- 非用户信息管理 -->
      <!-- <ul v-else class="c-l-list">
        <li
          class="c-l-item"
          v-for="(post, index) in postList"
          :key="index + post.id"
        >
          <div class="l-container" @click="goDetail(post.post_id)">
            <h4 class="con-title">{{ post.title }}</h4>
            <div class="con-memo">
              <p>{{ post.content }}</p>
            </div>
          </div>
        </li>
      </ul> -->
    </div>
    <div class="right">
      <div class="communities">
        <h2 class="r-c-title">今日火热频道排行榜</h2>

        <ul class="r-c-content">
          <li class="r-c-item" v-for="(item, index) in hotChannel" :key="index">
            <span class="index">{{ index + 1 }}</span>
            <i class="icon"></i>{{ item.title }}
          </li>
        </ul>
        <button class="view-all">查看所有</button>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import formatDate from "../components/formatFDate";

export default {
  name: "Home",
  components: {},
  props: ["searchInput"],
  data() {
    return {
      role: "",
      activeName: "all", //标签页默认选项
      community: {
        first: 1,
        second: 2,
        third: 3,
        fourth: 4,
        fifth: 5,
      },
      tabsShow: true,
      order: "time",
      type: "time",
      page: 1,
      postList: [],
      infoList: [],
      isUser: false,
      hotChannel: [
        { title: "专注跳投" },
        { title: "NBA赛事讲解" },
        { title: "专注跳投" },
        { title: "NBA赛事讲解" },
        { title: "专注跳投" },
        { title: "NBA赛事讲解" },
        { title: "专注跳投" },
        { title: "NBA赛事讲解" },
        { title: "专注跳投" },
        { title: "NBA赛事讲解" },
      ], //火热频道排行榜
    };
  },
  watch: {
    searchInput: {
      handler() {
        // console.log('home',this.searchInput);
        this.getSearchList();
      },
    },
  },
  methods: {
    getFormattedDate(timeStr) {
      return formatDate(timeStr);
    },
    selectOrder(order) {
      this.tabsShow = true;
      this.isUser = false;
      this.type = order;
      this.order = order;
      this.activeName = "all";
      this.getSearchList();
    },
    selectVerify(val) {
      this.tabsShow = true;
      this.isUser = false;
      this.type = val;
      this.activeName = "all";
      this.getPostAudit();
    },
    selectManager(val) {
      this.tabsShow = false;
      this.isUser = true;
      this.type = val;
      this.getUserList();
    },
    goPublish() {
      this.$router.push({ name: "Publish" });
    },
    goDetail(id) {
      this.$router.push({ name: "Content", params: { id: id } });
    },

    // 根据搜索内容查询list
    getSearchList(data) {
      this.$axios({
        method: "get",
        url: "/posts/getPostByTitle",
        params: {
          page: this.page,
          title: this.searchInput,
          communityId: this.community[data],
        },
      })
        .then((response) => {
          if (response.code == 1000) {
            this.postList = response.data;
            console.log("搜索内容查询list", this.postList);
          } else {
            console.log(response.msg);
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    // /posts/getPostAudit
    getPostAudit(data) {
      this.$axios({
        method: "get",
        url: "/posts/getPostAudit",
        params: {
          page: this.page,
          communityId: this.community[data],
        },
      })
        .then((response) => {
          if (response.code == 1000) {
            this.postList = response.data;
            console.log("博客审核", this.postList);
          } else {
            console.log(response.msg);
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    // 获取管理用户信息列表
    getUserList() {
      this.$axios({
        method: "get",
        url: "/admin/userList",
        params: {
          page: this.page,
        },
      })
        .then((response) => {
          if (response.code == 1000) {
            this.infoList = response.data;
            console.log("用户信息", this.infoList);
          } else {
            console.log(response.msg);
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    //
    getAuditSuccess(id) {
      this.$axios({
        method: "put",
        url: "/posts/PostAuditSuccess",
        params: {
          postId: id,
        },
      })
        .then(() => {
          this.getPostAudit(this.activeName);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getAuditFail(id) {
      this.$axios({
        method: "put",
        url: "/posts/PostAuditFail",
        params: {
          postId: id,
        },
      })
        .then(() => {
          this.getPostAudit(this.activeName);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    deleteUser(userId) {
      this.$axios({
        method: "put",
        url: "/admin/deleteUser",
        params: {
          userId: userId,
        },
      })
        .then((response) => {
          if (response.code == 1000) {
            this.getUserList();
          } else {
            console.log(response.msg);
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    recoverDel(userName) {
      this.$axios({
        method: "put",
        url: "/admin/recoverUser",
        params: {
          userName: userName,
        },
      })
        .then((response) => {
          if (response.code == 1000) {
            this.getUserList();
            console.log("删除用户信息", response.data);
          } else {
            console.log(response.msg);
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    vote(post_id, direction) {
      this.$axios({
        method: "post",
        url: "/vote",
        data: JSON.stringify({
          post_id: post_id,
          direction: direction,
        }),
      })
        .then((response) => {
          if (response.code == 1000) {
            console.log("vote success");
          } else {
            console.log(response.msg);
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    // 标签页点击事件
    handleClick(tab, event) {
      if (this.type == "blogVerify") {
        this.getPostAudit(tab.name);
      } else {
        this.getSearchList(tab.name);
      }

      console.log(tab.name, event);
    },

    //
    audit(data, id) {
      if (data) {
        this.getAuditSuccess(id);
      } else {
        this.getAuditFail(id);
      }
    },
  },
  mounted: function () {
    this.getSearchList();
    this.role = JSON.parse(localStorage.getItem("loginResult")).role;
    console.log("this.role", this.role);
  },
  computed: {
    timeOrder() {
      return this.type == "time";
    },
    scoreOrder() {
      return this.type == "score";
    },
    blogVerify() {
      return this.type == "blogVerify";
    },
    uInfoManager() {
      return this.type == "uInfoManager";
    },
  },
};
</script>

<style scoped lang="less">
.content {
  max-width: 100%;
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin: 48px auto 0;
  padding: 20px 24px;
  .left {
    width: 640px;
    padding-bottom: 10px;
    .c-l-title {
      font-size: 14px;
      font-weight: 500;
      line-height: 18px;
      color: #1a1a1b;
      text-transform: unset;
      padding-bottom: 10px;
    }
    .c-l-header {
      align-items: center;
      background-color: #ffffff;
      border: 1px solid #ccc;
      border-radius: 4px;
      box-sizing: border-box;
      display: -ms-flexbox;
      display: flex;
      -ms-flex-flow: row nowrap;
      flex-flow: row nowrap;
      height: 56px;
      -ms-flex-pack: start;
      justify-content: flex-start;
      margin-bottom: 16px;
      padding: 0 12px;
      .iconfont {
        margin-right: 4px;
      }
      .btn-iconfont {
        display: flex;
        display: -webkit-flex;
        cursor: pointer;
      }
      .active {
        background: #f6f7f8;
        color: #0079d3;
        fill: #0079d3;
        border-radius: 20px;
        height: 32px;
        line-height: 32px;
        margin-right: 8px;
        padding: 0 10px;
      }
      .new {
        font-size: 14px;
        margin-right: 18px;
      }
      .top {
        font-size: 14px;
        margin-right: 18px;
      }
      .btn-publish {
        width: 64px;
        height: 32px;
        line-height: 32px;
        background-color: #54b351;
        color: #ffffff;
        border: 1px solid transparent;
        border-radius: 4px;
        box-sizing: border-box;
        text-align: center;
        margin-left: auto;
        cursor: pointer;
      }
      .sort {
        margin-left: 300px;
        display: flex;
        color: #0079d3;
        display: -webkit-flex;
        align-items: center;
        .sort-triangle {
          width: 0;
          height: 0;
          border-top: 5px solid #0079d3;
          border-right: 5px solid transparent;
          border-bottom: 5px solid transparent;
          border-left: 5px solid transparent;
          margin-top: 5px;
          margin-left: 10px;
        }
      }
    }
    .c-l-list {
      .c-l-item {
        list-style: none;
        border-radius: 4px;
        padding-left: 40px;
        cursor: pointer;
        border: 1px solid #ccc;
        margin-bottom: 10px;
        background-color: rgba(255, 255, 255, 0.8);
        color: #878a8c;
        position: relative;
        .flex-container {
          padding: 15px;
          display: flex;
          justify-content: space-between;
          align-content: center;
          .flex {
            display: flex;
            align-content: center;
          }
          .flex-title {
            color: #000000;
            font-size: 18px;
            font-weight: 500;
            text-decoration: none;
            word-break: break-word;
            margin-right: 10px;
          }
          .btn {
            font-size: 15px;
            padding: 5px 10px;
            color: #fff;
            border-radius: 4px;
            background-color: #0079d3;
          }
        }
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
          background: #f8f9fa;
          .iconfont {
            margin-right: 0;
          }
          .down {
            transform: scaleY(-1);
          }
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
          .con-title {
            color: #000000;
            font-size: 18px;
            font-weight: 500;
            line-height: 22px;
            text-decoration: none;
            word-break: break-word;
            .con-button {
              float: right;
            }
          }
          .con-memo {
            margin-top: 10px;
            margin-bottom: 10px;
            p {
              white-space: nowrap;
              text-overflow: ellipsis;
              overflow: hidden;
            }
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
          .user-btn {
            font-size: 14px;
            display: flex;
            display: -webkit-flex;
            .btn-item {
              display: flex;
              display: -webkit-flex;
              margin-right: 10px;
              .iconfont {
                margin-right: 4px;
              }
            }
          }
        }
      }
    }
  }
  .right {
    width: 312px;
    margin-left: 24px;
    margin-top: 28px;
    .communities {
      background-color: #ffffff;
      color: #1a1a1b;
      border: 1px solid #ccc;
      border-radius: 4px;
      overflow: visible;
      word-wrap: break-word;
      margin-bottom: 20px;
      .r-c-title {
        background-image: linear-gradient(
          0deg,
          rgba(0, 0, 0, 0.3) 0,
          transparent
        );
        background-color: #0079d3;
        height: 80px;
        width: 100%;
        color: #fff;
        font-size: 20px;
        line-height: 120px;
        padding-left: 10px;
        box-sizing: border-box;
        text-align: center;
      }
      .r-c-content {
        .r-c-item {
          align-items: center;
          display: flex;
          display: -webkit-flex;
          height: 48px;
          padding: 0 12px;
          border-bottom: thin solid #edeff1;
          font-size: 14px;
          .index {
            width: 20px;
            color: #1c1c1c;
            font-size: 14px;
            font-weight: 500;
            line-height: 18px;
          }
          .icon {
            width: 32px;
            height: 32px;
            background-image: url("../assets/images/avatar06.png");
            background-repeat: no-repeat;
            background-size: cover;
            margin-right: 20px;
          }
          &:last-child {
            border-bottom: none;
          }
        }
      }
      .view-all {
        background-color: #0079d3;
        border: 1px solid transparent;
        border-radius: 4px;
        box-sizing: border-box;
        text-align: center;
        letter-spacing: 1px;
        text-decoration: none;
        font-size: 12px;
        font-weight: 700;
        letter-spacing: 0.5px;
        line-height: 24px;
        text-transform: uppercase;
        padding: 3px 0;
        width: 280px;
        color: #fff;
        margin: 20px 0 20px 16px;
      }
    }
    .r-trending {
      padding-top: 16px;
      width: 312px;
      background-color: #ffffff;
      color: #1a1a1b;
      fill: #1a1a1b;
      border: 1px solid #cccccc;
      border-radius: 4px;
      overflow: visible;
      word-wrap: break-word;
      .r-t-title {
        font-size: 10px;
        font-weight: 700;
        letter-spacing: 0.5px;
        line-height: 12px;
        text-transform: uppercase;
        background-color: #ffffff;
        border-radius: 3px 3px 0 0;
        color: #1a1a1b;
        display: -ms-flexbox;
        display: flex;
        fill: #1a1a1b;
        padding: 0 12px 12px;
      }
      .rank {
        padding: 12px;
        .r-t-cell {
          display: flex;
          display: -webkit-flex;
          align-items: center;
          justify-content: space-between;
          margin-bottom: 16px;
          .r-t-cell-info {
            display: flex;
          }
          .avatar {
            width: 32px;
            height: 32px;
            background: url("../assets/images/avatar.png") no-repeat;
            background-size: cover;
            margin-right: 10px;
          }
          .info {
            margin-right: 10px;
            .info-title {
              font-size: 12px;
              font-weight: 500;
              line-height: 16px;
              text-overflow: ellipsis;
              width: 144px;
            }
            .info-num {
              font-size: 12px;
              font-weight: 400;
              line-height: 16px;
              padding-bottom: 4px;
            }
          }
          .join-btn {
            width: 106px;
            height: 32px;
            line-height: 32px;
            background-color: #0079d3;
            color: #ffffff;
            border: 1px solid transparent;
            border-radius: 4px;
            box-sizing: border-box;
            text-align: center;
          }
        }
      }
    }
  }
}
</style>
