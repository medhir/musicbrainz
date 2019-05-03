package mbclient

func (c *MBClient) getTypeString(typeFilters []string) string {
	typesString := "("
	for i, filter := range typeFilters {
		typesString += "type:" + filter
		if i < len(typeFilters)-1 {
			typesString += " OR "
		}
	}
	typesString += ")"
	return typesString
}
