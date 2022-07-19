# golib


## GOPATH
```
GOPATH=/go/libs
#ここにパスを通すことでライブラリがインストールされる

go get ライブラリのパス(例　github.com/thoas/go-funk)
```

## go.mod、go.sum

goのライブラリ管理ツール

```
#初期化
go mod init プロジェクト名(golib)

#依存ライブラリのインストール　& ビルド
#ちなみにルートディレクトリでなくプログラムごとの階層に行き、下記コマンドを行うこと
#ルートディレクトリでコマンドを打っても再帰的に全てインストールされない
go build

#go.sumの内容からgo.modを作る
go mod tidy
```

エディタで自動保管する場合にはホスト機の中にライブラリが入っているので内部で入れ直すこと

### go sum

ライブラリの状態管理
composer.lockのようなもの

## migration tool

https://bitbucket.org/liamstask/goose/src/master/

<br>
create文

```
docker exec golib_app_1 goose -path=. create start sql
```

現状
```
 docker exec golib_app_1 goose -path=. status
```

migration実行
```
 docker exec golib_app_1 goose -path=. up
```

rollback
```
 docker exec golib_app_1 goose -path=. down
```

## gorm
goのORマッパー<br>
http://gorm.io/ja_JP/

## el
簡易ロガー<br>
https://github.com/go-easylog/el
