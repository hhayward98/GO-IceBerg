package main 

import (
	"database/sql"
	"fmt"
	"strings"
	"net/http"
	"html/template"
	"log"
	"time"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template



type CartItems struct {
	ItemID string
	Name string
	Image string
	Amount int
	Price float64
}

