package pgstorage

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func scanRouteRow(rows pgx.Row) (*models.Route, error) {
	var rule RoutingRules
	err := rows.Scan(
		&rule.ID,
		&rule.Name,
		&rule.SourceChatID,
		&rule.ReceiverID,
		&rule.Keywords,
		&rule.IsActive,
	)
	if err != nil {
		return nil, err
	}

	return &models.Route{
		ID:           rule.ID,
		Name:         rule.Name,
		SourceChatID: rule.SourceChatID,
		ReceiverID:   rule.ReceiverID,
		Keywords:     rule.Keywords,
		IsActive:     rule.IsActive,
	}, nil
}

func (storage *PGstorage) queryGetRoutes(ctx context.Context, query squirrel.SelectBuilder) ([]*models.Route, error) {
	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := storage.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []*models.Route
	for rows.Next() {
		var route *models.Route
		route, err = scanRouteRow(rows)
		if err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	return routes, nil
}

func (storage *PGstorage) GetAllRoutes(ctx context.Context) ([]*models.Route, error) {
	query := squirrel.Select("*").
		From("routing_rules").
		PlaceholderFormat(squirrel.Dollar)

	return storage.queryGetRoutes(ctx, query)
}

func (storage *PGstorage) GetActiveRoutesByReceiverID(ctx context.Context, receiverID string) ([]*models.Route, error) {
	query := squirrel.Select("*").
		From("routing_rules").
		Where(squirrel.Eq{
			"receiver_id": receiverID,
			"is_active":   true,
		}).
		PlaceholderFormat(squirrel.Dollar)

	return storage.queryGetRoutes(ctx, query)
}

func (storage *PGstorage) GetRouteByID(ctx context.Context, id string) (*models.Route, error) {
	query := squirrel.Select("*").
		From("routing_rules").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	route, err := scanRouteRow(storage.db.QueryRow(ctx, queryText, args...))
	if err != nil {
		return nil, err
	}
	return route, nil
}
