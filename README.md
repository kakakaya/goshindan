# goshindan
Goから診断メーカー（ <https://shindanmaker.com> ）へのアクセスをする

# コマンド
## shindan
診断メーカーでの診断を行い、結果を印字する。用例:

`% goshindan shindan -s 509717 -u kakakaya --append-url`

* `--shindan-id / -s` 診断メーカーのID。もし <https://shindanmaker.com/509717> を使って診断を行いたい場合、509717を指定する。
* `--username / -u` ユーザー名(「診断したい名前を入れて下さい」の部分)。
* `--append-url / --add-url` 診断メーカーのURLを末尾に付加して出力する。
