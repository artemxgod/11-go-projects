package checker

import (
	"bufio"
	"fmt"
	"os"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

)

func StartCheck() error {
	fmt.Println("Write the domain you want to check:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && scanner.Text() != "" {
		if err := checkDomain(scanner.Text()); err != nil {
			return err
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}
	
	return nil
}

func checkDomain(email string) error {
	if err := ValidateEmail(email); err != nil {
		fmt.Println("Email is not valid")
		return nil
	}

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	return nil
}

func ValidateEmail(email string) error {
	return validation.Validate(email, is.Email)
}