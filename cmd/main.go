package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type UserAccount struct {
	ID       string `json:"id"`
	UserName string `json:"user-name"`
	Password string `json:"password"`
}

var accounts = []UserAccount{
	{"1", "alim", "992001"},
	{"2", "pavel", "123456"},
	{"3", "anton", "random"},
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	router := gin.Default()
	router.GET("/accounts", getAccounts)
	router.GET("/accounts/:id", getAccountByID)
	router.POST("/accounts", postAccount)

	domain := viper.GetString("domain")
	port := viper.GetString("port")
	address := domain + ":" + port

	err := router.Run(address)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func getAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, accounts)
}

func getAccountByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range accounts {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "account not found"})
}

func postAccount(c *gin.Context) {
	var newUserAccount UserAccount

	if err := c.BindJSON(&newUserAccount); err != nil {
		log.Fatal(err.Error())
		return
	}

	accounts = append(accounts, newUserAccount)
	c.IndentedJSON(http.StatusCreated, newUserAccount)
}
