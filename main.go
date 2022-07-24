package main

import (
	"fmt"

	"net"

	// pkg/errors エラー処理のパッケージ。今はアーカイブになっている
	"github.com/pkg/errors"
)

// main 実行関数
func main() {
	// run() の返り値をerrに格納。 errがnilでないなら
	if err := run(); err != nil {
		// エラー文を表示
		fmt.Printf("%+v", err)
	}
}

func run() error {
	// スタート表示
	fmt.Println("start tcp listen ...")

	// listenポート、サーバー機能の作成
	// localhost:12345でサーバーを立てることになる
	listen, err := net.Listen("tcp", "localhost:12345")

	if err != nil {
		return errors.WithStack(err)
	}
	// 関数の最後に closeする
	defer listen.Close()

	// コネクションを受け付ける
	conn, err := listen.Accept()

	if err != nil {
		return errors.WithStack(err)
	}

	// 関数の最後に closeする
	defer conn.Close()

	fmt.Println(">>> start")

	// make()・・・スライス作成
	buf := make([]byte, 1024)

	// Readメソッドの返り値が 0 byte なら全て Read したとしておく
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}
		fmt.Println(string(buf[:n]))
	}

	fmt.Println("<<< end")

	return nil
}
