package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordanpatton/go-rpc-sandbox/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewServiceClient(conn)
	g := gin.Default()

	g.GET("calculate/:x/:operator/:y", func(ctx *gin.Context) {
		operator := ctx.Param("operator")

		x, err := strconv.ParseUint(ctx.Param("x"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter `x`."})
			return
		}

		y, err := strconv.ParseUint(ctx.Param("y"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter `y`."})
			return
		}

		req := &proto.CalculateRequest{Operator: operator, X: int64(x), Y: int64(y)}
		if response, err := client.Calculate(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"output": fmt.Sprint(response.Output)})
		}
	})

	if err := g.Run(":4002"); err != nil {
		log.Fatalf("Failed to run client: %v.", err)
	}
}
