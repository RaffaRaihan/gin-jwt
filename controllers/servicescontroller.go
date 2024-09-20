package controllers

import (
	"encoding/json"
	"main/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetService(c *gin.Context) {
	var services []db.Services

	db.DB.Find(&services)
	c.JSON(http.StatusOK, gin.H{
		"services": services,
	})
}

func ReadService(c *gin.Context) {
	var services db.Services
	id := c.Param("ID")

	if err := db.DB.First(&services, id).Error; err != nil{
		switch err{
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error": "services not found",
				})
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"services": services,
	})
}

func CreateService(c *gin.Context) {
	var services db.Services

	if err := c.ShouldBindJSON(&services); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Create(&services)
	c.JSON(http.StatusOK, gin.H{
		"services": services,
	})
}

func UpdateService(c *gin.Context) {
	var services db.Services
	id := c.Param("ID")

	if err := c.ShouldBindJSON(&services); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if db.DB.Model(&services).Where("id = ?", id).Updates(&services).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "gagal mengupdate data services",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"services": "data berhasil di perbaharui",
	})
}

func DeleteService(c *gin.Context) {
	var services db.Services

	var input struct{
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := input.Id.Int64()
	if db.DB.Delete(&services, id).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "gagal menghapus data services",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"services": "data berhasil di hapus",
	})
}