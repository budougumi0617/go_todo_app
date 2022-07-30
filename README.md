# go_todo_app
このリポジトリは書籍「詳解Go言語Webアプリケーション開発」のサンプルコードリポジトリです。


| 項目   | データ                                                              |
|------|------------------------------------------------------------------|
| 書籍名  | 詳解Go言語Webアプリケーション開発                                              |
| URL  | https://www.c-r.com/book/detail/1462                             |
| ISBN | 978-4-86354-372-0                                                |
| 正誤表  | https://github.com/budougumi0617/go_todo_app/blob/main/errata.md |

## 作成するアプリについて
このリポジトリで作成するWebアプリケーションは認証付きのTODOタスクを管理するAPIサーバーです。

最終的には次のエンドポイントを実装します。

| HTTPメソッド | パス         | 概要                       |
|----------|------------|--------------------------|
| POST     | `/regiser` | 新しいユーザーを登録する             |
| POST     | `/login`   | 登録済みユーザー情報でアクセストークンを取得する |
| POST     | `/tasks`   | アクセストークンを使ってタスクを登録する     |
| GET      | `/tasks`   | アクセストークンを使ってタスクを一覧する     |
| GET      | `/admin`   | 管理者権限のユーザーのみがアクセスできる     |

`Docker Compose`を利用してAPIサーバー、MySQL、Redisを起動します。    
主に実行するであろうコマンドは `Makefile` に事前定義されています。

```bash
$ make
build                Build docker image to deploy
build-local          Build docker image to local development
up                   Do docker compose up with hot reload
down                 Do docker compose down
logs                 Tail docker compose logs
ps                   Check container status
test                 Execute tests
dry-migrate          Try migration
migrate              Execute migration
generate             Generate codes
help                 Show options
```

### 動作確認方法
このリポジトリのコードがローカルで実行できるか確認する手順です。

#### サーバを起動する
事前にDockerイメージを作成しておきます。
```bash
$ make build-local
```
Docker Composeを使って各サービスを起動します。
```bash
$ make up
```
MySQLにマイグレーションを実行します。
```bash
$ make migrate
```
ユーザーを作成します。
```bash
$ curl -X POST localhost:18000/register -d '{"name": "budou", "password":"test", "role":"admin"}'
{"id":37}
```
ユーザーの認証情報を使っていくつかタスクを登録します。
```bash
$ curl -i -XPOST -H "Authorization: Bearer $(curl -XPOST localhost:18000/login -d '{"user_name": "budou", "password":"test"}' | jq ".access_token" | sed "s/\"//g")" localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1023  100   982  100    41   7756    323 --:--:-- --:--:-- --:--:--  8525
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 20 Jul 2022 17:21:03 GMT
Content-Length: 9

{"id":76}%

$ curl -XPOST -H "Authorization: Bearer $(curl -XPOST localhost:18000/login -d '{"user_name": "budou", "password":"test"}' | jq ".access_token" | sed "s/\"//g")" localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1023  100   982  100    41   8634    360 --:--:-- --:--:-- --:--:--  9560
{"id":77}%
```
タスクを表示して、登録済みのタスクが表示されれば期待通り動いています。
```bash
$ curl -XGET -H "Authorization: Bearer $(curl -XPOST localhost:18000/login -d '{"user_name": "budou", "password":"test"}' | jq ".access_token" | sed "s/\"//g")" localhost:18000/tasks | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1023  100   982  100    41   8158    340 --:--:-- --:--:-- --:--:--  9133
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   113  100   113    0     0  13450      0 --:--:-- --:--:-- --:--:-- 28250
[
  {
    "id": 76,
    "title": "Implement a handler",
    "status": "todo"
  },
  {
    "id": 77,
    "title": "Implement a handler",
    "status": "todo"
  }
]
```

各エンドポイントの仕様の詳細は書籍を参照ください。

## リポジトリの構成について
書籍では`CHAPTER 13`からハンズオン形式でAPIサーバを実装します。  
本リポジトリのルートディレクトリにあるサンプルコードはハンズオンを最後まで実施した状態のコードです。  
なお、指摘や修正があった場合はコードも修正されます。

`_chapterN`で始まるディレクトリとその中にある`sectionN`ディレクトリは該当CHAPTER/SECTION完了時のコード状態を保存しています。  
紙面上の断片的なコードで理解が難しい場合はご利用ください。既知の紙面上のコードの不備も修正済みです。

## 書籍・コードの内容について
- コードに対するご指摘はIssueを作成してください。
  - https://github.com/budougumi0617/go_todo_app/issues
- 書籍内容に対するご指摘・ご質問などはDiscussionsで受け付けています。issueを作成するかDiscussionを作成するか迷った際はDiscussionの作成をお願いします。
  - https://github.com/budougumi0617/go_todo_app/discussions
- 正誤表もこのリポジトリにあります。
  - [https://github.com/budougumi0617/go_todo_app/blob/main/errata.md](./errata.md)