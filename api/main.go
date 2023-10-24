package main

// Go言語：文法基礎まとめ | https://qiita.com/HiromuMasuda0228/items/65b9a593275f769f6b69
func main() {
	router := newRouter()
	router.Logger.Fatal(router.Start(":8080"))
}
