package pkg

type Project struct {
	Id    int
	Title string
}

type Issue struct {
	Id           int
	Title        string
	Description  string
	ProjectId    int
	StatusName   string
	PriorityName string
}

type Comment struct {
	Id        int
	IssueId   int
	Content   string
	CreatedAt string
}

type Status struct {
	Id           int
	Name         string
	DisplayOrder int
}

type Priority struct {
	Id           int
	Name         string
	DisplayOrder int
}

type Branch struct {
	Id   int
	Name string
	Url  string
}

type NewIssue struct {
	Id          int
	Title       string
	Description string
	ProjectID   int
	StatusID    int
	PriorityID  int
	BranchID    int
	CreatedAt   string
}
