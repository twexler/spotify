package spotify

import (
	"net/url"
	"strings"
)

// Save one or more tracks to the current user’s ‘Your Music’ library.
func (a *API) SaveTracks(ids ...string) error {
	v := url.Values{}
	v.Add("ids", strings.Join(ids, ","))

	return a.put("/me/tracks?"+v.Encode(), nil)
}

// Remove one or more tracks from the current user’s ‘Your Music’ library.
func (a *API) RemoveSavedTracks(ids ...string) error {
	v := url.Values{}
	v.Add("ids", strings.Join(ids, ","))

	return a.delete("/me/tracks?" + v.Encode())
}
