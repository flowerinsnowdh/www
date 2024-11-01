/*
 * Copyright (c) 2024 flowerinsnow
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package config

var Conf *Config

type Config struct {
	Bind        string       `toml:"bind"`
	MySQLConfig *MySQLConfig `toml:"mysql"`
	SiteConfig  *SiteConfig  `toml:"site"`
}

type MySQLConfig struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Schema   string `toml:"schema"`
}

type SiteConfig struct {
	Title        string       `toml:"title"`
	StaticDomain string       `toml:"static_domain"`
	WWWDomain    string       `toml:"www_domain"`
	BlogURL      string       `toml:"blog_url"`
	ICPNumber    string       `toml:"icp_number"`
	NISMSPNumber string       `toml:"nismsp_number"`
}
