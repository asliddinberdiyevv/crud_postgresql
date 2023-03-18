package database

import (
	"context"
	"posts/pkgs/models"

	"github.com/pkg/errors"
)

type PostDB interface {
	CreatePost(ctx context.Context, post *models.Post) error
	GetListPost(ctx context.Context) ([]*models.Post, error)
	GetPostByID(ctx context.Context, postID models.PostID) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, postID models.PostID) (bool, error)
	SearchByName(ctx context.Context, postName string) ([]*models.Post, error)
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

const listPostQuery = `
	SELECT post_id, post_name, post_like, post_star, created_at, deleted_at
	FROM posts
	WHERE deleted_at IS NULL;
`

func (d *database) GetListPost(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	if err := d.conn.SelectContext(ctx, &posts, listPostQuery); err != nil {
		return nil, errors.Wrap(err, "could not get posts")
	}
	return posts, nil
}

const getPostByIDQuery = `
	SELECT post_id, post_name, post_like, post_star, created_at
	FROM posts
	WHERE post_id = $1 AND deleted_at IS NULL;
`

func (d *database) GetPostByID(ctx context.Context, postID models.PostID) (*models.Post, error) {
	var post models.Post
	if err := d.conn.GetContext(ctx, &post, getPostByIDQuery, postID); err != nil {
		return nil, err
	}
	return &post, nil
}

const updatePostQuery = `
	UPDATE posts
	SET post_name = :post_name,
	 		post_like = :post_like,
			post_star = :post_star
	WHERE post_id = :post_id;
`

func (d *database) UpdatePost(ctx context.Context, post *models.Post) error {
	result, err := d.conn.NamedExecContext(ctx, updatePostQuery, post)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return errors.New("Post not found")
	}
	return nil
}

const DeletePostQuery = `
	UPDATE posts
	SET deleted_at = NOW()
	WHERE post_id = $1 AND deleted_at IS NULL;
`

func (d *database) DeletePost(ctx context.Context, postID models.PostID) (bool, error) {
	result, err := d.conn.ExecContext(ctx, DeletePostQuery, postID)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows > 0, nil
}

const SearchByNameQuery = `
	SELECT *
	FROM posts
	WHERE post_name like '%' || $1 || '%' AND deleted_at IS NULL;
	`

func (d *database) SearchByName(ctx context.Context, postName string) ([]*models.Post, error) {
	var posts []*models.Post
	if err := d.conn.SelectContext(ctx, &posts, SearchByNameQuery, postName); err != nil {
		return nil, errors.Wrap(err, "could not get posts")
	}
	return posts, nil
}
