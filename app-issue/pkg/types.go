package pkg

type Project struct {
	Id    int
	Title string
}

type Issue struct {
	Id          int
	Title       string
	Description string
	ProjectId   int
}

type Comment struct {
	Id        int
	IssueId   int
	Content   string
	CreatedAt string
}