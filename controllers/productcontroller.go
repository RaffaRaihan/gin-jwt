package controllers

import (
	"encoding/json"
	"main/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var product []db.Product

	db.DB.Find(&product)
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func Show(c *gin.Context) {
	var product db.Product
	id := c.Param("id")

	if err := db.DB.First(&product, id).Error; err != nil{
		switch err{
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error": "product not found",
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
		"product": product,
	})
}

func Create(c *gin.Context) {
	var product db.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func Update(c *gin.Context) {
	var product db.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if db.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "gagal mengupdate data product",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"product": "data berhasil di perbaharui",
	})
}

func Delete(c *gin.Context) {
	var product db.Product

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
	if db.DB.Delete(&product, id).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "gagal menghapus data product",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": "data berhasil di hapus",
	})
}