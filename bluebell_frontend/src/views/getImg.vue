<template>
    <div class="main">
        <div class="container">
            <el-row :gutter="20">
                <!-- 左- 筛选条件 -->

                <el-col :span="8">
                    <el-button type="success" @click="goUpload">去上传</el-button>
                    <el-form ref="form" :model="form" :label-position="labelPosition">
                        <el-form-item label="按名称">
                            <el-input v-model="form.name" style="max-width: 217px;"></el-input>
                            <el-button type="primary" @click="search(form.name, null)"
                                style="float: right;">筛选</el-button>

                        </el-form-item>
                        <el-form-item label="按标签">
                            <el-select v-model="form.tag" multiple filterable allow-create default-first-option
                                placeholder="请选择标签">
                                <el-option v-for="item in tagOptions" :key="item.tagId" :label="item.tagName"
                                    :value="item.tagId">
                                </el-option>
                            </el-select>
                            <el-button type="primary" @click="search(null, form.tag)"
                                style="float: right;">筛选</el-button>
                        </el-form-item>
                    </el-form>

                </el-col>

                <!-- 右-图片及标题 -->
                <el-col :span="16">
                    <div class="rightbox">

                        <div class="rightupbox">
                            <div class="imgbox" v-for="(item, index) in imgLists" :key="index">
                                <img :src="item.imgUrl" alt="">
                                <span>{{ item.title }}{{ item.type }}</span>
                            </div>
                        </div>
                    </div>
                    <div class="block" v-if="showPagination">
                        <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                            :current-page="currentPage" :page-sizes="[10, 20, 30, 40]" :page-size="pageSize"
                            layout="total, sizes, prev, pager, next, jumper" :total="400">
                        </el-pagination>
                    </div>
                </el-col>

            </el-row>

        </div>
    </div>
</template>

<script>
import qs from 'qs'

export default {
    name: "Upload",
    data() {
        return {
            labelPosition: 'top',
            form: {
                name: '',
                tag: [],
                year: ''
            },
            tagOptions: [],
            result: [],
            imgLists: [],
            currentPage: 1,
            showPagination: false,
            pageSize: 10,
            tags: []
        };
    },
    mounted() {
        this.fetchImageList();
        this.getAlltagList();
    },
    methods: {
        fetchImageList() {
            this.$axios({
                method: 'get',
                url: '/getImageList',
                params: {
                    page: this.currentPage,
                    pageNum: this.pageSize,
                    title: this.form.name,
                    projectId: 1,
                    version: 1,
                    tagIds: this.tags
                },
                paramsSerializer: params => {
                    return qs.stringify(params, { arrayFormat: 'repeat' });
                }
            }).then((res) => {
                if (res.data && res.data.length > 0) {
                    this.result = res.data;
                    this.getImageDetails();
                } else {
                    this.$message({
                        message: '没有图片了',
                        type: 'success'
                    });
                }
            }).catch((error) => {
                console.log(error);
            });
        },
        async getImageDetails() {
            const promises = this.result.map(async (item) => {
                const filename = `${item.imageId}${item.type}`;
                console.log(`Fetching image for filename: ${filename}`);

                try {
                    const imageDetailResponse = await this.$axios({
                        method: 'get',
                        url: `/view/${filename}`,
                        responseType: 'blob'
                    });

                    if (imageDetailResponse instanceof Blob) {
                        const imgUrl = URL.createObjectURL(imageDetailResponse);
                        console.log(`Image URL created: ${imgUrl}`);
                        item.imgUrl = imgUrl;
                    } else {
                        console.warn('Blob data expected but received something else.');
                    }
                } catch (error) {
                    console.error(`Error fetching image details for ${filename}:`, error);
                }
            });

            await Promise.all(promises);
            this.imgLists = this.result;
            this.showPagination = true
            console.log(this.imgLists);
        },
        getAlltagList() {
            // 获取标签
            this.$axios({
                method: 'get',
                url: '/tag/GetTags',
                params: {
                    page: 1,
                    pageNum: 10,
                },
            }).then((res) => {
                console.log(res.data);
                this.tagOptions = res.data
            }).catch((error) => {
                console.log(error);
            });
        },
        goUpload() {
            this.$router.push({ name: "Upload" });
        },
        handleSizeChange(val) {
            console.log(`每页 ${val} 条`);
            this.pageSize = val
            this.fetchImageList();
        },
        handleCurrentChange(val) {
            console.log(`当前页: ${val}`);
            this.currentPage = val
            this.fetchImageList();
        },
        search(name, tag) {
            this.form.name = name
            this.form.tag = tag
            // 过滤掉任何非字符串的元素，确保只包含有效的标签ID  
            this.tags = tag.filter(item => typeof item === 'string');
            console.log(name, this.tags);
            // this.tags = tag
            console.log(name, tag, this.tags);
            this.$axios({
                method: 'get',
                url: '/getImageList',
                params: {
                    page: this.currentPage,
                    pageNum: this.pageSize,
                    title: name,
                    projectId: 1,
                    version: 1,
                    tagIds: this.tags
                },
                paramsSerializer: params => {
                    return qs.stringify(params, { arrayFormat: 'repeat' });
                }
            }).then((res) => {
                if (res.data && res.data.length > 0) {
                    this.result = res.data;
                    this.getImageDetails();
                } else {
                    this.$message({
                        message: '搜索失败了',
                        type: 'success'
                    });
                }
            }).catch((error) => {
                console.log(error);
            });
        }
    }
};
</script>

<style lang="less" scoped>
.main {
    background: #f8f8f8;
    padding: 100px 0;
    height: calc(100vh - 200px);

    .container {
        width: 90%;
        background: #fff;
        margin: 0 auto;
        max-width: 1200px;
        padding: 20px;

        .rightbox {
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            height: calc(100vh - 272px);
            overflow-y: auto;
        }

        .rightupbox {
            // width: 300px;
            display: flex;
            flex-wrap: wrap;
        }

        .imgbox {
            margin-bottom: 10px;
            padding: 10px;
            text-align: center;
            display: flex;
            flex-direction: column;
            justify-content: space-between;
        }

        img {
            width: 130px;
            height: auto;
        }
    }
}
</style>