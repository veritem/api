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
	"github.com/veritem/api/pkg/github"
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

func (r *mutationResolver) CreateExperience(ctx context.Context, input model.CreateExperienceInput) (*model.Experience, error) {
	var startedAt *time.Time

	var endedAt *time.Time

	var err error

	startedAt, err = parseDate(input.StartedAt)

	if input.EndedAt != nil {
		endedAt, err = parseDate(*input.EndedAt)
	}

	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	var expreince = db.Experience{
		Name:      input.Name,
		StartedAt: *startedAt,
		EndedAt:   *endedAt,
		Roles:     input.Roles,
	}

	result := db.DB.Create(&expreince)

	if result.Error != nil {
		return nil, gqlerror.Errorf("Failed to create project!")
	}

	EndedAt := utils.FormatTime(expreince.EndedAt)

	return &model.Experience{
		ID:        expreince.ID,
		Name:      expreince.Name,
		StartedAt: utils.FormatTime(expreince.StartedAt),
		EndedAt:   &EndedAt,
		CreatedAt: utils.FormatTime(expreince.CreatedAt),
		UpdatedAt: utils.FormatTime(expreince.UpdatedAt),
		Roles:     expreince.Roles,
	}, nil
}

func (r *mutationResolver) CreateProject(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	var projectEcosystem db.ProjectEcosystem

	db.DB.Where("id = ? ", input.CategoryID).First(&projectEcosystem)

	if projectEcosystem.Name == "" {
		return nil, gqlerror.Errorf("Peoject Ecosystem not found!")
	}

	newProject := db.Project{
		ProjectEcosystemID: projectEcosystem.ID,
		Name:               input.Name,
		Description:        input.Description,
		ProjectURL:         input.ProjectURL,
		GithubURL:          input.GithubURL,
		Public:             input.IsPublic,
		Logo:               input.Logo,
	}

	db.DB.Create(&newProject)

	return &model.Project{
		ID:          newProject.ID,
		Name:        newProject.Name,
		Description: newProject.Description,
		ProjectURL:  newProject.ProjectURL,
		GithubURL:   newProject.GithubURL,
		CreatedAt:   utils.FormatTime(newProject.CreatedAt),
		UpdatedAt:   utils.FormatTime(newProject.UpdatedAt),
		IsPublic:    newProject.Public,
	}, nil
}

func (r *mutationResolver) CreateProjectEcosystem(ctx context.Context, input model.ProjectEcoInput) (*model.ProjectEcosystem, error) {
	ecosystem := db.ProjectEcosystem{
		Name: input.Name,
	}

	result := db.DB.Create(&ecosystem)

	if result.Error != nil {
		return nil, gqlerror.Errorf("Failed to create ecosystem!")
	}

	return &model.ProjectEcosystem{
		Name:      ecosystem.Name,
		ID:        ecosystem.ID,
		CreatedAt: utils.FormatTime(ecosystem.CreatedAt),
		UpdatedAt: utils.FormatTime(ecosystem.UpdatedAt),
	}, nil
}

func (r *queryResolver) Names(ctx context.Context) (*model.Name, error) {
	return &model.Name{First: "Verite", Last: "Makuza", Middle: "Mugabo", Username: "veritem"}, nil
}

func (r *queryResolver) Status(ctx context.Context) (string, error) {
	return "Building opensource for humans!", nil
}

func (r *queryResolver) Blogs(ctx context.Context) ([]*model.Blog, error) {
	const base = 10

	blogs := utils.GetBlogs()

	if len(blogs) == 0 {
		return nil, nil
	}

	response := make([]*model.Blog, 0)

	for index, item := range blogs {
		response = append(response, &model.Blog{
			ID:          fmt.Sprint(index + 1),
			LastUpdated: strconv.FormatInt(item.Published, base),
			URL:         item.URL,
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

	response := make([]*model.Skill, 0)

	for _, skill := range skills {
		response = append(response, &model.Skill{
			ID:          fmt.Sprint(skill.ID),
			Name:        skill.Name,
			Description: skill.Description,
			CreatedAt:   utils.FormatTime(skill.CreatedAt),
			UpdatedAt:   utils.FormatTime(skill.UpdatedAt),
		})
	}

	return response, nil
}

func (r *queryResolver) OpenSource(ctx context.Context) (*model.OpenSource, error) {
	return github.Contributions(), nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var projects []*db.Project

	dbQuery := db.DB.Find(&projects)

	if dbQuery.Error != nil {
		return nil, gqlerror.Errorf("Failed to get Projects")
	}

	var response = make([]*model.Project, 0)

	for _, item := range projects {
		response = append(response, &model.Project{
			Name:        item.Name,
			ID:          item.ID,
			Logo:        &item.Logo,
			Description: item.Description,
			CreatedAt:   utils.FormatTime(item.CreatedAt),
			UpdatedAt:   utils.FormatTime(item.UpdatedAt),
			IsPublic:    item.Public,
			GithubURL:   item.GithubURL,
			ProjectURL:  item.ProjectURL,
		})
	}

	return response, nil
}

func (r *queryResolver) Experiences(ctx context.Context) ([]*model.Experience, error) {
	var expreinces []db.Experience

	db.DB.Find(&expreinces)

	var response = make([]*model.Experience, 0)

	//nolint:gocritic // fixed
	for _, item := range expreinces {
		endedAt := utils.FormatTime(item.EndedAt)

		response = append(response, &model.Experience{
			Name:      item.ID,
			ID:        item.ID,
			StartedAt: utils.FormatTime(item.StartedAt),
			EndedAt:   &endedAt,
			CreatedAt: utils.FormatTime(item.CreatedAt),
			UpdatedAt: utils.FormatTime(item.UpdatedAt),
			Roles:     item.Roles,
		})
	}

	return response, nil
}

func (r *queryResolver) ProjectsEcosystems(ctx context.Context) ([]*model.ProjectEcosystem, error) {
	var projectsEco []db.ProjectEcosystem

	db.DB.Preload("Projects").Find(&projectsEco)

	var response = make([]*model.ProjectEcosystem, 0)

	for _, item := range projectsEco {
		response = append(response, &model.ProjectEcosystem{
			Name:      item.Name,
			ID:        item.ID,
			CreatedAt: utils.FormatTime(item.CreatedAt),
			UpdatedAt: utils.FormatTime(item.UpdatedAt),
			Projects:  convertProject(item.Projects),
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
func parseDate(date string) (*time.Time, error) {
	tm, err := time.Parse("02/2006", date)

	if err != nil {
		return nil, err
	}

	return &tm, nil
}
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
func convertProject(dbProject []db.Project) []*model.Project {
	var projects = make([]*model.Project, 0)

	var dbProjects []*db.Project

	for i := 0; i < len(dbProject); i++ {
		dbProjects = append(dbProjects, &dbProject[i])
	}

	for _, item := range dbProjects {
		projects = append(projects, &model.Project{
			Name:        item.Name,
			ID:          item.ID,
			Description: item.Description,
			CreatedAt:   utils.FormatTime(item.CreatedAt),
			UpdatedAt:   utils.FormatTime(item.UpdatedAt),
			IsPublic:    item.Public,
			GithubURL:   item.GithubURL,
			ProjectURL:  item.ProjectURL,
		})
	}

	return projects
}
