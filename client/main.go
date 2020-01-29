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

	client := proto.NewMathServiceClient(conn)
	g := gin.Default()

	g.GET("add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter `a`."})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter `b`."})
			return
		}

		req := &proto.MathRequest{A: int64(a), B: int64(b)}
		if response, err := client.Add(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		}
	})

	g.GET("divide/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter `a`."})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter `b`."})
			return
		}

		req := &proto.MathRequest{A: int64(a), B: int64(b)}
		if response, err := client.Divide(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		}
	})

	g.GET("multiply/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter `a`."})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter `b`."})
			return
		}

		req := &proto.MathRequest{A: int64(a), B: int64(b)}
		if response, err := client.Multiply(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		}
	})

	g.GET("subtract/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter `a`."})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter `b`."})
			return
		}

		req := &proto.MathRequest{A: int64(a), B: int64(b)}
		if response, err := client.Subtract(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		}
	})

	if err := g.Run(":4002"); err != nil {
		log.Fatalf("Failed to run server: %v.", err)
	}
}
