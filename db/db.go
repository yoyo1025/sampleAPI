package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connceted")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}

/*
このコードは、Go言語を使用してPostgreSQLデータベースに接続するためのものです。gormパッケージを使用しています。gormは、Go言語用の人気のあるORM（Object-Relational Mapping）ライブラリで、データベース操作をより簡単に行うことができます。

コードは大きく2つの関数から構成されています。

NewDB 関数
環境変数GO_ENVが"dev"に設定されている場合、godotenv.Load()を呼び出して.envファイルから環境変数を読み込みます。これにより、開発環境でのデータベース接続情報などの設定をファイルから読み込むことができます。
環境変数からPostgreSQLデータベースへの接続情報を読み取り、接続用のURLを組み立てます。
gorm.Openメソッドを使用して、組み立てたURLを使ってデータベースに接続を試みます。
接続に成功した場合、"Connected"とコンソールに出力し、*gorm.DBオブジェクトを返します。このオブジェクトを通じてデータベース操作を行うことができます。
もし接続に失敗した場合、エラーメッセージをログに記録してプログラムを終了します。
CloseDB 関数
gorm.DBオブジェクトから、実際のデータベース接続を表すsql.DBオブジェクトを取得します。
sql.DB.Closeメソッドを呼び出してデータベース接続を閉じます。エラーが発生した場合は、そのエラーをログに記録してプログラムを終了します。
このコードによって、Go言語のアプリケーションからPostgreSQLデータベースへの接続と切断を簡単に行うことができます。開発環境と本番環境での使い分けも、環境変数を使って柔軟に行うことが可能です。
*/