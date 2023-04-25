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

package v1beta1

// ConjurProvider configures an store to sync secrets using a Conjur backend.
type ConjurProvider struct {
	// ServiceURL is the URL of the Conjur service.
	ServiceURL string `json:"serviceUrl"`
	// ServiceUser is the host identity that will authenticate to the Conjur service.
	// Conjur host identities always begin with host/
	// For example, host/my-application
	ServiceUser string `json:"serviceUser"`
	// ServiceAPIKey is the API key of the host identity that will authenticate to the Conjur service.
	ServiceAPIKey string `json:"serviceApiKey"`
	// ServiceAccount is the Conjur account name chosen during initial deployment.
	ServiceAccount string `json:"serviceAccount"`
	// ServiceCertificate is the public SSL certificate used to verify the Conjur service.
	// This is only required if the Conjur service is using a self-signed certificate.
	// To retrieve the public SSL certificate from the Conjur service, you may use openssl:
	// openssl s_client -showcerts -connect <conjur-service-url>:443 </dev/null 2>/dev/null|openssl x509 -outform PEM
	ServiceCertificate string `json:"serviceCertificate,omitempty"`
}
