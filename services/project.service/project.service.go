package project_service

import (
	m "main/models"

	projectRepo "main/repositories/project.repository"
)

func CreateProject(project m.Project) error {

	err := projectRepo.CreateProject(project)

	if err != nil {
		return err
	}

	return nil
}

func GetProjectList() (m.Projects, error) {

	projects, err := projectRepo.GetProjectList()

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func GetProjectById(projectId string) (*m.Project, error) {

	project, err := projectRepo.GetProjectById(projectId)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func GetProyectListForLeader(leaderId string) (m.Projects, error) {

	projects, err := projectRepo.GetProyectListForLeader(leaderId)

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func GetProjectListForUser(userId string) (m.Projects, error) {

	projects, err := projectRepo.GetProjectListForUser(userId)

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func UpdateProject(project m.Project, projectId string) error {

	err := projectRepo.UpdateProject(project, projectId)

	if err != nil {
		return err
	}

	return nil
}

func GetProjectMembers(projectId string) (m.Users, error) {
	project, err := GetProjectById(projectId)
	if err != nil {
		return nil, err
	}

	members, err := projectRepo.GetProjectMembers(project.MembersId)

	if err != nil {
		return nil, err
	}

	return members, nil
}

func GetProjectPhases(projectId string) (m.Phases, error) {
	project, er := GetProjectById(projectId)
	if er != nil {
		return nil, er
	}

	phases, err := projectRepo.GetProjectPhases(project.Phases)
	if err != nil {
		return nil, err
	}

	return phases, nil
}

func AddMemberProject(projectId string, userId string) error {
	project, err := GetProjectById(projectId)
	if err != nil {
		return err
	}

	//verify if user is already in project
	if contains(project.MembersId, userId) {
		project.MembersId = append(project.MembersId, userId)
		err = UpdateProject(*project, projectId)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemovedMemberProject(projectId string, userId string) error {
	project, err := GetProjectById(projectId)
	if err != nil {
		return err
	}

	for i, v := range project.MembersId {
		if v == userId {
			project.MembersId = append(project.MembersId[:i], project.MembersId[i+1:]...)
			break
		}
	}

	err = UpdateProject(*project, projectId)
	if err != nil {
		return err
	}

	return nil
}

func GetProjectMembersNotInProject(projectId string) (m.Users, error) {
	project, err := GetProjectById(projectId)
	if err != nil {
		return nil, err
	}

	members, err := projectRepo.GetProjectMembersNotInProject(project.MembersId, projectId)

	if err != nil {
		return nil, err
	}

	return members, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
