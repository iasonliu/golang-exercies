<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>

<form method="post">

    <input type="email" name="username" placeholder="email"><br>
    <input type="text" name="firstname" placeholder="first name"><br>
    <input type="text" name="lastname" placeholder="last name"><br>
    <input type="submit">

</form>


{{if .First}}
USER NAME {{.UserName}}<br>
FIRST {{.First}}<br>
LAST {{.Last}}<br>
{{end}}

<br>
<h2>Go to <a href="/bar">the bar</a></h2>
</body>
</html>
