package dbassert

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func DBHas(t *testing.T, db *gorm.DB, model any) bool {
	r := db.Model(model).Where(model).First(model)

	err := r.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(fmt.Errorf("%w: while searching model=%v", err, model))
		return false
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf(fmt.Sprintf("model with given data not found. model=%v", model))
		return false
	}

	return true
}

func DBHasNot(t *testing.T, db *gorm.DB, model any) bool {
	r := db.Model(model).Where(model).First(model)

	err := r.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(fmt.Errorf("%w: while searching model=%v", err, model))
		return false
	}

	if err == nil {
		t.Errorf(fmt.Sprintf("model exists. model=%v", model))
		return false
	}

	return true
}
