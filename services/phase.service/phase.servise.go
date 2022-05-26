package phase_service

import (
	m "main/models"
	phaseRepo "main/repositories/phase.repository"
)

func CreatePhase(phase m.Phase) error {

	err := phaseRepo.CreatePhase(phase)

	if err != nil {
		return err
	}

	return nil
}

func GetPhaseList() (m.Phases, error) {

	phaseList, err := phaseRepo.GetPhaseList()

	if err != nil {
		return nil, err
	}

	return phaseList, nil
}

func GetPhaseById(phaseId string) (*m.Phase, error) {

	phase, err := phaseRepo.GetPhaseById(phaseId)

	if err != nil {
		return nil, err
	}

	return phase, nil
}

func UpdatePhase(phase m.Phase, phaseId string) error {

	err := phaseRepo.UpdatePhase(phase, phaseId)

	if err != nil {
		return err
	}

	return nil
}

func GetPhaseMembers(phaseId string) (m.Users, error) {
	phase, err := GetPhaseById(phaseId)
	if err != nil {
		return nil, err
	}

	members, err := phaseRepo.GetPhaseMembers(phase.MembersId)

	if err != nil {
		return nil, err
	}

	return members, nil
}

func AddMemberPhase(phaseId string, memberId string) error {
	phase, err := GetPhaseById(phaseId)
	if err != nil {
		return err
	}
	//verificar que el usuario no este en la lista de miembros
	if !contains(phase.MembersId, memberId) {
		phase.MembersId = append(phase.MembersId, memberId)
		err = UpdatePhase(*phase, phaseId)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveMemberPhase(phaseId string, memberId string) error {
	phase, err := GetPhaseById(phaseId)
	if err != nil {
		return err
	}

	for i, v := range phase.MembersId {
		if v == memberId {
			phase.MembersId = append(phase.MembersId[:i], phase.MembersId[i+1:]...)
			break
		}
	}

	err = UpdatePhase(*phase, phaseId)
	if err != nil {
		return err
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
