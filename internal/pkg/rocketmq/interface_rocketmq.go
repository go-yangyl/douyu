// Code generated by struct2interface; DO NOT EDIT.

package rocketmq

import (
	"context"
)

// ExampleInterface ...
type ExampleInterface interface {
	PushExampleMessage(ctx context.Context) error
}
