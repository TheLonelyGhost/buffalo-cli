package ci

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/gogen"
)

//go:embed templates/*
var templates embed.FS

// New generator for adding travis, gitlab, or circleci
func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	g.Transformer(genny.Replace("-no-pop", ""))
	g.Transformer(genny.Dot())

	var fname string
	switch opts.Provider {
	case "travis", "travis-ci":
		fname = "-dot-travis.yml.tmpl"
	case "gitlab", "gitlab-ci":
		if opts.App.WithPop {
			fname = "-dot-gitlab-ci.yml.tmpl"
		} else {
			fname = "-dot-gitlab-ci-no-pop.yml.tmpl"
		}
	case "circleci":
		fname = "-dot-circleci/config.yml.tmpl"
	default:
		return g, fmt.Errorf("could not find a template for %s", opts.Provider)
	}

	sub, err := fs.Sub(templates, "templates")
	if err != nil {
		return g, err
	}

	f, err := sub.Open(fname)
	if err != nil {
		return g, err
	}

	g.File(genny.NewFile(fname, f))

	data := map[string]interface{}{
		"opts": opts,
	}

	if opts.DBType == "postgres" {
		data["testDbUrl"] = "postgres://postgres:postgres@postgres:5432/" + opts.App.Name.File().String() + "_test?sslmode=disable"
	} else if opts.DBType == "mysql" {
		data["testDbUrl"] = "mysql://root:root@(mysql:3306)/" + opts.App.Name.File().String() + "_test?parseTime=true&multiStatements=true&readTimeout=1s"
	} else {
		data["testDbUrl"] = ""
	}

	helpers := template.FuncMap{}

	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)

	return g, nil
}
