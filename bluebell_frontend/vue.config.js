module.exports = {
  assetsDir: "static",
  devServer: {
    port: 9000, // 将端口号改为你想要的，例如 3000 
    proxy: {
      '/api/v1': {
        // target: 'http://127.0.0.1:8081/',
        // target: 'http://192.168.202.14:8081/',
        target: 'http://47.121.127.132:8081/', // 公网ip
        changeOrigin: true,
      }
    }
  },
  configureWebpack: {
    module: {
      rules: [
        {
          test: /\.mjs$/,
          include: /node_modules/,
          type: "javascript/auto"
        },

      ]

    }
  }
}
