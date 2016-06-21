package api

import (
	"fmt"

	"github.com/juju/errors"
)

// GetSafes returns a list of safe
func (o *OnlineAPI) GetSafes() (safes []OnlineGetSafe, err error) {
	if err = o.getWrapper(fmt.Sprintf("%s/storage/c14/safe", APIUrl), &safes); err != nil {
		err = errors.Annotate(err, "GetSafes")
	}
	return
}

// GetSafe returns a safe
func (o *OnlineAPI) GetSafe(uuid string) (safe OnlineGetSafe, err error) {
	// TODO: enable to use the name instead of only the UUID
	if err = o.getWrapper(fmt.Sprintf("%s/storage/c14/safe/%s", APIUrl, uuid), &safe); err != nil {
		err = errors.Annotate(err, "GetSafe")
	}
	return
}

// GetPlatforms returns a list of platform
func (o *OnlineAPI) GetPlatforms() (platform []OnlineGetPlatform, err error) {
	if err = o.getWrapper(fmt.Sprintf("%s/storage/c14/platform", APIUrl), &platform); err != nil {
		err = errors.Annotate(err, "GetPlatforms")
	}
	return
}

// GetPlatform returns a platform
func (o *OnlineAPI) GetPlatform(uuid string) (platform OnlineGetPlatform, err error) {
	// TODO: enable to use the name instead of only the UUID
	if err = o.getWrapper(fmt.Sprintf("%s/storage/c14/platform/%s", APIUrl, uuid), &platform); err != nil {
		err = errors.Annotate(err, "GetPlatform")
	}
	return
}

func (o *OnlineAPI) CreateSafe(name, desc string) (err error) {
	var (
		buff []byte
	)

	if buff, err = o.postWrapper(fmt.Sprintf("%s/storage/c14/safe", APIUrl), OnlinePostSafe{
		Name:        name,
		Description: desc,
	}); err != nil {
		err = errors.Annotate(err, "GetPlatform")
		return
	}
	fmt.Println(string(buff))
	return
}