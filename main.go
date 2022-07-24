package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

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
	fmt.Println("TCPセツゾク　ヲ　カイシ　シマス ...")

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

	fmt.Println(">>> スタート　シマス")

	scanner := bufio.NewScanner(conn)
	// ヘッダのContent-Length のバイト数を取得する
	var contentLength int

	// 一行ずつ処理
	for scanner.Scan() {
		// リクエストヘッダーを表示する
		// Text()からの返り値が空文字であれば空と判断する
		line := scanner.Text()
		if line == "" {
			break
		}

		if strings.HasPrefix(line, "Content-Length") {
			contentLength, err := strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1]))
			if err != nil {
				return errors.WithStack(err)
			}
		}
		fmt.Println(line)
	}

	// non-EOF errorがある場合
	if scanner.Err() != nil {
		return scanner.Err()
	}

	// リクエストボディ
	buf := make([]byte, contentLength)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("BODY: ", string(buf))

	if scanner.Err() != nil {
		return scanner.Err()
	}

	// // make()・・・スライス作成

	// // Readメソッドの返り値が 0 byte なら全て Read したとしておく
	// for {
	// 	n, err := conn.Read(buf)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	if err != nil {
	// 		return errors.WithStack(err)
	// 	}
	// 	fmt.Println(string(buf[:n]))
	// }

	fmt.Println("<<< セツゾク オワリマス")

	return nil
}
