package main

/*
CloudSession ... Stores a Beetroot Cloud session.
*/
type CloudSession struct {
	sessionID string
	closed    bool

	//
	autoQueue bool
}

/*
NewCloudSession ... Creates a new session on a Beetroot Cloud server.
*/
func NewCloudSession(server string) (*CloudSession, error) {
	session := new(CloudSession)
	session.closed = false

	return session, nil
}

/*
UploadLibrary ... Uploads the library to the
func (session *CloudSession) UploadLibrary(library *[]string) {

}

/*
Close ... Ends the session on the server.
*/
func (session *CloudSession) Close() {
	session.closed = true

	// TODO: send a request to close the client and destroy HTTP client
}
