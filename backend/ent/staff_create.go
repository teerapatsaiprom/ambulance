// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/ambu/app/ent/predicament"
	"github.com/ambu/app/ent/staff"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// StaffCreate is the builder for creating a Staff entity.
type StaffCreate struct {
	config
	mutation *StaffMutation
	hooks    []Hook
}

// SetStaffEmail sets the staff_email field.
func (sc *StaffCreate) SetStaffEmail(s string) *StaffCreate {
	sc.mutation.SetStaffEmail(s)
	return sc
}

// SetStaffName sets the staff_name field.
func (sc *StaffCreate) SetStaffName(s string) *StaffCreate {
	sc.mutation.SetStaffName(s)
	return sc
}

// SetStaffPassword sets the staff_password field.
func (sc *StaffCreate) SetStaffPassword(s string) *StaffCreate {
	sc.mutation.SetStaffPassword(s)
	return sc
}

// AddPredicamentIDs adds the predicament edge to Predicament by ids.
func (sc *StaffCreate) AddPredicamentIDs(ids ...int) *StaffCreate {
	sc.mutation.AddPredicamentIDs(ids...)
	return sc
}

// AddPredicament adds the predicament edges to Predicament.
func (sc *StaffCreate) AddPredicament(p ...*Predicament) *StaffCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPredicamentIDs(ids...)
}

// Mutation returns the StaffMutation object of the builder.
func (sc *StaffCreate) Mutation() *StaffMutation {
	return sc.mutation
}

// Save creates the Staff in the database.
func (sc *StaffCreate) Save(ctx context.Context) (*Staff, error) {
	if _, ok := sc.mutation.StaffEmail(); !ok {
		return nil, &ValidationError{Name: "staff_email", err: errors.New("ent: missing required field \"staff_email\"")}
	}
	if _, ok := sc.mutation.StaffName(); !ok {
		return nil, &ValidationError{Name: "staff_name", err: errors.New("ent: missing required field \"staff_name\"")}
	}
	if v, ok := sc.mutation.StaffName(); ok {
		if err := staff.StaffNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "staff_name", err: fmt.Errorf("ent: validator failed for field \"staff_name\": %w", err)}
		}
	}
	if _, ok := sc.mutation.StaffPassword(); !ok {
		return nil, &ValidationError{Name: "staff_password", err: errors.New("ent: missing required field \"staff_password\"")}
	}
	if v, ok := sc.mutation.StaffPassword(); ok {
		if err := staff.StaffPasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "staff_password", err: fmt.Errorf("ent: validator failed for field \"staff_password\": %w", err)}
		}
	}
	var (
		err  error
		node *Staff
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StaffMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StaffCreate) SaveX(ctx context.Context) *Staff {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *StaffCreate) sqlSave(ctx context.Context) (*Staff, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *StaffCreate) createSpec() (*Staff, *sqlgraph.CreateSpec) {
	var (
		s     = &Staff{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: staff.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: staff.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.StaffEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: staff.FieldStaffEmail,
		})
		s.StaffEmail = value
	}
	if value, ok := sc.mutation.StaffName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: staff.FieldStaffName,
		})
		s.StaffName = value
	}
	if value, ok := sc.mutation.StaffPassword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: staff.FieldStaffPassword,
		})
		s.StaffPassword = value
	}
	if nodes := sc.mutation.PredicamentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   staff.PredicamentTable,
			Columns: []string{staff.PredicamentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: predicament.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}
