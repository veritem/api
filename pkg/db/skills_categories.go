package db

type SkillsCategory struct {
	Model
	Name        string
	Description string
	Skills      []Skill `gorm:"foreignKey:SkillsCategoryID"`
}
