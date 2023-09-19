package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func init() {
	v := validator.New(validator.WithRequiredStructEnabled())
	Setup(v)
}

func TestSetupValid(t *testing.T) {
	fmt.Println(validatorInstance)
}

type PointTest struct {
	Location string `json:"location" validate:"required,location"`
}

func TestLocationValidator(t *testing.T) {
	param := PointTest{
		Location: "11836",
	}
	err := validatorInstance.Struct(param)
	if err != nil {
		t.Error(FormatError(err))
	}
}
