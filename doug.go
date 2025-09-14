package doug

import "go.uber.org/dig"

// Modifies the state of the container.
type Option func(*dig.Container) error

// Configures a container with the provided options.
func Configure(c *dig.Container, opts ...Option) error {
	for _, apply := range opts {
		if err := apply(c); err != nil {
			return err
		}
	}

	return nil
}

// Invoke runs the given function after instantiating its dependencies.
func Invoke(function any, opts ...dig.InvokeOption) Option {
	return func(c *dig.Container) error {
		return c.Invoke(function, opts...)
	}
}

// Provide teaches the container how to build values of one or more types and
// expresses their dependencies.
func Provide(constructor any, opts ...dig.ProvideOption) Option {
	return func(c *dig.Container) error {
		return c.Provide(constructor, opts...)
	}
}

// Supply adds the value of one or more types to the container.
func Supply[T any](value T, opts ...dig.ProvideOption) Option {
	return Provide(func() T { return value }, opts...)
}
