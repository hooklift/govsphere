package vim

type ManagedObject struct {
	mor *ManagedObjectReference
}

func (mo *ManagedObject) currentProperty(property string) (interface{}, error) {
	return nil, nil
}
