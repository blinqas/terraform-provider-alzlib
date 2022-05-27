package provider

import "encoding/json"

func flattenJSON(stringMap interface{}) string {
	if stringMap != nil {
		value := stringMap.(map[string]interface{})
		jsonString, err := flattenJsonToString(value)
		if err == nil {
			return jsonString
		}
	}

	return ""
}

func flattenJsonToString(input map[string]interface{}) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	result, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
