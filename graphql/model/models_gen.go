// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ConfUser struct {
	ID       string    `json:"id"`
	UserName string    `json:"userName"`
	Email    *string   `json:"email"`
	Phone    *string   `json:"phone"`
	Meetups  []*Meetup `json:"meetups"`
}

type Meetup struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	User        *ConfUser `json:"user"`
}