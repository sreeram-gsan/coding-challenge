package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/sreeram-gsan/coding-challenge/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// Struct to store the metadata of inventory table.
type item struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Quantity  int32   `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

// Struct to return response message.
type ResponseMessage struct {
	Message string
}

// getItems      godoc
// @Summary      Get all items from inventory.
// @Description  Responds with the list of all items as JSON.
// @Tags         items
// @Produce      json
// @Success      200  {array}  item
// @Router       /items [get]
func getItems(c *gin.Context) {
	var items = []item{}
	getDBConnection().Table("inventory").Find(&items)
	c.IndentedJSON(http.StatusOK, items)
}

// getItemById      godoc
// @Summary      Get a single item from inventory by id.
// @Description  Returns the item whose id value matches the provided id.
// @Tags         items
// @Produce      json
// @Param        id  path integer true  "search item by id"
// @Success      200  {object} item
// @Router       /items/{id} [get]
func getItemById(c *gin.Context) {
	var result_item = item{}
	id := c.Param("id")
	res := getDBConnection().Table("inventory").First(&result_item, "id = ?", id)

	// If record is not found, return a message
	err := res.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusOK, ResponseMessage{Message: "Record not found!"})
	} else {
		c.IndentedJSON(http.StatusOK, result_item)
	}
}

// addItems             godoc
// @Summary      Add a new item to the inventory.
// @Description  Takes a item JSON and store in DB. Return saved JSON.
// @Tags         items
// @Produce      json
// @Param        item  body item  true  "item JSON"
// @Success      200   {object}  item
// @Router       /items [post]
func addItems(c *gin.Context) {
	var newItem item

	// Binding the JSON received to the item object
	err := c.BindJSON(&newItem)
	if err != nil {
		return
	}

	// createInventoryTable(db)
	res := getDBConnection().Table("inventory").Create(newItem)
	create_err := res.Error

	if create_err != nil && strings.Contains(create_err.Error(), "1062") {
		c.IndentedJSON(http.StatusOK, ResponseMessage{Message: "Duplicate entry!"})
	} else {
		c.IndentedJSON(http.StatusCreated, newItem)
	}
}

// deleteItemById             godoc
// @Summary      Delete an existing item from inventory by id.
// @Description  Takes an id and deletes its value in DB.
// @Tags         items
// @Produce      json
// @Param        id  path integer true  "delete item by id"
// @Success      200   {object}  item
// @Router       /items/{id} [delete]
func deleteItemById(c *gin.Context) {
	var result_item item
	id := c.Param("id")
	res := getDBConnection().Table("inventory").First(&result_item, "id = ?", id)

	// If record is not found, return a message
	err := res.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusOK, ResponseMessage{Message: "Record not found!"})
	} else {
		getDBConnection().Table("inventory").Delete(&item{}, id)
		c.IndentedJSON(http.StatusOK, ResponseMessage{Message: "Record deleted!"})
	}
}

// patchItems             godoc
// @Summary      Update an existing item in inventory.
// @Description  Takes a item JSON and updates its value in DB.
// @Tags         items
// @Produce      json
// @Param        item  body item  true  "item JSON"
// @Param        id  path integer true  "update item by id"
// @Success      200   {object}  item
// @Router       /items/{id} [patch]
func patchItems(c *gin.Context) {
	var updatedItem item
	var existingItem item
	id := c.Param("id")

	// Binding the JSON received to the item object
	err := c.BindJSON(&updatedItem)
	if err != nil {
		return
	}

	// If record is not found, return a message
	res := getDBConnection().Table("inventory").First(&existingItem, "id = ?", id)
	patch_err := res.Error
	if errors.Is(patch_err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusOK, ResponseMessage{Message: "Record not found!"})
	} else {
		getDBConnection().Table("inventory").Save(&updatedItem)
		c.IndentedJSON(http.StatusCreated, updatedItem)
	}
}

// getItemsAsCSV             godoc
// @Summary      Download CSV with all items from inventory.
// @Description  Download CSV with all items from inventory.
// @Tags         items
// @Produce      text/csv
// @Success      200   {text}  item
// @Router       /items/csv [GET]
func getItemsAsCSV(c *gin.Context) {
	// Getting http writer from gin context.
	var w http.ResponseWriter = c.Writer

	// Getting all items from Inventory DB.
	var items = []item{}
	getDBConnection().Table("inventory").Find(&items)

	// Preparing Buffer and Writer to write to CSV.
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)

	//Writing the headers
	wr.Write([]string{"Id", "Name", "Quantity", "UnitPrice"})
	wr.Flush()

	// Iterating through list of items.
	for i := 0; i < len(items); i++ {
		item := items[i]
		v := reflect.ValueOf(item)
		values := make([]string, v.NumField())

		// Iterating every row and converting the column value to string.
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				values[i] = strconv.FormatInt(v.Field(i).Int(), 10)
			case reflect.Float32, reflect.Float64:
				values[i] = fmt.Sprintf("%f", v.Field(i).Float())
			case reflect.String:
				values[i] = v.Field(i).String()
			}
		}
		wr.Write(values)
		// Writes the csv writer data to buffer.
		wr.Flush()
	}

	// Returning a response as a CSV file.
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=inventory.csv")
	w.Write(b.Bytes())
}

// @title     Inventory API
// @version         1.0
// @description     An inventory management API with Go using Gin framework.
// @termsOfService  https://sreeramganesan.com

// @contact.name   Sreeram Ganesan
// @contact.url    https://sreeramganesan.com
// @contact.email  srga8641@colorado.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /inventory/v1
func main() {
	// Load environment variables.
	LoadEnv()

	// Create router and register all end points.
	router := gin.Default()

	// Allow CORS. Only for demo purpose.
	router.Use(CORSMiddleware())

	// Route for Swagger docs.
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	prefix := os.Getenv("SERVER_ENDPOINT_PREFIX_V1")
	router.GET(prefix+"/items", getItems)
	router.GET(prefix+"/items/:id", getItemById)
	router.GET(prefix+"/items/csv", getItemsAsCSV)
	router.POST(prefix+"/items", addItems)
	router.PATCH(prefix+"/items/:id", patchItems)
	router.DELETE(prefix+"/items/:id", deleteItemById)
	router.Run(os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT"))
}
