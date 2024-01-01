# dbassert

Package to check if model exists in DB or not via DBHas, DBHasNot functions

**Examples of assertions:**

```go
package test

import (
	"fmt"
	"github.com/vsvp21/dbassert/v1"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"models"
)

func TestSmth(t *testing.T) {
	dsn := "yourdsn"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// check if exists in database 
	dbassert.DBHas(t, db, &models.User{Name: "John", Age: 15})
	// check if not exists in database
	dbassert.DBHasNot(t, db,  &models.User{Name: "John", Age: 15})
}
```
