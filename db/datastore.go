package db

// DataPuller is the behaviour implemented by our chosen way of storing data.
type DataPuller interface {
	GetUsers() Users
	GetUser(userId string) (User, bool)
	ValidateServiceFeaturePermission(sfp []string) bool
}

// Storage implements DataPuller and holds the structure for our service.permission.feature model.
type Storage struct {
	ServicesStructure Services
}

// NewStorage is the factory function for the Storage struct.
func NewStorage() *Storage {
	return &Storage{
		ServicesStructure: InitServiceStructure(),
	}
}

// InitStorage initializes the database by unmarshalling the datastore.json file.
// If you want the file embedded from the datastore.json automatically, pass an empty string.
// If you want to specify a certain file, pass the name of the file which should be stored in the ./db directory.
func (s *Storage) InitStorage(fileName string) {
	if fileName == Embed {
		s.JSONUnmarshalEmbed()
		return
	}
	s.JSONUnmarshalFile(fileName)
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

// ValidateServiceFeaturePermission checks that the service.feature.permission model is correct.
func (s *Storage) ValidateServiceFeaturePermission(sfp []string) bool {
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
