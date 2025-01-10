package contexts

import (
	"fmt"
	"github.com/DAR-Innovations/city-zen/internal/modules/auth/types"
	"github.com/gofiber/fiber/v3"
	"reflect"
)

const USER_CONTEXT = "USER_CONTEXT"

func GetUserClaimsFromCtx(c fiber.Ctx) (*types.UserClaims, error) {
	ctx := c.Locals(USER_CONTEXT)
	if ctx == nil {
		return nil, fmt.Errorf("no User context found")
	}

	claims, ok := ctx.(*types.UserClaims)
	if !ok {
		wrappedErr := fmt.Errorf("error getting User context: type assertion failed, from <%v> to <%v>",
			reflect.TypeOf(ctx),
			reflect.TypeOf((*types.UserClaims)(nil)))
		return nil, wrappedErr
	}

	return claims, nil
}

func SetUserCtx(c fiber.Ctx, claims *types.UserClaims) {
	c.Locals(USER_CONTEXT, claims)
}
