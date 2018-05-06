package main

import (
	"bytes"
	"image/jpeg"

	"github.com/sheng/air"
	"github.com/sheng/cameron"
)

func main() {
	air.GET("/", index)
	air.GET("/identicons/:Name", identicon)
	air.Serve()
}

func index(req *air.Request, res *air.Response) error {
	return res.HTML(`
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
	buf := &bytes.Buffer{}
	jpeg.Encode(
		buf,
		cameron.Identicon([]byte(req.Params["Name"]), 540, 50),
		&jpeg.Options{
			Quality: 100,
		},
	)
	return res.Blob("image/jpeg", buf.Bytes())
}
