// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package purpose

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Purpose {
	return predicate.Purpose(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Purpose {
	return predicate.Purpose(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldLTE(FieldID, id))
}

// PrimaryPurposeEQ applies the EQ predicate on the "primary_purpose" field.
func PrimaryPurposeEQ(v PrimaryPurpose) predicate.Purpose {
	return predicate.Purpose(sql.FieldEQ(FieldPrimaryPurpose, v))
}

// PrimaryPurposeNEQ applies the NEQ predicate on the "primary_purpose" field.
func PrimaryPurposeNEQ(v PrimaryPurpose) predicate.Purpose {
	return predicate.Purpose(sql.FieldNEQ(FieldPrimaryPurpose, v))
}

// PrimaryPurposeIn applies the In predicate on the "primary_purpose" field.
func PrimaryPurposeIn(vs ...PrimaryPurpose) predicate.Purpose {
	return predicate.Purpose(sql.FieldIn(FieldPrimaryPurpose, vs...))
}

// PrimaryPurposeNotIn applies the NotIn predicate on the "primary_purpose" field.
func PrimaryPurposeNotIn(vs ...PrimaryPurpose) predicate.Purpose {
	return predicate.Purpose(sql.FieldNotIn(FieldPrimaryPurpose, vs...))
}

// HasNode applies the HasEdge predicate on the "node" edge.
func HasNode() predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodeWith applies the HasEdge predicate on the "node" edge with a given conditions (other predicates).
func HasNodeWith(preds ...predicate.Node) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		step := newNodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Purpose) predicate.Purpose {
	return predicate.Purpose(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Purpose) predicate.Purpose {
	return predicate.Purpose(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Purpose) predicate.Purpose {
	return predicate.Purpose(sql.NotPredicates(p))
}
