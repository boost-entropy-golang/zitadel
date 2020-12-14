package view

import (
	proj_model "github.com/caos/zitadel/internal/project/model"
	"github.com/caos/zitadel/internal/project/repository/view"
	"github.com/caos/zitadel/internal/project/repository/view/model"
	"github.com/caos/zitadel/internal/view/repository"
	"time"
)

const (
	projectMemberTable = "management.project_members"
)

func (v *View) ProjectMemberByIDs(projectID, userID string) (*model.ProjectMemberView, error) {
	return view.ProjectMemberByIDs(v.Db, projectMemberTable, projectID, userID)
}

func (v *View) ProjectMembersByProjectID(projectID string) ([]*model.ProjectMemberView, error) {
	return view.ProjectMembersByProjectID(v.Db, projectMemberTable, projectID)
}

func (v *View) SearchProjectMembers(request *proj_model.ProjectMemberSearchRequest) ([]*model.ProjectMemberView, uint64, error) {
	return view.SearchProjectMembers(v.Db, projectMemberTable, request)
}

func (v *View) ProjectMembersByUserID(userID string) ([]*model.ProjectMemberView, error) {
	return view.ProjectMembersByUserID(v.Db, projectMemberTable, userID)
}

func (v *View) PutProjectMember(project *model.ProjectMemberView, sequence uint64, eventTimestamp time.Time) error {
	err := view.PutProjectMember(v.Db, projectMemberTable, project)
	if err != nil {
		return err
	}
	return v.ProcessedProjectMemberSequence(sequence, eventTimestamp)
}

func (v *View) PutProjectMembers(project []*model.ProjectMemberView, sequence uint64, eventTimestamp time.Time) error {
	err := view.PutProjectMembers(v.Db, projectMemberTable, project...)
	if err != nil {
		return err
	}
	return v.ProcessedProjectMemberSequence(sequence, eventTimestamp)
}

func (v *View) DeleteProjectMember(projectID, userID string, eventSequence uint64, eventTimestamp time.Time) error {
	err := view.DeleteProjectMember(v.Db, projectMemberTable, projectID, userID)
	if err != nil {
		return nil
	}
	return v.ProcessedProjectMemberSequence(eventSequence, eventTimestamp)
}

func (v *View) DeleteProjectMembersByProjectID(projectID string) error {
	return view.DeleteProjectMembersByProjectID(v.Db, projectMemberTable, projectID)
}

func (v *View) GetLatestProjectMemberSequence() (*repository.CurrentSequence, error) {
	return v.latestSequence(projectMemberTable)
}

func (v *View) ProcessedProjectMemberSequence(eventSequence uint64, eventTimestamp time.Time) error {
	return v.saveCurrentSequence(projectMemberTable, eventSequence, eventTimestamp)
}

func (v *View) UpdateProjectMemberSpoolerRunTimestamp() error {
	return v.updateSpoolerRunSequence(projectMemberTable)
}

func (v *View) GetLatestProjectMemberFailedEvent(sequence uint64) (*repository.FailedEvent, error) {
	return v.latestFailedEvent(projectMemberTable, sequence)
}

func (v *View) ProcessedProjectMemberFailedEvent(failedEvent *repository.FailedEvent) error {
	return v.saveFailedEvent(failedEvent)
}
