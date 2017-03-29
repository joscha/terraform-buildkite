package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/joscha/terraform-buildkite/buildkite"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: buildkite.Provider,
    })
}
