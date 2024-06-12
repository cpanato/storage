// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package hashesentry

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the hashesentry type in the database.
	Label = "hashes_entry"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldExternalReferenceID holds the string denoting the external_reference_id field in the database.
	FieldExternalReferenceID = "external_reference_id"
	// FieldNodeID holds the string denoting the node_id field in the database.
	FieldNodeID = "node_id"
	// FieldHashAlgorithmType holds the string denoting the hash_algorithm_type field in the database.
	FieldHashAlgorithmType = "hash_algorithm_type"
	// FieldHashData holds the string denoting the hash_data field in the database.
	FieldHashData = "hash_data"
	// EdgeExternalReference holds the string denoting the external_reference edge name in mutations.
	EdgeExternalReference = "external_reference"
	// EdgeNode holds the string denoting the node edge name in mutations.
	EdgeNode = "node"
	// Table holds the table name of the hashesentry in the database.
	Table = "hashes_entries"
	// ExternalReferenceTable is the table that holds the external_reference relation/edge.
	ExternalReferenceTable = "hashes_entries"
	// ExternalReferenceInverseTable is the table name for the ExternalReference entity.
	// It exists in this package in order to avoid circular dependency with the "externalreference" package.
	ExternalReferenceInverseTable = "external_references"
	// ExternalReferenceColumn is the table column denoting the external_reference relation/edge.
	ExternalReferenceColumn = "external_reference_id"
	// NodeTable is the table that holds the node relation/edge.
	NodeTable = "hashes_entries"
	// NodeInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodeInverseTable = "nodes"
	// NodeColumn is the table column denoting the node relation/edge.
	NodeColumn = "node_id"
)

// Columns holds all SQL columns for hashesentry fields.
var Columns = []string{
	FieldID,
	FieldExternalReferenceID,
	FieldNodeID,
	FieldHashAlgorithmType,
	FieldHashData,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// HashAlgorithmType defines the type for the "hash_algorithm_type" enum field.
type HashAlgorithmType string

// HashAlgorithmType values.
const (
	HashAlgorithmTypeUNKNOWN     HashAlgorithmType = "UNKNOWN"
	HashAlgorithmTypeMD5         HashAlgorithmType = "MD5"
	HashAlgorithmTypeSHA1        HashAlgorithmType = "SHA1"
	HashAlgorithmTypeSHA256      HashAlgorithmType = "SHA256"
	HashAlgorithmTypeSHA384      HashAlgorithmType = "SHA384"
	HashAlgorithmTypeSHA512      HashAlgorithmType = "SHA512"
	HashAlgorithmTypeSHA3_256    HashAlgorithmType = "SHA3_256"
	HashAlgorithmTypeSHA3_384    HashAlgorithmType = "SHA3_384"
	HashAlgorithmTypeSHA3_512    HashAlgorithmType = "SHA3_512"
	HashAlgorithmTypeBLAKE2B_256 HashAlgorithmType = "BLAKE2B_256"
	HashAlgorithmTypeBLAKE2B_384 HashAlgorithmType = "BLAKE2B_384"
	HashAlgorithmTypeBLAKE2B_512 HashAlgorithmType = "BLAKE2B_512"
	HashAlgorithmTypeBLAKE3      HashAlgorithmType = "BLAKE3"
	HashAlgorithmTypeMD2         HashAlgorithmType = "MD2"
	HashAlgorithmTypeADLER32     HashAlgorithmType = "ADLER32"
	HashAlgorithmTypeMD4         HashAlgorithmType = "MD4"
	HashAlgorithmTypeMD6         HashAlgorithmType = "MD6"
	HashAlgorithmTypeSHA224      HashAlgorithmType = "SHA224"
)

func (hat HashAlgorithmType) String() string {
	return string(hat)
}

// HashAlgorithmTypeValidator is a validator for the "hash_algorithm_type" field enum values. It is called by the builders before save.
func HashAlgorithmTypeValidator(hat HashAlgorithmType) error {
	switch hat {
	case HashAlgorithmTypeUNKNOWN, HashAlgorithmTypeMD5, HashAlgorithmTypeSHA1, HashAlgorithmTypeSHA256, HashAlgorithmTypeSHA384, HashAlgorithmTypeSHA512, HashAlgorithmTypeSHA3_256, HashAlgorithmTypeSHA3_384, HashAlgorithmTypeSHA3_512, HashAlgorithmTypeBLAKE2B_256, HashAlgorithmTypeBLAKE2B_384, HashAlgorithmTypeBLAKE2B_512, HashAlgorithmTypeBLAKE3, HashAlgorithmTypeMD2, HashAlgorithmTypeADLER32, HashAlgorithmTypeMD4, HashAlgorithmTypeMD6, HashAlgorithmTypeSHA224:
		return nil
	default:
		return fmt.Errorf("hashesentry: invalid enum value for hash_algorithm_type field: %q", hat)
	}
}

// OrderOption defines the ordering options for the HashesEntry queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByExternalReferenceID orders the results by the external_reference_id field.
func ByExternalReferenceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExternalReferenceID, opts...).ToFunc()
}

// ByNodeID orders the results by the node_id field.
func ByNodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodeID, opts...).ToFunc()
}

// ByHashAlgorithmType orders the results by the hash_algorithm_type field.
func ByHashAlgorithmType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHashAlgorithmType, opts...).ToFunc()
}

// ByHashData orders the results by the hash_data field.
func ByHashData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHashData, opts...).ToFunc()
}

// ByExternalReferenceField orders the results by external_reference field.
func ByExternalReferenceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExternalReferenceStep(), sql.OrderByField(field, opts...))
	}
}

// ByNodeField orders the results by node field.
func ByNodeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeStep(), sql.OrderByField(field, opts...))
	}
}
func newExternalReferenceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExternalReferenceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ExternalReferenceTable, ExternalReferenceColumn),
	)
}
func newNodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
	)
}
