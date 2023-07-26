
# generate-groups.sh all \
#     github.com/ubombar/namespace-request-controller/pkg/generated \
#     github.com/ubombar/namespace-request-controller/pkg/apis \
#     namespacerequest:v1alpha1 \
#     --output-base "${GOPATH}/src" \
#     --go-header-file hack/boilerplate.go.txt

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${GOPATH}src/github.com/kubernetes/code-generator

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.

"${CODEGEN_PKG}"/generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/ubombar/namespace-request-controller/pkg/generated github.com/ubombar/namespace-request-controller/pkg/apis \
  namespacerequest:v1alpha1 \
  --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt

# CODEGEN_PKG=${GOPATH}src/github.com/kubernetes/code-generator
# SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

# source "${CODEGEN_PKG}/kube_codegen.sh"

# # generate the code with:
# # --output-base    because this script should also be able to run inside the vendor dir of
# #                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
# #                  instead of the $GOPATH directly. For normal projects this can be dropped.

# kube::codegen::gen_helpers \
#     --input-pkg-root github.com/ubombar/namespace-request-controller/pkg/apis \
#     --output-base "$(dirname "${BASH_SOURCE[0]}")/../../.." \
#     --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"

# kube::codegen::gen_client \
#     --with-watch \
#     --input-pkg-root github.com/ubombar/namespace-request-controller/pkg/apis \
#     --output-pkg-root github.com/ubombar/namespace-request-controller/pkg/generated \
#     --output-base "$(dirname "${BASH_SOURCE[0]}")/../../.." \
#     --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"