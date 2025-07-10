// Copyright (c) 2024 Tulir Asokan
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package database

import (
	"context"

	"go.mau.fi/util/dbutil"
	"maunium.net/go/mautrix/id"
)

const (
	getPseudoMappingsQuery = `
		SELECT room_id, user_id, sender_id
		FROM pseudo_mappings
	`
	deletePseudoMappingQuery = `
		DELETE FROM pseudo_mappings WHERE room_id = $1 AND user_id = $2
	`
	upsertPseudoMappingQuery = `
		INSERT INTO pseudo_mappings (room_id, user_id, sender_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (room_id, user_id) DO UPDATE
			SET sender_id = $3
	`
)

type PseudoMappingQuery struct {
	*dbutil.QueryHelper[*PseudoMapping]
}

func (pmq *PseudoMappingQuery) GetAll(ctx context.Context) ([]*PseudoMapping, error) {
	return pmq.QueryMany(ctx, getPseudoMappingsQuery)
}

func (pmq *PseudoMappingQuery) Upsert(ctx context.Context, pm *PseudoMapping) error {
	return pmq.Exec(ctx, upsertPseudoMappingQuery, pm.sqlVariables()...)
}

func (pmq *PseudoMappingQuery) Delete(ctx context.Context, roomID id.RoomID, userID id.UserID) error {
	return pmq.Exec(ctx, deletePseudoMappingQuery, roomID, userID)
}

type PseudoMapping struct {
	ID          id.RoomID          `json:"room_id"`
	UserID      id.UserID          `json:"user_id"`
	SenderID    id.SenderID        `json:"sender_id"`
}

func (p *PseudoMapping) sqlVariables() []any {
	return []any{
		p.ID,
		p.UserID,
		p.SenderID,
	}
}

func (p *PseudoMapping) Scan(row dbutil.Scannable) (*PseudoMapping, error) {
	// var createdAt int64
	err := row.Scan(&p.ID, &p.UserID, &p.SenderID)
	if err != nil {
		return nil, err
	}
	return p, nil
}
