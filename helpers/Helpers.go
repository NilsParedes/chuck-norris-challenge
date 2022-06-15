package helpers

import "github.com/nilsparedes/chuck-norris-challenge/entities"

func FilterUniqueItems(items []entities.Item) []entities.Item {
	allKeys := make(map[string]bool)
	var uniqueItems []entities.Item

	for _, item := range items {
		if value := allKeys[item.Id]; !value {
			allKeys[item.Id] = true
			uniqueItems = append(uniqueItems, item)
		}
	}

	return uniqueItems
}
