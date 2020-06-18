package entity

// メニュー構造体
type menu struct {
	ID int // ID
	name string // 商品名
	price int // 値段
	text string // キャプション
}

// オーダーの構造体
type order struct {
	date string // 日付
	time string // 時刻
	ID int // ID
	name string // 商品名
	amount int // 数量
	totalprice int // 総額
}