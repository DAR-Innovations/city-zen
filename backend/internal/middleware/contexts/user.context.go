package contexts

import (
	"fmt"
	"github.com/DAR-Innovations/city-zen/internal/modules/auth/types"
	"github.com/gofiber/fiber/v3"
	"reflect"
)

const EMPLOYEE_CONTEXT = "EMPLOYEE_CONTEXT"

func GetEmployeeClaimsFromCtx(c fiber.Ctx) (*types.EmployeeClaims, error) {
	ctx := c.Locals(EMPLOYEE_CONTEXT)
	if ctx == nil {
		return nil, fmt.Errorf("no employee context found")
	}

	claims, ok := ctx.(*types.EmployeeClaims)
	if !ok {
		wrappedErr := fmt.Errorf("error getting employee context: type assertion failed, from <%v> to <%v>",
			reflect.TypeOf(ctx),
			reflect.TypeOf((*types.EmployeeClaims)(nil)))
		return nil, wrappedErr
	}

	return claims, nil
}

func SetEmployeeCtx(c fiber.Ctx, claims *types.EmployeeClaims) {
	c.Locals(EMPLOYEE_CONTEXT, claims)
}
