package checker

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

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

	_, domain, _ := strings.Cut(email, "@")

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Println(err)
	}

	hasSPF, spfRecord = findRecord(txtRecords, "v=spf1")

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		fmt.Println(err)
	}

	hasDMARC, dmarcRecord = findRecord(dmarcRecords, "v=DMARC1")


	fmt.Printf("- Has MX: %v\n- Has SPF: %v\n" + 
		"- SPF record: %s\n- Has DMARC: %v\n- DMARC record: %s\n", hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

	return nil
}

func findRecord(records []string, prefix string) (bool, string) {
	for _, record := range records {
		if strings.HasPrefix(record, prefix) {
			return true, record
		}
	}

	return false, ""
}

func ValidateEmail(email string) error {
	return validation.Validate(email, is.Email)
}