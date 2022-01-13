package helper

import (
	"strings"
)

//to export a function we capitalize the name of the function
func UserInputValidation(firstname string, lastname string, email string, ticketcount uint, remainingtickets uint) (bool, bool, bool) {
	namevalid := (len(firstname) >= 2 && len(lastname) >= 2)
	emailvalid := strings.Contains(email, "@")
	usertickets := ticketcount > 0 && ticketcount <= remainingtickets //ticket must be greater than 0
	return namevalid, emailvalid, usertickets
}

//we should export it so it is available for use in other files
