package vim

func ToManagedEntities(mors []*ManagedObjectReference) []*ManagedEntity {
	var e []*ManagedEntity
	for _, r := range mors {
		e = append(e, &ManagedEntity{
			ExtensibleManagedObject: &ExtensibleManagedObject{
				ManagedObject: &ManagedObject{
					ManagedObjectReference: r,
				},
			},
		})
	}

	return e
}
