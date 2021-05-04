package model

type Category struct {
	Id             int32  `xorm:"pk autoincr comment('自增ID') BIGINT"`
	Name           string `xorm:"NOT NULL varchar(20)"`
	ParentCategory string
	Level          int  `xorm:"default 1 NOT NULL INT"`
	IsTab          bool `xorm:"default false NOT NULL INT"`
}

type Brands struct {
	Id   int32  `xorm:"pk autoincr comment('自增ID') BIGINT"`
	Name string `xorm:"varchar(20)"`
	Log  string `xorm:"varchar(20)"`
}

type Goods struct {
	Id          int32 `xorm:"pk autoincr comment('自增ID') BIGINT"`
	Category    Category
	Brands      Brands
	CategoryID  int
	BrandsID    int
	OnSale      bool  `xorm:"default true"`
	GoodsId     int32 `xorm:"not null comment('唯一货号')"`
	Name        string
	ClickNum    int `xorm:"default 0"`
	SoldNum     int `xorm:"default 0"`
	FavNum      int `xorm:"default 0"`
	Stocks      int `xorm:"default 0"`
	MarketPrice int `xorm:"default 0"`
	ShopPrice   int `xorm:"default 0"`
	GoodsBrief  string
	ShipFree    bool
	IsNew       bool `xorm:"default false comment('IsNew')"`
	IsHot       bool `xorm:"default false comment('IsHot')"`
}

type GoodsCategoryBrands struct {
	Id         int32 `xorm:"pk autoincr comment('自增ID') BIGINT"`
	Category   Category
	Brands     Brands
	CategoryID int
	BrandsID   int
}

type Banner struct {
	Id    int32  `xorm:"pk autoincr comment('自增ID') BIGINT"`
	Image string `xorm:"comment('Image')"`
	Url   string `xorm:"comment('Url')"`
	Index int    `xorm:"comment('Index')"`
}
