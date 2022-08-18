# Named return values

戻り値に名前をつけられる。関数の最初で定義したものとみなされる。

# Short variable declarations

var宣言の代わりに:=の代入分で暗黙的な方宣言が可能

```go

var i int=3
i:=3
```

# Type inference

明示的に型を指定しない場合、右側の変数から型推論される

# Methods continued

レシーバを伴うメソッド宣言できるのはレシーバ型が同じパッケージにある必要がある

# Interfaces are implemented implicitly

明示的に実装することを宣言しなくていい