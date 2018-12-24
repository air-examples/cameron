package main

import (
	"bytes"
	"image/jpeg"

	"github.com/aofei/air"
	"github.com/aofei/cameron"
)

func main() {
	a := air.Default
	a.DebugMode = true
	a.GET("/", index)
	a.GET("/identicons/:Name", identicon)
	a.Serve()
}

func index(req *air.Request, res *air.Response) error {
	return res.WriteHTML(`
<!DOCTYPE html>
<html>
  <head>
    <title>Cameron</title>
    <meta name="description" content="Cameron - An avatar generator for Go.">
  </head>

  <body>
    <h1>Cameron - An avatar generator for Go.</h1>

    <h2>Identicons</h2>
    <ul>
      <li><a href="/identicons/Robb Stark">Robb Stark's identicon</a></li>
      <li><a href="/identicons/Jon Snow">Jon Snow's identicon</a></li>
      <li><a href="/identicons/Sansa Stark">Sansa Stark's identicon</a></li>
      <li><a href="/identicons/Arya Stark">Arya Stark's identicon</a></li>
      <li><a href="/identicons/Bran Stark">Bran Stark's identicon</a></li>
      <li><a href="/identicons/Rickon Stark">Rickon Stark's identicon</a></li>
    </ul>
  </body>
</html>
`)
}

func identicon(req *air.Request, res *air.Response) error {
	buf := bytes.Buffer{}
	jpeg.Encode(
		&buf,
		cameron.Identicon(
			[]byte(req.Param("Name").Value().String()),
			540,
			60,
		),
		&jpeg.Options{
			Quality: 100,
		},
	)

	res.Header.Set("Content-Type", "image/jpeg")

	return res.Write(bytes.NewReader(buf.Bytes()))
}
