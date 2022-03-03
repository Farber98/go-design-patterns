package proxy

import "errors"

/*
PROXY:
- Wraps an object to hide some of its characteristics (remote object, very heavy object, restricted access object)

OBJECTIVE:
- Hide an object behind the proxy so the features can be hidden , restricted and so on.
- Provide a new abstraction layer that is easy to work with, and can be changed easily

EXAMPLE: Remote proxy going to be a FIFO cache of objects before accessing DB.

ACCEPTANCE CRITERIA:
- All access to the DB os users will be done through the Proxy type.
- A stack of n number of recent users will be kept in the Proxy.
- If a user alread exists in the stack, it won't query the database and will return the stored one.
- If the queried user doesn't exist in the stack, it will query the database, remove the oldest user in the stack if it's full, store the new one, and return it.

*/

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

type UserListProxy struct {
	SomeDatabase           UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	return User{}, errors.New("not implemented yet")
}
