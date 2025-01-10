package data

type Issue struct {
	BaseEntity
	Name        string  `gorm:"size:255;not null"`
	Description string  `gorm:"type:text;not null"`
	IsCompleted bool    `gorm:"default:false"`
	Longitude   float64 `gorm:"type:decimal(9,6)"`
	Latitude    float64 `gorm:"type:decimal(9,6)"`
	AuthorID    uint    `gorm:"index;not null"`
	Author      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ImageURL    string  `gorm:"size:255"`
}

type VolunteerTask struct {
	BaseEntity
	IssueID     uint   `gorm:"index;not null"`
	Issue       Issue  `gorm:"foreignKey:IssueID"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"size:50;default:'PENDING';check:status IN ('PENDING', 'IN PROGRESS', 'DONE')"`
	Urgency     string `gorm:"size:50;check:urgency IN ('HIGH', 'MEDIUM', 'LOW')"`
	VolunteerID string `gorm:"index;not null"`
	User        User   `gorm:"foreignKey:VolunteerID"`
	Complexity  string `gorm:"size:50;check:complexity IN ('HIGH', 'MEDIUM', 'LOW')"`
}

type DepartmentTask struct {
	BaseEntity
	IssueID      uint       `gorm:"index;not null"`
	Issue        Issue      `gorm:"foreignKey:IssueID"`
	Title        string     `gorm:"size:255;not null"`
	Description  string     `gorm:"type:text"`
	TaskTypeID   uint       `gorm:"index;not null"`
	TaskType     TaskType   `gorm:"foreignKey:TaskTypeID"`
	Status       string     `gorm:"size:50;default:'PENDING';check:status IN ('PENDING', 'IN PROGRESS', 'DONE')"`
	Urgency      string     `gorm:"size:50;check:urgency IN ('HIGH', 'MEDIUM', 'LOW')"`
	Complexity   string     `gorm:"size:50;check:complexity IN ('HIGH', 'MEDIUM', 'LOW')"`
	DepartmentID uint       `gorm:"index;not null"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
}
