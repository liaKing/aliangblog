<template>
    <div class="main">
        <div class="container">
            <!-- :headers="headers" multiple :file-list="fileList" -->
            <el-upload style="flex: 1; margin-right: 30px;" class="upload-demo" drag :limit="1" action=""
                :http-request="httpRequest" :on-remove="handleRemove" :before-remove="beforeRemove"
                :before-upload="beforeAvatarUpload">
                <!-- :on-success="handleSuccess" -->
                <img v-if="uploadedImageUrl" :src="uploadedImageUrl" class="uploaded-image">
                <div v-else>
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                    <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
                </div>
            </el-upload>

            <el-form style="flex: 1.5;" ref="form" :model="form" :label-position="labelPosition">
                <el-form-item label="标题">
                    <!-- <el-input v-model="form.name" placeholder="请输入"></el-input> -->
                    <span class="imgtitle" v-if="form.name">{{ form.name }}</span>
                    <span class="imgtitle" v-else>请上传图片</span>
                </el-form-item>
                <el-form-item label="标签">
                    <!-- <el-select v-model="form.tag" placeholder="请选择">
                        <el-option label="男人" value="man"></el-option>
                        <el-option label="女人" value="woman"></el-option>
                        <el-option label="老人" value="oldman"></el-option>
                        <el-option label="孩子" value="kid"></el-option>
                    </el-select> -->
                    <!-- @close="handleClose(tag)" closable-->
                    已选标签：
                    <div class="midtag">
                        <el-tag type="success" :key="tag.tagId" v-for="tag in selectedTags" :disable-transitions="false"
                            @close="handleClose(tag)" closable>
                            {{ tag.tagName }}
                        </el-tag>
                    </div>


                    可选标签：
                    <div class="midtag">
                        <el-tag :type="tag.tagType" :key="tag.tagId" v-for="tag in optionalTags"
                            :disable-transitions="false" @click="handleClick(tag)">
                            {{ tag.tagName }}
                        </el-tag>
                    </div>

                    新增自定义标签：
                    <el-input class="input-new-tag" v-if="inputVisible" v-model="inputValue" ref="saveTagInput"
                        size="small" @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm">
                    </el-input>
                    <el-button v-else class="button-new-tag" size="small" @click="showInput">+ 新增 Tag</el-button>

                </el-form-item>
                <!-- <el-form-item label="分支">
                    <el-select v-model="form.branch" placeholder="请选择">
                        <el-option label="分支一" value="branch1"></el-option>
                        <el-option label="分支二" value="branch2"></el-option>
                    </el-select>
                </el-form-item> -->
                <div class="twobtn">
                    <el-button type="primary" @click="onSubmit">完成</el-button>
                    <el-button @click="console">取消</el-button>
                </div>
            </el-form>
        </div>

    </div>
</template>

<script>

export default {
    name: "Upload",
    data() {
        return {
            // headers: { Authorization: "Bearer " + JSON.parse(localStorage.getItem("loginResult")).accessToken },
            // action: "/api/v1/image/upload",
            // fileList: [],
            // fileList: [{ name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100' }],
            uploadedImageUrl: '',  // 存放上传后返回的图片 URL  
            labelPosition: 'top',
            imageInfo: '',
            form: {
                name: '',
                tag: '',
            },
            optionalTags: [],
            selectedTags: [],
            inputVisible: false,
            inputValue: '',
            // tagType: 'info'
        };
    },
    computed: {},
    created() { },
    methods: {
        async httpRequest(item) {
            console.log(item);
            let formData = new FormData();
            formData.append('image', item.file);
            formData.append('teamId', 1);
            formData.append('projectId', 1);
            this.form.name = item.file.name.split('.')[0]
            formData.append('title', this.form.name);
            try {
                const uploadResponse = await this.$axios({
                    method: 'post',
                    url: '/image/upload',
                    data: formData,
                    headers: {
                        'Authorization': "Bearer " + JSON.parse(localStorage.getItem("loginResult")).accessToken
                    }
                });
                this.imageInfo = uploadResponse.image
                console.log(uploadResponse);
                const imageDetailResponse = await this.$axios({
                    method: 'get',
                    url: `${uploadResponse.url}`,  // Ensure this is the correct URL path to fetch the image  
                    responseType: 'blob'
                });
                if (imageDetailResponse instanceof Blob) {
                    this.uploadedImageUrl = URL.createObjectURL(imageDetailResponse);
                } else {
                    console.warn('Blob data expected but received something else.');
                }
                // 获取标签
                this.$axios({
                    method: 'get',
                    url: '/tag/GetTags',
                    params: {
                        page: 1,
                        pageNum: 10,
                    },
                }).then((res) => {
                    // this.optionalTags = res.data
                    this.optionalTags = res.data.map(tag => ({
                        ...tag,
                        tagType: 'info',
                    }));
                    console.log(res.data);

                }).catch((error) => {
                    console.log(error);
                });
            } catch (error) {
                console.log(error);
            }
        },
        beforeAvatarUpload(file) {
            const isJPG = file.type === 'image/jpeg';
            const isLt2M = file.size / 1024 / 1024 < 2;

            if (!isJPG) {
                this.$message.error('上传图片只能是 JPG 格式!');
            }
            if (!isLt2M) {
                this.$message.error('上传图片大小不能超过 2MB!');
            }
            return isJPG && isLt2M;
        },
        beforeRemove(file, fileList) {
            console.log(file, fileList);
            return this.$confirm(`确定移除 ${file.name}？`);
        },
        handleRemove(file, fileList) {
            console.log(file, fileList);
            this.uploadedImageUrl = ''
            this.form.name = ''
        },
        handleClick(tag) {
            console.log(tag);
            tag.tagType = tag.tagType === 'info' ? 'success' : 'info';
            // 如果tag.tagType === 'info' 就从choosedTags数组中删除这个tag
            // 如果tag.tagType === 'success' 就从choosedTags数组中新增这个tag
            if (tag.tagType === 'info') {
                // 从 selectedTags 中删除该 tag  
                this.selectedTags = this.selectedTags.filter(t => t.tagId !== tag.tagId);
            } else if (tag.tagType === 'success') {
                // 将该 tag 添加到 selectedTags 中  
                this.selectedTags.push(tag);
            }
        },
        handleClose(tag) {
            console.log(tag);
            tag.tagType = 'info'
            this.selectedTags.splice(this.selectedTags.indexOf(tag), 1);
            // 如果是用户当场自定义的标签 删除后显示在可选标签内
            if (tag.type === 'diy') {
                this.optionalTags.push(tag);
            }
        },
        showInput() {
            this.inputVisible = true;
            // this.$nextTick(_ => { // _
            //     this.$refs.saveTagInput.$refs.input.focus();
            // });
        },

        handleInputConfirm() {
            let inputValue = this.inputValue;
            if (inputValue) {
                // 创建自定义标签
                this.$axios({
                    method: 'post',
                    url: '/tag/createTag',
                    data: JSON.stringify({
                        tagName: inputValue
                    })
                }).then((res) => {
                    const result = {
                        ...res.data,
                        tagType: 'success',
                        type: 'diy'
                    };
                    this.selectedTags.push(result);
                }).catch((error) => {
                    console.log(error);
                });
            }
            this.inputVisible = false;
            this.inputValue = '';
        },

        onSubmit() {
            console.log('submit!', this.selectedTags, this.imageInfo);
            // 收集tagid和imgid
            const tagIds = this.selectedTags.map(tag => tag.tagId).join(',');
            console.log(tagIds, this.imageInfo.imageId);
            if (tagIds || this.imageInfo.imageId) {
                this.$axios({
                    method: 'post',
                    url: '/tag/createImageTagLink ',
                    data: JSON.stringify({
                        tagId: tagIds,
                        imageId: this.imageInfo.imageId
                    })
                }).then((res) => {
                    console.log(res);
                    if (res.code === 1000) {
                        this.$message({
                            message: '上传并创建成功',
                            type: 'success'
                        });
                        this.$router.push({ name: "Home" });
                    }

                }).catch((error) => {
                    console.log(error);
                });
            } else {
                this.$message({
                    message: '请上传图片',
                    type: 'warning'
                });
            }


        },
        console() {
            this.$router.push({ name: "Home" });
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
        // height: 100%;
        background: #fff;
        margin: 0 auto;
        max-width: 1200px;
        padding: 20px;
        display: flex;
        justify-content: space-around;

        // 上传图片样式
        .uploaded-image {
            max-width: 100%;
            height: auto;
            // margin-top: 10px;
        }

        ::v-deep .el-upload-dragger {
            min-height: 350px;
            display: flex;
            align-items: center;
            justify-content: center;
        }

        ::v-deep .el-form-item {
            margin-bottom: 10px;
        }

        .el-tag {
            cursor: pointer;
        }

        .el-tag+.el-tag {
            margin-left: 10px;
        }

        .button-new-tag {
            margin-left: 10px;
            height: 32px;
            line-height: 30px;
            padding-top: 0;
            padding-bottom: 0;
        }

        .input-new-tag {
            width: 90px;
            margin-left: 10px;
            vertical-align: bottom;
        }

        .imgtitle {
            margin-left: 10px;
            color: #409EFF;
        }

        .midtag {
            height: 50px;
            background-color: aliceblue;
            border-radius: 5px;
            padding: 8px;
            overflow-y: auto;
        }

        .twobtn {
            padding-top: 10px;
            display: flex;
            justify-content: center;
        }
    }
}
</style>