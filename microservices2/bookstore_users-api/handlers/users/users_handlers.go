package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/helpers/customerr"
	"bookstore_users-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {

		restErr := customerr.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//todo handle error
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {

	// 	//TODO handle json error
	// 	log.Println(err)
	// 	return
	// }

	result, err := services.CreateUser(user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func FindUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implement me")
}
