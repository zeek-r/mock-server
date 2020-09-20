package path

import "strings"

// GetResponsePath returns the path for response file
func GetResponsePath(mockDir, httpMethod, path string) string {
	method := strings.ToLower(httpMethod)
	responsePath := strings.Split(path, "v1")[1]
	return "./" + mockDir + "/" + method + "/" + responsePath + ".json"
}
