// Copyright 2020 Fairwinds
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License

package targetprocess

// DateTime currently does nothing special but will represent the TP DateTime objects
type DateTime string

// CustomField are user defined in TargetProcess and are arbitrarily set for many different
// resource types
type CustomField struct {
	ID    int32       `json:"Id,omitempty"`
	Type  string      `json:",omitempty"`
	Name  string      `json:",omitempty"`
	Value interface{} `json:",omitempty"`
}

// EntityState contains metadata for the state of an Entity. Collection of EntityStates
// form Workflow for Entity. For example, Bug has four EntityStates by default: Open, Fixed, Invalid and Done
type EntityState struct {
	ID              int32   `json:"Id,omitempty"`
	Name            string  `json:",omitempty"`
	NumericPriority float64 `json:",omitempty"`
}

// AssignedTeams is used in UserStories and potentially other places.
// The AssignedTeams section of a UserStory response replaces the deprecated usage of the Team field
type AssignedTeams struct {
	Items []TeamAssignment
}

// TeamAssignment has it's own unique Id and also includes a reference to the team, which also has an Id
type TeamAssignment struct {
	ID        int32    `json:"Id,omitempty"`
	StartDate DateTime `json:",omitempty"`
	EndDate   DateTime `json:",omitempty"`
	Team      *Team    `json:",omitempty"`
}