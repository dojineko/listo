Listo
=====

Adaptable snippet manager on Alfred workflow 

<img src="./icon.png" alt="image" width="320">

"listo" means...

- list (to Alfread)
- with Squirrel [in Japanese, リスと(risu-to)]
- wise [in Spanish]

illustration by いらすとや

----

## :apple: つかいかた

### データを検索する

デフォルトでは `⌘` + `Shift` + `L` で listo が起動します。  
検索したいキーワードを入力することで、追加済みのデータ(ストレージ)から検索を行います。

#### 初期状態
- 入力されたキーワードを元にすべてのストレージから検索を行います。
- マッチするストレージの行がリストに表示されます。

```sh
# 追加済みのストレージから りんご を含む行を絞り込む
りんご
```

#### `@` から始まるキーワード (ストレージの絞込)
- 続くキーワードでストレージの絞込を行います。
- マッチする名前のストレージがリストに表示されます。

```sh
# hoge.tsv から ゴリラ を含む行を絞り込む
@hoge.tsv ゴリラ
```

#### `:` から始まるキーワード (行の選択と列の絞込)
- ストレージの絞込と組み合わせて、行選択します。
- 行内の項目がリストに表示されます。
- 更にキーワードを入力すると列の絞込を行います。

```sh
# hoge.tsv の 10 行名から ラッパ を含む列を絞り込む
@hoge.tsv :10 ラッパ
```

### データを追加する

Alfredを表示し `listo-install` と入力します。    
TSV ファイル名を入力し、確定するとストレージにコピーして追加します。

```sh
# デスクトップに有る hoge.tsv を追加する
listo-install ~/Desktop/hoge.tsv
```

### データを削除する
Alfredを表示し `listo-remove` と入力します。  
追加済みのストレージの一覧から削除したいデータを選択することで削除できます。

```sh
# hoge.tsv を削除する
listo-remove hoge.tsv
```
