package pgstorage

import (
	"context"
	"fmt"
	"log"

	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/Iversy/unified-message-hub/internal/utils"
	"github.com/Masterminds/squirrel"
	"github.com/samber/lo"
)

func (storage *PGstorage) CreateMessage(ctx context.Context, messageInfos []*models.Message) error {
	messageGroups := groupByShard(messageInfos, func(m *models.Message) uint64 {
		return uint64(m.ChatId)
	})
	log.Println("Starting creating messages...")
	for shardSchema, messages := range messageGroups {
		query := storage.upsertQueryMessageAudit(messages, shardSchema)
		queryText, args, err := query.ToSql()
		if err != nil {
			return fmt.Errorf("shard %s: %w", shardSchema, err)
		}

		_, err = storage.db.Exec(ctx, queryText, args...)
		if err != nil {
			return fmt.Errorf("shard %s: %w", shardSchema, err)
		}
	}
	log.Println("Done creating messages...")

	return nil
}

func (storage *PGstorage) upsertQueryMessageAudit(
	messageInfos []*models.Message,
	shardSchema string,
) squirrel.Sqlizer {
	infos := lo.Map(messageInfos, func(info *models.Message, _ int) *MessageAudit {
		return &MessageAudit{
			SourcePlatform: uint64(info.Client),
			SourceChatID:   uint64(info.ChatId),
			SenderID:       info.Sender,
			MessageText:    info.Text,
			Received_at:    info.Timestamp,
		}
	})

	tableName := fmt.Sprintf("%s.message_audit", shardSchema)

	q := squirrel.Insert(tableName).Columns(utils.GetStructTag(MessageAudit{})...).
		PlaceholderFormat(squirrel.Dollar)

	for _, info := range infos {
		q = q.Values(
			info.SourcePlatform,
			info.SourceChatID,
			info.SenderID,
			info.MessageText,
			info.Received_at,
		)
	}
	return q
}
func (storage *PGstorage) UpsertRoute(ctx context.Context, routes []*models.Route) error {
	routeGroups := groupByShard(routes, func(r *models.Route) uint64 {
		return uint64(r.SourceChatID)
	})
	log.Println("Starting upsert routes...")
	for shardSchema, routes := range routeGroups {
		query := storage.upsertQueryRoutingRules(routes, shardSchema)
		queryText, args, err := query.ToSql()
		if err != nil {
			return fmt.Errorf("shard %s: %w", shardSchema, err)
		}

		_, err = storage.db.Exec(ctx, queryText, args...)
		if err != nil {
			return fmt.Errorf("shard %s: %w", shardSchema, err)
		}
	}
	log.Println("Done upserting routes...")

	return nil
}

func (storage *PGstorage) upsertQueryRoutingRules(
	routes []*models.Route,
	shardSchema string,
) squirrel.Sqlizer {
	infos := lo.Map(routes, func(info *models.Route, _ int) *RoutingRules {
		return &RoutingRules{
			ID:           info.ID,
			Name:         info.Name,
			SourceChatID: info.SourceChatID,
			ReceiverID:   info.ReceiverID,
			Keywords:     info.Keywords,
			IsActive:     info.IsActive,
		}
	})

	tableName := fmt.Sprintf("%s.routing_rules", shardSchema)

	q := squirrel.Insert(tableName).Columns(utils.GetStructTag(RoutingRules{})...).
		PlaceholderFormat(squirrel.Dollar)

	for _, info := range infos {
		q = q.Values(
			info.ID,
			info.Name,
			info.SourceChatID,
			info.ReceiverID,
			info.Keywords,
			info.IsActive,
		)
	}

	q = q.Suffix(fmt.Sprintf(`
        ON CONFLICT (id) 
        DO UPDATE SET 
            name = EXCLUDED.name,
            source_chat_id = EXCLUDED.source_chat_id,
            receiver_id = EXCLUDED.receiver_id,
            keywords = EXCLUDED.keywords,
            is_active = EXCLUDED.is_active
    `))

	return q
}
