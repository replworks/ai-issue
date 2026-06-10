package domain

type IssueDraft struct {
	Title string
	Body  string
}

type PublishableIssue struct {
	Title      string
	Body       string
	Repository string
}

type Publisher struct {
	Username string
}
