<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
<ul>
    {{ range .Wisdom }}
    <li>{{.Name | uc}} -  {{uc .Motto}}</li>
    {{end}}
</ul>
<ul>
    {{ range .Transport }}
        <li>{{.Manufacturer | ft}} - {{.Model}} - {{.Doors}}</li>
    {{end}}
</ul>

<h1>{{.Timenow | ftime}}</h1>
</body>
</html>