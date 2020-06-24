/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:validation:Enum=WebFolder;S3
type Type string

const (
	WebFolder Type = "WebFolder"
	S3        Type = "S3"
)

type WebFolderMeta struct {
	BaseURL string `json:"baseUrl,omitempty"`
}

type S3Meta struct {
	AccessKey  string `json:"accessKey,omitempty"`
	SecretKey  string `json:"secretKey,omitempty"`
	Region     string `json:"region,omitempty"`
	PathPrefix string `json:"pathPrefix,omitempty"`
}

type Source struct {
	// Specifies storage backend to use with darkroom.
	// Valid values are:
	// - "WebFolder" (default): simple storage backend to serve images from a hosted image source;
	// - "S3": storage backend to serve images from an S3 bucket;
	Type Type `json:"type"`

	// +optional
	WebFolderMeta `json:",inline"`

	// +optional
	S3Meta `json:",inline"`
}

// DarkroomSpec defines the desired state of Darkroom
type DarkroomSpec struct {
	// +optional
	Version string `json:"version"`

	Source Source `json:"source"`
	// +kubebuilder:validation:MinItems=1
	SubDomains []string `json:"subDomains"`
}

// DarkroomStatus defines the observed state of Darkroom
type DarkroomStatus struct {
	// +optional
	Domains []string `json:"domains,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.version",name=VERSION,type=string
// +kubebuilder:printcolumn:JSONPath=".spec.source.type",name=TYPE,type=string
// +kubebuilder:resource:shortName=dr

// Darkroom is the Schema for the darkrooms API
type Darkroom struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DarkroomSpec   `json:"spec,omitempty"`
	Status DarkroomStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DarkroomList contains a list of Darkroom
type DarkroomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Darkroom `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Darkroom{}, &DarkroomList{})
}
