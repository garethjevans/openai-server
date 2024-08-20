package server

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

func EmbeddingHandler(c *gin.Context) {
	var req openai.EmbeddingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	handleEmbedding(&req, c)
}

func handleEmbedding(req *openai.EmbeddingRequest, c *gin.Context) {
	res := openai.EmbeddingResponse{
		Object: "list",
		Data: []openai.Embedding{
			{Object: "embedding",
				Index: 0,
				Embedding: []float32{-0.0030878915,
					0.008527344,
					0.0036363655,
					-0.05539084,
					-0.04473669,
					0.03561223,
					0.015806785,
					0.0021955732,
					-0.0052868193,
					0.011660523,
					-0.012861462,
					0.026071802,
					0.00896344,
					0.002973836,
					0.0025729635}},
		},
		Model: req.Model,
	}
	rand.Seed(time.Now().UnixNano())

	//for i := 0; i < req.N; i++ {
	//	completionStr := RandomContent()
	//
	//	res.Choices = append(res.Choices, openai.ChatCompletionChoice{
	//		Message: openai.ChatCompletionMessage{
	//			Role:    openai.ChatMessageRoleAssistant,
	//			Content: completionStr,
	//		},
	//		Index: i,
	//	})
	//}
	inputTokens := numTokens(req.Input[0])

	res.Usage = openai.Usage{
		PromptTokens: inputTokens,
		TotalTokens:  inputTokens,
	}
	c.JSON(http.StatusOK, res)
}
