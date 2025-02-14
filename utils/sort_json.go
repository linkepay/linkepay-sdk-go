package utils

import (
	"encoding/json"
	"fmt"
	"sort"
)

// ToJSON converts any struct to a JSON string
func ToJSON(v interface{}) (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// ToSortedJSON converts any struct to a JSON string with sorted keys
func ToSortedJSON(v interface{}) (string, error) {
	// First marshal the data to get a JSON bytes
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	// Unmarshal into a map to sort the keys
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &jsonMap); err != nil {
		return "", err
	}

	// Marshal again with sorted keys
	sortedBytes, err := json.Marshal(sortMap(jsonMap))
	if err != nil {
		return "", err
	}

	return string(sortedBytes), nil
}

// SortJSON takes a JSON string and returns a sorted JSON string
func SortJSON(jsonStr string) (string, error) {
	var jsonMap map[string]interface{}

	// Unmarshal the JSON string into a map
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	// Sort the map and marshal back to JSON
	sortedBytes, err := json.Marshal(sortMap(jsonMap))
	if err != nil {
		return "", fmt.Errorf("failed to marshal sorted JSON: %v", err)
	}

	return string(sortedBytes), nil
}

// sortMap recursively sorts map keys
func sortMap(m map[string]interface{}) map[string]interface{} {
	sorted := make(map[string]interface{})
	keys := make([]string, 0, len(m))

	// Collect all keys
	for k := range m {
		keys = append(keys, k)
	}

	// Sort keys
	sort.Strings(keys)

	// Add sorted keys back to map
	for _, k := range keys {
		if subMap, ok := m[k].(map[string]interface{}); ok {
			// Recursively sort nested maps
			sorted[k] = sortMap(subMap)
		} else if subSlice, ok := m[k].([]interface{}); ok {
			// Handle arrays by sorting any maps within them
			sortedSlice := make([]interface{}, len(subSlice))
			for i, v := range subSlice {
				if subMap, ok := v.(map[string]interface{}); ok {
					sortedSlice[i] = sortMap(subMap)
				} else {
					sortedSlice[i] = v
				}
			}
			sorted[k] = sortedSlice
		} else {
			sorted[k] = m[k]
		}
	}

	return sorted
}
