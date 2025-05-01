package chat

import (
	"app/app/model"
	"app/app/request"
	"context"
	"errors"
	"log"
	"time"
)

func (s *Service) Create(ctx context.Context, req request.CreateChat) (*model.Chat, error) {

	customermap := &model.UserMap{}
	ex, err := s.db.NewSelect().Model(customermap).
		Where("(user_id = ? AND ref_id = ?) OR (user_id = ? AND ref_id = ?)", req.Sender, req.Recipient, req.Recipient, req.Sender).
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("you still not friend")
	}

	err = s.db.NewSelect().Model(customermap).
		Where("(user_id = ? AND ref_id = ?) OR (user_id = ? AND ref_id = ?)", req.Sender, req.Recipient, req.Recipient, req.Sender).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	chat := &model.Chat{
		Sender:    req.Sender,
		UserMapId: customermap.ID,
		Recipient: req.Recipient,
		Type:      req.Type,
		Detail:    req.Detail,
		IsRead:    false,
	}

	if _, err := s.db.NewInsert().Model(chat).Exec(ctx); err != nil {
		return nil, err
	}

	log.Printf("success")

	// err := s.firebaseplay.CreateChat(ctx, chat)
	// if err != nil {
	// 	return nil, err
	// }
	return chat, nil
}

func (s *Service) GetMessage(ctx context.Context, limit, page int, userID, RecipientID string) ([]model.Chat, int, error) {
	chat := []model.Chat{}
	if limit <= 0 {
		limit = 20 // Default to 20 messages per page if limit is invalid
	}
	if page <= 0 {
		page = 1 // Default to page 1 if page is invalid
	}
	query := s.db.NewSelect().Model(&chat).
		Where("sender = ? or sender = ?", userID, RecipientID).
		Where("recipient = ? or recipient = ?", userID, RecipientID)
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err

	}

	offset := total - (limit * page)

	query.Limit(limit).Offset(offset)

	if err := query.Order("created_at ASC").Scan(ctx); err != nil {
		return nil, 0, err
	}
	return chat, total, nil
}

func (s *Service) GetUnreadMessage(ctx context.Context, user, recipient string) (int, error) {
	query := s.db.NewSelect().Model(&model.Chat{}).
		Where("sender = ?", recipient).
		Where("recipient = ?", user).
		Where("is_read = ?", false)
	total, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *Service) Addfriend(ctx context.Context, req request.Addfriend, user string) (*model.UserMap, error) {
	customermap := &model.UserMap{}
	ex, err := s.db.NewSelect().Model(customermap).
		Where("(user_id = ? AND ref_id = ?) OR (user_id = ? AND ref_id = ?)", req.UserID, user, user, req.UserID).
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if ex {
		return nil, errors.New("user already friend")
	}

	newMap := &model.UserMap{
		UserID: user,
		RefID:  req.UserID,
	}

	_, err = s.db.NewInsert().Model(newMap).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return newMap, nil
}

func (s *Service) ReadMessage(ctx context.Context, user, recipient string) error {
	_, err := s.db.NewUpdate().Model(&model.Chat{}).
		Set("is_read = ?", true).
		Where("sender = ?", recipient).
		Where("recipient = ?", user).
		Where("created_at <= ?", time.Now().Unix()).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Removefriend(ctx context.Context, req request.Addfriend, user string) (*model.UserMap, error) {
	customermap := &model.UserMap{}
	ex, err := s.db.NewSelect().Model(customermap).
		Where("(user_id = ? AND ref_id = ?) OR (user_id = ? AND ref_id = ?)", req.UserID, user, user, req.UserID).
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("user not friend")
	}

	_, err = s.db.NewDelete().Model(customermap).
		Where("(user_id = ? AND ref_id = ?) OR (user_id = ? AND ref_id = ?)", req.UserID, user, user, req.UserID).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return customermap, nil
}
