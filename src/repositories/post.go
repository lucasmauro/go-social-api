package repositories

import (
	"api/src/models"
	"database/sql"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db}
}

func (repository PostRepository) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(postID), nil
}

func (repository PostRepository) GetPost(postID uint64) (models.Post, error) {
	lines, err := repository.db.Query(`
		SELECT p.*, u.nickname 
		FROM posts p
		INNER JOIN users u
		ON u.id = p.author_id
		WHERE p.id = ?`,
		postID,
	)

	if err != nil {
		return models.Post{}, err
	}

	var post models.Post

	if lines.Next() {
		if err := lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.UpVotes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repository PostRepository) GetPosts(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.nickname FROM posts p 
		INNER JOIN users u ON u.id = p.author_id 
		INNER JOIN followers f ON p.author_id = f.user_id 
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 DESC`,
		userID,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.UpVotes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository PostRepository) Update(postID uint64, post models.Post) error {
	statement, err := repository.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

func (repository PostRepository) Delete(postID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repository PostRepository) GetUserPosts(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		SELECT p.*, u.nickname FROM posts p
		JOIN users u ON u.id = p.author_id
		WHERE p.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.UpVotes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository PostRepository) UpVotePost(postID uint64) error {
	statement, err := repository.db.Prepare("UPDATE posts SET upvotes = upvotes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repository PostRepository) DownVotePost(postID uint64) error {
	statement, err := repository.db.Prepare(`
		UPDATE posts SET upvotes =
		CASE 
			WHEN upvotes > 0 THEN upvotes - 1
			ELSE 0 
		END
		WHERE id = ?
	`)
	if err != nil {
		return err
	}

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
