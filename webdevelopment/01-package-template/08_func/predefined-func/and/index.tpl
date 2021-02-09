<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>predefined global functions</title>
</head>
<body>
EXAMPLE #1
{{range .}}
    {{.}}
{{end}}

EXAMPLE #2
{{if .}}
    EXAMPLE #2 - {{.}}
{{end}}

EXAMPLE #3
{{range .}}
    {{if .Name}}
        EXAMPLE #3 - {{.Name}}
    {{end}}
{{end}}

EXAMPLE #4
{{range .}}
    {{if and .Name .Admin}}
        EXAMPLE #4 - Name: {{.Name}}
        EXAMPLE #4 - Motto: {{.Motto}}
        EXAMPLE #4 - Admin: {{.Admin}}
    {{end}}
{{end}}
</body>
</html>