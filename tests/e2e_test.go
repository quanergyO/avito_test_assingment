package tests

import (
	"avito_test_assingment/internal/cache"
	"avito_test_assingment/internal/handler"
	"avito_test_assingment/internal/repository"
	"avito_test_assingment/internal/repository/postgres"
	"avito_test_assingment/internal/service"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func TestHandler_PostBanner(t *testing.T) {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs", err)
		os.Exit(1)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "qwerty",
	})
	if err != nil {
		slog.Error("Error: failed to init db connection ", err)
		os.Exit(1)
	}

	redisDB, err := cache.NewRedis(cache.Config{
		Host: viper.GetString("redis.host"),
		Port: viper.GetString("redis.port"),
		DB:   viper.GetInt("redis.DB"),
	})
	if err != nil {
		slog.Error("Error: failed to init redis connection ", err)
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, redisDB)
	handlers := handler.NewHandler(services)

	testTable := []struct {
		name                     string
		signInBody               string
		requestBody              string
		expectedStatusCodeVerify int
		expectedStatusCode       int
		expectedResponseBody     string
	}{
		{
			name:                     "Admin Post Ok",
			signInBody:               `{"username": "admin", "password": "admin"}`,
			requestBody:              `{"feature_id": 1,"tag_ids": [1, 4],"content": {"name": "test","surname": "Test2","age": 26},"is_active": true}`,
			expectedStatusCodeVerify: 200,
			expectedStatusCode:       200,
			expectedResponseBody:     `{"id":1}`,
		},
		{
			name:                     "Bad admin pass",
			signInBody:               `{"username": "admin", "password": "admin1"}`,
			requestBody:              `{"feature_id": 1, "tag_ids":[1, 4]}`,
			expectedStatusCodeVerify: 401,
			expectedStatusCode:       401,
			expectedResponseBody:     "mock",
		},
	}

	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", handlers.SignIn)
		auth.POST("/sign-up", handlers.SignUp)
	}
	api := r.Group("/api/v1", handlers.UserIdentity)
	{
		api.POST("/banner", handlers.BannerPost)
	}

	req, err := http.NewRequest("POST", "/auth/sign-up", bytes.NewBufferString(`{"username": "admin", "password": "admin", "role":2}`))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, err = http.NewRequest("POST", "/auth/sign-up", bytes.NewBufferString(`{"username": "user", "password": "user", "role":1}`))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			req, err = http.NewRequest("POST", "/auth/sign-in", bytes.NewBufferString(testCase.signInBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			w = httptest.NewRecorder()
			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCodeVerify, w.Code)
			token := strings.Split(w.Body.String(), ":")[1]
			token = token[1:]
			token = token[:len(token)-2]

			req, err = http.NewRequest("POST", "/api/v1/banner", bytes.NewBufferString(testCase.requestBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Authorization", "Bearer "+token)
			w = httptest.NewRecorder()

			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedResponseBody != "mock" {
				require.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}

	_ = db.Close()
}

func TestHandler_GetBanner(t *testing.T) {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs", err)
		os.Exit(1)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "qwerty",
	})
	if err != nil {
		slog.Error("Error: failed to init db connection ", err)
		os.Exit(1)
	}

	redisDB, err := cache.NewRedis(cache.Config{
		Host: viper.GetString("redis.host"),
		Port: viper.GetString("redis.port"),
		DB:   viper.GetInt("redis.DB"),
	})
	if err != nil {
		slog.Error("Error: failed to init redis connection ", err)
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, redisDB)
	handlers := handler.NewHandler(services)

	testTable := []struct {
		name                     string
		signInBody               string
		expectedStatusCodeVerify int
		expectedStatusCode       int
		expectedResponseBody     string
	}{
		{
			name:                     "Admin Get Ok",
			signInBody:               `{"username": "admin", "password": "admin"}`,
			expectedStatusCodeVerify: 200,
			expectedStatusCode:       200,
			expectedResponseBody:     "mock",
		},
		{
			name:                     "Bad admin pass",
			signInBody:               `{"username": "admin", "password": "admin123"}`,
			expectedStatusCodeVerify: 401,
			expectedStatusCode:       401,
			expectedResponseBody:     "mock",
		},
	}

	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", handlers.SignIn)
		auth.POST("/sign-un", handlers.SignUp)
	}
	api := r.Group("/api/v1", handlers.UserIdentity)
	{
		api.GET("/banner", handlers.BannerGet)
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/auth/sign-in", bytes.NewBufferString(testCase.signInBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCodeVerify, w.Code)
			token := strings.Split(w.Body.String(), ":")[1]
			token = token[1:]
			token = token[:len(token)-2]
			req, err = http.NewRequest("GET", "/api/v1/banner", nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Authorization", "Bearer "+token)

			w = httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedResponseBody != "mock" {
				assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}

	_ = db.Close()
}

func TestHandler_GetUserBanner(t *testing.T) {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs", err)
		os.Exit(1)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "qwerty",
	})
	if err != nil {
		slog.Error("Error: failed to init db connection ", err)
		os.Exit(1)
	}

	redisDB, err := cache.NewRedis(cache.Config{
		Host: viper.GetString("redis.host"),
		Port: viper.GetString("redis.port"),
		DB:   viper.GetInt("redis.DB"),
	})
	if err != nil {
		slog.Error("Error: failed to init redis connection ", err)
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, redisDB)
	handlers := handler.NewHandler(services)

	testTable := []struct {
		name                     string
		signInBody               string
		requestBody              string
		expectedStatusCodeVerify int
		expectedStatusCode       int
		expectedResponseBody     string
	}{
		{
			name:                     "Admin Get Ok",
			signInBody:               `{"username": "admin", "password": "admin"}`,
			requestBody:              `{"feature_id": 1, "tag_ids":[1, 4]}`,
			expectedStatusCodeVerify: 200,
			expectedStatusCode:       200,
			expectedResponseBody:     "mock",
		},
		{
			name:                     "Bad admin pass",
			signInBody:               `{"username": "admin", "password": "admin1"}`,
			requestBody:              `{"feature_id": 1, "tag_ids":[1, 4]}`,
			expectedStatusCodeVerify: 401,
			expectedStatusCode:       401,
			expectedResponseBody:     "mock",
		},
		{
			name:                     "User get OK",
			signInBody:               `{"username": "user", "password": "user"}`,
			requestBody:              `{"feature_id": 1, "tag_ids":[1, 4]}`,
			expectedStatusCodeVerify: 200,
			expectedStatusCode:       200,
			expectedResponseBody:     "mock",
		},
	}

	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", handlers.SignIn)
		auth.POST("/sign-un", handlers.SignUp)
	}
	api := r.Group("/api/v1", handlers.UserIdentity)
	{
		api.GET("/user_banner", handlers.UserBannerGet)
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/auth/sign-in", bytes.NewBufferString(testCase.signInBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCodeVerify, w.Code)
			token := strings.Split(w.Body.String(), ":")[1]
			token = token[1:]
			token = token[:len(token)-2]

			req, err = http.NewRequest("GET", "/api/v1/user_banner", bytes.NewBufferString(testCase.requestBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Authorization", "Bearer "+token)
			w = httptest.NewRecorder()

			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedResponseBody != "mock" {
				require.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}

	_ = db.Close()
}

func TestHandler_BannerIdPatch(t *testing.T) {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs", err)
		os.Exit(1)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "qwerty",
	})
	if err != nil {
		slog.Error("Error: failed to init db connection ", err)
		os.Exit(1)
	}

	redisDB, err := cache.NewRedis(cache.Config{
		Host: viper.GetString("redis.host"),
		Port: viper.GetString("redis.port"),
		DB:   viper.GetInt("redis.DB"),
	})
	if err != nil {
		slog.Error("Error: failed to init redis connection ", err)
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, redisDB)
	handlers := handler.NewHandler(services)

	testTable := []struct {
		name                     string
		signInBody               string
		requestBode              string
		idForDelete              string
		expectedStatusCodeVerify int
		expectedStatusCode       int
		expectedResponseBody     string
	}{
		{
			name:                     "Admin Patch Ok",
			signInBody:               `{"username": "admin", "password": "admin"}`,
			requestBode:              `{"tag_ids": [3, 4],"content": {"newField": "updated","now": true}}`,
			idForDelete:              "1",
			expectedStatusCodeVerify: 200,
			expectedStatusCode:       200,
			expectedResponseBody:     `{"status":"ok"}`,
		},
		{
			name:                     "Bad admin pass",
			signInBody:               `{"username": "admin", "password": "admin1"}`,
			idForDelete:              "1",
			expectedStatusCodeVerify: 401,
			expectedStatusCode:       401,
			expectedResponseBody:     "mock",
		},
	}

	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", handlers.SignIn)
		auth.POST("/sign-un", handlers.SignUp)
	}
	api := r.Group("/api/v1", handlers.UserIdentity)
	{
		api.PATCH("/banner/:id", handlers.BannerIdPatch)
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/auth/sign-in", bytes.NewBufferString(testCase.signInBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCodeVerify, w.Code)

			token := strings.Split(w.Body.String(), ":")[1]
			token = token[1:]
			token = token[:len(token)-2]

			req, err = http.NewRequest("PATCH", "/api/v1/banner/1", bytes.NewBufferString(testCase.requestBode))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Authorization", "Bearer "+token)

			w = httptest.NewRecorder()

			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedResponseBody != "mock" {
				require.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}

	_ = db.Close()
}

func TestHandler_BannerIdDelete(t *testing.T) {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs", err)
		os.Exit(1)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "qwerty",
	})
	if err != nil {
		slog.Error("Error: failed to init db connection ", err)
		os.Exit(1)
	}

	redisDB, err := cache.NewRedis(cache.Config{
		Host: viper.GetString("redis.host"),
		Port: viper.GetString("redis.port"),
		DB:   viper.GetInt("redis.DB"),
	})
	if err != nil {
		slog.Error("Error: failed to init redis connection ", err)
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, redisDB)
	handlers := handler.NewHandler(services)

	testTable := []struct {
		name                     string
		signInBody               string
		idForDelete              string
		expectedStatusCodeVerify int
		expectedStatusCode       int
		expectedResponseBody     string
	}{
		{
			name:                     "Admin Delete Ok",
			signInBody:               `{"username": "admin", "password": "admin"}`,
			idForDelete:              "1",
			expectedStatusCodeVerify: 200,
			expectedStatusCode:       200,
			expectedResponseBody:     `{"status":"ok"}`,
		},
		{
			name:                     "Bad admin pass",
			signInBody:               `{"username": "admin", "password": "admin1"}`,
			idForDelete:              "1",
			expectedStatusCodeVerify: 401,
			expectedStatusCode:       401,
			expectedResponseBody:     "mock",
		},
	}

	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", handlers.SignIn)
		auth.POST("/sign-un", handlers.SignUp)
	}
	api := r.Group("/api/v1", handlers.UserIdentity)
	{
		api.DELETE("/banner/:id", handlers.BannerIdDelete)
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/auth/sign-in", bytes.NewBufferString(testCase.signInBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCodeVerify, w.Code)

			token := strings.Split(w.Body.String(), ":")[1]
			token = token[1:]
			token = token[:len(token)-2]

			req, err = http.NewRequest("DELETE", "/api/v1/banner/1", nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Authorization", "Bearer "+token)

			w = httptest.NewRecorder()

			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedResponseBody != "mock" {
				require.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}

	_ = db.Close()
}
