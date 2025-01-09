package data

type Employee struct {
	BaseEntity
	Name         string `gorm:"size:255;not null"`
	Phone        string `gorm:"size:20;unique;not null"`
	IsVerified   bool   `gorm:"default:false"`
	DepartmentID *uint  `gorm:"index"` // Nullable foreign key
	Department   *Department
	Role         string `gorm:"size:50;not null;check:role IN ('ADMIN', 'EMPLOYEE')"`
	Password     string `gorm:"size:255;not null"`
}

type Department struct {
	BaseEntity
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	Phone       string `gorm:"size:20"`
	Employees   []Employee
}

type TaskType struct {
	BaseEntity
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
}

type DepartmentTaskType struct {
	BaseEntity
	DepartmentID uint       `gorm:"index;not null"`
	Department   Department `gorm:"foreignKey:DepartmentID;constraint:OnDelete:CASCADE"`
	TaskTypeID   uint       `gorm:"index;not null"`
	TaskType     TaskType   `gorm:"foreignKey:TaskTypeID;constraint:OnDelete:CASCADE"`
}
