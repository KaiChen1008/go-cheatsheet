package context

import "context"

type Key struct{}

func WithValue(ctx context.Context) context.Context {
	// best practice for key: do not use string directly, use a self-defined type
	// e.g. there are two self-defined types, A & B -> A("key") != B("key")
	return context.WithValue(ctx, Key{}, "val")
}
