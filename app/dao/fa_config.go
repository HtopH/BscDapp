// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"BscDapp/app/dao/internal"
)

// faConfigDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type faConfigDao struct {
	internal.FaConfigDao
}

var (
	// FaConfig is globally public accessible object for table fa_config operations.
	FaConfig = faConfigDao{
		internal.FaConfig,
	}
)

// Fill with you ideas below.
