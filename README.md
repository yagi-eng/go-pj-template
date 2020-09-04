# how to start this project

```
# MySQL準備
$ docker-compose up -d
# migration実行
$ go run tools/migrate.go
# アプリ実行
$ go run server.go
```

# go-pj-template
Goをいい感じに始められるひな形です。
<br/>Goがインストールしてあることを前提としています。
<br/>簡単なAPIサーバを立てる時や、Goを始めたばかりでどうプロジェクトを作成していけばいいかわからない時に有効だと思います。

本格的なGoプロジェクトを作る時には不向きです。
<br/>コードを読むとおわかり頂けるのですが、`db *gorm.DB`を引きずりまわしているのでその点が辛みです。
<br/>本格的なプロジェクトを作成する際は、`wire`などのDIライブラリを使うと捗ると思います。

`go.mod`に記載されているライブラリに関しては、以下のQiitaを参照ください。
<br/>`wire`や`gomock`などはインポートはしましたが、ひな形では使っていません。
<br/>記念に入れておいただけなので、`go mod tidy`すると消えます。

[Goの新規プロジェクト作成方法とおすすめライブラリ8.5選。ひな形もあるよ。]()

# 本プロジェクトに含まれる内容
- MySQL接続
- DB migrationのサンプル
- DB Read/Writeのサンプル
- echoを使ったルーティング、ロギング、CORS設定
- `godotenv`ライブラリを使った`.env`ファイルの読み込み
