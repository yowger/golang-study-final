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
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

	projection := bson.M{
		"_id": 0,
	}
	cursor, err := h.Collection.Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	defer cursor.Close(ctx)

	fmt.Println("cursor", cursor)
	fmt.Println()

	var products []models.Product
	for cursor.Next(ctx) {
		fmt.Println("next cursor ->", cursor)
		fmt.Println()
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			fmt.Println("cursor error ->", cursor)
			fmt.Println("cursor error description ->", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		products = append(products, product)
	}

	return c.JSON(http.StatusOK, products)
}
