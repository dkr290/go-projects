package handlers

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/helpers/customerr"
	"bookstore_users-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserid(userIdParam string) (int64, *customerr.RestError) {

	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, customerr.NewBadRequestErr("invalid user id")

	}

	return userId, nil
}

func GetUser(c *gin.Context) {

	userId, idErr := getUserid(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr.Message)
		return

	}

	result, err := services.GetUser(userId)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.MarshallFn(c.GetHeader("X-Public") == "true"))

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

	c.JSON(http.StatusCreated, result.MarshallFn(c.GetHeader("X-Public") == "true"))

}

func UpdateUser(c *gin.Context) {

	userId, idErr := getUserid(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr.Message)
		return

	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {

		restErr := customerr.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.MarshallFn(c.GetHeader("X-Public") == "true"))

}

func DeleteUser(c *gin.Context) {

	userId, idErr := getUserid(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr.Message)
		return

	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {

	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
