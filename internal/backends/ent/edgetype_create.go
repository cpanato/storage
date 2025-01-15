// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/edgetype"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/nodelist"
)

// EdgeTypeCreate is the builder for creating a EdgeType entity.
type EdgeTypeCreate struct {
	config
	mutation *EdgeTypeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDocumentID sets the "document_id" field.
func (etc *EdgeTypeCreate) SetDocumentID(u uuid.UUID) *EdgeTypeCreate {
	etc.mutation.SetDocumentID(u)
	return etc
}

// SetNillableDocumentID sets the "document_id" field if the given value is not nil.
func (etc *EdgeTypeCreate) SetNillableDocumentID(u *uuid.UUID) *EdgeTypeCreate {
	if u != nil {
		etc.SetDocumentID(*u)
	}
	return etc
}

// SetProtoMessage sets the "proto_message" field.
func (etc *EdgeTypeCreate) SetProtoMessage(s *sbom.Edge) *EdgeTypeCreate {
	etc.mutation.SetProtoMessage(s)
	return etc
}

// SetType sets the "type" field.
func (etc *EdgeTypeCreate) SetType(e edgetype.Type) *EdgeTypeCreate {
	etc.mutation.SetType(e)
	return etc
}

// SetNodeID sets the "node_id" field.
func (etc *EdgeTypeCreate) SetNodeID(u uuid.UUID) *EdgeTypeCreate {
	etc.mutation.SetNodeID(u)
	return etc
}

// SetToNodeID sets the "to_node_id" field.
func (etc *EdgeTypeCreate) SetToNodeID(u uuid.UUID) *EdgeTypeCreate {
	etc.mutation.SetToNodeID(u)
	return etc
}

// SetID sets the "id" field.
func (etc *EdgeTypeCreate) SetID(u uuid.UUID) *EdgeTypeCreate {
	etc.mutation.SetID(u)
	return etc
}

// SetDocument sets the "document" edge to the Document entity.
func (etc *EdgeTypeCreate) SetDocument(d *Document) *EdgeTypeCreate {
	return etc.SetDocumentID(d.ID)
}

// SetFromID sets the "from" edge to the Node entity by ID.
func (etc *EdgeTypeCreate) SetFromID(id uuid.UUID) *EdgeTypeCreate {
	etc.mutation.SetFromID(id)
	return etc
}

// SetFrom sets the "from" edge to the Node entity.
func (etc *EdgeTypeCreate) SetFrom(n *Node) *EdgeTypeCreate {
	return etc.SetFromID(n.ID)
}

// SetToID sets the "to" edge to the Node entity by ID.
func (etc *EdgeTypeCreate) SetToID(id uuid.UUID) *EdgeTypeCreate {
	etc.mutation.SetToID(id)
	return etc
}

// SetTo sets the "to" edge to the Node entity.
func (etc *EdgeTypeCreate) SetTo(n *Node) *EdgeTypeCreate {
	return etc.SetToID(n.ID)
}

// AddNodeListIDs adds the "node_lists" edge to the NodeList entity by IDs.
func (etc *EdgeTypeCreate) AddNodeListIDs(ids ...uuid.UUID) *EdgeTypeCreate {
	etc.mutation.AddNodeListIDs(ids...)
	return etc
}

// AddNodeLists adds the "node_lists" edges to the NodeList entity.
func (etc *EdgeTypeCreate) AddNodeLists(n ...*NodeList) *EdgeTypeCreate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return etc.AddNodeListIDs(ids...)
}

// Mutation returns the EdgeTypeMutation object of the builder.
func (etc *EdgeTypeCreate) Mutation() *EdgeTypeMutation {
	return etc.mutation
}

// Save creates the EdgeType in the database.
func (etc *EdgeTypeCreate) Save(ctx context.Context) (*EdgeType, error) {
	etc.defaults()
	return withHooks(ctx, etc.sqlSave, etc.mutation, etc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (etc *EdgeTypeCreate) SaveX(ctx context.Context) *EdgeType {
	v, err := etc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etc *EdgeTypeCreate) Exec(ctx context.Context) error {
	_, err := etc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etc *EdgeTypeCreate) ExecX(ctx context.Context) {
	if err := etc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etc *EdgeTypeCreate) defaults() {
	if _, ok := etc.mutation.DocumentID(); !ok {
		v := edgetype.DefaultDocumentID()
		etc.mutation.SetDocumentID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etc *EdgeTypeCreate) check() error {
	if _, ok := etc.mutation.ProtoMessage(); !ok {
		return &ValidationError{Name: "proto_message", err: errors.New(`ent: missing required field "EdgeType.proto_message"`)}
	}
	if _, ok := etc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "EdgeType.type"`)}
	}
	if v, ok := etc.mutation.GetType(); ok {
		if err := edgetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "EdgeType.type": %w`, err)}
		}
	}
	if _, ok := etc.mutation.NodeID(); !ok {
		return &ValidationError{Name: "node_id", err: errors.New(`ent: missing required field "EdgeType.node_id"`)}
	}
	if _, ok := etc.mutation.ToNodeID(); !ok {
		return &ValidationError{Name: "to_node_id", err: errors.New(`ent: missing required field "EdgeType.to_node_id"`)}
	}
	if len(etc.mutation.FromIDs()) == 0 {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required edge "EdgeType.from"`)}
	}
	if len(etc.mutation.ToIDs()) == 0 {
		return &ValidationError{Name: "to", err: errors.New(`ent: missing required edge "EdgeType.to"`)}
	}
	return nil
}

func (etc *EdgeTypeCreate) sqlSave(ctx context.Context) (*EdgeType, error) {
	if err := etc.check(); err != nil {
		return nil, err
	}
	_node, _spec := etc.createSpec()
	if err := sqlgraph.CreateNode(ctx, etc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	etc.mutation.id = &_node.ID
	etc.mutation.done = true
	return _node, nil
}

func (etc *EdgeTypeCreate) createSpec() (*EdgeType, *sqlgraph.CreateSpec) {
	var (
		_node = &EdgeType{config: etc.config}
		_spec = sqlgraph.NewCreateSpec(edgetype.Table, sqlgraph.NewFieldSpec(edgetype.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = etc.conflict
	if id, ok := etc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := etc.mutation.ProtoMessage(); ok {
		_spec.SetField(edgetype.FieldProtoMessage, field.TypeBytes, value)
		_node.ProtoMessage = value
	}
	if value, ok := etc.mutation.GetType(); ok {
		_spec.SetField(edgetype.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if nodes := etc.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.DocumentTable,
			Columns: []string{edgetype.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DocumentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.FromTable,
			Columns: []string{edgetype.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NodeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.ToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.ToTable,
			Columns: []string{edgetype.ToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ToNodeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.NodeListsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   edgetype.NodeListsTable,
			Columns: edgetype.NodeListsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nodelist.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EdgeType.Create().
//		SetDocumentID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EdgeTypeUpsert) {
//			SetDocumentID(v+v).
//		}).
//		Exec(ctx)
func (etc *EdgeTypeCreate) OnConflict(opts ...sql.ConflictOption) *EdgeTypeUpsertOne {
	etc.conflict = opts
	return &EdgeTypeUpsertOne{
		create: etc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EdgeType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (etc *EdgeTypeCreate) OnConflictColumns(columns ...string) *EdgeTypeUpsertOne {
	etc.conflict = append(etc.conflict, sql.ConflictColumns(columns...))
	return &EdgeTypeUpsertOne{
		create: etc,
	}
}

type (
	// EdgeTypeUpsertOne is the builder for "upsert"-ing
	//  one EdgeType node.
	EdgeTypeUpsertOne struct {
		create *EdgeTypeCreate
	}

	// EdgeTypeUpsert is the "OnConflict" setter.
	EdgeTypeUpsert struct {
		*sql.UpdateSet
	}
)

// SetType sets the "type" field.
func (u *EdgeTypeUpsert) SetType(v edgetype.Type) *EdgeTypeUpsert {
	u.Set(edgetype.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *EdgeTypeUpsert) UpdateType() *EdgeTypeUpsert {
	u.SetExcluded(edgetype.FieldType)
	return u
}

// SetNodeID sets the "node_id" field.
func (u *EdgeTypeUpsert) SetNodeID(v uuid.UUID) *EdgeTypeUpsert {
	u.Set(edgetype.FieldNodeID, v)
	return u
}

// UpdateNodeID sets the "node_id" field to the value that was provided on create.
func (u *EdgeTypeUpsert) UpdateNodeID() *EdgeTypeUpsert {
	u.SetExcluded(edgetype.FieldNodeID)
	return u
}

// SetToNodeID sets the "to_node_id" field.
func (u *EdgeTypeUpsert) SetToNodeID(v uuid.UUID) *EdgeTypeUpsert {
	u.Set(edgetype.FieldToNodeID, v)
	return u
}

// UpdateToNodeID sets the "to_node_id" field to the value that was provided on create.
func (u *EdgeTypeUpsert) UpdateToNodeID() *EdgeTypeUpsert {
	u.SetExcluded(edgetype.FieldToNodeID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.EdgeType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(edgetype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EdgeTypeUpsertOne) UpdateNewValues() *EdgeTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(edgetype.FieldID)
		}
		if _, exists := u.create.mutation.DocumentID(); exists {
			s.SetIgnore(edgetype.FieldDocumentID)
		}
		if _, exists := u.create.mutation.ProtoMessage(); exists {
			s.SetIgnore(edgetype.FieldProtoMessage)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EdgeType.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EdgeTypeUpsertOne) Ignore() *EdgeTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EdgeTypeUpsertOne) DoNothing() *EdgeTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EdgeTypeCreate.OnConflict
// documentation for more info.
func (u *EdgeTypeUpsertOne) Update(set func(*EdgeTypeUpsert)) *EdgeTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EdgeTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *EdgeTypeUpsertOne) SetType(v edgetype.Type) *EdgeTypeUpsertOne {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *EdgeTypeUpsertOne) UpdateType() *EdgeTypeUpsertOne {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.UpdateType()
	})
}

// SetNodeID sets the "node_id" field.
func (u *EdgeTypeUpsertOne) SetNodeID(v uuid.UUID) *EdgeTypeUpsertOne {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.SetNodeID(v)
	})
}

// UpdateNodeID sets the "node_id" field to the value that was provided on create.
func (u *EdgeTypeUpsertOne) UpdateNodeID() *EdgeTypeUpsertOne {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.UpdateNodeID()
	})
}

// SetToNodeID sets the "to_node_id" field.
func (u *EdgeTypeUpsertOne) SetToNodeID(v uuid.UUID) *EdgeTypeUpsertOne {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.SetToNodeID(v)
	})
}

// UpdateToNodeID sets the "to_node_id" field to the value that was provided on create.
func (u *EdgeTypeUpsertOne) UpdateToNodeID() *EdgeTypeUpsertOne {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.UpdateToNodeID()
	})
}

// Exec executes the query.
func (u *EdgeTypeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EdgeTypeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EdgeTypeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EdgeTypeUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: EdgeTypeUpsertOne.ID is not supported by MySQL driver. Use EdgeTypeUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EdgeTypeUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EdgeTypeCreateBulk is the builder for creating many EdgeType entities in bulk.
type EdgeTypeCreateBulk struct {
	config
	err      error
	builders []*EdgeTypeCreate
	conflict []sql.ConflictOption
}

// Save creates the EdgeType entities in the database.
func (etcb *EdgeTypeCreateBulk) Save(ctx context.Context) ([]*EdgeType, error) {
	if etcb.err != nil {
		return nil, etcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(etcb.builders))
	nodes := make([]*EdgeType, len(etcb.builders))
	mutators := make([]Mutator, len(etcb.builders))
	for i := range etcb.builders {
		func(i int, root context.Context) {
			builder := etcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EdgeTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, etcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = etcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, etcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, etcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (etcb *EdgeTypeCreateBulk) SaveX(ctx context.Context) []*EdgeType {
	v, err := etcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etcb *EdgeTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := etcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etcb *EdgeTypeCreateBulk) ExecX(ctx context.Context) {
	if err := etcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EdgeType.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EdgeTypeUpsert) {
//			SetDocumentID(v+v).
//		}).
//		Exec(ctx)
func (etcb *EdgeTypeCreateBulk) OnConflict(opts ...sql.ConflictOption) *EdgeTypeUpsertBulk {
	etcb.conflict = opts
	return &EdgeTypeUpsertBulk{
		create: etcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EdgeType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (etcb *EdgeTypeCreateBulk) OnConflictColumns(columns ...string) *EdgeTypeUpsertBulk {
	etcb.conflict = append(etcb.conflict, sql.ConflictColumns(columns...))
	return &EdgeTypeUpsertBulk{
		create: etcb,
	}
}

// EdgeTypeUpsertBulk is the builder for "upsert"-ing
// a bulk of EdgeType nodes.
type EdgeTypeUpsertBulk struct {
	create *EdgeTypeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.EdgeType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(edgetype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EdgeTypeUpsertBulk) UpdateNewValues() *EdgeTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(edgetype.FieldID)
			}
			if _, exists := b.mutation.DocumentID(); exists {
				s.SetIgnore(edgetype.FieldDocumentID)
			}
			if _, exists := b.mutation.ProtoMessage(); exists {
				s.SetIgnore(edgetype.FieldProtoMessage)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EdgeType.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EdgeTypeUpsertBulk) Ignore() *EdgeTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EdgeTypeUpsertBulk) DoNothing() *EdgeTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EdgeTypeCreateBulk.OnConflict
// documentation for more info.
func (u *EdgeTypeUpsertBulk) Update(set func(*EdgeTypeUpsert)) *EdgeTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EdgeTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *EdgeTypeUpsertBulk) SetType(v edgetype.Type) *EdgeTypeUpsertBulk {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *EdgeTypeUpsertBulk) UpdateType() *EdgeTypeUpsertBulk {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.UpdateType()
	})
}

// SetNodeID sets the "node_id" field.
func (u *EdgeTypeUpsertBulk) SetNodeID(v uuid.UUID) *EdgeTypeUpsertBulk {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.SetNodeID(v)
	})
}

// UpdateNodeID sets the "node_id" field to the value that was provided on create.
func (u *EdgeTypeUpsertBulk) UpdateNodeID() *EdgeTypeUpsertBulk {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.UpdateNodeID()
	})
}

// SetToNodeID sets the "to_node_id" field.
func (u *EdgeTypeUpsertBulk) SetToNodeID(v uuid.UUID) *EdgeTypeUpsertBulk {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.SetToNodeID(v)
	})
}

// UpdateToNodeID sets the "to_node_id" field to the value that was provided on create.
func (u *EdgeTypeUpsertBulk) UpdateToNodeID() *EdgeTypeUpsertBulk {
	return u.Update(func(s *EdgeTypeUpsert) {
		s.UpdateToNodeID()
	})
}

// Exec executes the query.
func (u *EdgeTypeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EdgeTypeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EdgeTypeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EdgeTypeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
