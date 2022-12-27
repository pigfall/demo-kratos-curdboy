package schema

import (
	"entgo.io/ent"	
	"entgo.io/ent/schema/field"	
	"entgo.io/ent/schema/edge"	
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner",User.Type).Ref("cars").Unique(),
	}
}
