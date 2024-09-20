package controllers

import (
	"encoding/json"
	"main/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAc(c *gin.Context) {
	var ac []db.Ac

	db.DB.Find(&ac)
	c.JSON(http.StatusOK, gin.H{
		"ac": ac,
	})
}

func ReadAc(c *gin.Context) {
	var ac db.Ac
	id := c.Param("ID")

	if err := db.DB.First(&ac, id).Error; err != nil{
		switch err{
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error": "ac not found",
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
		"ac": ac,
	})
}

func CreateAc(c *gin.Context) {
	var ac db.Ac

	if err := c.ShouldBindJSON(&ac); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Create(&ac)
	c.JSON(http.StatusOK, gin.H{
		"ac": ac,
	})
}

func UpdateAc(c *gin.Context) {
	var ac db.Ac
	id := c.Param("ID")

	if err := c.ShouldBindJSON(&ac); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if db.DB.Model(&ac).Where("id = ?", id).Updates(&ac).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "gagal mengupdate data ac",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"ac": "data berhasil di perbaharui",
	})
}

func DeleteAc(c *gin.Context) {
	var ac db.Ac

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
	if db.DB.Delete(&ac, id).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "gagal menghapus data ac",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ac": "data berhasil di hapus",
	})
}