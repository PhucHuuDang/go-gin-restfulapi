package handler

//
// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strconv"
//
// 	"golang-restapi/forms"
// 	"golang-restapi/models"
//
// 	"github.com/gin-gonic/gin"
//
// 	jwt "github.com/golang-jwt/jwt/v4"
// )
//
// type AuthController struct{}
//
// var authModel = new(models.AuthModel)
//
// func (ctk AuthController) TokenValid(c *gin.Context) {
// 	tokenAuth, err := authModel.ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized, gin.H{"message": "Please login first"})
// 		return
// 	}
//
// 	userID, err := authModel.VerifyToken(c.Request)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}
//
// }
