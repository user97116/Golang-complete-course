package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
)

var redisClient *redis.Client

func init() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func rateLimitMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the client IP address
		clientIP := r.RemoteAddr

		// Check if the client has exceeded the rate limit
		ctx := context.Background()
		key := fmt.Sprintf("rate_limit:%s", clientIP)
		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if count > 10 {
			http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
			return
		}

		// Call the next handler
		next(w, r, ps)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", rateLimitMiddleware(indexHandler))
	http.ListenAndServe(":8080", router)
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Hello, World!")
}
