<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>LIST</title>
</head>
<body>

<h1> LIST All USERS</h1>

{{ range .}}
    USER NAME {{.UserName}}<br>
    FIRST {{.First}}<br>
    LAST {{.Last}}<br>
{{ end}}
<br>
<h2>Go to <a href="/bar">bar</a></h2><br>
<h2>Go to <a href="/logout">logout</a></h2>
</body>
</html>