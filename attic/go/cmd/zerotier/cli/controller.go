/*
 * Copyright (c)2013-2021 ZeroTier, Inc.
 *
 * Use of this software is governed by the Business Source License included
 * in the LICENSE.TXT file in the project's root directory.
 *
 * Change Date: 2026-01-01
 *
 * On the date above, in accordance with the Business Source License, use
 * of this software will be governed by version 2.0 of the Apache License.
 */
/****/

package cli

func Controller(basePath string, authTokenGenerator func() string, args []string, jsonOutput bool) int {
	if len(args) < 1 {
		Help()
		return 1
	}

	switch args[0] {
	}

	return 0
}
