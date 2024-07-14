package user

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
)

var users = []*User{
    {
        Id: "sharath",
        AddUserBody: AddUserBody{
            Name: "Sharath",
            Age:  29,
            Address: Address{
                State:  "OH",
                City:   "Solon",
                Street: "Maplewood Drive",
                Hno:    "7533",
                },
            },
        },
}

func GetUsers() gin.HandlerFunc {
    return func(context *gin.Context) {
        context.JSON(200, gin.H{
            "data": users,
        })
    }
}

func GetUser() gin.HandlerFunc {
    return func(context *gin.Context) {
        var userParams UserRouteParam
        err := context.ShouldBindUri(&userParams)
        if err != nil {
            context.JSON(500, gin.H{
                "message": fmt.Sprintf("error parsing url params: %v", err),
            })
        }
        var user *User
        for _, u := range users {
            if u.Id == userParams.ID {
                user = u
                break
            }
        }
        if user != nil {
            context.JSON(200, gin.H{
                "data": user,
            })
            //			context.Done()
            return
        }
        log.Println(userParams)
        context.AbortWithStatus(404)
    }
}

func AddUser() gin.HandlerFunc {
    return func(context *gin.Context) {
        var user AddUserBody
        err := context.ShouldBindBodyWithJSON(&user)
        if err != nil {
            context.JSON(500, gin.H{
                "message": fmt.Sprintf("error parsing body: %v", err),
            }) 
        }
        users = append(users, NewUser(user))
    }
}

func UpdateUser() gin.HandlerFunc {
    return func(context *gin.Context) {
        var userBody AddUserBody
        var userParams UserRouteParam
        err := context.ShouldBindUri(&userParams)
        if err != nil {
            context.JSON(500, gin.H{
                "message": fmt.Sprintf("error parsing url params: %v", err),
            })
        }
        err = context.ShouldBindBodyWithJSON(&userBody)
        if err != nil {
            context.JSON(500, gin.H{
                "message": fmt.Sprintf("error parsing body: %v", err),
            }) 
        }
        for _, u := range users {
            if u.Id == userParams.ID {
                u.Update(userBody)
            }
        }
    }
}

func DeleteUser() gin.HandlerFunc {
    return func(context *gin.Context) {
        var userParams UserRouteParam
        err := context.ShouldBindUri(&userParams)
        if err != nil {
            context.JSON(500, gin.H{
                "message": fmt.Sprintf("error parsing url params: %v", err),
            })
        }
        var idx int = -1
        for i, u := range users {
            fmt.Print(u.Id, "->", userParams.ID)
            if u.Id == userParams.ID {
                idx = i
                break
            }
        }
        if idx == -1 {
            context.JSON(404, gin.H{
                "message": fmt.Sprintf("no user with id: %v", userParams.ID),
            }) 
        }
        users = append(users[:idx], users[idx+1:]...)
    }
}
