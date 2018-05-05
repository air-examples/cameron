package main

import (
	"bytes"
	"image/jpeg"

	"github.com/sheng/air"
	"github.com/sheng/cameron"
)

func main() {
	air.GET("/", index)
	air.GET("/identicon", identicon)
	air.Serve()
}

func index(req *air.Request, res *air.Response) error {
	return res.HTML(`
<!DOCTYPE html>
<html>
  <head>
    <title>Cameron</title>
  </head>

  <body>
    <ul>
      <li><a href="/identicon?data=Robb+Stark">Identicon (Robb Stark)</a></li>
      <li><a href="/identicon?data=Jon+Snow">Identicon (Jon Snow)</a></li>
      <li><a href="/identicon?data=Sansa+Stark">Identicon (Sansa Stark)</a></li>
      <li><a href="/identicon?data=Arya+Stark">Identicon (Arya Stark)</a></li>
      <li><a href="/identicon?data=Bran+Stark">Identicon (Bran Stark)</a></li>
      <li><a href="/identicon?data=Rickon+Stark">Identicon (Rickon Stark)</a></li>
    </ul>
  </body>
</html>
`)
}

func identicon(req *air.Request, res *air.Response) error {
	buf := &bytes.Buffer{}
	jpeg.Encode(
		buf,
		cameron.NewIdenticon([]byte(req.Params["data"]), 540, 50),
		&jpeg.Options{
			Quality: 100,
		},
	)
	return res.Blob("image/jpeg", buf.Bytes())
}
