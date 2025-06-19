package models

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=saarthi-database.cqfqse2kiuzm.us-east-1.rds.amazonaws.com user=postgres password=yourpassword dbname=saarthi-DataBase port=5432 sslmode=disable"

    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Auto migrate all models
    err = database.AutoMigrate(
        &User{},
        &Role{},
        &Permission{},
        &StudentProgress{},
        &Dashboard{},
        &Diploma{},
        &UpcomingClass{},
        &Query{},
        &CallSession{},
        &Command{},
        // &CareerPath{},
        &Bank{},
        &CreditCard{},
        &Loan{},
        &Transaction{},
        &UserSecurity{},
        &Notification{},
        &Analytics{},
        // &Attendance{},
        // &Course{},
    )

    if err != nil {
        log.Fatal("Failed to auto migrate models: ", err)
    }

    DB = database
}
