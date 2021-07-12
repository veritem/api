package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/veritem/api/pkg/db"
	"github.com/veritem/api/pkg/graph/generated"
	"github.com/veritem/api/pkg/graph/model"
)

func (r *mutationResolver) CreateSkillCategory(ctx context.Context, input model.SkillsCategoryInput) (*model.SkillsCategory, error) {
	skillsCat := db.SkillsCategory{
		Name:        input.Name,
		Description: input.Description,
	}

	result := db.DB.Create(&skillsCat)

	if result.Error != nil {
		return nil, gqlerror.Errorf("Failed to create skill category: " + result.Error.Error())
	}

	response := &model.SkillsCategory{
		ID:          fmt.Sprint(skillsCat.ID),
		CreatedAt:   skillsCat.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   skillsCat.UpdatedAt.Format(time.RFC3339),
		Name:        skillsCat.Name,
		Description: skillsCat.Description,
	}

	return response, nil
}

func (r *mutationResolver) CreateSkill(ctx context.Context, input model.SkillInput) (*model.Skill, error) {
	skillCatId, err := strconv.ParseUint(input.SkillsCategoryID, 10, 64)

	if err != nil {
		return nil, gqlerror.Errorf("Failed to parse skillscategry id")
	}

	var skillCategory db.SkillsCategory

	db.DB.Where("id = ? ", skillCatId).First(&skillCategory)

	log.Println(skillCategory)

	if skillCategory.Name == "" {
		return nil, gqlerror.Errorf("Skill category not found!")
	}

	skill := db.Skill{
		Name:             input.Name,
		Description:      input.Description,
		SkillsCategoryID: uint(skillCatId),
	}

	result := db.DB.Create(&skill)

	if result.Error != nil {
		fmt.Println("Failed create skill")
	}
	skill_category := &model.SkillsCategory{
		ID:          fmt.Sprint(skillCategory.ID),
		CreatedAt:   skillCategory.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   skillCategory.UpdatedAt.Format(time.RFC3339),
		Name:        skillCategory.Name,
		Description: skillCategory.Description,
	}
	response := &model.Skill{
		ID:          fmt.Sprint(skill.ID),
		Name:        skill.Name,
		Description: skill.Description,
		CreatedAt:   skill.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   skill.UpdatedAt.Format(time.RFC3339),
		Category:    skill_category,
	}

	return response, nil
}

func (r *mutationResolver) GenerateSecret(ctx context.Context, input model.ScretInput) (*model.Secret, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Names(ctx context.Context) (*model.Name, error) {
	return &model.Name{First: "Verite", Last: "Makuza", Middle: "Mugabo", Username: "veritem"}, nil
}

func (r *queryResolver) Status(ctx context.Context) (string, error) {
	return "Building stuffs to help people learn and grow!", nil
}

func (r *queryResolver) Blogs(ctx context.Context) ([]*model.Blog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Socials(ctx context.Context) ([]*model.Social, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SkillsCategries(ctx context.Context) ([]*model.SkillsCategory, error) {
	var categories []*model.SkillsCategory
	result := db.DB.Statement.Find(&categories)

	if result.Error != nil {
		return nil, gqlerror.Errorf("Failed to get categories!" + result.Error.Error())
	}

	return categories, nil
}

func (r *queryResolver) Skills(ctx context.Context) ([]*model.Skill, error) {
	var skills []*model.Skill
	result := db.DB.Find(&skills)

	if result.Error != nil {
		return nil, gqlerror.Errorf("Failed to get skills!" + result.Error.Error())
	}
	return skills, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
