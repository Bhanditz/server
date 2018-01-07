package users

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/LibreRead/server/error"
	"github.com/LibreRead/server/session"
	"github.com/gin-gonic/gin"
)

func GetSignIn(c *gin.Context) {

	DB := c.MustGet("DB")

	email := session.GetEmailFromSession(c)
	if email != nil {
		c.Redirect(302, "/")
	}

	rows, err := DB.(*sql.DB).Query("select email from user where id = ?", 1)
	error.CheckError(err)

	defer rows.Close()
	var cEmail string
	if rows.Next() {
		err := rows.Scan(&cEmail)
		error.CheckError(err)
	}
	fmt.Println(cEmail)

	enableSignUp := false
	if cEmail == "" {
		enableSignUp = true
	}

	demoLabel := false
	if os.Getenv("LIBREREAD_DEMO_SERVER") == "1" {
		demoLabel = true
	}

	c.HTML(200, "signin.html", gin.H{
		"enableSignUp": enableSignUp,
		"demoLabel":    demoLabel,
	})
}
