package pgstorage

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/Iversy/unified-message-hub/internal/utils"
	"github.com/Masterminds/squirrel"
	"github.com/samber/lo"
)

func (storage *PGstorage) CreateMessage(ctx context.Context, messageInfos []*models.Message) error {
	query := storage.upsertQueryMessageAudit(messageInfos)
	queryText, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = storage.db.Exec(ctx, queryText, args...)
	if err != nil {
		return err
	}
	return nil
}

func (storage *PGstorage) upsertQueryMessageAudit(messageInfos []*models.Message) squirrel.Sqlizer {
	infos := lo.Map(messageInfos, func(info *models.Message, _ int) *MessageAudit {
		return &MessageAudit{
			SourcePlatform: uint64(info.Client),
			SourceChatID:   uint64(info.ChatId),
			SenderID:       info.Sender,
			MessageText:    info.Text,
			Received_at:    info.Timestamp,
		}
	})

	tableName := "message_audit"
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
	query := storage.upsertQueryRoutingRules(routes)
	queryText, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = storage.db.Exec(ctx, queryText, args...)
	if err != nil {
		return err
	}
	return nil
}

func (storage *PGstorage) upsertQueryRoutingRules(messageInfos []*models.Route) squirrel.Sqlizer {
	infos := lo.Map(messageInfos, func(info *models.Route, _ int) *RoutingRules {
		return &RoutingRules{
			ID:           info.ID,
			Name:         info.Name,
			SourceChatID: info.SourceChatID,
			ReceiverID:   info.ReceiverID,
			Keywords:     info.Keywords,
			IsActive:     info.IsActive,
		}
	})

	tableName := "routing_rules"
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
	q = q.Suffix(`
		ON CONFLICT (id) 
		DO UPDATE SET 
			name = EXCLUDED.name,
			source_chat_id = EXCLUDED.source_chat_id,
			receiver_id = EXCLUDED.receiver_id,
			keywords = EXCLUDED.keywords,
			is_active = EXCLUDED.is_active
	`)

	return q
}
