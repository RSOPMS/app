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
	ProjectName  string
	StatusName   string
	PriorityName string
	StatusId     int
	PriorityId   int
	BranchId     int
	BranchName   string
	BranchUrl    string
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
