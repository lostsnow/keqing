package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"

	"github.com/lostsnow/keqing/data"
	"github.com/lostsnow/keqing/pkg/character"
	"github.com/lostsnow/keqing/pkg/weapon"
)

func main() {
	err := character.Init()
	if err != nil {
		log.Fatalf("i18n generate failed: %s", err)
	}

	cfg := &packages.Config{
		Mode:  packages.NeedImports,
		Tests: false,
	}

	pkgs, err := packages.Load(cfg)
	if err != nil {
		log.Fatalf("i18n generate load package failed: %s", err)
	}

	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}

	pkgID := pkgs[0].ID
	idx := strings.LastIndex(pkgID, "/")
	pkg := pkgID[idx+1:]

	tmpl, err := template.ParseFS(data.TPL, "tpl/i18n_strings.tpl")
	if err != nil {
		log.Fatalf("i18n generate parse template failed: %s", err)
	}

	trans, err := getTranslations()
	if err != nil {
		log.Fatalf("i18n generate get translations failed: %s", err)
	}

	m := struct {
		Package      string
		Translations []string
	}{
		Package:      pkg,
		Translations: trans,
	}

	out := "./i18n_strings.go"

	f, err := os.Create(out)
	if err != nil {
		log.Fatalf("i18n generate create file failed: %s", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, m)
	if err != nil {
		log.Printf("i18n generate execute failed: %s", err)
	} else {
		log.Printf("i18n generate successfully: %s\n", out)
	}
}

func getTranslations() ([]string, error) {
	trans := make([]string, 0)

	err := character.Init()
	if err != nil {
		return nil, fmt.Errorf("get character i18n strings failed: %w", err)
	}

	for _, c := range character.ObjectMap() {
		trans = append(trans, c.ID)
	}

	err = weapon.Init()
	if err != nil {
		return nil, fmt.Errorf("get weapon i18n strings failed: %w", err)
	}

	for _, c := range weapon.ObjectMap() {
		trans = append(trans, c.ID)
	}

	if len(trans) == 0 {
		return nil, fmt.Errorf("i18n strings empty")
	}

	sort.Strings(trans)

	return trans, nil
}
