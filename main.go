package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args[0]はコマンド名(バイナリ名)になる

	// コマンドパラメーターを受け取る場合はos.Args[1]以降を取得して使う
	for _, o := range os.Args[1:] {
		fmt.Print(o)
	}
}
