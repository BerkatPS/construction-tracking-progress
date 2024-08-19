package models

import "time"

type User struct {
	ID                      int64            `json:"id"`
	Username                string           `json:"username"`
	Email                   string           `json:"email"`
	Password                string           `json:"password"`
	Role                    string           `json:"role"`
	RefreshToken            string           `json:"refresh_token"`
	Projects                []Project        `json:"projects"`                  // One-to-Many
	AssignedTasks           []Task           `json:"assigned_tasks"`            // One-to-Many
	SentMessages            []Message        `json:"sent_messages"`             // One-to-Many
	ApprovedExpenses        []Expense        `json:"approved_expenses"`         // One-to-Many
	UploadedDocuments       []Document       `json:"uploaded_documents"`        // One-to-Many
	ConductedQualityChecks  []QualityCheck   `json:"conducted_quality_checks"`  // One-to-Many
	ReportedSafetyIncidents []SafetyIncident `json:"reported_safety_incidents"` // One-to-Many
	CreatedReports          []Report         `json:"created_reports"`           // One-to-Many
}

type Project struct {
	ID              int64            `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Budget          float64          `json:"budget"`
	Status          string           `json:"status"`
	ManagerID       int64            `json:"manager_id"`
	Manager         *User            `json:"manager"`          // Many-to-One
	Tasks           []Task           `json:"tasks"`            // One-to-Many
	Expenses        []Expense        `json:"expenses"`         // One-to-Many
	Documents       []Document       `json:"documents"`        // One-to-Many
	Messages        []Message        `json:"messages"`         // One-to-Many
	QualityChecks   []QualityCheck   `json:"quality_checks"`   // One-to-Many
	SafetyIncidents []SafetyIncident `json:"safety_incidents"` // One-to-Many
	Reports         []Report         `json:"reports"`          // One-to-Many
}

type Task struct {
	ID           int64     `json:"id"`
	ProjectID    int64     `json:"project_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	AssignedToID int64     `json:"assigned_to_id"`
	Project      *Project  `json:"project"`     // Many-to-One
	AssignedTo   *User     `json:"assigned_to"` // Many-to-One
}

type Expense struct {
	ID           int64     `json:"id"`
	ProjectID    int64     `json:"project_id"`
	Description  string    `json:"description"`
	Amount       float64   `json:"amount"`
	Date         time.Time `json:"date"`
	ApprovedBy   int64     `json:"approved_by"`
	Project      *Project  `json:"project"`       // Many-to-One
	ApprovedUser *User     `json:"approved_user"` // Many-to-One
}

type Document struct {
	ID           int64     `json:"id"`
	ProjectID    int64     `json:"project_id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	URL          string    `json:"url"`
	UploadedBy   int64     `json:"uploaded_by"`
	UploadDate   time.Time `json:"upload_date"`
	Project      *Project  `json:"project"`       // Many-to-One
	UploadedUser *User     `json:"uploaded_user"` // Many-to-One
}

type Message struct {
	ID        int64     `json:"id"`
	SenderID  int64     `json:"sender_id"`
	ProjectID int64     `json:"project_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	Sender    *User     `json:"sender"`  // Many-to-One
	Project   *Project  `json:"project"` // Many-to-One
}

type QualityCheck struct {
	ID          int64     `json:"id"`
	ProjectID   int64     `json:"project_id"`
	InspectorID int64     `json:"inspector_id"`
	Date        time.Time `json:"date"`
	Status      string    `json:"status"`
	Comments    string    `json:"comments"`
	Project     *Project  `json:"project"`   // Many-to-One
	Inspector   *User     `json:"inspector"` // Many-to-One
}

type SafetyIncident struct {
	ID          int64     `json:"id"`
	ProjectID   int64     `json:"project_id"`
	ReporterID  int64     `json:"reporter_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Severity    string    `json:"severity"`
	Status      string    `json:"status"`
	Project     *Project  `json:"project"`  // Many-to-One
	Reporter    *User     `json:"reporter"` // Many-to-One
}

type Report struct {
	ID           int64     `json:"id"`
	ProjectID    int64     `json:"project_id"`
	Type         string    `json:"type"`
	Content      string    `json:"content"`
	CreatedBy    int64     `json:"created_by"`
	CreationDate time.Time `json:"creation_date"`
	Project      *Project  `json:"project"` // Many-to-One
	Creator      *User     `json:"creator"` // Many-to-One
}

type QualityReport struct {
	ID                int64            `json:"id"`
	ProjectID         int64            `json:"project_id"`
	InspectorID       int64            `json:"inspector_id"`
	GeneratedByID     int64            `json:"generated_by_id"`
	GeneratedBy       *User            `json:"generated_by"` // Many-to-One
	Project           *Project         `json:"project"`      // Many-to-One
	Inspector         *User            `json:"inspector"`    // Many-to-One
	StartDate         time.Time        `json:"start_date"`
	EndDate           time.Time        `json:"end_date"`
	GeneratedDate     time.Time        `json:"generated_date"`
	QualityChecks     []QualityCheck   `json:"quality_checks"`  // List of quality checks included in the report
	OverallStatus     string           `json:"overall_status"`  // Overall status of quality checks (e.g., Passed, Failed, Mixed)
	Comments          string           `json:"comments"`        // General comments on the report
	AttachedDocuments []Document       `json:"attached_documents"` // Any related documents
}