
package authlibrary

type UserProfile struct {
    // Define fields for user profile
}

type UserRepository interface {
    GetProfile(username string) (UserProfile, error)
    // ... other user-related methods
}
