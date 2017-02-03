# savetofile - a consul-template plugin

Give credit where credit is due, this is based off of https://gist.github.com/tam7t/1b45125ae4de13b3fc6fd0455954c08e I just wanted it in a repo, with releases so I can install the binary through automated means.

I also wanted to tweak it a bit to be just about saving the contents to a file and allow passing of the UID and GID instead of the user.

## Example Usage

```
{{ with secret "pki/example-dot-com" "common_name=testing.example.com" }}
{{ .Data.serial_number }}
{{if .Data.certificate}}{{ .Data.certificate | plugin "/opt/bin/savetofile" "/opt/ssl/apiserver.pem" "0" }}{{end}}
{{if .Data.private_key}}{{ .Data.private_key | plugin "/opt/bin/savetofile" "/opt/ssl/apiserver-key.pem" "0" }}{{end}}
{{ end }}
```
