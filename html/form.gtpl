<html>

<head>
    <title>Upload file</title>
</head>

<body>
    <form enctype="multipart/form-data" action="http://{{.Host}}:{{.Port}}/upload" method="post">
        <p> Video:
            <input type="file" name="video" />
        </p>


        <p> Subtitles:
            <input type="file" name="subtitle" />
        </p>

        <input type="submit" value="upload" />
    </form>
</body>

</html>