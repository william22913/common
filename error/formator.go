package error

import (
	"strings"

	"github.com/william22913/common/bundles"
)

func NewErrorFormator(
	bundles bundles.Bundles,
) Formator {
	return &formator{
		bundles: bundles,
	}
}

type formator struct {
	bundles bundles.Bundles
}

func (f formator) ReformatErrorMessage(err error) DefaultErrorResponse {

	result := DefaultErrorResponse{
		DefaultMessage: DefaultMessage{
			Type: "error",
		},
	}

	switch errs := err.(type) {
	case *UnbundledErrorMessages:
		result.Error = DefaultError{
			Status: errs.status,
			Code:   errs.code.Error(),
		}

		param := make(map[string]interface{})

		if errs.function != nil {
			tempParam := errs.function(errs.param...)

			for key := range tempParam {
				param[key] = tempParam[key].Param
				val, _ := tempParam[key].Param.(string)

				if tempParam[key].IsConverted {
					param[key] = f.bundles.ReadMessageBundle("common.constanta", strings.ToUpper(val), errs.language, nil)
				}
			}
		}

		result.Error.Reason = f.bundles.ReadMessageBundle("common.error", errs.Error(), errs.language, param)
	default:
		result.Error = DefaultError{
			Status: 500,
			Reason: err.Error(),
		}
	}

	return result
}
