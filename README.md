# Discontinued

Due to change of shindanmaker.com, this project no longer maintained.

# goshindan

[![Travis](https://img.shields.io/travis/kakakaya/goshindan.svg)](https://travis-ci.org/kakakaya/goshindan)
[![Coveralls](https://img.shields.io/coveralls/kakakaya/goshindan.svg)](https://coveralls.io/github/kakakaya/goshindan)
[![GoDoc](https://godoc.org/github.com/kakakaya/goshindan?status.png)](https://godoc.org/github.com/kakakaya/goshindan)
[![Go Report Card](https://goreportcard.com/badge/github.com/kakakaya/goshindan)](https://goreportcard.com/report/github.com/kakakaya/goshindan)

Goから診断メーカー（ <https://shindanmaker.com> ）へのアクセスをするコマンドラインツール及びライブラリ

# ライブラリ
GoDocを見てください。

[![GoDoc](https://godoc.org/github.com/kakakaya/goshindan?status.png)](https://godoc.org/github.com/kakakaya/goshindan)


# コマンドラインツール
## shindan
診断メーカーでの診断を行い、結果を印字する。用例:

`% goshindan shindan -s 509717 -u kakakaya --append-url`

* `--shindan-id / -s` 診断メーカーのID。もし <https://shindanmaker.com/509717> を使って診断を行いたい場合、509717を指定する。
* `--username / -u` ユーザー名(「診断したい名前を入れて下さい」の部分)。
* `--append-url / --add-url` 診断メーカーのURLを末尾に付加して出力する。

# Screencast

[![asciicast](https://asciinema.org/a/8ukl2p62u48748uhqb0fkq7fl.png)](https://asciinema.org/a/8ukl2p62u48748uhqb0fkq7fl)
