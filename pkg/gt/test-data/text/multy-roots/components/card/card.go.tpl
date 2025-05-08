{{define "card"}}
    ## Hello
    Card {{ .25 | formatPercent }}
    {{ template "button-blue" }}
    {{ template "button-red" }}
{{end}}
