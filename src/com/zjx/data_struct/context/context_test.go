package context

import (
	"context"
	"fmt"
	"testing"
)

// TestContext
func TestContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "123", "456")
	str := ctx.Value("123").(string)
	fmt.Printf("ok, str : %v", str)
}
