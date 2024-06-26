package productcontroller

import (
	"net/http"
	"github.com/AhmadHafidz1316/goAPI/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {
	var products models.Product

	id := c.Param("id")

	if err := models.DB.First(&products, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Data tidak di temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Product": products})
}

func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"Product": product})
}

func Update(c *gin.Context) {
	var product models.Product

	id := c.Param("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ? ", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Data berhasil di perbarui"})
}

func Delete(c *gin.Context) {

	var product models.Product

	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Data tidak di temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
			return
		}
	}

	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Data dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Data berhasil di hapus"})

}
