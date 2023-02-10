<html>
<head>
    <title>upload test</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
    <input type="file" name="upload_file" />
    <input type="hidden" name="token" value="{{.}}" />
    <input type="submit" value="upload /">
</form>
</body>
</html>