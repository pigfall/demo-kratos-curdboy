package schema

import (
	"entgo.io/ent"	
	"entgo.io/ent/schema/field"	
	
)

// Dept holds the schema definition for the Dept entity.
type Dept struct {
	ent.Schema
}

// Fields of the Dept.
func (Dept) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Dept.
func (Dept) Edges() []ent.Edge {
	return nil
}
