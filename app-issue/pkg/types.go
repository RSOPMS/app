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
	StatusId     int
	PriorityId   int
	BranchId     int
	CreatedAt    string
}

type Comment struct {
	Id        int
	IssueId   int
	Content   string
	CreatedAt string
}

type Status struct {
	Id   int
	Name string
}

type Priority struct {
	Id   int
	Name string
}

type Branch struct {
	Id   int
	Name string
}
