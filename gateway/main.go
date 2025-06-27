package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Прокси для users
    router.Any("/users/*proxyPath", func(c *gin.Context) {
        proxyRequest(c, "http://user-service:5000", c.Param("proxyPath"))
    })

    // Прокси для products
    router.Any("/products", func(c *gin.Context) {
        proxyRequest(c, "http://product-service:8000", "/products")
    })

    router.Any("/products/*proxyPath", func(c *gin.Context) {
        proxyRequest(c, "http://product-service:8000", "/products"+c.Param("proxyPath"))
    })

    // Прокси для analytics
    router.Any("/analytics/*proxyPath", func(c *gin.Context) {
        proxyRequest(c, "http://analytics-service:5000", "/analytics"+c.Param("proxyPath"))
    })

    router.Run(":8080")
}

func proxyRequest(c *gin.Context, targetURL, proxyPath string) {
    // Простейший вариант — проксируем все методы методом Forward
    req, err := http.NewRequest(c.Request.Method, targetURL+proxyPath, c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
        return
    }

    req.Header = c.Request.Header

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request"})
        return
    }
    defer resp.Body.Close()

    c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}
