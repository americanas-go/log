package log

import "context"

func ToContext(ctx context.Context) context.Context {
	return l.ToContext(ctx)
}

func FromContext(ctx context.Context) Logger {
	return l.FromContext(ctx)
}
