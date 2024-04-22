package controller

import (
	"godesk/app/model"
	"gorm.io/gorm"
)

// Config 配置
type Config struct {
	Base
}

func NewConfig() *Config {
	return &Config{}
}

// GetConfig 获取配置项
func (c *Config) GetConfig() Res {
	var list []model.Config
	if err := c.db().Find(&list).Error; err != nil {
		c.log(err.Error())
		return c.error("未查询到数据")
	}
	return c.success(list)
}

// AddConfig 添加配置项
func (c *Config) AddConfig(config model.Config) Res {
	if err := c.db().Create(&config).Error; err != nil {
		c.log(err.Error())
		return c.error("添加失败")
	}
	return c.success(nil)
}

// SaveConfig 修改配置项
func (c *Config) SaveConfig(list []model.Config) Res {
	err := c.db().Transaction(func(tx *gorm.DB) error {
		for _, config := range list {
			if err := tx.Save(&config).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		c.log(err.Error())
		return c.error("更新失败")
	}
	return c.success(nil)
}

// DelConfig 删除配置项
func (c *Config) DelConfig(config model.Config) Res {
	if err := c.db().Delete(&config).Error; err != nil {
		c.log(err.Error())
		return c.error("删除失败")
	}
	return c.success(nil)
}
