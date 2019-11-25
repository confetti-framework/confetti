package foundation

import (
	"lanvard/support"
	"reflect"
)

type bindings map[string]interface{}
type instances map[string]interface{}

type Container struct {
	// The container's bindings.
	bindings bindings

	// The registered type aliases.
	aliases map[string]interface{}

	// The registered aliases keyed by the abstract name.
	abstractAliases map[string]map[string]interface{}

	// The container's shared instances.
	instances instances
}

func NewContainer() Container {
	containerStruct := Container{}
	containerStruct.bindings = make(bindings)

	return containerStruct
}

// Determine if the given abstract type has been bound.
func (c *Container) Bound(abstract string) bool {
	_, bind := c.bindings[abstract]
	_, instance := c.instances[abstract]

	if bind || instance || c.IsAlias(abstract) {
		return true
	}

	return false
}

// Determine if a given string is an alias.
func (c *Container) IsAlias(name string) bool {
	if _, ok := c.aliases[name]; ok {
		return true
	}

	return false
}

// Register a binding with the container.
func (c *Container) Bind(abstract interface{}, concrete interface{}) {
	abstractString := support.Name(abstract)

	c.bindings[abstractString] = concrete
}

// Register a shared binding in the container.
func (c *Container) Singleton(abstract interface{}, concrete interface{}) {
	c.Bind(abstract, concrete)
}

// Register an existing instance as shared in the container.
func (c Container) Instance(abstract interface{}, instance interface{}) {
	abstractName := support.Name(abstract)

	c.removeAbstractAlias(abstractName)

	_, ok := c.aliases[abstractName]
	if ok {
		delete(c.aliases, abstractName)
	}

	if c.instances == nil {
		c.instances = make(instances)
	}

	c.instances[abstractName] = instance
}

// Get the container's bindings.
func (c Container) GetBindings() bindings {
	return c.bindings
}

// Resolve the given type from the container.
func (c *Container) Make(abstract interface{}) interface{} {
	return c.resolve(abstract)
}

// Resolve the given type from the container.
func (c *Container) resolve(abstract interface{}) interface{} {
	var abstractString string

	if support.Type(abstract) == reflect.String {
		abstractString = abstract.(string)
	} else {
		abstractString = reflect.TypeOf(abstract).Elem().String()
	}

	if support.Type(abstract) == reflect.String && c.IsAlias(abstract.(string)) {
		return c.aliases[abstract.(string)]
	}

	object, present := c.bindings[abstractString]

	if present {
		return object
	}

	panic("Can't resole container with: " + abstractString)
}

// Remove an alias from the contextual binding alias cache.
func (c Container) removeAbstractAlias(abstract string) {
	if _, ok := c.aliases[abstract]; !ok {
		return
	}

	panic("Todo, implement removeAbstractAlias")
}
