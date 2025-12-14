package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlExec := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repo.DB.ExecContext(ctx, sqlExec, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	sqlExec := "SELECT id, email, comment FROM comments WHERE id=?"
	rows, err := repo.DB.QueryContext(ctx, sqlExec, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id ")
	}

}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repo *commentRepositoryImpl) Update(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlExec := "UPDATE comments SET email=?, comment=? WHERE id=?"
	result, err := repo.DB.ExecContext(ctx, sqlExec, comment.Email, comment.Comment, comment.Id)
	if err != nil {
		return comment, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return comment, err
	}
	if rowsAffected == 0 {
		return comment, errors.New("No rows updated")
	}
	return comment, nil
}

func (repo *commentRepositoryImpl) Delete(ctx context.Context, id int32) (entity.Comment, error) {
	comment, err := repo.FindById(ctx, id)
	if err != nil {
		return comment, err
	}
	sqlExec := "DELETE FROM comments WHERE id=?"
	result, err := repo.DB.ExecContext(ctx, sqlExec, id)
	if err != nil {
		return comment, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return comment, err
	}
	if rowsAffected == 0 {
		return comment, errors.New("No rows deleted")
	}

	return comment, nil
}
