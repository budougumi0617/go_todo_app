# 「詳解Go言語Webアプリケーション開発」正誤表

『詳解Go言語Webアプリケーション開発』は多くのレビューアのご協力により出版されましたが、誤字や内容に不備があった場合、すべて [@budugumi0617](https://github.com/budougumi0617) の責任です。  
誤りに気づかれた方は、[本リポジトリのDiscussionsにてご指摘](https://github.com/budougumi0617/go_todo_app/discussions) お願いいたします。


## 第1刷（2022年8月1日発行）
**P203 リスト19.1 「handler」パッケージから「store」パッケージを使う**  
https://github.com/budougumi0617/go_todo_app/blob/v1.0.1/_chapter19/section75/handler/add_task.go#L50
```diff
+    err := at.Repo.AddTask(ctx, at.DB, t)
  // 中略
-  }{ID: id}
+  }{ID: t.id}
   RespondJSON(ctx, w, rsp, http.StatusOK)
}
```
`t.id`ではなく、`t.ID` に修正。

**P203 リスト19.2　「handler.ListTask」型に対するリファクタリングの差分**  
https://github.com/budougumi0617/go_todo_app/blob/v1.0.1/_chapter19/section75/handler/list_task.go#L13

```diff
 type ListTask struct {
-  Store *store.TaskStore
+  DB *sqlx.DB
+  Repo store.Repository
 }
```
`ListTask`型の`Repo`フィールドを`store.Repository`型ではなく、`*store.Repository` 型のフィールドに修正。

**P204 リスト19.3 リファクタリング後の「NewMux」関数**  
https://github.com/budougumi0617/go_todo_app/blob/v1.0.1/_chapter19/section75/mux.go#L27-L29
```go
  at := &handler.AddTask{DB: db, Repo: r, Validator: v}
  mux.Post("/tasks", at.ServeHTTP)
  lt := &handler.ListTask{DB: db, Repo: r}
  mux.Get("/tasks", lt.ServeHTTP)
  return mux, cleanup, nil
}
```

`at`変数と`lt`変数の初期化を`Repo: r`ではなく、`Repo: &r`に修正。

**P209 リスト19.10　「handler/add_task_test.go」の修正**  
https://github.com/budougumi0617/go_todo_app/blob/v1.0.1/_chapter19/section78/handler/add_task_test.go#L53
```diff
+         moq := &AddTaskServiceMock{}
+         moq.AddTaskFunc = func(
+           ctx context.Context, title string
+         ) (*entity.Task, error) {
+           if tt.want.status ==
```

`ctx context.Context, title string` ではなく、`ctx context.Context, title string,`に修正。  
（行末にカンマを追加）

**P218 リスト19.24　「POST /register」エンドポイントの追加**  
サンプルコード例は差分表記が正しいです。

```go
mux.Get("/tasks", lt.ServeHTTP)
ru := &handler.RegisterUser{
  Service: &service.RegisterUser{DB: db, Repo: &r},
  Validator: v,
}
mux.Post("/register", ru.ServeHTTP)
```
ではなく、
```diff
  mux.Get("/tasks", lt.ServeHTTP)
+ ru := &handler.RegisterUser{
+   Service: &service.RegisterUser{DB: db, Repo: &r},
+   Validator: v,
+ }
+ mux.Post("/register", ru.ServeHTTP)
```
という装飾に修正。

# 第2刷（2022年8月19日発行）
**P89 リスト9.1　無名関数は状態を持てる**  
`リスト9.1`ではなく、`リスト9.2`に修正。  
[@YuyaAbo](https://github.com/YuyaAbo) さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/22) ありがとうございました。


**P112 リスト12.6　ログミドルウェアの実装例**  
`WriteHeader`メソッド内の「`r.status = status`」を「`r.status = statusCode`」に修正。  
[@litencatt](https://github.com/litencatt)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/33) ありがとうございました（2022/08/15）

**P148 テストとコードカバレッジ取得の自動実行**  
「GitHub Actiuons上で実行したテスト結果」ではなく、「GitHub Actions上で実行したテスト結果」に修正。  
[@kdnakt](https://twitter.com/kdnakt)さんご指摘ありがとうございました（2022/08/06）


**P163 リスト16.16　「httptest」パッケージを使った擬似的なHTTPリクエストのテスト**  
「`{"status": " ok"}`」ではなく、「`{"status": "ok"}`」に修正。  
[@kdnakt](https://twitter.com/kdnakt)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/29)ご指摘ありがとうございました（2022/08/10）

**P167 リスト17.3　「store/store.go」に実装したタスクの簡易管理方法**  
`Tasks`変数は`&TaskStore{Tasks: map[int]*entity.Task{}}`ではなく、`&TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}`に修正。  
`TaskStore`構造体の`LastID`プロパティは`int`ではなく、`entity.TaskID`に修正。  
`TaskStore`構造体の`Tasks`プロパティは`map[int]*entity.Task`ではなく、`map[entity.TaskID]*entity.Task`に修正。  
`func (ts *TaskStore) Add(t *entity.Task) (int, error) {`ではなく、`func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {`に修正。
`func (ts *TaskStore) Get(id int) (*entity.Task, error) {`ではなく、`func (ts *TaskStore) Get(id entity.TaskID) (*entity.Task, error) {`
[@mizutec](https://twitter.com/mizutec)さん[ご指摘](https://twitter.com/mizutec/status/1555043156865208320)ありがとうございました（2022/08/06）  
[@Mo3g4u](https://github.com/Mo3g4u)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/25)ありがとうございました（2022/08/06）

**P173 リスト17.7　「handler/add_task.go」のタスクを追加する実装**  
`}{ID: id}` ではなく、`}{ID: int(id)}`に修正。（2022/08/06）

**P175 リスト17.8　ファイルを使った入出力の検証**  
`Tasks: map[int]*entity.Task{},` ではなく、`Tasks: map[entity.TaskID]*entity.Task{},`に修正。（2022/08/06）

**P187 SECTION-072 MySQL実行環境の構築**  
「GitHub Actionis上でMySQLコンテナを起動します。」ではなく、「GitHub Actions上でMySQLコンテナを起動します。」に修正。  
「GitHub Actionisではサービスコンテナという方法で」ではなく、「GitHub Actionsではサービスコンテナという方法で」に修正。  
[@kdnakt](https://twitter.com/kdnakt)さんご指摘ありがとうございました（2022/08/15）

**P232 リスト20.13　「go:embed」ディレクティブの動作確認**  
2回目の`t.Errorf`の引数は「`("want %s, but got %s", want, rawPubKey)`」ではなく、「`Errorf("want %s, but got %s", want, rawPrivKey)`」に修正。（2022/08/28）

```go
want = []byte("-----BEGIN PRIVATE KEY-----")
if !bytes.Contains(rawPrivKey, want) {
  t.Errorf("want %s, but got %s", want, rawPrivKey)
}
```

**P244 SECTION-085 ユーザーログインエンドポイントの実装**  
「`LoginServiceインターフェースはauth/service.goに追記し，`」ではなく「`LoginServiceインターフェースはhandler/service.goに追記し，`」に修正。  
[@manaty226](https://github.com/manaty226)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/27)ご指摘ありがとうございました（2022/08/10）

**P250 リスト20.32　「POST /login」エンドポイントを追加する**  
「`JWTer: jwter,`」ではなく、「`TokenGenerator: jwter,`」に修正。  
[@yskgc](https://github.com/yskgc)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/31)ご指摘ありがとうございました（2022/08/12）