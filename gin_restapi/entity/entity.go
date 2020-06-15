package entity

// メニュー構造体
type menu struct {
	ID int // ID
	name string // 商品名
	price int // 値段
	text string // キャプション
}

// 注文の構造体
type order struct {
	ID int // ID
	Name string // 商品名
	Amount int // 数量
	Price int // 総額
	Date string // 注文時刻
}