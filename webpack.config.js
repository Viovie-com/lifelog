var path = require('path')

module.exports = {
    mode: 'development',
    context: path.resolve(__dirname, 'web/js'),
    entry: './app.js',
    output: {
        path: path.join(__dirname, 'web/public/js'),
        filename: 'app.js'
    },
    resolve: {
        alias: {
            vue: 'vue/dist/vue.esm.js'
        }
    }
}
