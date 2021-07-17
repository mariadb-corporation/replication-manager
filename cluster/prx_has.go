// replication-manager - Replication Manager Monitoring and CLI for MariaDB and MySQL
// Copyright 2017 Signal 18 SARL
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <svaroqui@gmail.com>
// This source code is licensed under the GNU General Public License, version 3.
// Redistribution/Reuse of this code is permitted under the GNU v3 license, as
// an additional term, ALL code must carry the original Author(s) credit in comment form.
// See LICENSE in this directory for the integral text.
package cluster

import (
	"os"
)

func (proxy *Proxy) hasCookie(key string) bool {
	if _, err := os.Stat(proxy.Datadir + "/@" + key); os.IsNotExist(err) {
		return false
	}
	return true
}

func (proxy *Proxy) HasProvisionCookie() bool {
	return proxy.hasCookie("cookie_prov")
}

func (proxy *Proxy) HasWaitStartCookie() bool {
	return proxy.hasCookie("cookie_waitstart")
}

func (proxy *Proxy) HasWaitStopCookie() bool {
	return proxy.hasCookie("cookie_waitstop")
}

func (proxy *Proxy) HasRestartCookie() bool {
	return proxy.hasCookie("cookie_restart")
}

func (proxy *Proxy) HasReprovCookie() bool {
	return proxy.hasCookie("cookie_reprov")
}

func (proxy *Proxy) IsRunning() bool {
	return !proxy.IsDown()
}

func (proxy *Proxy) IsDown() bool {
	if proxy.State == stateFailed || proxy.State == stateSuspect || proxy.State == stateErrorAuth {
		return true
	}
	return false
}
