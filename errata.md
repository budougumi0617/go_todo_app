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
**P48 COLUMN 「internal」パッケージ**  
「たとえば、`example.com/root/internal`パッケージに`exported`な`Hooga`型が宣言されていた場合でも、」ではなく、
「たとえば、`example.com/root/internal`パッケージに`exported`な`Hoge`型が宣言されていた場合でも、」に修正。  
[@yakkun](https://github.com/yakkun)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/80) ありがとうございました（2023/01/29）

**P50 SECTION-019 モジュールとパッケージ**  
「`たとえば、あるGitHubeにあるリポジトリが1モジュールで、`」ではなく、「`たとえば、あるGitHubにあるリポジトリが1モジュールで、`」に修正。  
[@hajipy](https://github.com/hajipy) さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/89) ありがとうございました（2024/09/01）

**P51 SECTION-020 Go Modules**  
「`極力、古いバージョンが選択されるので「パッケージX」は「パッケージX v1.0.1」が選ばれます。`」ではなく、「`「パッケージX」はたとえv1.0.9がリリースされていても、「パッケージX v1.0.8」が選ばれます。`」に修正。  
[@hajipy](https://github.com/hajipy) さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/90) ありがとうございました（2024/09/01）

**P77 リスト8.3　「errors.New」関数と「fmt.Errorf」関数**  
「`return errors.New("GetAuthor: id is invalid")`」ではなく、「`return nil, errors.New("GetAuthor: id is invalid")`」に修正。  
[@fuuukeee3](https://github.com/fuuukeee3) さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/42) ありがとうございました（2022/09/02）


**P77 リスト8.3　「errors.New」関数と「fmt.Errorf」関数**  
「`return "", fmt.Errorf("GetAuthor: %v", err)`」ではなく、「`return "", fmt.Errorf("GetAuthorName: %v", err)`」に修正。  
[@katutoshi](https://github.com/katutoshi) さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/42) ありがとうございました（2022/09/02）

**P77 SECTION-033 エラーについて**  
「`Goではこのようチェーンして`」ではなく、「`Goではこのようにチェーンして`」に修正。  
[WomenWhoGoTokyo](https://github.com/WomenWhoGoTokyo/book-reading-party)のみなさんご指摘ありがとうございました（2022/10/02）

**P77 リスト8.4 チェーンされた最終的なエラーの出力**  
「`GetBookSummary: GetAuthorName: GetUser: id is invalid`」ではなく、「`GetBookSummary: GetAuthorName: GetAuthor: id is invalid`」に修正。  
[@katutoshi](https://github.com/katutoshi) さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/42) ありがとうございました（2022/09/02）

**P78 リスト8.5　「sql.ErrNoRows」と比較しても真になることはない**  
**P79 リスト8.6　Go 1.13からのエラーラッピング**  
「`func (r *Repo) GetBook(t BookTitle))(*Book, error) {`」ではなく、「`func (r *Repo) GetBook(t BookTitle)(*Book, error) {`」に修正。  
[WomenWhoGoTokyo](https://github.com/WomenWhoGoTokyo/book-reading-party)のみなさんご指摘ありがとうございました（2022/10/02）

**P81 リスト8.9　「errors.As」関数を使ってRDBMSのエラー情報を得る**  
「`return fmt.Errorf("store: cannot save book_id %d: %w", ErrAlreadyExists)`」ではなく、「`return fmt.Errorf("store: cannot save book_id %d: %w", book.ID, ErrAlreadyExists)`」に修正。  
[WomenWhoGoTokyo](https://github.com/WomenWhoGoTokyo/book-reading-party)のみなさんご指摘ありがとうございました（2022/10/02）

**P89 リスト9.1　無名関数は状態を持てる**  
`リスト9.1`ではなく、`リスト9.2`に修正。  
[@YuyaAbo](https://github.com/YuyaAbo) さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/22) ありがとうございました。


**P112 リスト12.6　ログミドルウェアの実装例**  
`WriteHeader`メソッド内の「`r.status = status`」を「`r.status = statusCode`」に修正。  
[@litencatt](https://github.com/litencatt)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/33) ありがとうございました（2022/08/15）


**P148 テストとコードカバレッジ取得の自動実行**  
「GitHub Actiuons上で実行したテスト結果」ではなく、「GitHub Actions上で実行したテスト結果」に修正。  
[@kdnakt](https://twitter.com/kdnakt)さんご指摘ありがとうございました（2022/08/06）

**P148 テストとコードカバレッジ取得の自動実行**  
「`.github/workflows/`ディレクトリ配下に`test.yml`を作成します。」ではなく、「`.github/workflows/`ディレクトリ配下に`test.yml`を作成します。また、ルートディレクトリに`.octocov.yml`を作成します」に修正。  
`.octocov.yml`のファイル内容は次のとおりです。  
https://github.com/budougumi0617/go_todo_app/blob/v1.0.4/.octocov.yml
```yaml
coverage:
  paths:
    - coverage.out
codeToTestRatio:
  code:
    - '**/*.go'
    - '!**/*_test.go'
  test:
    - '**/*_test.go'
testExecutionTime:
  if: true
diff:
  datastores:
    - artifact://${GITHUB_REPOSITORY}
comment:
  if: is_pull_request
report:
  if: is_default_branch
  datastores:
    - artifact://${GITHUB_REPOSITORY}
```

[@ac0mz](https://github.com/ac0mz)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/43) ありがとうございました（2022/09/04）

**P148 テストとコードカバレッジ取得の自動実行**  
macOSでgolangci-lintコマンドを実行時に「`can't extract issues from gofmt diff output`」のような主旨のエラーメッセージがでると[報告](https://github.com/budougumi0617/go_todo_app/discussions/74))されています。  
https://github.com/golangci/golangci-lint/issues/3087 を参考に`brew install diffutils`を実行して`diff`コマンドを置き換えてください。

[@nnabeyang](https://github.com/nnabeyang)さん[ご助言](https://github.com/budougumi0617/go_todo_app/discussions/74) ありがとうございました（2022/12/12）

**P163 リスト16.16　「httptest」パッケージを使った擬似的なHTTPリクエストのテスト**  
「`{"status": " ok"}`」ではなく、「`{"status": "ok"}`」に修正。  
[@kdnakt](https://twitter.com/kdnakt)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/29)ご指摘ありがとうございました（2022/08/10）

**P169 リスト17.4　HTTPハンドラー中で面倒なJSONレスポンス作成を簡略化**  
`import`文中の「`"log"`」を削除  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/61) ありがとうございました（2022/10/02）

**P167 リスト17.3　「store/store.go」に実装したタスクの簡易管理方法**  
`Tasks`変数は`&TaskStore{Tasks: map[int]*entity.Task{}}`ではなく、`&TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}`に修正。  
`TaskStore`構造体の`LastID`プロパティは`int`ではなく、`entity.TaskID`に修正。  
`TaskStore`構造体の`Tasks`プロパティは`map[int]*entity.Task`ではなく、`map[entity.TaskID]*entity.Task`に修正。  
`func (ts *TaskStore) Add(t *entity.Task) (int, error) {`ではなく、`func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {`に修正。
`func (ts *TaskStore) Get(id int) (*entity.Task, error) {`ではなく、`func (ts *TaskStore) Get(id entity.TaskID) (*entity.Task, error) {`
[@mizutec](https://twitter.com/mizutec)さん[ご指摘](https://twitter.com/mizutec/status/1555043156865208320)ありがとうございました（2022/08/06）  
[@Mo3g4u](https://github.com/Mo3g4u)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/25)ありがとうございました（2022/08/06）

**P172 SECTION-067タスクを登録するエンドポイントの実装**  
「リクエストの処理が正常が完了する場合」ではなく、「リクエストの処理が正常に完了する場合」に修正。  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/66) ありがとうございました（2022/10/02）

**P173 リスト17.7　「handler/add_task.go」のタスクを追加する実装**  
```go
err := at.Validator.Struct(b)
if err != nil {
```
ではなく、「`if err := at.Validator.Struct(b); err != nil {`」に修正。  
[@kdnakt](https://twitter.com/kdnakt)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/36)ご指摘ありがとうございました（2022/09/02）

**P173 リスト17.7　「handler/add_task.go」のタスクを追加する実装**  
`}{ID: id}` ではなく、`}{ID: int(id)}`に修正。（2022/08/06）

**P175 リスト17.8　ファイルを使った入出力の検証**  
`Tasks: map[int]*entity.Task{},` ではなく、`Tasks: map[entity.TaskID]*entity.Task{},`に修正。（2022/08/06）

**P175 リスト17.8　ファイルを使った入出力の検証**  
「`rspFile: "testdata/add_task/bad_req_rsp.json.golden",`」ではなく、「`rspFile: "testdata/add_task/bad_rsp.json.golden",`」に修正。  
[@fuuukeee3](https://github.com/fuuukeee3)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/51) ありがとうございました（2022/09/15）

**P176 リスト17.9　ゴールデンテストで利用しているJSONファイルの中身**  
「`// handler/testdata/add_task/bad_req_rsp.json.goldenの中身`」ではなく、「`// handler/testdata/add_task/bad_rsp.json.goldenの中身`」に修正。  
[@fuuukeee3](https://github.com/fuuukeee3)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/51) ありがとうございました（2022/09/15）


**P185 リスト18.4　MySQLコンテナの設定の追加前後を比較した「docker-compose.yml」の差分**  
「`TODO_DB_DATABASE: todo`」ではなく、「`TODO_DB_NAME: todo`」に修正。  
[@ac0mz](https://github.com/ac0mz)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/39) ありがとうございました（2022/09/04）

**P186 ローカルマシン上のMySQLコンテナにマイグレーションを実施する**  
「`make migrate コマンドを実行してマイグレーションを行います。`」という1文に下記の脚注リンクを追加。  
https://github.com/budougumi0617/go_todo_app/blob/v1.0.7/Makefile  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/62) ありがとうございました（2022/10/02）

**P187 ローカルマシン上のMySQLコンテナにマイグレーションを実施する**  
「`コメントを変更するなどしてから make drymigrateコマンドを実行するとマイグレーション用のDDLが生成される様子がわかります。`」という1文に下記の脚注リンクを追加。  
https://github.com/budougumi0617/go_todo_app/blob/v1.0.7/Makefile  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/62) ありがとうございました（2022/10/02）

**P187 SECTION-072 MySQL実行環境の構築**  
「GitHub Actionis上でMySQLコンテナを起動します。」ではなく、「GitHub Actions上でMySQLコンテナを起動します。」に修正。  
「GitHub Actionisではサービスコンテナという方法で」ではなく、「GitHub Actionsではサービスコンテナという方法で」に修正。  
[@kdnakt](https://twitter.com/kdnakt)さんご指摘ありがとうございました（2022/08/15）

**P192 リスト18.13 設定情報からDBへの接続を開く**  
`fmt.Sprint`関数を使って`sql.Open`関数にわたす接続用文字列を生成していますが、MySQLとの接続では`go-sql-driver`の`Config.FormatDSN`メソッドを利用できます。  
https://pkg.go.dev/github.com/go-sql-driver/mysql#Config.FormatDSN   
[@nnabeyang](https://github.com/nnabeyang)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/73) ありがとうございました（2022/12/12）

**P192 リスト18.13 設定情報からDBへの接続を開く**  
`sql.Open`関数の戻り値に対するエラーチェックで`return`する際の戻り値を「`nil, nil, err`」ではなく、「`nil, func() {}, err`」に修正。

合わせてP204 リスト19.4の差分を
```diff
url := fmt.Sprintf("http://%s", l.Addr().String()) log.Printf("start with: %v", url)
- mux := NewMux()
+ mux, cleanup, err := NewMux(ctx, cfg)
+ if err != nil {
+
+}
+ defer cleanup()
s := NewServer(l, mux)
```
ではなく、
```diff
url := fmt.Sprintf("http://%s", l.Addr().String()) log.Printf("start with: %v", url)
- mux := NewMux()
+ mux, cleanup, err := NewMux(ctx, cfg)
+ // エラーが返ってきてもcleanup関数は実行する
+ defer cleanup()
+ if err != nil {
+
+}
s := NewServer(l, mux)
```

に修正。

[@yamagit01](https://github.com/yamagit01)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/81)ありがとうございました（2023/01/29）

**P198 リスト18.19　「ListTasks」メソッドが期待されるデータを取得できるか検証**  
「`t.Fatalf("unexected error: %v", err)`」ではなく、「`t.Fatalf("unexpected error: %v", err)`」に修正。  
[@youta32449999](https://github.com/youta32449999)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/54) ありがとうございました（2022/09/15）

**P200 リスト18.22　「github.com/DATA-DOG/go-sqlmock」を使ったRDBMSを用いないテスト**  
「`t.Cleanup(func() {  db.Close() })`」ではなく、「`t.Cleanup(func() { _ = db.Close() })`」に修正。  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/64) ありがとうございました（2022/10/02）

**P200 リスト18.22　「github.com/DATA-DOG/go-sqlmock」を使ったRDBMSを用いないテスト**  
「`).WithArgs(okTask.Title, okTask.Status, c.Now(), c.Now()).`」ではなく、「` ).WithArgs(okTask.Title, okTask.Status, okTask.Created, okTask.Modified).`」に修正。  
[@halllllll](https://github.com/halllllll)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/76) ありがとうございました（2023/01/15）

**P204 HTTPハンドラーからRDBMSを使った永続化を行う**  
「`ru 関数の実装もNewMux関数のシグネチャの変更に伴ってリスト19.4のように変更しました。`」のあとに「`単純な動作確認のみのmux_test.goはここでファイルごと削除します。`」を追加。  
[@mom0tomo](https://github.com/mom0tomo)さん、[@mi-bear](https://github.com/mi-bear)さん、[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/83)ありがとうございました（2023/03/20）

**P216 「service」パッケージの実装**  
「`store` パッケージを利用して実際の登録データを組み立てる処理がリスト19.21の実装です。」の後に「リスト19.21中で利用する`UserRegister`インターフェースの定義はサンプルコードリポジトリの[`_chapter19/section79/service/interface.go`](https://github.com/budougumi0617/go_todo_app/blob/main/_chapter19/section79/service/interface.go)を参考にご用意ください。」  
[@youta32449999](https://github.com/youta32449999)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/55) ありがとうございました（2022/09/15）

**P218 リスト19.23　「store/repository.go」ファイルに定義する汎用エラー定義**  
「`// ErrCodeMySQLDuplicateEntry はMySQL系ののDUPLICATEエラーコード`」ではなく、「`// ErrCodeMySQLDuplicateEntry はMySQL系のDUPLICATEエラーコード`」に修正。  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/66) ありがとうございました（2022/10/02）

**P226 「KVS」型に対するテストを実装する**  
「Redisの接続情をの差分を吸収するため」ではなく、「Redisの接続情の差分を吸収するため」に修正。  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/66) ありがとうございました（2022/10/02）

**P232 リスト20.13　「go:embed」ディレクティブの動作確認**  
2回目の`t.Errorf`の引数は「`("want %s, but got %s", want, rawPubKey)`」ではなく、「`Errorf("want %s, but got %s", want, rawPrivKey)`」に修正。  
[@kdnakt](https://twitter.com/kdnakt)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/41)ご指摘ありがとうございました（2022/09/02）

```go
want = []byte("-----BEGIN PRIVATE KEY-----")
if !bytes.Contains(rawPrivKey, want) {
  t.Errorf("want %s, but got %s", want, rawPrivKey)
}
```

**P233 リスト20.14　「JWTer」構造体と初期化関数の定義**  
「`NewJWTer`」関数の実装を次のように修正。

```go
func NewJWTer(s Store, c clock.Clocker) (*JWTer, error) {
	j := &JWTer{Store: s}
	privkey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}
	pubkey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}
	j.PrivateKey = privkey
	j.PublicKey = pubkey
	j.Clocker = c
	return j, nil
}
```

[@kdnakt](https://twitter.com/kdnakt)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/44)ご指摘ありがとうございました（2022/09/04）

**P236 フィクスチャ関数の実装**  
「OSSを利用してしてダミーデータを生成」ではなく、「OSSを利用してダミーデータを生成」に修正。  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/66) ありがとうございました（2022/10/02）

**P236 「POST /tasks」でタスクを追加するときはユーザー情報をタスクに残す**  
「アクセストークンからユーザーIDをできるようになったので、」ではなく、「アクセストークンからユーザーIDを取得できるようになったので、」に修正。  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/66) ありがとうございました（2022/10/02）

**P244 SECTION-085 ユーザーログインエンドポイントの実装**  
「`LoginServiceインターフェースはauth/service.goに追記し，`」ではなく「`LoginServiceインターフェースはhandler/service.goに追記し，`」に修正。  
[@manaty226](https://github.com/manaty226)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/27)ご指摘ありがとうございました（2022/08/10）

**P244 リスト20.24　ログインを受け付けるハンドラーの実装**  
「`RespondJSON(r.Context(), w, rsp, http.StatusOK)`」ではなく「`RespondJSON(ctx, w, rsp, http.StatusOK)`」に修正。  
[@youta32449999](https://github.com/youta32449999)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/57) ありがとうございました（2022/09/16）

**P245 リスト20.26　「*handler.Login.ServeHTTP」メソッドに対する実装**  
「`rspFile: "testdata/login/bad_req_rsp.json.golden",`」ではなく、「`rspFile: "testdata/login/bad_rsp.json.golden",`」に修正。  
[@fuuukeee3](https://github.com/fuuukeee3)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/51) ありがとうございました（2022/09/15）

**P247 リスト20.27　「TestLogin_ServeHTTP」関数で利用しているJSONファイル**  
「`// handler/testdata/login/bad_req_rsp.json.golden`」ではなく、「`// handler/testdata/login/bad_rsp.json.golden`」に修正。  
[@fuuukeee3](https://github.com/fuuukeee3)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/51) ありがとうございました（2022/09/15）

**P247 「handler.Login」型に対するテストコードの実装**  
「`なお、handler/testdata/login/bad_req_rsp.json.goldenファイルの中身のエラーを見ると`」ではなく、「`なお、handler/testdata/login/bad_rsp.json.goldenファイルの中身のエラーを見ると`」に修正。  
[@fuuukeee3](https://github.com/fuuukeee3)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/51) ありがとうございました（2022/09/15）

**P248 リスト20.28　ログイン情報の検証とアクセストークンの生成を行う**  
「`if err = u.ComparePassword(pw); err != nil {`」ではなく、「`if err := u.ComparePassword(pw); err != nil {`」に修正。  
[@fuuukeee3](https://github.com/fuuukeee3)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/59) ありがとうございました（2022/09/15）

**P250 リスト20.32　「POST /login」エンドポイントを追加する**  
「`JWTer: jwter,`」ではなく、「`TokenGenerator: jwter,`」に修正。  
[@yskgc](https://github.com/yskgc)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/31)ご指摘ありがとうございました（2022/08/12）

**P250 リスト20.32　「POST /login」エンドポイントを追加する**  
「`jwter, err := auth.NewJWTer(rcli)`」ではなく、「`jwter, err := auth.NewJWTer(rcli, clocker)`」に修正。  
[@ac0mz](https://github.com/ac0mz)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/44) ありがとうございました（2022/09/04）

**P256 リスト20.40　「/tasks」エンドポイントに「AuthMiddleware」を適用する**
リスト20.40を差分で記載すると次のようなコード修正になります。
```diff
   at := &handler.AddTask{
     Service:   &service.AddTask{DB: db, Repo: &r},
     Validator: v,
   }
-  mux.Post("/tasks", at.ServeHTTP)
   lt := &handler.ListTask{
     Service: &service.ListTask{DB: db, Repo: &r},
   }
-  mux.Get("/tasks", lt.ServeHTTP)

+  mux.Route("/tasks", func(r chi.Router) {
+    r.Use(handler.AuthMiddleware(jwter))
+    r.Post("/", at.ServeHTTP)
+    r.Get("/", lt.ServeHTTP)
+  })
```
[@smirror](https://github.com/smirror)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/77) ありがとうございました（2023/01/14）


**P258 リスト20.42　「user_id」カラムへ対応した「*store.Repository.AddTask」メソッド**  
ページ脚注に「リスト20.42の変更に対応するテストコードの修正は https://github.com/budougumi0617/go_todo_app/blob/v1.0.7/store/task_test.go を参照のこと」を追記。  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/67) ありがとうございました（2022/10/02）

**P259 リスト20.44　ユーザーIDを使ってタスクを検索する**  
ページ脚注に「リスト20.44の変更に対応する`TaskLister`の修正は https://github.com/budougumi0617/go_todo_app/blob/v1.0.7/service/interface.go を参照のこと」を追記。  
[@y-magavel](https://github.com/y-magavel)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/67) ありがとうございました（2022/10/02）

**P262 「admin」ロールのユーザーのみがアクセス可能なエンドポイントを作成する**  
「ミドルウェアの適用順序に気をつけながら実装したのがapply_adminです。」ではなく、「ミドルウェアの適用順序に気をつけながら実装したのがリスト20.47の実装です。リスト20.47では、アクセストークンから取得したユーザーIDとロールを`http.Request`型の値に含まれる`context.Context`型の値に埋め込む`handler.AuthMiddleware`を実行したあとに、`handler.AdminMiddleware`を実行する順序でミドルウェアを適用しています。」に修正・追記。  
[@ac0mz](https://github.com/ac0mz)さん[ご指摘](https://github.com/budougumi0617/go_todo_app/discussions/45) ありがとうございました（2022/09/04）
