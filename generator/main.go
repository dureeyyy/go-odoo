package main

// import (
// 	"github.com/dureeyyy/go-odoo/generator/cmd"
// )

// func main() {
// 	cmd.Execute()
// }

import (
	"fmt"
	"log"

	odoo "github.com/dureeyyy/go-odoo"
)

func main() {
	c, err := odoo.NewClient(&odoo.ClientConfig{
		Admin:    "admin@talinolabs.com",
		Password: "J7rGSzVUoYfrloK",
		Database: "seafoodcityodoo-seafoodcityodoo-production-13146340",
		URL:      "https://seafoodcityodoo-seafoodcityodoo.odoo.com/",
	})
	if err != nil {
		log.Fatal(err)
	}

	criteria := odoo.NewCriteria()

	model, err := c.FindAccountAnalyticLine(criteria)
	if err != nil {
		log.Fatalf("Error finding account: %v", err)
	}

	fmt.Println(model.SoLine)

}
