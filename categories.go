// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/elastic/package-registry/util"
)

type Category struct {
	Id    util.CategoryID `yaml:"id" json:"id"`
	Title string          `yaml:"title" json:"title"`
	Count int             `yaml:"count" json:"count"`
}

// categoriesHandler is a dynamic handler as it will also allow filtering in the future.
func categoriesHandler(packagesBasePath, cacheTime string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cacheHeaders(w, cacheTime)

		packages, err := util.GetPackages(packagesBasePath)
		if err != nil {
			notFound(w, err)
			return
		}
		packageList := map[string]util.Package{}
		// Get unique list of newest packages
		for _, p := range packages {

			// Skip internal packages
			if p.Internal {
				continue
			}

			// Check if the version exists and if it should be added or not.
			if pp, ok := packageList[p.Name]; ok {
				// If the package in the list is newer, do nothing. Otherwise delete and later add the new one.
				if pp.IsNewer(p) {
					continue
				}
			}
			packageList[p.Name] = p
		}

		categories := map[util.CategoryID]*Category{}

		for _, p := range packageList {
			for _, c := range p.Categories {
				if _, ok := categories[c]; !ok {
					categories[c] = &Category{
						Id:    c,
						Title: util.CategoryTitles[c],
						Count: 0,
					}
				}

				categories[c] = &Category{
					Id:    c,
					Title: util.CategoryTitles[c],
					Count: categories[c].Count + 1,
				}

			}
		}

		data, err := getCategoriesOutput(categories)
		if err != nil {
			notFound(w, err)
			return
		}
		jsonHeader(w)
		fmt.Fprint(w, string(data))
	}
}

func getCategoriesOutput(categories map[util.CategoryID]*Category) ([]byte, error) {
	var outputCategories []*Category
	for k := range categories {
		outputCategories = append(outputCategories, categories[k])
	}
	sort.SliceStable(outputCategories, func(i, j int) bool { return outputCategories[i].Title < outputCategories[j].Title })

	return json.MarshalIndent(outputCategories, "", "  ")
}
