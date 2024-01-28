package types

// Entities:
// Projects:

// Attributes: Project ID, Project Name, Description, Start Date, End Date, Project Manager, Team Members.
// Tasks:

// Attributes: Task ID, Task Name, Description, Due Date, Priority, Status, Assigned To (User), Project ID (Foreign Key).
// Users:

// Attributes: User ID, Name, Email, Role (Admin, Project Manager, Team Member), Password.
// Comments:

// Attributes: Comment ID, Text, User ID (Who posted the comment), Task ID (Foreign Key).
// Attachments:

// Attributes: Attachment ID, File Name, File Type, Uploaded By (User), Task ID (Foreign Key).

type Project struct {
	Id          int
	Name        string
	Description string
	StartDate   string
	EndDate     string
	// Manager     User
	// TeamMembers []User
}

type Task struct {
	Id          int
	Name        string
	Description string
	DueDate     string
	Priority    string
	Status      string
	// AssignedTo  User
	ProjectId int
}

type Comment struct {
	Id     int
	Text   string
	UserId int
	TaskId int
}

type Attachment struct {
	Id       int
	FileName string
	FileType string
	// UploadedBy User
	TaskId int
}
