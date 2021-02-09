<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
<ul>
    {{$x := .Name}}
    {{$y := .Age}}
    <li>{{$x}} -  {{$y}}</li>
</ul>
</body>
</html>