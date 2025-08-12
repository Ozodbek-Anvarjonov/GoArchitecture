package app

import (
	"Architecture/config"
	_ "Architecture/docs"
	"Architecture/internal/infrastructure/repository"
	"Architecture/internal/interfaces"
	"Architecture/internal/usecase"
	"Architecture/pkg/logger"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default: // linux
		err = exec.Command("xdg-open", url).Start()
	}
	if err != nil {
		fmt.Println("Brauzer ochishda xatolik:", err)
	}
}

func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Log.Level)

	// Postgres ulanish
	db, err := sql.Open("postgres", cfg.PG.URL)
	if err != nil {
		l.Fatal(fmt.Errorf("db ulanishda xatolik: %w", err))
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		l.Fatal(fmt.Errorf("db ping xatolik: %w", err))
	}

	// Gin router
	router := gin.Default()

	// Biznes logika
	userRepo := repository.NewUserRepo(db)
	userService := usecase.NewUserService(userRepo)

	// Routerga handlerlar qo‘shish
	interfaces.NewRouter(router, userService)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Swagger avtomatik ochish
	go openBrowser(fmt.Sprintf("http://localhost:%d/swagger/index.html", cfg.HTTP.Port))

	// Server ishga tushirish
	port := fmt.Sprintf(":%d", cfg.HTTP.Port)
	go func() {
		if err := router.Run(port); err != nil {
			l.Fatal(fmt.Errorf("server ishga tushmadi: %w", err))
		}
	}()
	l.Info(fmt.Sprintf("Server listening on port %s", port))

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	l.Info("Shutting down server...")
}
