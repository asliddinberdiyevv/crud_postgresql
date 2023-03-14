package database

import (
	"context"
	"posts/pkgs/models"

	"github.com/pkg/errors"
)

// UserDB persist Users.
type PostDB interface {
	CreatePost(ctx context.Context, user *models.Post) error
}

const createPostQuery = `
	INSERT INTO posts (post_name)
	VALUES (:post_name)
	RETURNING post_id
`

func (d *database) CreatePost(ctx context.Context, post *models.Post) error {
	rows, err := d.conn.NamedQueryContext(ctx, createPostQuery, post)
	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return errors.Wrap(err, "could not create post")
	}

	rows.Next()

	if err := rows.Scan(&post.ID); err != nil {
		return errors.Wrap(err, "could not get created postID")
	}

	return nil
}

// const getUserByIDQuery = `
// 	SELECT user_id, email, username, password_hash, created_at
// 	FROM users
// 	WHERE user_id = $1 AND deleted_at IS NULL;
// `
// func (d *database) GetUserByID(ctx context.Context, userID models.UserID) (*models.User, error) {
// 	var user models.User
// 	if err := d.conn.GetContext(ctx, &user, getUserByIDQuery, userID); err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
// const getUserByUsernameQuery = `
// 	SELECT user_id, email, username, password_hash, created_at
// 	FROM users
// 	WHERE username = $1 AND deleted_at IS NULL;
// `
// func (d *database) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
// 	var user models.User
// 	if err := d.conn.GetContext(ctx, &user, getUserByUsernameQuery, username); err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
// const listUsersQuery = `
// 	SELECT user_id, email, username, password_hash, created_at
// 	FROM users
// 	WHERE deleted_at IS NULL;
// `
// func (d *database) ListUsers(ctx context.Context) ([]*models.User, error) {
// 	var users []*models.User
// 	if err := d.conn.SelectContext(ctx, &users, listUsersQuery); err != nil {
// 		return nil, errors.Wrap(err, "could not get users")
// 	}
// 	return users, nil
// }
// const updateUserQuery = `
// 	UPDATE users
// 	SET username = :username,
// 	 		password_hash = :password_hash
// 	WHERE user_id = :user_id;
// `
// func (d *database) UpdateUser(ctx context.Context, user *models.User) error {
// 	result, err := d.conn.NamedExecContext(ctx, updateUserQuery, user)
// 	if err != nil {
// 		return err
// 	}

// 	rows, err := result.RowsAffected()
// 	if err != nil || rows == 0 {
// 		return errors.New("User not found")
// 	}

// 	return nil
// }
// const DeleteUserQuery = `
// 	UPDATE users
// 	SET deleted_at = NOW(),
// 			email = CONCAT(email, '-DELETED-', uuid_generate_v4()),
// 			username = CONCAT(username, '-DELETED-', uuid_generate_v4())
// 	WHERE user_id = $1 AND deleted_at IS NULL;
// `
// func (d *database) DeleteUser(ctx context.Context, userID models.UserID) (bool, error) {
// 	result, err := d.conn.ExecContext(ctx, DeleteUserQuery, userID)
// 	if err != nil {
// 		return false, err
// 	}

// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		return false, err
// 	}

// 	return rows > 0, nil
// }
