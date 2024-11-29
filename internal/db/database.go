package db

import (
    "fmt"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDatabase() {
    host := os.Getenv("DB_HOST")
    if host == "" {
        host = "localhost"
    }
    
    user := os.Getenv("DB_USER")
    if user == "" {
        user = "postgres"
    }
    
    password := os.Getenv("DB_PASSWORD")
    if password == "" {
        password = "postgres"
    }
    
    dbname := os.Getenv("DB_NAME")
    if dbname == "" {
        dbname = "numeris"
    }
    
    port := os.Getenv("DB_PORT")
    if port == "" {
        port = "5432"
    }
    
    sslmode := os.Getenv("DB_SSLMODE")
    if sslmode == "" {
        sslmode = "disable"
    }

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        host, user, password, dbname, port, sslmode,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    DB = db
    fmt.Println("Database connected successfully!")
}
