package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// ""
	"github.com/programmer-for-good/flashcardApi/controllers"
	"github.com/programmer-for-good/flashcardApi/models"
)

func main() {
	db, err := Connect()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Card{})

	router := gin.Default()
	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Replace with the origin you want to allow
	router.Use(cors.New(config))

	cardsController := &controllers.CardsController{DB: db}

	router.POST("/Register", cardsController.RegisterUser)
	router.POST("/Login", cardsController.LoginUser)
	router.GET("/cards", cardsController.Index)
	router.GET("/cards/:id", cardsController.Show)
	router.POST("/flashcards", cardsController.CreateCard)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

// package main

// import (
// 	"encoding/json"
// 	"net/http"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Card struct {
// 	Term       string `json:"term"`
// 	Definition string `json:"definition"`
// }

// func main() {

// 	// Set up the database connection string
// 	dsn := "user:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local"

// 	// Connect to the database
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to connect to database!")
// 	}

// 	// Migrate the schema
// 	db.AutoMigrate(&Flashcard{})
// 	http.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {
// 		cards := []Card{
// 			{Term: "JavaScript", Definition: "A high-level, interpreted programming language used to create interactive effects within web browsers."},
// 			{Term: "React", Definition: "A JavaScript library for building user interfaces."},
// 			{Term: "Node.js", Definition: "A JavaScript runtime built on Chrome's V8 JavaScript engine."},
// 		}
// 		// // Enable CORS
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		// Set the response type to JSON
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(cards)
// 	})

// 	http.ListenAndServe(":8080", nil)
// }

// // Define a Flashcard struct to map to the flashcards table in the database
// type Flashcard struct {
// 	gorm.Model
// 	Term       string
// 	Definition string
// }
