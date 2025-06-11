package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Маршрут для пользователей
    router.Any("/users/*proxyPath", func(c *gin.Context) {
        proxyRequest(c, "http://user-service:5000", c.Param("proxyPath"))
    })

    // Маршрут для продуктов
    router.Any("/products/*proxyPath", func(c *gin.Context) {
        proxyRequest(c, "http://product-service:8080", c.Param("proxyPath"))
    })

    // Другие маршруты аналогично...

    router.Run(":8000")
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
