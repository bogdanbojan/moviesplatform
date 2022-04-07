package db

// DataPuller is the behaviour implemented by our chosen way of storing data.
type DataPuller interface {
	GetUsers() Users
	GetUser(userId string) (User, bool)
	ValidateServicePermission(sfp []string) bool
}

// Storage implements DataPuller and holds the structure for our service.permission.feature model.
type Storage struct {
	ServicesStructure Services
}

// InitStorage initializes the database by unmarshalling the datastore.json file.
func InitStorage() {
	JsonUnmarshal()
}

// NewStorage is the factory function for the Storage struct.
func NewStorage() *Storage {
	InitStorage()
	return &Storage{
		ServicesStructure: InitServiceStructure(),
	}
}

// GetUser pulls a particular user out of our database, based on his userId.
func (s *Storage) GetUser(userId string) (User, bool) {
	v, ok := UsersCollection[userId]
	return v, ok
}

// GetUsers pulls all the users from our database.
func (s *Storage) GetUsers() Users {
	return UsersCollection
}

// ValidateServicePermission checks that the service.feature.permission model is correct.
func (s *Storage) ValidateServicePermission(sfp []string) bool {
	_, ok := s.ServicesStructure[sfp[0]][sfp[1]][sfp[2]]
	return ok
}

type Services map[string]ServiceFeatures

type ServiceFeatures map[string]ServicePermissions

type ServicePermissions map[string]struct{}

// InitServiceStructure constructs our service.feature.permission model.
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
