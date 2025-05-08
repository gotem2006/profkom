package models

import "github.com/google/uuid"

type (
	PostChatRequest struct {
		Title string `json:"title"`
		Users []int  `json:"users"`
	}
	PostChatResponse struct {
		ChatID uuid.UUID `json:"chat_id"`
	}
	PostMessageRequest struct {
		Content string    `json:"content"`
		ChatID  uuid.UUID `json:"omitempty"`
		UserID  int       `json:"omitempty"`
	}
	PostMessageResponse struct {
		ID        uuid.UUID `json:"id"`
		Content   string    `json:"content"`
		ChatID    string    `json:"chat_id,omitempty"`
		UserID    int       `json:"user_id,omitempty"`
		CreatedAt int64     `json:"created_at"`
		UpdatedAt int64     `json:"updated_at"`
	}
	CheckAccessToChat struct {
		UserID int
		ChatID uuid.UUID
	}
	Chat struct {
		ID       uuid.UUID             `json:"chat_id"`
		Title    string                `json:"title"`
		Messages []PostMessageResponse `json:"messages"`
	}
	GetChatsRequest struct {
		UserID int
	}
	GetChatsResponse struct {
		Chats []Chat `json:"chats"`
	}
)
