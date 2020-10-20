package schema

import (

	"time"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Predicament holds the schema definition for the Predicament entity.
type Predicament struct {
	ent.Schema
}

// Fields of the Predicament.
func (Predicament) Fields() []ent.Field {
	return []ent.Field{
		field.Time("Added_Time").Default(time.Now),
	}
}

// Edges of the Predicament.
func (Predicament) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("car", Car.Type).Ref("predicament").Unique(),
		edge.From("statuscar", Statuscar.Type).Ref("predicament").Unique(),
		edge.From("staff", Staff.Type).Ref("predicament").Unique(),
		edge.From("user", User.Type).Ref("predicament").Unique(),
	}
}
