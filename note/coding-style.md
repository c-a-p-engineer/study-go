## コーディング規約

### インデント
* タブを使用
  * 半角スペースは使用しない

### 文字数
* 1行あたりの文字数の制限はない
  * 長すぎる場合は開業しても良い

### コメント
* `/* */` ブロックコメント
* `//` 行コメント
* 連続したコメントは縦に揃える
```go
type T struct {
	name	string	// name of the object
	value	int			// its value
}
```

### 演算子
* 演算子の前後にスペースを入れない。
  * 演算子の優先順位を明確にする時のみスペースを入れる。括弧は使用しない
```go
x << 8 + y << 16
```

### 自動整形
以下のコマンドを使用することで go が自動的に整形をしてくれます
```shell
go fmt hoge.go
```

## 命名規則
* ローワーキャメルケース
  * hogePiyo
* アッパーキャメルケース
  * HogePiyo

### ディレクトリ名
* 決まりなし？

### ファイル名
* 決まりなし？

### 関数、構造体
* 内部向け
    * ローワーキャメルケース
```go
// packageの内に公開する場合
func hogePiyo(filename string) (string, error) {}
```
* 外部向け
    * アッパーキャメルケース

```go
// packageの外に公開する場合
func HogePiyo(filename string) (string, error) {}
```

### type
* ローワーキャメルケース

### コンストラクタ
* NewBook（New + 生成対象の構造体名）

### インターフェイス
* 1つのメソッドを持つインターフェイスの場合
  * Reader（Read メソッドだけを持つインタフェース）

## 参考情報
* [Effective GO - The Go Programming Language](https://golang.org/doc/effective_go)