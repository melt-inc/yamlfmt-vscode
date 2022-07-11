module github.com/melt-inc/yamlfmt-vscode

// TODO(jr) this is just for local dev, remove when published
replace github.com/jamesrom/yamlfmt v0.0.0 => C:\Projects\yamlfmt

// Once https://github.com/kubernetes-sigs/yaml/pull/76 lands, remove this
replace sigs.k8s.io/yaml v1.3.0 => github.com/natasha41575/yaml-1 v1.3.1-0.20220514005426-0e00b683066c

go 1.18

require github.com/jamesrom/yamlfmt v0.0.0

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	golang.org/x/text v0.3.7 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
