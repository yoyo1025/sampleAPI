package main

import (
	"fmt"
	"sampleAPI/db"
	"sampleAPI/model"
)

// migrationのコマンドを実行するときはmigrationの処理を
// プログラムのエントリポイントに配置したいのでmain関数を使用し、mainパッケージに含める

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{}) // データベースに反映させたいモデル構造を渡す
}