package graph

import "github.com/codeedu/fc2-graphql/graph/model"

type Resolver struct {
	Categories []*model.Category
	Courses    []*model.Course
	Chapters   []*model.Chapter
}
