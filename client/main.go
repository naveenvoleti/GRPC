package main

import (
	pb "example.com/grpc/grpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const (
	address = "localhost:54321"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	c := pb.NewAddServiceClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	g := gin.Default()
	g.Static("/assets", "./assets")
	g.LoadHTMLGlob("index.html")
	g.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	g.GET("/chat/:a", func(ctx *gin.Context) {
		a := ctx.Param("a")
		//b := strings.Replace(a, '/', "%2F", 1)
		req := &pb.Request{A: a}
		if r, err := c.Chat(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(r.Result),
			})
			fmt.Printf("Reuslt: %s\n", r.Result)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/chat", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
