package controllers

import (
	"encoding/json"
	"main/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoles(c *gin.Context) {
	var roles []db.Roles

	db.DB.Find(&roles)
	c.JSON(http.StatusOK, gin.H{
		"roles": roles,
	})
}

func ReadRoles(c *gin.Context) {
	var roles db.Roles
	id := c.Param("ID")

	if err := db.DB.First(&roles, id).Error; err != nil{
		switch err{
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error": "roles not found",
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
		"roles": roles,
	})
}

func CreateRoles(c *gin.Context) {
	var roles db.Roles

	if err := c.ShouldBindJSON(&roles); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Create(&roles)
	c.JSON(http.StatusOK, gin.H{
		"roles": roles,
	})
}

func UpdateRoles(c *gin.Context) {
	var roles db.Roles
	id := c.Param("ID")

	if err := c.ShouldBindJSON(&roles); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if db.DB.Model(&roles).Where("id = ?", id).Updates(&roles).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "gagal mengupdate data roles",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"roles": "data berhasil di perbaharui",
	})
}

func DeleteRoles(c *gin.Context) {
	var roles db.Roles

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
	if db.DB.Delete(&roles, id).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "gagal menghapus data roles",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"roles": "data berhasil di hapus",
	})
}