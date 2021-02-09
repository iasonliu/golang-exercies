<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
<ul>
    {{ range $person := .}}
    <li>{{$person.Name}} -  {{$person.Age}}</li>
    {{end}}
</ul>
</body>
</html>