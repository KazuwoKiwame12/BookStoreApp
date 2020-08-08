# api of book_store_application
このリポジトリは、本棚アプリのapiを作成し提供する。

## セットアップ
1. .env.exampleをコピーし、.envを用意する
```bash
$ cp .env.example
```

2. go.modにある依存関係を取得する
```bash
$ go mod download
```

## 注意事項
ファイルの変更をcommitする前に、以下のコマンドを行い
エラーがないか確認すること
参考文献: https://golang.org/cmd/vet/
```bash
$ go tool vet [ディレクトリ]
```
