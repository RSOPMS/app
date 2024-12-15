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
