// Package aggreagtemyqiita is logic package
package aggreagtemyqiita

type pageItem struct {
	ID string `json:"id"`
}

type pageItems []pageItem

type pageDetailItem struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	User           qiitaUser `json:"user"`
	Body           string    `json:"body"`
	RenderedBody   string    `json:"rendered_body"`
	LikesCount     int       `json:"likes_count"`
	CommentsCount  int       `json:"comments_count"`
	PageViewsCount int       `json:"page_views_count"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
	URL            string    `json:"url"`
	Tags           []struct {
		Name string `json:"name"`
	} `json:"tags"`
}

func (item *pageDetailItem) likeRatio() float64 {
	return float64(item.LikesCount) / float64(item.PageViewsCount)
}

type qiitaUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ItemsCount  int    `json:"items_count"`
}
