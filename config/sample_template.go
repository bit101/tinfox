// Package config has config related functions.
package config

var htmlTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <title>${TITLE}</title>
    <link rel="stylesheet" href="styles/main.css">
  </head>
  <body>
    <script type="text/javascript" src="src/main.js"></script>
  </body>
</html>
`

var cssTemplate = `
h1 {
  font-family: Arial;
  font-size: 24px;
}
`

var jsTemplate = `
var div = document.createElement("h1");
div.innerText = "Hello, world!";
document.body.appendChild(div);
`

var jsonTemplate = `
{
  "name": "HTML",
  "description": "Barebones HTML Project with JavaScript and CSS",
  "tokens": [
    {
      "name": "TITLE",
      "default": "Hello world"
    }
  ],
  "preMessage": "This is just a barebones HTML project.",
 	"postMessage": "Go into the new project dir and open up 'index.html' in a browser."
}
`
