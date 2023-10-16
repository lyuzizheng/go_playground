package json

import (
	"fmt"
	"strings"
)

var strToBeEncoded = `
{
  "title": "Scaler History",
  "message": "The below are recent 10 scaler history: \n\n",
  "fields": [{{range $val := .}}"**{{ .PSM }}**\n{{ .Region }}-{{ .Cluster }}-{{ .Env }}\nInstance: **{{ .InstanceFrom }} -> {{ .InstanceTo }}**\nCores: **{{ .CoresFrom }} -> {{ .CoresTo }}**\nStart: {{ .CreatedAt }}\nEnd: {{ .UpdatedAt }}\nStatus: {{ .Status }}\n",{{end}}"END_OF_ALL_FIELDS"]
}
`

func JsonEncode() string {

	withoutquotation := strings.ReplaceAll(strToBeEncoded, "\"", "\\\"")
	withoutNewLineChar := strings.ReplaceAll(withoutquotation, "\\n", "\\\\n")
	removeNewLine := strings.ReplaceAll(withoutNewLineChar, "\n", "")

	fmt.Println(removeNewLine)
	return removeNewLine

}
