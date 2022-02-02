module github.com/Orange-OpenSource/nifikop

go 1.15

require (
	emperror.dev/errors v0.4.2
	github.com/antihax/optional v1.0.0
	github.com/banzaicloud/k8s-objectmatcher v1.4.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/erdrix/nigoapi v0.0.0-20211026131235-b258e0b217af
	github.com/go-logr/logr v0.4.0
	github.com/imdario/mergo v0.3.12
	github.com/jarcoal/httpmock v1.0.6
	github.com/jetstack/cert-manager v1.6.1
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.15.0
	github.com/pavel-v-chernykh/keystore-go v2.1.0+incompatible
	github.com/stretchr/testify v1.7.0
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	sigs.k8s.io/controller-runtime v0.10.1
)

//replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
// 	github.com/erdrix/nigoapi v0.0.0-20200824133217-ce90b74151a2
