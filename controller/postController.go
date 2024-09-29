package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/happynet78/goblogbackend/database"
	"github.com/happynet78/goblogbackend/models"
	"github.com/happynet78/goblogbackend/util"
	"math"
	"strconv"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("unable to parse body")
	}
	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Congratulation!, Your post is live",
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(models.Blog{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": getblog,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id = ?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("unable to parse body")
	}
	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"message": "post updated successfully",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)
	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id = ?", id).Preload("user").Find(blog)

	return c.JSON(blog)
}
