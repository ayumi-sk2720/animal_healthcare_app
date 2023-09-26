# 概要
DockerでVue3の環境を構築
ペットの健康管理アプリケーションを作成する

# 立ち上げまでの流れ
## コンテナ起動
```
$ make up-build
```

## コンテナに入る
```
$ make exec-vue
```

## vueフォルダに移動してvueを立ち上げる
```
$ cd vue
$ yarn run serve
```

## 起動確認
* [https://localhost:3000](https://localhost:3000) にアクセスすると画面が表示される

## コマンドについて
Makefileを参照すること