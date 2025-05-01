package user

import (
	"app/app/enum"
	"app/app/model"
	"app/app/request"
	"app/app/response"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Create(ctx context.Context, req request.CreateUser) (*model.User, error) {
	// Check if the user already exists
	ex, err := s.db.NewSelect().Model(&model.User{}).Where("username = ?", req.Username).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if ex {
		return nil, errors.New("username already exists")
	}

	ex, err = s.db.NewSelect().Model(&model.User{}).Where("email = ?", req.Email).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if ex {
		return nil, errors.New("email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the user
	user := &model.User{
		Username:    req.Username,
		Email:       req.Email,
		Password:    string(hashedPassword),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DisplayName: req.DisplayName,
		RoleID:      1,
		Status:      enum.STATUS_ACTIVE,
	}
	if _, err := s.db.NewInsert().Model(user).Exec(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) List(ctx context.Context, limit, page int, search, roleID, status, planType string) ([]model.User, int, error) {
	Offset := (page - 1) * limit
	users := []model.User{}
	query := s.db.NewSelect().Model(&users)
	if search != "" {
		query.Where("first_name LIKE ? OR display_name LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if roleID != "" {
		query.Where("role_id = ?", roleID)
	}
	if status != "" {
		query.Where("status = ?", status)
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	// if planType != "" {
	// 	query.Where("plan_type = ?", planType)
	// }
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}
	query.Limit(limit).Offset(Offset).
		Order("created_at ASC")
	if err := query.Scan(ctx); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *Service) Get(ctx context.Context, id string) (*model.User, error) {

	user := model.User{ID: id}
	if err := s.db.NewSelect().Model(&user).WherePK().Scan(ctx); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) Update(ctx context.Context, req request.UpdateUser, id string) (*model.User, error) {

	user := model.User{}
	if err := s.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	// Hash the password
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.DisplayName = req.DisplayName
	user.SetUpdateNow()

	if _, err := s.db.NewUpdate().Model(&user).Where("id = ?", id).Exec(ctx); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {

	user := model.User{ID: id}
	if _, err := s.db.NewDelete().Model(&user).WherePK().Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Service) ListFriend(ctx context.Context, limit, page int, id, search string) ([]response.ListFriendResponse, int, error) {
	users := []response.ListFriendResponse{}
	query := s.db.NewSelect().
		TableExpr("user_maps AS um").
		ColumnExpr(
			`u.id AS id,
		u.username AS username,
		u.first_name AS firstname,
		u.last_name AS lastname,
		u.email AS email,
		u.display_name AS display_name,
		u.role_id AS role_id,
		u.status AS status,
		c1.type AS type,
		c1.detail AS detail,
		c1.created_at AS message_time,
		COUNT(DISTINCT c.id) AS unread`,
		).
		Join(`LEFT JOIN users AS u 
		ON u.id = CASE 
					WHEN um.user_id = ? THEN um.ref_id
					WHEN um.ref_id = ? THEN um.user_id
				 END`,
			id, id,
		).
		Join("LEFT JOIN chats AS c ON (c.sender = u.id AND c.recipient = ? AND c.is_read = ?)", id, false).
		Join(`LEFT JOIN LATERAL (
			SELECT * 
			FROM chats
			WHERE user_map_id = um.id
			ORDER BY created_at DESC
			LIMIT 1
		) AS c1 ON true`).
		Where("? IN (um.user_id, um.ref_id)", id).
		Where("um.user_id != um.ref_id").
		GroupExpr("u.id, um.created_at,c1.type,c1.detail,c1.created_at").
		Order("message_time DESC NULLS LAST")

	if search != "" {
		query.Where("u.display_name LIKE ?", "%"+search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query.Limit(limit).Offset((page - 1) * limit)
	}

	if err := query.Order("message_time DESC NULLS LAST").
		Scan(ctx, &users); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
