package cmd

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	middleware "project-name/cmd/middleware"
	routes "project-name/cmd/routes"
	config "project-name/config"
	utils "project-name/migrate"

	"project-name/internal/aws"
	"project-name/internal/logger"
	repo "project-name/internal/repository"
	service "project-name/internal/service"

	sql "project-name/internal/database"

	"github.com/aws/aws-sdk-go/service/s3"
	gin "github.com/gin-gonic/gin"
)

var (
	port      string
	mode      string
	dsn       string
	secret    string
	s3session *s3.S3
	err       error
)

var migrate = flag.Bool("m", false, "for migration")

func Setup() {
	router := gin.New()
	router.MaxMultipartMemory = 1 << 20
	v1 := router.Group("/api/v1")
	v1.Use(gin.Logger())
	v1.Use(gin.Recovery())
	router.Use(middleware.CORS())

	db, err := sql.New(dsn)
	if err != nil {
		log.Println("Error Connecting to DB: ", err)
	}

	defer db.Close()
	conn := db.GetConn()

	// AWS Repository
	awsRepo := repo.NewAWSRepo(s3session)

	// Deposit Repository
	depositRepo := repo.NewDepositRepo(conn)

	// Home Service
	homeSrv := service.NewHomeService()

	// Deposit Service
	depositSrv := service.NewDepositService(awsRepo, depositRepo)

	// Routes
	routes.HomeRoute(v1, homeSrv)
	routes.DepositRoute(v1, depositSrv)
	routes.ErrorRoute(router)

	// Start Server
	router.Run(":" + port)
}

func init() {
	flag.Parse()

	port = config.AppConfig.ADDR
	if port == "" {
		port = "10"
	}

	secret = config.AppConfig.SECRET_KEY_TOKEN
	if secret == "" {
		log.Println("Please provide a secret key token")
	}

	mode = config.AppConfig.MODE
	if mode == "development" {
		loadDev()
	}

	if mode == "production" {
		loadDev()
	}

	aws_access_key := config.AppConfig.AWS_ACCESS_KEY
	aws_secret_key := config.AppConfig.AWS_SECRET_KEY

	// S3 Session
	s3session, err = aws.NewAWSSession(aws_access_key, aws_secret_key, "")
	if err != nil {
		log.Printf("Error occured: %v", err)
	}

	fmt.Println("Migration: ", *migrate)

	if *migrate {
		if err := utils.Migration(dsn); err != nil {
			log.Printf("Error occurrred when running migration: %v", err)
		}
	}
}

func loadDev() {
	gin.SetMode(gin.DebugMode)

	getDSN()
}

func loadProd() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.NewLogger())
	gin.DisableConsoleColor()

	getDSN()
}

func getDSN() {
	host := config.AppConfig.POSTGRES_HOST
	username := config.AppConfig.POSTGRES_USERNAME
	passwd := config.AppConfig.POSTGRES_PASSWORD
	dbname := config.AppConfig.POSTGRES_DBNAME
	port := config.AppConfig.POSTGRES_PORT

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, passwd, dbname, port)
	if dsn == "" {
		log.Println("DSN cannot be empty")
	}
}

// var _ = loadProd
