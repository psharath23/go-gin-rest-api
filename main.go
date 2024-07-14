package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "go-rest-api/src/user"
)

func main() {
    r := gin.Default()

    v1 := r.Group("/api/v1")
    {
        userRoute := v1.Group("/user")
        {
            userRoute.GET("/getUsers", user.GetUsers())
            userRoute.GET("/getUser/:id", user.GetUser())
            userRoute.POST("/addUser", user.AddUser())
            userRoute.PUT("/updateUser/:id", user.UpdateUser())
            userRoute.DELETE("/deleteUser/:id", user.DeleteUser())
        }
    }
    err := r.Run()
    if err != nil {
        log.Fatal(err)
    }
}
