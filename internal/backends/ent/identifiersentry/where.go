// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package identifiersentry

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldLTE(FieldID, id))
}

// NodeID applies equality check predicate on the "node_id" field. It's identical to NodeIDEQ.
func NodeID(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldEQ(FieldNodeID, v))
}

// SoftwareIdentifierValue applies equality check predicate on the "software_identifier_value" field. It's identical to SoftwareIdentifierValueEQ.
func SoftwareIdentifierValue(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldEQ(FieldSoftwareIdentifierValue, v))
}

// NodeIDEQ applies the EQ predicate on the "node_id" field.
func NodeIDEQ(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldEQ(FieldNodeID, v))
}

// NodeIDNEQ applies the NEQ predicate on the "node_id" field.
func NodeIDNEQ(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldNEQ(FieldNodeID, v))
}

// NodeIDIn applies the In predicate on the "node_id" field.
func NodeIDIn(vs ...string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldIn(FieldNodeID, vs...))
}

// NodeIDNotIn applies the NotIn predicate on the "node_id" field.
func NodeIDNotIn(vs ...string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldNotIn(FieldNodeID, vs...))
}

// NodeIDGT applies the GT predicate on the "node_id" field.
func NodeIDGT(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldGT(FieldNodeID, v))
}

// NodeIDGTE applies the GTE predicate on the "node_id" field.
func NodeIDGTE(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldGTE(FieldNodeID, v))
}

// NodeIDLT applies the LT predicate on the "node_id" field.
func NodeIDLT(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldLT(FieldNodeID, v))
}

// NodeIDLTE applies the LTE predicate on the "node_id" field.
func NodeIDLTE(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldLTE(FieldNodeID, v))
}

// NodeIDContains applies the Contains predicate on the "node_id" field.
func NodeIDContains(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldContains(FieldNodeID, v))
}

// NodeIDHasPrefix applies the HasPrefix predicate on the "node_id" field.
func NodeIDHasPrefix(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldHasPrefix(FieldNodeID, v))
}

// NodeIDHasSuffix applies the HasSuffix predicate on the "node_id" field.
func NodeIDHasSuffix(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldHasSuffix(FieldNodeID, v))
}

// NodeIDIsNil applies the IsNil predicate on the "node_id" field.
func NodeIDIsNil() predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldIsNull(FieldNodeID))
}

// NodeIDNotNil applies the NotNil predicate on the "node_id" field.
func NodeIDNotNil() predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldNotNull(FieldNodeID))
}

// NodeIDEqualFold applies the EqualFold predicate on the "node_id" field.
func NodeIDEqualFold(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldEqualFold(FieldNodeID, v))
}

// NodeIDContainsFold applies the ContainsFold predicate on the "node_id" field.
func NodeIDContainsFold(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldContainsFold(FieldNodeID, v))
}

// SoftwareIdentifierTypeEQ applies the EQ predicate on the "software_identifier_type" field.
func SoftwareIdentifierTypeEQ(v SoftwareIdentifierType) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldEQ(FieldSoftwareIdentifierType, v))
}

// SoftwareIdentifierTypeNEQ applies the NEQ predicate on the "software_identifier_type" field.
func SoftwareIdentifierTypeNEQ(v SoftwareIdentifierType) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldNEQ(FieldSoftwareIdentifierType, v))
}

// SoftwareIdentifierTypeIn applies the In predicate on the "software_identifier_type" field.
func SoftwareIdentifierTypeIn(vs ...SoftwareIdentifierType) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldIn(FieldSoftwareIdentifierType, vs...))
}

// SoftwareIdentifierTypeNotIn applies the NotIn predicate on the "software_identifier_type" field.
func SoftwareIdentifierTypeNotIn(vs ...SoftwareIdentifierType) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldNotIn(FieldSoftwareIdentifierType, vs...))
}

// SoftwareIdentifierValueEQ applies the EQ predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueEQ(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldEQ(FieldSoftwareIdentifierValue, v))
}

// SoftwareIdentifierValueNEQ applies the NEQ predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueNEQ(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldNEQ(FieldSoftwareIdentifierValue, v))
}

// SoftwareIdentifierValueIn applies the In predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueIn(vs ...string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldIn(FieldSoftwareIdentifierValue, vs...))
}

// SoftwareIdentifierValueNotIn applies the NotIn predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueNotIn(vs ...string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldNotIn(FieldSoftwareIdentifierValue, vs...))
}

// SoftwareIdentifierValueGT applies the GT predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueGT(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldGT(FieldSoftwareIdentifierValue, v))
}

// SoftwareIdentifierValueGTE applies the GTE predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueGTE(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldGTE(FieldSoftwareIdentifierValue, v))
}

// SoftwareIdentifierValueLT applies the LT predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueLT(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldLT(FieldSoftwareIdentifierValue, v))
}

// SoftwareIdentifierValueLTE applies the LTE predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueLTE(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldLTE(FieldSoftwareIdentifierValue, v))
}

// SoftwareIdentifierValueContains applies the Contains predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueContains(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldContains(FieldSoftwareIdentifierValue, v))
}

// SoftwareIdentifierValueHasPrefix applies the HasPrefix predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueHasPrefix(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldHasPrefix(FieldSoftwareIdentifierValue, v))
}

// SoftwareIdentifierValueHasSuffix applies the HasSuffix predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueHasSuffix(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldHasSuffix(FieldSoftwareIdentifierValue, v))
}

// SoftwareIdentifierValueEqualFold applies the EqualFold predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueEqualFold(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldEqualFold(FieldSoftwareIdentifierValue, v))
}

// SoftwareIdentifierValueContainsFold applies the ContainsFold predicate on the "software_identifier_value" field.
func SoftwareIdentifierValueContainsFold(v string) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.FieldContainsFold(FieldSoftwareIdentifierValue, v))
}

// HasNode applies the HasEdge predicate on the "node" edge.
func HasNode() predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodeWith applies the HasEdge predicate on the "node" edge with a given conditions (other predicates).
func HasNodeWith(preds ...predicate.Node) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(func(s *sql.Selector) {
		step := newNodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IdentifiersEntry) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IdentifiersEntry) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.IdentifiersEntry) predicate.IdentifiersEntry {
	return predicate.IdentifiersEntry(sql.NotPredicates(p))
}
