<!DOCTYPE html>

<html>


<h1>Home</h1>

<a href="/form">Upload a new file and create a room!</a>

{{range .Rooms}}

<p>
    <a href="/r/{{.Index}}">{{.Connected}} watching {{.Name}}</a>
</p>

{{end}}

</html>