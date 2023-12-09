package service

const (
	// the permission of read post
	ReadPost int = 1 << iota
	// the permission of update post
	UpdatePost
	// the permission of create post
	CreatePost
	// the permission of delete post
	DeletePost
	// the permission of like post
	LikePost
	/*
	 * the permission of diss like post
	 */
	DissPost
	// the permission of read comment
	ReadComment
	// the permission of update comment
	UpdateComment
	// the permission of create comment
	CreateComment
	// the permission of delete comment
	DeleteComment
	// the permission of like comment
	LikeComment
	// the permission of diss like comment
	DissComment
	// the permission of reply comment
	ReplyComment
	// the permision of enable to be devolop
	Develop
	// publish book
	PublishBook
)

type Permission struct {
	ReadPost bool `json:"readPost"`
	// the permission of update post
	UpdatePost bool `json:"updatePost"`
	// the permission of create post
	CreatePost bool `json:"createPost"`
	// the permission of delete post
	DeletePost bool `json:"deletePost"`
	// the permission of like post
	LikePost bool `json:"likePost"`
	/*
	 * the permission of diss like post
	 */
	DissPost bool `json:"dissPost"`
	// the permission of read comment
	ReadComment bool `json:"readComment"`
	// the permission of update comment
	UpdateComment bool `json:"updateComment"`
	// the permission of create comment
	CreateComment bool `json:"createComment"`
	// the permission of delete comment
	DeleteComment bool `json:"deleteComment"`
	// the permission of like comment
	LikeComment bool `json:"likeComment"`
	// the permission of diss like comment
	DissComment bool `json:"dissComment"`
	// the permission of reply comment
	ReplyComment bool `json:"replyComment"`
	// the permision of enable to be devolop
	Develop bool `json:"develop"`
}

func ParsePermission(permission int) *Permission {
	return &Permission{
		ReadPost:      permission&ReadPost != 0,
		UpdatePost:    permission&UpdatePost != 0,
		CreatePost:    permission&CreatePost != 0,
		DeletePost:    permission&DeletePost != 0,
		LikePost:      permission&LikePost != 0,
		DissPost:      permission&DissPost != 0,
		ReadComment:   permission&ReadComment != 0,
		UpdateComment: permission&UpdateComment != 0,
		CreateComment: permission&CreateComment != 0,
		DeleteComment: permission&DeleteComment != 0,
		LikeComment:   permission&LikeComment != 0,
		DissComment:   permission&DissComment != 0,
		ReplyComment:  permission&ReplyComment != 0,
		Develop:       permission&Develop != 0,
	}
}

func ComposePermission(permission ...int) int {
	result := 0
	for _, value := range permission {
		result = result | value
	}
	return result
}

func CreateRole(){
	
}