package cmds

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/FuturFusion/migration-manager/internal/util"
	"github.com/FuturFusion/migration-manager/shared/api"
)

var supportedTargetTypes = []string{"incus"}

type CmdTarget struct {
	Global *CmdGlobal
}

func (c *CmdTarget) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "target"
	cmd.Short = "Interact with migration targets"
	cmd.Long = `Description:
  Interact with migration targets

  Configure migration targets for use by the migration manager.
`

	// Add
	targetAddCmd := cmdTargetAdd{global: c.Global}
	cmd.AddCommand(targetAddCmd.Command())

	// List
	targetListCmd := cmdTargetList{global: c.Global}
	cmd.AddCommand(targetListCmd.Command())

	// Remove
	targetRemoveCmd := cmdTargetRemove{global: c.Global}
	cmd.AddCommand(targetRemoveCmd.Command())

	// Update
	targetUpdateCmd := cmdTargetUpdate{global: c.Global}
	cmd.AddCommand(targetUpdateCmd.Command())

	// Workaround for subcommand usage errors. See: https://github.com/spf13/cobra/issues/706
	cmd.Args = cobra.NoArgs
	cmd.Run = func(cmd *cobra.Command, args []string) { _ = cmd.Usage() }

	return cmd
}

// Add the target.
type cmdTargetAdd struct {
	global *CmdGlobal

	flagTrustedServerCertificateFingerprint string
}

func (c *cmdTargetAdd) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "add [type] <name> <IP|FQDN|URL>"
	cmd.Short = "Add a new target"
	cmd.Long = `Description:
  Add a new target

  Adds a new target for the migration manager to use. The "type" argument is optional,
  and defaults to "incus" if not specified.

  Depending on the target type, you may be prompted for additional information required
  to connect to the target.
`

	cmd.RunE = c.Run
	cmd.Flags().StringVar(&c.flagTrustedServerCertificateFingerprint, "trusted-cert-fingerprint", "", "Trusted SHA256 fingerprint of the target's TLS certificate")

	return cmd
}

func (c *cmdTargetAdd) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 2, 3)
	if exit {
		return err
	}

	targetType := "incus"
	targetName := ""
	targetEndpoint := ""

	// Set variables.
	if len(args) == 3 {
		if !slices.Contains(supportedTargetTypes, strings.ToLower(args[0])) {
			return fmt.Errorf("Unsupported target type %q; must be one of %q", args[0], supportedTargetTypes)
		}

		targetType = strings.ToLower(args[0])
		targetName = args[1]
		targetEndpoint = args[2]
	} else {
		targetName = args[0]
		targetEndpoint = args[1]
	}

	// Add the target.
	switch targetType {
	case "incus":
		incusProperties := api.IncusProperties{
			Endpoint:                            targetEndpoint,
			TrustedServerCertificateFingerprint: c.flagTrustedServerCertificateFingerprint,
		}

		t := api.Target{
			TargetPut: api.TargetPut{
				Name: targetName,
			},
			TargetType: api.TARGETTYPE_INCUS,
		}

		var importLimit int64 = 50
		importLimit, err = c.global.Asker.AskInt(fmt.Sprintf("How many instances can be concurrently imported? [default=%d]: ", importLimit), 0, 1024, strconv.Itoa(int(importLimit)), nil)
		if err != nil {
			return err
		}

		incusProperties.ImportLimit = int(importLimit)

		var createLimit int64 = 10
		createLimit, err = c.global.Asker.AskInt(fmt.Sprintf("How many instances can be concurrently created? [default=%d]: ", createLimit), 0, 1024, strconv.Itoa(int(createLimit)), nil)
		if err != nil {
			return err
		}

		incusProperties.CreateLimit = int(createLimit)

		authType, err := c.global.Asker.AskChoice("Use OIDC or TLS certificates to authenticate to target? [default=oidc]: ", []string{"oidc", "tls"}, "oidc")
		if err != nil {
			return err
		}

		// Only TLS certs require additional prompting at the moment; we'll grab OIDC tokens below when we verify the target.
		if authType == "tls" {
			tlsCertPath, err := c.global.Asker.AskString("Please enter the absolute path to client TLS certificate: ", "", validateAbsFilePathExists)
			if err != nil {
				return err
			}

			contents, err := os.ReadFile(tlsCertPath)
			if err != nil {
				return err
			}

			incusProperties.TLSClientCert = string(contents)

			tlsKeyPath, err := c.global.Asker.AskString("Please enter the absolute path to client TLS key: ", "", validateAbsFilePathExists)
			if err != nil {
				return err
			}

			contents, err = os.ReadFile(tlsKeyPath)
			if err != nil {
				return err
			}

			incusProperties.TLSClientKey = string(contents)
		}

		t.Properties, err = json.Marshal(incusProperties)
		if err != nil {
			return err
		}

		// Insert into database.
		content, err := json.Marshal(t)
		if err != nil {
			return err
		}

		resp, err := c.global.doHTTPRequestV1("/targets", http.MethodPost, "", content)
		if err != nil {
			return err
		}

		metadata := make(map[string]string)
		err = json.Unmarshal(resp.Metadata, &metadata)
		if err != nil {
			return err
		}

		connectivityStatus := api.ExternalConnectivityStatus(metadata["ConnectivityStatus"])

		if connectivityStatus == api.EXTERNALCONNECTIVITYSTATUS_TLS_CONFIRM_FINGERPRINT {
			return fmt.Errorf("Successfully added new target %q, but received an untrusted TLS server certificate with fingerprint %s. Please update the target to correct the issue.", t.Name, metadata["certFingerprint"])
		} else if connectivityStatus == api.EXTERNALCONNECTIVITYSTATUS_WAITING_OIDC {
			return fmt.Errorf("Successfully added new target %q; please visit %s to complete OIDC authorization.", t.Name, metadata["OIDCURL"])
		} else if connectivityStatus != api.EXTERNALCONNECTIVITYSTATUS_OK {
			return fmt.Errorf("Successfully added new target %q, but connectivity check reported an issue: %s. Please update the target to correct the issue.", t.Name, connectivityStatus)
		}

		cmd.Printf("Successfully added new target %q.\n", t.Name)
	}

	return nil
}

// List the targets.
type cmdTargetList struct {
	global *CmdGlobal

	flagFormat string
}

func (c *cmdTargetList) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "list"
	cmd.Short = "List available targets"
	cmd.Long = `Description:
  List the available targets
`

	cmd.RunE = c.Run
	cmd.Flags().StringVarP(&c.flagFormat, "format", "f", "table", `Format (csv|json|table|yaml|compact), use suffix ",noheader" to disable headers and ",header" to enable if demanded, e.g. csv,header`)
	cmd.PreRunE = func(cmd *cobra.Command, _ []string) error {
		return validateFlagFormat(cmd.Flag("format").Value.String())
	}

	return cmd
}

func (c *cmdTargetList) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 0, 0)
	if exit {
		return err
	}

	// Get the list of all targets.
	resp, err := c.global.doHTTPRequestV1("/targets", http.MethodGet, "recursion=1", nil)
	if err != nil {
		return err
	}

	targets := []api.Target{}

	err = responseToStruct(resp, &targets)
	if err != nil {
		return err
	}

	// Render the table.
	header := []string{"Name", "Type", "Endpoint", "Connectivity Status", "Auth Type", "Trusted TLS Cert SHA256 Fingerprint"}
	data := [][]string{}

	for _, t := range targets {
		switch t.TargetType {
		case api.TARGETTYPE_INCUS:
			incusProperties := api.IncusProperties{}
			err := json.Unmarshal(t.Properties, &incusProperties)
			if err != nil {
				return err
			}

			authType := "OIDC"
			if incusProperties.TLSClientKey != "" {
				authType = "TLS"
			}

			data = append(data, []string{t.Name, string(t.TargetType), incusProperties.Endpoint, string(incusProperties.ConnectivityStatus), authType, incusProperties.TrustedServerCertificateFingerprint})
		default:
			return fmt.Errorf("Unsupported target type %s", t.TargetType)
		}
	}

	sort.Sort(util.SortColumnsNaturally(data))

	return util.RenderTable(cmd.OutOrStdout(), c.flagFormat, header, data, targets)
}

// Remove the target.
type cmdTargetRemove struct {
	global *CmdGlobal
}

func (c *cmdTargetRemove) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "remove <name>"
	cmd.Short = "Remove target"
	cmd.Long = `Description:
  Remove target
`

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdTargetRemove) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 1, 1)
	if exit {
		return err
	}

	name := args[0]

	// Remove the target.
	_, err = c.global.doHTTPRequestV1("/targets/"+name, http.MethodDelete, "", nil)
	if err != nil {
		return err
	}

	cmd.Printf("Successfully removed target %q.\n", name)
	return nil
}

// Update the target.
type cmdTargetUpdate struct {
	global *CmdGlobal
}

func (c *cmdTargetUpdate) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "update <name>"
	cmd.Short = "Update target"
	cmd.Long = `Description:
  Update target
`

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdTargetUpdate) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 1, 1)
	if exit {
		return err
	}

	name := args[0]

	// Get the existing target.
	resp, err := c.global.doHTTPRequestV1("/targets/"+name, http.MethodGet, "", nil)
	if err != nil {
		return err
	}

	tgt := api.Target{}

	err = responseToStruct(resp, &tgt)
	if err != nil {
		return err
	}

	// Prompt for updates, depending on the target type.
	origTargetName := ""
	newTargetName := ""
	switch tgt.TargetType {
	case api.TARGETTYPE_INCUS:
		incusProperties := api.IncusProperties{}
		err := json.Unmarshal(tgt.Properties, &incusProperties)
		if err != nil {
			return err
		}

		origTargetName = tgt.Name

		tgt.Name, err = c.global.Asker.AskString("Target name [default="+tgt.Name+"]: ", tgt.Name, nil)
		if err != nil {
			return err
		}

		incusProperties.Endpoint, err = c.global.Asker.AskString("Endpoint [default="+incusProperties.Endpoint+"]: ", incusProperties.Endpoint, nil)
		if err != nil {
			return err
		}

		importLimit := int64(incusProperties.ImportLimit)
		importLimit, err = c.global.Asker.AskInt(fmt.Sprintf("How many instances can be concurrently imported? [default=%d]: ", importLimit), 0, 1024, strconv.Itoa(int(importLimit)), nil)
		if err != nil {
			return err
		}

		incusProperties.ImportLimit = int(importLimit)

		createLimit := int64(incusProperties.CreateLimit)
		createLimit, err = c.global.Asker.AskInt(fmt.Sprintf("How many instances can be concurrently created? [default=%d]: ", createLimit), 0, 1024, strconv.Itoa(int(createLimit)), nil)
		if err != nil {
			return err
		}

		incusProperties.CreateLimit = int(createLimit)

		updateAuth, err := c.global.Asker.AskBool("Update configured authentication? (yes/no) [default=no]: ", "no")
		if err != nil {
			return err
		}

		if updateAuth {
			// Clear out existing auth.
			incusProperties.TLSClientKey = ""
			incusProperties.TLSClientCert = ""
			incusProperties.OIDCTokens = nil

			authType, err := c.global.Asker.AskChoice("Use OIDC or TLS certificates to authenticate to target? [default=oidc]: ", []string{"oidc", "tls"}, "oidc")
			if err != nil {
				return err
			}

			// Only TLS certs require additional prompting at the moment; we'll grab OIDC tokens below when we verify the target.
			if authType == "tls" {
				tlsCertPath, err := c.global.Asker.AskString("Please enter the absolute path to client TLS certificate: ", "", validateAbsFilePathExists)
				if err != nil {
					return err
				}

				contents, err := os.ReadFile(tlsCertPath)
				if err != nil {
					return err
				}

				incusProperties.TLSClientCert = string(contents)

				tlsKeyPath, err := c.global.Asker.AskString("Please enter the absolute path to client TLS key: ", "", validateAbsFilePathExists)
				if err != nil {
					return err
				}

				contents, err = os.ReadFile(tlsKeyPath)
				if err != nil {
					return err
				}

				incusProperties.TLSClientKey = string(contents)
			}
		}

		incusProperties.TrustedServerCertificateFingerprint, err = c.global.Asker.AskString("Manually-set trusted TLS cert SHA256 fingerprint ["+incusProperties.TrustedServerCertificateFingerprint+"]: ", incusProperties.TrustedServerCertificateFingerprint, validateSHA256Format)
		if err != nil {
			return err
		}

		tgt.Properties, err = json.Marshal(incusProperties)
		if err != nil {
			return err
		}

		newTargetName = tgt.Name

	default:
		return fmt.Errorf("Unsupported target type %s; must be one of %q", tgt.TargetType, supportedTargetTypes)
	}

	// Update the target.
	content, err := json.Marshal(tgt.TargetPut)
	if err != nil {
		return err
	}

	resp, err = c.global.doHTTPRequestV1("/targets/"+origTargetName, http.MethodPut, "", content)
	if err != nil {
		return err
	}

	metadata := make(map[string]string)
	err = json.Unmarshal(resp.Metadata, &metadata)
	if err != nil {
		return err
	}

	connectivityStatus := api.ExternalConnectivityStatus(metadata["ConnectivityStatus"])

	if connectivityStatus == api.EXTERNALCONNECTIVITYSTATUS_TLS_CONFIRM_FINGERPRINT {
		return fmt.Errorf("Successfully updated target %q, but received an untrusted TLS server certificate with fingerprint %s. Please update the target to correct the issue.", newTargetName, metadata["certFingerprint"])
	} else if connectivityStatus == api.EXTERNALCONNECTIVITYSTATUS_WAITING_OIDC {
		return fmt.Errorf("Successfully updated target %q; please visit %s to complete OIDC authorization.", newTargetName, metadata["OIDCURL"])
	} else if connectivityStatus != api.EXTERNALCONNECTIVITYSTATUS_OK {
		return fmt.Errorf("Successfully updated target %q, but connectivity check reported an issue: %s. Please update the target to correct the issue.", newTargetName, connectivityStatus)
	}

	cmd.Printf("Successfully updated target %q.\n", newTargetName)

	return nil
}
