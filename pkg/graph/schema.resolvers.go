package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/veritem/api/pkg/db"
	"github.com/veritem/api/pkg/graph/generated"
	"github.com/veritem/api/pkg/graph/model"
	"github.com/veritem/api/pkg/utils"
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
		CreatedAt:   utils.FormatTime(skillsCat.CreatedAt),
		UpdatedAt:   utils.FormatTime(skillsCat.UpdatedAt),
		Name:        skillsCat.Name,
		Description: skillsCat.Description,
	}

	return response, nil
}

func (r *mutationResolver) CreateSkill(ctx context.Context, input model.SkillInput) (*model.Skill, error) {
	var skillCategory db.SkillsCategory

	db.DB.Where("id = ? ", input.SkillsCategoryID).First(&skillCategory)

	if skillCategory.Name == "" {
		return nil, gqlerror.Errorf("Skill category not found!")
	}

	skill := db.Skill{
		Name:             input.Name,
		Description:      input.Description,
		SkillsCategoryID: input.SkillsCategoryID,
	}

	result := db.DB.Omit("skills_and_category.*").Create(&skill)

	if result.Error != nil {
		fmt.Println("Failed create skill")
	}

	response := &model.Skill{
		ID:          fmt.Sprint(skill.ID),
		Name:        skill.Name,
		Description: skill.Description,
		CreatedAt:   utils.FormatTime(skill.CreatedAt),
		UpdatedAt:   utils.FormatTime(skill.UpdatedAt),
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
	blogs := utils.GetBlogs()

	if len(blogs) == 0 {
		return nil, nil
	}

	var response []*model.Blog

	for index, item := range blogs {
		response = append(response, &model.Blog{
			ID:          fmt.Sprint(index + 1),
			LastUpdated: strconv.FormatInt(item.Published, 10),
			URL:         item.Url,
			Title:       item.Title,
			Image:       item.Image,
			Summary:     item.Summary,
		})
	}
	return response, nil
}

func (r *queryResolver) Socials(ctx context.Context) ([]*model.Social, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SkillsCategories(ctx context.Context) ([]*model.SkillsCategory, error) {
	var categories []db.SkillsCategory

	result := db.DB.Preload("Skills").Find(&categories)

	if result.Error != nil {
		return nil, gqlerror.Errorf("Failed to get categories: " + result.Error.Error())
	}

	response := make([]*model.SkillsCategory, 0)

	for _, skillCategory := range categories {
		response = append(response, &model.SkillsCategory{
			ID:          fmt.Sprint(skillCategory.ID),
			Name:        skillCategory.Name,
			Description: skillCategory.Description,
			CreatedAt:   utils.FormatTime(skillCategory.CreatedAt),
			UpdatedAt:   utils.FormatTime(skillCategory.UpdatedAt),
			Skills:      convertSkills(skillCategory.Skills),
		})
	}

	return response, nil
}

func (r *queryResolver) Skills(ctx context.Context) ([]*model.Skill, error) {
	var skills []db.Skill

	result := db.DB.Find(&skills)

	if result.Error != nil {
		return nil, gqlerror.Errorf("Failed to get skills!" + result.Error.Error())
	}

	response := make([]*model.Skill, 0, len([]*model.Skill{}))

	for _, skill := range skills {
		response = append(response, &model.Skill{
			ID:          fmt.Sprint(skill.ID),
			Name:        skill.Name,
			Description: skill.Description,
			CreatedAt:   skill.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   skill.UpdatedAt.Format(time.RFC3339),
		})
	}

	return response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func convertSkills(skills []db.Skill) []*model.Skill {

	skillsModels := make([]*model.Skill, 0)

	for _, item := range skills {
		skillsModels = append(skillsModels, &model.Skill{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			CreatedAt:   utils.FormatTime(item.CreatedAt),
			UpdatedAt:   utils.FormatTime(item.UpdatedAt),
		})
	}
	return skillsModels
}
