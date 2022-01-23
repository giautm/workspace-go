// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"giautm.dev/awesome/ent/schema"
	"giautm.dev/awesome/ent/schema/pulid"
	"giautm.dev/awesome/ent/todo"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoMixin := schema.Todo{}.Mixin()
	todoMixinFields0 := todoMixin[0].Fields()
	_ = todoMixinFields0
	todoMixinFields1 := todoMixin[1].Fields()
	_ = todoMixinFields1
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescCreateTime is the schema descriptor for create_time field.
	todoDescCreateTime := todoMixinFields0[0].Descriptor()
	// todo.DefaultCreateTime holds the default value on creation for the create_time field.
	todo.DefaultCreateTime = todoDescCreateTime.Default.(func() time.Time)
	// todoDescUpdateTime is the schema descriptor for update_time field.
	todoDescUpdateTime := todoMixinFields0[1].Descriptor()
	// todo.DefaultUpdateTime holds the default value on creation for the update_time field.
	todo.DefaultUpdateTime = todoDescUpdateTime.Default.(func() time.Time)
	// todo.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	todo.UpdateDefaultUpdateTime = todoDescUpdateTime.UpdateDefault.(func() time.Time)
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoFields[0].Descriptor()
	// todo.TextValidator is a validator for the "text" field. It is called by the builders before save.
	todo.TextValidator = todoDescText.Validators[0].(func(string) error)
	// todoDescPriority is the schema descriptor for priority field.
	todoDescPriority := todoFields[2].Descriptor()
	// todo.DefaultPriority holds the default value on creation for the priority field.
	todo.DefaultPriority = todoDescPriority.Default.(int)
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoMixinFields1[0].Descriptor()
	// todo.DefaultID holds the default value on creation for the id field.
	todo.DefaultID = todoDescID.Default.(func() pulid.ID)
}
