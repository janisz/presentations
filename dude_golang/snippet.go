commonTags := labelsToTags(app.Labels) // HLxxx
var intents []RegistrationIntent
for _, d := range definitions {
        intents = append(intents, RegistrationIntent{
                Name: app.labelsToName(d.Labels, nameSeparator),
                Port: task.Ports[d.Index],
                Tags: append(commonTags, labelsToTags(d.Labels)...), // HLxxx
        })
}

func labelsToTags(labels map[string]string) []string {
	tags := []string{} // HLxxx
	for key, value := range labels {
		if value == "tag" {
			tags = append(tags, key)
		}
	}
	return tags
}