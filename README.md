# 環境変数
```
export GOPATH=$(pwd)
export CGO_LDFLAGS="`mecab-config --libs`"
export CGO_CFLAGS="-I`mecab-config --inc-dir`"
```

# 辞書のインストール
```
git clone --depth 1 https://github.com/neologd/mecab-ipadic-neologd.git
cd mecab-ipadic-neologd
./bin/install-mecab-ipadic-neologd -n
```

# 実行
```
dep ensure
go run main.go
```
