/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package create

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"sigs.k8s.io/yaml"
)

// chartName is a regular expression for testing the supplied name of a chart.
// This regular expression is probably stricter than it needs to be. We can relax it
// somewhat. Newline characters, as well as $, quotes, +, parens, and % are known to be
// problematic.
var chartName = regexp.MustCompile("^[a-zA-Z0-9._-]+$")

const (
	// ChartfileName is the default Chart file name.
	ChartfileName = "Chart.yaml"
	// ValuesfileName is the default values file name.
	ValuesfileName = "values.yaml"
	// SchemafileName is the default values schema file name.
	SchemafileName = "values.schema.json"
	// TemplatesDir is the relative directory name for templates.
	TemplatesDir = "templates"
	// ChartsDir is the relative directory name for charts dependencies.
	ChartsDir = "charts"
	// TemplatesTestsDir is the relative directory name for tests.
	TemplatesTestsDir = TemplatesDir + sep + "tests"
	// IgnorefileName is the name of the Helm ignore file.
	IgnorefileName = ".helmignore"
	// IngressFileName is the name of the example ingress file.
	IngressFileName = TemplatesDir + sep + "ingress.yaml"
	// DeploymentName is the name of the example deployment file.
	DeploymentName = TemplatesDir + sep + "deployment.yaml"
	// ServiceName is the name of the example service file.
	ServiceName = TemplatesDir + sep + "service.yaml"
	// ServiceAccountName is the name of the example serviceaccount file.
	ServiceAccountName = TemplatesDir + sep + "serviceaccount.yaml"
	// HorizontalPodAutoscalerName is the name of the example hpa file.
	HorizontalPodAutoscalerName = TemplatesDir + sep + "hpa.yaml"
	// NotesName is the name of the example NOTES.txt file.
	NotesName = TemplatesDir + sep + "NOTES.txt"
	// HelpersName is the name of the example helpers file.
	HelpersName = TemplatesDir + sep + "_helpers.tpl"
	// TestConnectionName is the name of the example test file.
	TestConnectionName = TemplatesTestsDir + sep + "test-connection.yaml"
)

// maxChartNameLength is lower than the limits we know of with certain file systems,
// and with certain Kubernetes fields.
const maxChartNameLength = 250

const sep = string(filepath.Separator)

// Stderr is an io.Writer to which error messages can be written
//
// In Helm 4, this will be replaced. It is needed in Helm 3 to preserve API backward
// compatibility.
var Stderr io.Writer = os.Stderr

type Chart struct {
	Name string `yaml:"name"`
}

func GetChartName(chart string) (string, error) {
	// Get name from Chart.yaml
	var c Chart
	chartyaml := filepath.Join(chart, ChartfileName)
	yamlFile, err := ioutil.ReadFile(chartyaml)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c.Name, err
}

func Create(resource, chart, filename string)	(string, error)  {
	resourcemap := map[string]string{
		"ingress": defaultIngress,
		"deployment": defaultDeployment,
		"service": defaultService,
		"hpa": defaultHorizontalPodAutoscaler,
		"sa": defaultServiceAccount,
		"secret": defaultSecret,
	}

	chartname, err := GetChartName(chart)
	if err != nil {
		log.Fatalf("Error retriving Chart name from file, %s", err)
	}

	content := transform(resourcemap[resource], chartname)
	if len(content) == 0 {
		log.Fatalf("resource of type %s not supported", resource)
	}

	name := resource
	if filename != "" {
		name = filename
	}
	filepath := filepath.Join(chart, "/templates", name + "-" + resource + ".yaml")

	if _, err := os.Stat(filepath); err == nil {
		// There is no handle to a preferred output stream here.
		fmt.Fprintf(Stderr, "WARNING: File %q already exists. Overwriting.\n", filepath)
	}
	if err := writeFile(filepath, content); err != nil {
		return filepath, err
	}
	fmt.Printf("Successfully created %s at %s", resource, filepath)
	return "", nil
}

// transform performs a string replacement of the specified source for
// a given key with the replacement string
func transform(src, replacement string) []byte {
	return []byte(strings.ReplaceAll(src, "<CHARTNAME>", replacement))
}

func writeFile(name string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(name, content, 0644)
}