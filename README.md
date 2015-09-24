# envtpl
_a port of https://github.com/andreasjansson/envtpl to Go_

With this template file

something.conf.tpl
```
 foo = {{ .FOO }}
 bar = "{{ .BAR }}"
```

Running
```
$ FOO=123 BAR=abc envtpl something.conf.tpl outputfile.conf
```

will generate
```
 foo = 123
 bar = "abc"
```
