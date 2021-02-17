# savetofile - a consul-template plugin

Give credit where credit is due, this is based off of https://gist.github.com/tam7t/1b45125ae4de13b3fc6fd0455954c08e I just wanted it in a repo, with releases so I can install the binary through automated means.

I also wanted to tweak it a bit to be just about saving the contents to a file and allow passing of the UID and GID instead of the username along with a few other minor tweaks.

## Usage

```bash
savetofile <mode> <path> <uid> <gid> <perm> <contents>
```

Valid Values for Mode:
 * create
 * create-nl (will add a `\n` after the contents)
 * append
 * append-nl (will add a `\n` after the contents)

## Example Consul Template Usage

```
{{ with secret "pki/issue/consul" "common_name=consul.example.com" "ttl=1h" }}
  {{ if .Data.ca_chain }}
  {{ range $key, $value := .Data.ca_chain }}
    {{ if eq $key 0 }}
      {{ $value | plugin "/usr/local/bin/savetofile" "create-nl" "/etc/consul/ssl/ca.crt" "100" "1000" "644" }}
    {{ else }}
      {{ $value | plugin "/usr/local/bin/savetofile" "append-nl" "/etc/consul/ssl/ca.crt" "100" "1000" "644" }}
    {{ end }}
  {{ end }}
  {{ end }}
  
  {{ if .Data.certificate }}
  {{ .Data.certificate | plugin "/usr/local/bin/savetofile" "create-nl" "/etc/consul/ssl/consul.crt" "100" "1000" "644" }}
  {{ end }}

  {{ if .Data.private_key }}
  {{ .Data.private_key | plugin "/usr/local/bin/savetofile" "create-nl" "/etc/consul/ssl/consul.key" "100" "1000" "600" }}
  {{ end }}
{{ end }}

```

## License

MIT
