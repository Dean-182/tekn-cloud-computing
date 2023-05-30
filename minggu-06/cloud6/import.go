import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
)
// Konfigurasi MySQL
dbUser := "root"
dbPass := ""
dbName := "coba"
dbHost := "localhost"

mysqlConfig := mysql.Config{
	User:   dbUser,
	Passwd: dbPass,
	DBName: dbName,
	Net:    "tcp",
	Addr:   dbHost + ":3306",
}

// Konfigurasi MongoDB
mongoURI := "mongodb://localhost:27017"
