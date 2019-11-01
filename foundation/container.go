package foundation

import (
	"laravelgo/support"
	"reflect"
)

type bindings map[string]interface{}

type ContainerStruct struct {
	bindings        bindings
	aliases         map[string]interface{}
	abstractAliases map[string]map[string]interface{}
	instances       map[string]interface{}
}

func Container() ContainerStruct {
	return ContainerStruct{}
}

func (c *ContainerStruct) Bound(abstract string) bool {
	_, bind := c.bindings[abstract]
	_, instance := c.instances[abstract]

	if bind || instance || c.IsAlias(abstract) {
		return true
	}

	return false
}
func (c *ContainerStruct) IsAlias(name string) bool {
	if _, ok := c.aliases[name]; ok {
		return true
	}

	return false
}

func (c *ContainerStruct) Bind(abstract interface{}, concrete interface{}) {
	if c.bindings == nil {
		c.bindings = make(bindings)
	}
	abstractString := support.Name(abstract)

	c.bindings[abstractString] = concrete
}

func (c *ContainerStruct) Singleton(abstract interface{}, concrete interface{}) {
	c.Bind(abstract, concrete)
}

// Register an existing instance as shared in the container.
func (c ContainerStruct) Instance(abstract interface{}, instance interface{}) {
	abstractName := support.Name(abstract)

	c.removeAbstractAlias(abstractName)

	_, ok := c.aliases[abstractName]
	if ok {
		delete(c.aliases, abstractName)
	}

	if c.instances == nil {
		c.instances = make(map[string]interface{})
	}

	c.instances[abstractName] = instance
}

func (c ContainerStruct) GetBindings() bindings {
	return c.bindings
}

// Resolve the given type from the container.
func (c *ContainerStruct) Make(abstract interface{}) interface{} {
	return c.resolve(abstract)
}

// Resolve the given type from the container.
func (c *ContainerStruct) resolve(abstract interface{}) interface{} {

	abstractString := reflect.TypeOf(abstract).Elem().String()

	object, present := c.bindings[abstractString]

	if present {
		return object
	}

	panic("Can't resole container")
}

func (c ContainerStruct) removeAbstractAlias(abstract string) {
	if _, ok := c.aliases[abstract]; ! ok {
		return
	}

	panic("Todo, implement removeAbstractAlias")
}
