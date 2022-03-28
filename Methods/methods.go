package Methods

import (
	"kursach/CORS"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateAutoInput struct {
	Marka string `json:"marka"`
	Model string `json:"model"`
}

func GetAllAuto(c *gin.Context) {
	gormDB := CORS.ConnectDB()
	var autos []CORS.Auto
	gormDB.Find(&autos)
	c.JSON(http.StatusOK, autos)
}

func AddAuto(c *gin.Context) {
	gormDB := CORS.ConnectDB()
	var input CORS.Auto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	auto := CORS.Auto{Marka: input.Marka, Model: input.Model}
	gormDB.Create(&auto)
	c.JSON(http.StatusOK, gin.H{"message": "Auto added"})

}
func UpdateAuto(c *gin.Context) {
	gormDB := CORS.ConnectDB()
	var auto CORS.Auto
	if err := gormDB.Where("id = ?", c.Param("id")).First(&auto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateAutoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gormDB.Model(&auto).Update("marka", input.Marka)
	gormDB.Model(&auto).Update("model", input.Model)
	c.JSON(http.StatusOK, gin.H{"message": "Auto updated"})

}
func DelAuto(c *gin.Context) {
	gormDB := CORS.ConnectDB()
	var auto CORS.Auto
	if err := gormDB.Where("id = ?", c.Param("id")).First(&auto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	gormDB.Delete(&auto)

	c.JSON(http.StatusOK, gin.H{"autos": true})
}
