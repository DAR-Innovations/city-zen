package data

type UserReport struct {
	BaseEntity
	IssueID     uint   `gorm:"index"`
	Issue       Issue  `gorm:"foreignKey:IssueID"`
	Description string `gorm:"type:text"`
	ImageURL    string `gorm:"size:255"`
	ReportedBy  uint   `gorm:"index"`
	Reporter    User   `gorm:"foreignKey:ReportedBy"`
}

type DepartmentReport struct {
	BaseEntity
	IssueID     uint       `gorm:"index"`
	Issue       Issue      `gorm:"foreignKey:IssueID"`
	Description string     `gorm:"type:text"`
	ImageURL    string     `gorm:"size:255"`
	ReportedBy  uint       `gorm:"index"`
	Reporter    Department `gorm:"foreignKey:ReportedBy"`
}
