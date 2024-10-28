/*
 * Copyright (c) 2024 flowerinsnow
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package object

type IndexPageConfig struct {
	Title         string
	StaticDomain  string
	WWWDomain     string
	BlogURL       string
	ICPNumber     string
	NISMSPNumber  string
}

type IndexPageVariables struct {
	IndexPageConfig
	Status414     string
}
