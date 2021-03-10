<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>BAR</title>
</head>
<body>

<h1>Welcome to the bar. What can I get you to drink?</h1>

{{if .First}}
    USER NAME {{.UserName}}<br>
    PASSWORD {{.Password}}<br>
    ROLE {{.Role}}<br>
    FIRST {{.First}}<br>
    LAST {{.Last}}<br>
{{end}}
<br>
<h2>Go to <a href="/list">list</a></h2><br>
<h2>Go to <a href="/logout">logout</a></h2>
</body>
</html>