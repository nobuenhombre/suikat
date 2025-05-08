{{define "body"}}
    {{ .Content }}
    {{ .25 | formatPercent }}
{{end}}
