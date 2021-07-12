package db

type SkillsCategory struct {
	Model
	Name        string
	Description string
	Skills      []Skill `gorm:"many2many:skills_categories"`
}
