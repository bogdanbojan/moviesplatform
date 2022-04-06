package db

type DataPuller interface {
	GetUsers() Users
	GetUser(userId string) (User, bool)
	ValidateServicePermission(sfp []string) bool
}

type Storage struct {
	ServicesStructure map[string]ServiceFeatures
}

func InitStorage() {
	JsonUnmarshal()
}

func NewStorage() *Storage {
	InitStorage()
	return &Storage{
		ServicesStructure: InitServiceStructure(),
	}
}

func (s *Storage) GetUser(userId string) (User, bool) {
	v, ok := UsersCollection[userId]
	return v, ok
}

func (s *Storage) GetUsers() Users {
	return UsersCollection
}

func (s *Storage) ValidateServicePermission(sfp []string) bool {
	_, ok := s.ServicesStructure[sfp[0]][sfp[1]][sfp[2]]
	return ok
}

type Services map[string]ServiceFeatures

type ServiceFeatures map[string]ServicePermissions

type ServicePermissions map[string]struct{}

func InitServiceStructure() Services {
	blockbustersFeatures := ServiceFeatures{
		"director":        ServicePermissions{"direct": {}, "instructActors": {}, "argue": {}},
		"cinematographer": ServicePermissions{"shoot": {}, "changeLens": {}, "changeCamera": {}},
		"producer":        ServicePermissions{"changeBudget": {}, "changeSalary": {}, "addActor": {}},
	}

	commercialsFeatures := ServiceFeatures{
		"artist":   ServicePermissions{"createConcept": {}, "creativitySwitch": {}},
		"producer": ServicePermissions{"getDeals": {}, "onboardPeople": {}},
		"manager":  ServicePermissions{"adviseBoard": {}, "cancelMeetings": {}, "scheduleMeetings": {}},
	}

	shortsFeatures := ServiceFeatures{
		"actor":    ServicePermissions{"act": {}, "readScript": {}, "cryOnCommand": {}},
		"investor": ServicePermissions{"scandal": {}, "modifyBudget": {}},
		"director": ServicePermissions{"act": {}, "invest": {}, "direct": {}},
	}

	s := Services{
		"blockbusters": blockbustersFeatures,
		"commercials":  commercialsFeatures,
		"shorts":       shortsFeatures,
	}

	return s
}
