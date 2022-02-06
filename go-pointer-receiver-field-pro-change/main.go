package main

import (
	"fmt"
)

type samp interface {
  sampleFunc(string, int, int)
  samplePrint()
}

type sample struct {
  a string
  b int
  c int
}

func (sa *sample) sampleFunc(s string, i int, c int) {
  sa.a = s
  sa.b = i
  sa.c = c
}

func (sa *sample) samplePrint() {
  fmt.Printf("%v, %v, %v\n", sa.a, sa.b, sa.c)
  // data. 30, 40
}

func newSample() *sample {
  return &sample{}
}

func main() {

  // インターフェースでメソッドがちゃんと実装されているかを
  // 確かめることができる
  // インターフェースを介することで直接値を変更できない

  var sa samp = newSample()
  sa.sampleFunc("data", 30, 40)
  sa.samplePrint()
}

/*
外部パッケージに公開する時は
インターフェースとインスタンス生成関数を公開しそこから取得する

インターフェースとることで外部から使う際は構造体の中身を隠す

インターフェースを無理して使わなくとも良い
*/