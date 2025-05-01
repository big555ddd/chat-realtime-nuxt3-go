package groupChat

import (
	"app/app/model"
	"app/app/request"
	"app/app/response"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/uptrace/bun"
)

func (s *Service) Create(ctx context.Context, req request.CreateGroup) (*model.GroupChat, error) {
	groupChat := &model.GroupChat{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}

	if _, err := s.db.NewInsert().Model(groupChat).Exec(ctx); err != nil {
		return nil, err
	}
	var members []model.GroupMember

	if req.UserID != nil {
		for _, userID := range req.UserID {
			member := model.GroupMember{
				GroupID:  groupChat.ID,
				UserID:   userID,
				IsActive: true,
			}
			members = append(members, member)

		}
		if _, err := s.db.NewInsert().Model(&members).Exec(ctx); err != nil {
			return nil, err
		}
	}

	return groupChat, nil
}

// func (s *Service) List(ctx context.Context, req request.GroupQuery) ([]response.GroupChatResponse, int, error) {
// 	offset := (req.Page - 1) * req.Limit
// 	gc := []response.GroupChatResponse{}
// 	query := s.db.NewSelect().TableExpr("group_chats AS gc").
// 		Column("gc.id", "gc.name", "gc.description", "gc.is_active", "gc.created_at", "gc.last_message_time", "gc.last_message_sender", "gc.last_message_type", "gc.last_message_detail").
// 		ColumnExpr("u2.display_name AS last_message_display_name").
// 		ColumnExpr("u2.first_name AS last_message_first_name").
// 		ColumnExpr(`
//         json_agg(
//             json_build_object(
//                 'id', u.id,
//                 'username', u.username,
//                 'first_name', u.first_name,
//                 'display_name', u.display_name,
//                 'last_name', u.last_name,
//                 'is_active', gm.is_active,
//                 'time_read', gm.time_read,
//                 'message_count', (
//                     CASE
//                         WHEN gc.last_message_sender = u.id THEN 0
//                         ELSE (
//                             SELECT COUNT(gcm.group_id)
//                             FROM group_chat_maps AS gcm
//                             WHERE gcm.group_id = gc.id
//                               AND gcm.created_at > gm.time_read
//                               AND gcm.created_at <= gc.last_message_time
//                         )
//                     END
//                 )
//             )
//         ) AS members`).
// 		Join("LEFT JOIN group_members AS gm ON gm.group_id = gc.id").
// 		Join("LEFT JOIN users AS u ON u.id = gm.user_id").
// 		Join("LEFT JOIN users AS u2 ON u2.id = gc.last_message_sender").
// 		Where("gc.is_active = true").
// 		Group("gc.id", "u2.display_name", "u2.first_name")
// 	if req.Search != "" {
// 		search := fmt.Sprint("%" + strings.ToLower(req.Search) + "%")
// 		query.Where("LOWER(gc.name) Like ?", search)
// 	}

// 	if req.UserID != "" {
// 		query.Where("gm.user_id =?", req.UserID)
// 	}

// 	total, err := query.Count(ctx)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	order := fmt.Sprintf("%s %s", req.SortBy, req.OrderBy)

// 	if err := query.Offset(offset).Limit(req.Limit).Order(order).Scan(ctx, &gc); err != nil {
// 		return nil, 0, err
// 	}

// 	return gc, total, nil
// }

func (s *Service) List(ctx context.Context, req request.GroupQuery) ([]response.GroupChatResponse, int, error) {
	groupChats, total, err := s.queryGroupChats(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	groupIDs := extractGroupIDs(groupChats)
	members, err := s.queryMembers(ctx, groupIDs)
	if err != nil {
		return nil, 0, err
	}

	if err := s.calculateMessageCounts(ctx, groupChats, members); err != nil {
		return nil, 0, err
	}

	return groupChats, total, nil
}

func (s *Service) queryGroupChats(ctx context.Context, req request.GroupQuery) ([]response.GroupChatResponse, int, error) {
	query := s.db.NewSelect().TableExpr("group_chats AS gc").
		Column("gc.id", "gc.name", "gc.description", "gc.is_active", "gc.created_at", "gc.last_message_time", "gc.last_message_sender", "gc.last_message_type", "gc.last_message_detail").
		Where("gc.is_active = true")

	if req.Search != "" {
		search := fmt.Sprintf("%%%s%%", strings.ToLower(req.Search))
		query.Where("LOWER(gc.name) LIKE ?", search)
	}

	if req.UserID != "" {
		query.Where("EXISTS (SELECT 1 FROM group_members AS gm WHERE gm.group_id = gc.id AND gm.user_id = ?)", req.UserID)
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	order := fmt.Sprintf("%s %s", req.SortBy, req.OrderBy)
	groupChats := []response.GroupChatResponse{}
	if err := query.Limit(req.Limit).Order(order).Scan(ctx, &groupChats); err != nil {
		return nil, 0, err
	}

	return groupChats, total, nil
}

func (s *Service) queryMembers(ctx context.Context, groupIDs []string) ([]response.MemberInfo, error) {
	var members []response.MemberInfo
	if err := s.db.NewSelect().TableExpr("group_members AS gm").
		Column("gm.group_id", "gm.user_id", "gm.time_read", "gm.is_active").
		ColumnExpr("u.first_name, u.last_name, u.username, u.display_name").
		Join("LEFT JOIN users AS u ON u.id = gm.user_id").
		Where("gm.group_id IN (?)", bun.In(groupIDs)).
		Scan(ctx, &members); err != nil {
		return nil, err
	}

	return members, nil
}

func (s *Service) calculateMessageCounts(ctx context.Context, groupChats []response.GroupChatResponse, members []response.MemberInfo) error {
	for i, group := range groupChats {
		group.Member = []response.MemberChatResponse{}
		for _, member := range members {
			if member.GroupID == group.ID {
				messageCount, err := s.queryMessageCount(ctx, group.ID, member.TimeRead, group.LastMessageTime)
				if err != nil {
					return err
				}

				group.Member = append(group.Member, response.MemberChatResponse{
					ID:           member.UserID,
					Firstname:    member.FirstName,
					Lastname:     member.LastName,
					Username:     member.Username,
					DisplayName:  member.DisplayName,
					IsActive:     member.IsActive,
					TimeRead:     member.TimeRead,
					MessageCount: messageCount,
				})
			}
		}
		groupChats[i] = group
	}

	return nil
}

func (s *Service) queryMessageCount(ctx context.Context, groupID string, timeRead int64, lastMessageTime int64) (int64, error) {
	var messageCount int64
	if err := s.db.NewSelect().TableExpr("group_chat_maps AS gcm").
		ColumnExpr("COUNT(gcm.group_id)").
		Where("gcm.group_id = ?", groupID).
		Where("gcm.created_at > ?", timeRead).
		Where("gcm.created_at <= ?", lastMessageTime).
		Scan(ctx, &messageCount); err != nil {
		return 0, err
	}

	return messageCount, nil
}

func extractGroupIDs(groupChats []response.GroupChatResponse) []string {
	groupIDs := make([]string, len(groupChats))
	for i, group := range groupChats {
		groupIDs[i] = group.ID
	}
	return groupIDs
}

func (s *Service) Get(ctx context.Context, id string) (*response.GroupChatResponse, error) {
	// Step 1: Query group chat by ID
	groupChat, err := s.queryGroupChatByID(ctx, id)
	if err != nil {
		return nil, err
	}

	members, err := s.queryMembers(ctx, []string{groupChat.ID})
	if err != nil {
		return nil, err
	}

	groupChats := []response.GroupChatResponse{*groupChat}

	if err := s.calculateMessageCounts(ctx, groupChats, members); err != nil {
		return nil, err
	}

	return &groupChats[0], nil
}

// queryGroupChatByID queries a single group chat by ID
func (s *Service) queryGroupChatByID(ctx context.Context, id string) (*response.GroupChatResponse, error) {
	groupChat := &response.GroupChatResponse{}
	err := s.db.NewSelect().TableExpr("group_chats AS gc").
		Column("gc.id", "gc.name", "gc.description", "gc.is_active", "gc.created_at", "gc.last_message_time", "gc.last_message_sender", "gc.last_message_type", "gc.last_message_detail").
		Where("gc.is_active = true").
		Where("gc.id = ?", id).
		Scan(ctx, groupChat)
	if err != nil {
		return nil, err
	}

	return groupChat, nil
}

func (s *Service) Update(ctx context.Context, req request.UpdateGroup, id string) (*model.GroupChat, error) {
	ex, err := s.db.NewSelect().TableExpr("group_chats").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("group chat not found")
	}

	gc := model.GroupChat{}
	err = s.db.NewSelect().TableExpr("group_chats").Where("id =?", id).Scan(ctx, &gc)
	if err != nil {
		return nil, err
	}

	groupchat := model.GroupChat{
		ID:              id,
		Name:            req.Name,
		Description:     req.Description,
		IsActive:        req.IsActive,
		LastMessageTime: gc.CreatedAt,
	}

	groupchat.SetCreated(gc.CreatedAt)
	groupchat.SetUpdateNow()

	if _, err := s.db.NewUpdate().Model(&groupchat).WherePK().Exec(ctx); err != nil {
		return nil, err
	}

	if req.UserID != nil {
		_, err := s.db.NewDelete().TableExpr("group_members").Where("group_id =?", groupchat.ID).Exec(ctx)
		if err != nil {
			return nil, err
		}
		var members []model.GroupMember
		for _, userID := range req.UserID {
			member := model.GroupMember{
				GroupID:  groupchat.ID,
				UserID:   userID,
				IsActive: true,
			}
			members = append(members, member)
		}
		if _, err := s.db.NewInsert().Model(&members).Exec(ctx); err != nil {
			return nil, err
		}
	}

	return &groupchat, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	ex, err := s.db.NewSelect().TableExpr("group_chats").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return err
	}
	if !ex {
		return errors.New("group chat not found")
	}

	_, err = s.db.NewDelete().TableExpr("group_members").Where("group_id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().TableExpr("group_chats").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SentMessage(ctx context.Context, req request.SentMessage) (*model.GroupChatMap,error) {
	ex, err := s.db.NewSelect().TableExpr("group_members").
		Where("group_id = ? AND user_id = ?", req.GroupID, req.Sender).
		Exists(ctx)
	if err != nil {
		return nil,err
	}
	if !ex {
		return nil,errors.New("sender not found in group")
	}
	gcm := &model.GroupChatMap{
		GroupID: req.GroupID,
		UserID:  req.Sender,
		Type:    req.Type,
		Detail:  req.Detail,
	}

	gcm.SetCreatedNow()
	gcm.SetUpdateNow()
	_, err = s.db.NewInsert().Model(gcm).Exec(ctx)
	if err != nil {
		return nil,err
	}
	tm := time.Now().Unix()

	_, err = s.db.NewUpdate().TableExpr("group_members").
		Set("time_read = ?", tm).
		Where("group_id = ? AND user_id = ?", req.GroupID, req.Sender).
		Exec(ctx)
	if err != nil {
		return nil,err
	}
	_, err = s.db.NewUpdate().TableExpr("group_chats").
		Set("last_message_time = ?", tm).
		Set("last_message_sender = ?", req.Sender).
		Set("last_message_type = ?", req.Type).
		Set("last_message_detail = ?", req.Detail).
		Where("id = ?", req.GroupID).Exec(ctx)

	return gcm,err
}

func (s *Service) ListMessages(ctx context.Context, GroupId string, req request.GroupQuery) ([]response.GroupChatMap, int, error) {
	offset := (req.Page - 1) * req.Limit
	gcm := []response.GroupChatMap{}

	query := s.db.NewSelect().TableExpr("group_chat_maps AS gcm").
		Column("gcm.type", "gcm.detail", "gcm.created_at").
		ColumnExpr("u.id AS user_id").
		ColumnExpr("u.first_name AS first_name").
		ColumnExpr("u.display_name AS display_name").
		ColumnExpr("u.last_name AS last_name").
		ColumnExpr("gm.is_active AS is_active").
		Join("LEFT JOIN users AS u ON gcm.user_id = u.id").
		Join("LEFT JOIN group_members AS gm ON gm.group_id = ? AND u.id = gm.user_id", GroupId).
		Where("gcm.group_id = ?", GroupId)

	total, err := query.Count(ctx)

	if err != nil {
		return nil, 0, err
	}
	order := fmt.Sprintf("%s %s", req.SortBy, req.OrderBy)
	if err := query.Offset(offset).Limit(req.Limit).Order(order).Scan(ctx, &gcm); err != nil {
		return nil, 0, err
	}

	return gcm, total, nil
}
