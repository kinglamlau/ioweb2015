// +build !appengine

// TODO: implement using Bolt: https://github.com/boltdb/bolt

package main

import (
	"errors"
	"time"

	"golang.org/x/net/context"
)

// RunInTransaction runs f in a transaction.
// It calls f with a transaction context tc that f should use for all operations.
func RunInTransaction(c context.Context, f func(tc context.Context) error) error {
	// TODO: implement transaction support
	return f(c)
}

// storeCredentials saves OAuth2 credentials cred in a presistent DB.
// cred must have userID set to a non-zero value.
func storeCredentials(c context.Context, cred *oauth2Credentials) error {
	// TODO: implement
	return nil
}

// getCredentials fetches credentials from a persistent DB.
// A user must be present in the context.
func getCredentials(c context.Context) (*oauth2Credentials, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

// storeUserPushInfo saves user push configuration in a persistent DB.
// info must have userID set to a non-zero value.
func storeUserPushInfo(c context.Context, p *userPush) error {
	// TODO: implement
	return nil
}

// getUserPushInfo fetches user push configuration from a persistent DB.
// If the configuration does not exist yet, a default one is returned.
// Default configuration has all notification settings disabled.
// A user must be present in the context.
func getUserPushInfo(c context.Context) (*userPush, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

func storeEventData(c context.Context, d *eventData) error {
	return errors.New("not implemented")
}

func getLatestEventData(c context.Context) (*eventData, error) {
	return nil, errors.New("not implemented")
}

func storeChanges(c context.Context, d *dataChanges) error {
	return errors.New("not implemented")
}

// getChangesSince queries DB for all changes occurred since time t
// and returns them all combined in one dataChanges result.
// In a case where multiple changes have been introduced in the same data items,
// older changes will be overwritten by the most recent ones.
// At most 1000 changes will be returned.
// Resulting dataChanges.Changed time will be set to the most recent one.
func getChangesSince(c context.Context, t time.Time) (*dataChanges, error) {
	return nil, errors.New("not implemented")
}
