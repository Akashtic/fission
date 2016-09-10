/*
Copyright 2016 The Fission Authors.

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

package fission

type (
	// Metadata is used as the general identifier for all kinds of
	// resources managed by the controller.  In general, when a
	// resource is updated, the Name remains the same, but the UID
	// changes.  In other words, the UID identifies a particular
	// value of a resource.
	Metadata struct {
		Name string
		Uid  string
	}

	// Function is a unit of executable code.  Though it's called
	// a function, the code may have more than one function; it's
	// usually some sort of module or package.
	Function struct {
		Metadata
		Environment Metadata
		Code        string
	}

	// Environment identifies the language and OS specific
	// resources that a function depends on.  For now this
	// includes only the function run container image.  Later,
	// this will also include build containers, as well as support
	// tools like debuggers, profilers, etc.
	Environment struct {
		Metadata
		RunContainerImageUrl string
	}

	// HTTPTrigger maps URL patterns to functions.  Function.UID
	// is optional; if absent, the latest version of the function
	// will automatically be selected.
	HTTPTrigger struct {
		Metadata
		UrlPattern string
		Function   Metadata
	}
)
