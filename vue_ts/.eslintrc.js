module.exports = {
  root: true, // このフォルダがrootだという指定。つまり親フォルダの設定ファイルを探しに行かない
  env: {
    node: true // 事前定義されているグローバル変数をNode.jsグローバル変数にする。`browser`なども指定可能。
  },
  'extends': [ // vue用の拡張。基本これに従っておけばいいのだと思う。
    'plugin:vue/essential',
    '@vue/standard',
    '@vue/typescript'
  ],
  rules: { 
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off', // productionの場合はerrorレベルで出力。それ以外は普通に出力
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off'
  },
  parserOptions: {
    parser: '@typescript-eslint/parser' // typescript用のparserを指定
  }
}