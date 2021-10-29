package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-api/db"

	"github.com/gin-gonic/gin"
);

func main() {
	client:= db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func ()  {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	r := gin.Default()
	r.POST("/post",func(c *gin.Context) {
		err := createPost(client, ctx);
		if err != nil {
			panic(err)
		} 
		c.JSON(200,gin.H{
			"mssage":"post created",
		})
	})
	r.GET("/post",func(c *gin.Context) {
		err := getPostById(client, ctx, "ckvc4j8lb0154q70s58ovb5t0");
		if err != nil {
			panic(err)
		} 
		c.JSON(200,gin.H{
			"mssage":"post Fetched",
		})
	})
	r.GET("/ping",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"mssage":"pong",
		})
	})
	r.Run(":3000")
}
func getPostById(client *db.PrismaClient, ctx context.Context, id string) error{
	// find a single post
    post, err := client.Post.FindUnique(
        db.Post.ID.Equals(id),
    ).Exec(ctx)
	fmt.Print(post);
    if err != nil {
        return err
    }

    result, _ := json.MarshalIndent(post, "", "  ")
    fmt.Printf("post: %s\n", result)
	return nil
}
func createPost(client *db.PrismaClient, ctx context.Context) error {
	


	createPost, err := client.Post.CreateOne(
		db.Post.Title.Set("Hi from Prisma"),
		db.Post.Published.Set(true),
		db.Post.Desc.Set("Prisma is a database toolkit and makes databases easy."),
	).Exec(ctx)
	if err != nil {
		return err
	}
	result, _ := json.MarshalIndent(createPost,""," ")
	fmt.Printf("Created post: %s\n",result)
	return nil
}
