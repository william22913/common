package error_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/william22913/common/bundles"
	"github.com/william22913/common/dto/validator"
	errors "github.com/william22913/common/error"
)

func TestSomething(t *testing.T) {

	bundles, err := bundles.NewBundles("i18n", "id-ID")
	assert.Nil(t, err)
	assert.NotNil(t, bundles)

	validator := validator.NewValidator()

	errF := errors.NewErrorFormator(bundles)

	dto := TestDTO{
		Name: "gsagasg",
		Enum: "A",
	}

	err = validator.BasicValidatorByTag(&dto, "id-ID", "insert")
	if err != nil {
		fmt.Println(errF.ReformatErrorMessage(err))
	}

}
