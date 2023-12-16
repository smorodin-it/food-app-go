package constants

import "time"

const (
	// Auth constants
	AuthSignKey              = "3be9067a-0869-4291-ae46-1e943f05642a"
	AuthTokenDuration        = time.Minute * 5
	AuthRefreshTokenDuration = time.Hour * 24 * 30
	AuthRefreshTokenField    = "refreshToken"
	AuthAccessTokenField     = "token"

	// Tables and fields
	//TableUsers          = "users"
	//TKUsersId           = "id"
	//TKUsersEmail        = "email"
	//TKUsersRefreshToken = "refresh_token"
	//TKUsersPasswordHash = "password_hash"

	// Global table keys
	//TKCreatedAt = "created_at"
	//TKUpdatedAt = "updated_at"

	//	Token fields

	// Validation
	//MinLengthString   = 2
	//MinLengthPassword = 8
	//MaxLengthPassword = 20
	//MaxLengthEmail    = 50
)
