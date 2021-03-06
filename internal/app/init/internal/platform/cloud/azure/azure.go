/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package azure

import (
	"github.com/talos-systems/talos/pkg/userdata"
)

const (
	// AzureUserDataEndpoint is the local endpoint for the user data.
	// By specifying format=text and drilling down to the actual key we care about
	// we get a base64 encoded userdata response
	AzureUserDataEndpoint = "http://169.254.169.254/metadata/instance/compute/customData?api-version=2019-06-01&format=text"
	// AzureHostnameEndpoint is the local endpoint for the hostname.
	AzureHostnameEndpoint = "http://169.254.169.254/metadata/instance/compute/name?api-version=2019-06-01&format=text"
	// AzureInternalEndpoint is the Azure Internal Channel IP
	// https://blogs.msdn.microsoft.com/mast/2015/05/18/what-is-the-ip-address-168-63-129-16/
	AzureInternalEndpoint = "http://168.63.129.16"
)

// Azure is the concrete type that implements the platform.Platform interface.
type Azure struct{}

// Name implements the platform.Platform interface.
func (a *Azure) Name() string {
	return "Azure"
}

// UserData implements the platform.Platform interface.
func (a *Azure) UserData() (*userdata.UserData, error) {
	if err := linuxAgent(); err != nil {
		return nil, err
	}

	return userdata.Download(AzureUserDataEndpoint, userdata.WithHeaders(map[string]string{"Metadata": "true"}), userdata.WithFormat("base64"))
}

// Prepare implements the platform.Platform interface and handles initial host preparation.
func (a *Azure) Prepare(data *userdata.UserData) (err error) {
	return nil
}

func hostname() (err error) {

	// TODO get this sorted; assuming we need to set appropriate headers
	return err

	/*
		resp, err := http.Get(AzureHostnameEndpoint)
		if err != nil {
			return
		}
		// nolint: errcheck
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("download user data: %d", resp.StatusCode)
		}

		dataBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		if err = unix.Sethostname(dataBytes); err != nil {
			return
		}

		return nil
	*/
}

// Install implements the platform.Platform interface and handles additional system setup.
func (a *Azure) Install(data *userdata.UserData) (err error) {
	return hostname()
}
