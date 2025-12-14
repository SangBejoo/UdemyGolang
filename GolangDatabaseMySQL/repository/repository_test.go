package repository

import (
	GolangDatabaseMySQL "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(GolangDatabaseMySQL.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(GolangDatabaseMySQL.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(GolangDatabaseMySQL.GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for id, comment := range comments {
		fmt.Println(id, comment)
	}
}

func TestUpdate(t *testing.T) {
	commentRepository := NewCommentRepository(GolangDatabaseMySQL.GetConnection())
	comment, err := commentRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	comment.Comment = "Test Update"
	result, err := commentRepository.Update(context.Background(), comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	commentRepository := NewCommentRepository(GolangDatabaseMySQL.GetConnection())
	comment, err := commentRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	deleteComment, err := commentRepository.Delete(context.Background(), comment.Id)
	if err != nil {
		panic(err)
	}

	fmt.Println("deleted Comment", deleteComment)
	_, err = commentRepository.FindById(context.Background(), 1)
	if err == nil {
		panic("comment still exist")
	}
}
