package aggreagtemyqiita

import (
	"testing"
)

func Test_pageDetailItem_likeRatio(t *testing.T) {
	type fields struct {
		ID             string
		Title          string
		User           qiitaUser
		Body           string
		RenderedBody   string
		LikesCount     int
		CommentsCount  int
		PageViewsCount int
		CreatedAt      string
		UpdatedAt      string
		URL            string
		Tags           []struct {
			Name string `json:"name"`
		}
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"Normal", fields{LikesCount: 10, PageViewsCount: 200}, 0.05},
		{"Normal ZeroDivision", fields{LikesCount: 10, PageViewsCount: 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &pageDetailItem{
				ID:             tt.fields.ID,
				Title:          tt.fields.Title,
				User:           tt.fields.User,
				Body:           tt.fields.Body,
				RenderedBody:   tt.fields.RenderedBody,
				LikesCount:     tt.fields.LikesCount,
				CommentsCount:  tt.fields.CommentsCount,
				PageViewsCount: tt.fields.PageViewsCount,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
				URL:            tt.fields.URL,
				Tags:           tt.fields.Tags,
			}
			if got := item.likeRatio(); got != tt.want {
				t.Errorf("likeRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
