package model

// Migrate 数据库迁移
func Migrate() {
	DB.AutoMigrate(&Goods{})
}
