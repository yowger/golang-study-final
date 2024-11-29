package handlers

import (
	"context"
	"echo-mongo-api/models"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ProductHandler struct {
	Collection *mongo.Collection
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product models.Product
	if err := c.Bind(&product); err != nil {
		fmt.Println("error binding product", err)
		fmt.Println("Product", product)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	fmt.Println("still here baby")

	result, err := h.Collection.InsertOne(ctx, product)
	if err != nil {
		fmt.Println("error inserting product", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": result.InsertedID})
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := h.Collection.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	defer cursor.Close(ctx)

	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return c.JSON(http.StatusOK, products)
}
