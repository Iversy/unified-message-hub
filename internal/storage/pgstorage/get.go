package pgstorage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (storage *PGstorage) queryShardRoutes(ctx context.Context, shardSchema string, query squirrel.SelectBuilder) ([]*models.Route, error) {
	query = query.From(fmt.Sprintf("%s.routing_rules", shardSchema))

	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	log.Println("querying shard routes", queryText)
	rows, err := storage.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []*models.Route
	for rows.Next() {
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

		routes = append(routes, &models.Route{
			ID:           rule.ID,
			Name:         rule.Name,
			SourceChatID: rule.SourceChatID,
			ReceiverID:   rule.ReceiverID,
			Keywords:     rule.Keywords,
			IsActive:     rule.IsActive,
		})
	}

	return routes, nil
}

func (storage *PGstorage) queryAllShardsRoutes(ctx context.Context, query squirrel.SelectBuilder) ([]*models.Route, error) {
	var allRoutes []*models.Route
	var mu sync.Mutex
	var wg sync.WaitGroup
	errChan := make(chan error, 512)
	log.Println("Starting queryAllShardsRoutes")
	for i := 1; i <= 512; i++ {
		wg.Add(1)
		go func(shardNum int) {
			defer wg.Done()

			shardSchema := fmt.Sprintf("schema_%03d", shardNum)
			routes, err := storage.queryShardRoutes(ctx, shardSchema, query)

			if err != nil {
				errChan <- fmt.Errorf("shard %s: %w", shardSchema, err)
				return
			}

			mu.Lock()
			allRoutes = append(allRoutes, routes...)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	close(errChan)

	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return allRoutes, fmt.Errorf("errors in %d shards: %v", len(errs), errs)
	}

	return allRoutes, nil
}

// single shard query
func (storage *PGstorage) GetActiveRoutesBySourceChatID(ctx context.Context, chatID int) ([]*models.Route, error) {
	shardSchema := getShardSchema(uint64(chatID))

	query := squirrel.Select("*").
		Where(squirrel.Eq{
			"source_chat_id": chatID,
			"is_active":      true,
		}).
		PlaceholderFormat(squirrel.Dollar)

	return storage.queryShardRoutes(ctx, shardSchema, query)
}

func (storage *PGstorage) GetRouteByID(ctx context.Context, id int) (*models.Route, error) {
	//query := squirrel.Select("*").
	//	Where(squirrel.Eq{"id": id}).
	//	PlaceholderFormat(squirrel.Dollar)

	found := make(chan *models.Route, 1)
	errChan := make(chan error, 512)
	var wg sync.WaitGroup

	for i := 1; i <= 512; i++ {
		wg.Add(1)
		go func(shardNum int) {
			defer wg.Done()

			shardSchema := fmt.Sprintf("schema_%03d", shardNum)
			query := squirrel.Select("*").
				From(fmt.Sprintf("%s.routing_rules", shardSchema)).
				Where(squirrel.Eq{"id": id}).
				PlaceholderFormat(squirrel.Dollar)

			queryText, args, err := query.ToSql()
			if err != nil {
				errChan <- err
				return
			}

			var rule RoutingRules
			err = storage.db.QueryRow(ctx, queryText, args...).Scan(
				&rule.ID,
				&rule.Name,
				&rule.SourceChatID,
				&rule.ReceiverID,
				&rule.Keywords,
				&rule.IsActive,
			)

			if err == nil {
				select {
				case found <- &models.Route{
					ID:           rule.ID,
					Name:         rule.Name,
					SourceChatID: rule.SourceChatID,
					ReceiverID:   rule.ReceiverID,
					Keywords:     rule.Keywords,
					IsActive:     rule.IsActive,
				}:
				default:
					// Already found in another shard
				}
			} else if !errors.Is(err, pgx.ErrNoRows) {
				errChan <- fmt.Errorf("shard %s: %w", shardSchema, err)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(found)
		close(errChan)
	}()

	select {
	case route := <-found:
		return route, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// cross-shard query
func (storage *PGstorage) GetActiveRoutesByReceiverID(ctx context.Context, receiverID int) ([]*models.Route, error) {
	query := squirrel.Select("*").
		Where(squirrel.Eq{
			"receiver_id": receiverID,
			"is_active":   true,
		}).
		PlaceholderFormat(squirrel.Dollar)

	return storage.queryAllShardsRoutes(ctx, query)
}

// GetAllRoutes - Cross-shard query
func (storage *PGstorage) GetAllRoutes(ctx context.Context) ([]*models.Route, error) {
	query := squirrel.Select("*").
		PlaceholderFormat(squirrel.Dollar)

	return storage.queryAllShardsRoutes(ctx, query)
}
