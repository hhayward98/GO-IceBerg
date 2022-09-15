package RegisterC

import (
	"database/sql"
	"fmt"
	"strings"
	"log"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)


type RegisterDetails struct {
	Email string
	Username string
	Password string
	ConfPass string
}