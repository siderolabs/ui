OPERATING_SYSTEM := $(shell uname -s | tr "[:upper:]" "[:lower:]")

all:
	wails build -p -x linux/amd64
	wails build -p -x darwin/amd64
	wails build -p -x windows/amd64

build:
	wails build

debug:
	wails build -d

package:
	wails build -p

run: debug
	./build/ui

./bin/kustomize:
	mkdir ./bin
	curl -L https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv3.5.5/kustomize_v3.5.5_$(OPERATING_SYSTEM)_amd64.tar.gz | tar -xz -C ./bin kustomize
	chmod +x $@

./bin/clusterctl:
	curl -L -o ./bin/clusterctl https://github.com/kubernetes-sigs/cluster-api/releases/download/v0.3.6/clusterctl-$(OPERATING_SYSTEM)-amd64
	chmod +x $@

./bin/kubeconfig:
	talosctl --context arges kubeconfig ./bin

arges: ./bin/kustomize ./bin/clusterctl ./bin/kubeconfig
	#./bin/kustomize build github.com/talos-systems/arges//examples/bootstrap?ref=master | kubectl --kubeconfig ./bin/kubeconfig apply -f -
	kubectl --kubeconfig ./bin/kubeconfig apply -f ./hack/arges.yaml
	./bin/clusterctl --kubeconfig ./bin/kubeconfig --config ./hack/clusterctl.yaml init --control-plane "-" --bootstrap "talos" --infrastructure "metal" --target-namespace arges-system
	kubectl --kubeconfig ./bin/kubeconfig apply -f ./hack/environment.yaml

setup:
	talosctl cluster create --name=arges --workers=0 -p 69:69/udp,8081:8081/tcp,9091:9091/tcp --endpoint=$(ENDPOINT)
	$(MAKE) arges

clean:
	rm -rf ./bin ./build
	talosctl cluster destroy --name=arges
