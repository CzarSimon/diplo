package main

// saveUser saves a new user if it does not already exist.
func (env *Env) saveUser(userID string) error {
	return env.UserRepo.SaveUser(userID)
}

// removeUser removes a user if it exists.
func (env *Env) removeUser(userID string) error {
	return env.UserRepo.RemoveUser(userID)
}

// checkUserExists checks if a user exists
func (env *Env) checkUserExists(userID string) error {
	_, err := env.UserRepo.FindUser(userID)
	return err
}
