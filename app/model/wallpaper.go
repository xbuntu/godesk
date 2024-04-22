package model

const TableNameWallpaper = "wallpaper"
const TableNameWallpaperCate = "wallpaper_cate"

type Wallpaper struct {
	ID    int    `gorm:"column:id;primaryKey" json:"id"`
	CID   int    `gorm:"column:cid" json:"cid"`
	Title string `gorm:"column:title" json:"title"`
	Tag   string `gorm:"column:tag" json:"tag"`
	Url   string `gorm:"column:url" json:"url"`
}

func (*Wallpaper) TableName() string {
	return TableNameWallpaper
}

type WallpaperCate struct {
	ID    int    `gorm:"column:id;primaryKey" json:"id"`
	Title string `gorm:"column:title" json:"title"`
}

func (*WallpaperCate) TableName() string {
	return TableNameWallpaperCate
}
