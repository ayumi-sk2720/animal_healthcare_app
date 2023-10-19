package main

// [参考] Mysql,Go,Echo,GORM,sql-migrateでCRUDなAPIを作成するまで | https://qiita.com/yusuke1120925/items/a6e05178cf6e6bf1f030
func main() {
	router := newRouter()
	router.Logger.Fatal(router.Start(":8080"))
}
