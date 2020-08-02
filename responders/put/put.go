package putresponder

import (
	jsonreader "github.com/rolorin/pd-apollo-loki/utility/jsonReader"
)

func Put(filePath string) interface{} {
	return jsonreader.JSONReader(filePath)
}
