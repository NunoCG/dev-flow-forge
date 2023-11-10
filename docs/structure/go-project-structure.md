# Golang projects code structure and organization

### Table of contents:

These are all directories that are placed in the root of the project.

1. [/api](#api)
2. [/assets](#assets)
3. [/build](#build)
4. [/cmd](#cmd)
5. [/docs](#docs)
6. [/githooks](#githooks)
7. [/init](#init)
8. [/internal](#internal)
9. [/pkg](#pkg)
10. [/scripts](#scripts)
11. [/test](#test)
12. [/third_party](#third_party)
13. [/tools](#tools)
14. [/vendor](#vendor)
15. [/web](#web)
16. [/website](#test)

<br />

---

- ### /api

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

Examples:

https://github.com/kubernetes/kubernetes/tree/master/api
https://github.com/moby/moby/tree/master/api

---

- ### /assets

Other assets to go along with our repository (images, logos, etc).

---

- ### /build

Packaging and Continuous Integration.

Put cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the /build/package directory.

Put CI (travis, circle, drone) configurations and scripts in the /build/ci directory. Note that some of the CI tools (e.g., Travis CI) are very picky about the location of their config files. Try putting the config files in the /build/ci directory linking them to the location where the CI tools expect them when possible (don't worry if it's not and if keeping those files in the root directory makes your life easier :-)).

Examples:

https://github.com/cockroachdb/cockroach/tree/master/build

---

- ### /cmd

Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., /cmd/myapp).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the `/pkg` directory. If the code is not reusable or if you don't want others to reuse it, put that code in the `/internal` directory. You'll be surprised what others will do, so be explicit about your intentions!

It's common to have a small main function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.

Examples:

https://github.com/vmware-tanzu/velero/tree/main/cmd (just a really small main function with everything else in packages)
https://github.com/moby/moby/tree/master/cmd
https://github.com/prometheus/prometheus/tree/main/cmd
https://github.com/influxdata/influxdb/tree/master/cmd
https://github.com/kubernetes/kubernetes/tree/master/cmd
https://github.com/dapr/dapr/tree/master/cmd
https://github.com/ethereum/go-ethereum/tree/master/cmd

---

- ### /docs

`Design` and `User` documentation (in addition to your godoc generated documentation).

Examples:

https://github.com/gohugoio/hugo/tree/master/docs
https://github.com/openshift/origin/tree/master/docs
https://github.com/dapr/dapr/tree/master/docs

---

- ### /githooks

Git hooks.

---

- ### /init

System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

---

- ### /internal

Private application and library code. This is the code we don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 release notes for more details. Note that we are not limited to the top level internal directory. You can have more than one internal directory at any level of our project tree.

We can optionally add a bit of extra structure to our internal packages to separate our shared and non-shared internal code. It's not required (especially for smaller projects), but it's nice to have visual clues showing the intended package use. Our actual application code can go in the /internal/app directory (e.g., /internal/app/myapp) and the code shared by those apps in the /internal/pkg directory (e.g., /internal/pkg/myprivlib).

Examples:

https://github.com/hashicorp/terraform/tree/main/internal
https://github.com/influxdata/influxdb/tree/master/internal
https://github.com/perkeep/perkeep/tree/master/internal
https://github.com/jaegertracing/jaeger/tree/main/internal
https://github.com/moby/moby/tree/master/internal
https://github.com/satellity/satellity/tree/main/internal

Examples `/internal/pkg` :

https://github.com/hashicorp/waypoint/tree/main/internal/pkg

---

- ### /pkg

Library code that's ok to use by external applications (e.g., /pkg/mypubliclib). Other projects will import these libraries expecting them to work, so think twice before you put something here. :-)

Note that the internal directory is a better way to ensure your private packages are not importable because it's enforced by Go. The `/pkg` directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. The I'll take pkg over internal blog post by Travis Jeffery provides a good overview of the pkg and internal directories and when it might make sense to use them.

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and directories making it easier to run various Go tools (as mentioned in these talks: Best Practices for Industrial Programming from GopherCon EU 2018, GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps and GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go).

Note that this is not a universally accepted pattern and for every popular repo that uses it we can find 10 that don't. It's up to us to decide if we want to use this pattern or not. Regardless of whether or not it's a good pattern more people will know what we mean than not. It might a bit confusing for some of the new Go devs, but it's a pretty simple confusion to resolve and that's one of the goals for this project layout repo.

Ok not to use it if our app project is really small and where an extra level of nesting doesn't add much value (unless we really want to). Think about it when it's getting big enough and our root directory gets pretty busy (especially if we have a lot of non-Go app components).

The `/pkg` directory origins: The old Go source code used to use pkg for its packages and then various Go projects in the community started copying the pattern (see this Brad Fitzpatrick's tweet for more context).

Examples:

https://github.com/jaegertracing/jaeger/tree/master/pkg
https://github.com/istio/istio/tree/master/pkg
https://github.com/GoogleContainerTools/kaniko/tree/master/pkg
https://github.com/google/gvisor/tree/master/pkg
https://github.com/google/syzkaller/tree/master/pkg
https://github.com/perkeep/perkeep/tree/master/pkg
https://github.com/minio/minio/tree/master/pkg
https://github.com/heptio/ark/tree/master/pkg
https://github.com/argoproj/argo/tree/master/pkg
https://github.com/heptio/sonobuoy/tree/master/pkg
https://github.com/helm/helm/tree/master/pkg
https://github.com/kubernetes/kubernetes/tree/master/pkg
https://github.com/kubernetes/kops/tree/master/pkg
https://github.com/moby/moby/tree/master/pkg
https://github.com/grafana/grafana/tree/master/pkg
https://github.com/influxdata/influxdb/tree/master/pkg
https://github.com/cockroachdb/cockroach/tree/master/pkg
https://github.com/derekparker/delve/tree/master/pkg
https://github.com/etcd-io/etcd/tree/master/pkg
https://github.com/oklog/oklog/tree/master/pkg
https://github.com/flynn/flynn/tree/master/pkg
https://github.com/jesseduffield/lazygit/tree/master/pkg
https://github.com/gopasspw/gopass/tree/master/pkg
https://github.com/sosedoff/pgweb/tree/master/pkg
https://github.com/GoogleContainerTools/skaffold/tree/master/pkg
https://github.com/knative/serving/tree/master/pkg
https://github.com/grafana/loki/tree/master/pkg
https://github.com/bloomberg/goldpinger/tree/master/pkg
https://github.com/Ne0nd0g/merlin/tree/master/pkg
https://github.com/jenkins-x/jx/tree/master/pkg
https://github.com/DataDog/datadog-agent/tree/master/pkg
https://github.com/dapr/dapr/tree/master/pkg
https://github.com/cortexproject/cortex/tree/master/pkg
https://github.com/dexidp/dex/tree/master/pkg
https://github.com/pusher/oauth2_proxy/tree/master/pkg
https://github.com/pdfcpu/pdfcpu/tree/master/pkg
https://github.com/weaveworks/kured/tree/master/pkg
https://github.com/weaveworks/footloose/tree/master/pkg
https://github.com/weaveworks/ignite/tree/master/pkg
https://github.com/tmrts/boilr/tree/master/pkg
https://github.com/kata-containers/runtime/tree/master/pkg
https://github.com/okteto/okteto/tree/master/pkg
https://github.com/solo-io/squash/tree/master/pkg

---

- ### /scripts

Scripts to perform various build, install, analysis, etc operations.

These scripts keep the root level Makefile small and simple.

Examples:

https://github.com/kubernetes/helm/tree/master/scripts
https://github.com/cockroachdb/cockroach/tree/master/scripts
https://github.com/hashicorp/terraform/tree/master/scripts

---

- ### /test

Additional external test apps and test data. We can structure the `/test` directory anyway we want. For bigger projects it makes sense to have a data subdirectory. For example, we can have `/test/data` or `/test/testdata` if we need Go to ignore what's in that directory. Note that Go will also ignore directories or files that begin with "." or "\_", so we have more flexibility in terms of how we name our test data directory.

Examples:

https://github.com/openshift/origin/tree/master/test (test data is in the `/testdata` subdirectory)

---

- ### /third_party

External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).

---

- ### /tools

Supporting tools for this project. Note that these tools can import code from the `/pkg` and `/internal` directories.

Examples:

https://github.com/istio/istio/tree/master/tools
https://github.com/openshift/origin/tree/master/tools
https://github.com/dapr/dapr/tree/master/tools

---

- ### /vendor

Application dependencies (managed manually or by our favorite dependency management tool like the recent built-in, modules feature).

Don't commit our application dependencies if we are building a library.

Note that since 1.13 Go also enabled the module proxy feature (using https://proxy.golang.org as their module proxy server by default). Read more about it here to see if it fits all of your requirements and constraints. If it does, then you won't need the 'vendor' directory at all.

---

- ### /web

Web application specific components: static web assets, server side templates and SPAs.

---

- ### /website

This is the place to put our project's website data.

Examples:

https://github.com/hashicorp/vault/tree/master/website
https://github.com/perkeep/perkeep/tree/master/website
