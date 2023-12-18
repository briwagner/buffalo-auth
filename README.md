# Forked: Auth Generator for Buffalo

From <a href="https://github.com/gobuffalo/buffalo-auth">Buffalo Auth</a>

**Why is this forked?**

This repo includes <a href="https://github.com/gobuffalo/buffalo-auth/pull/43">some proposed changes</a> for account recovery, which never made it into the main branch. Also I modified some of the default routes to match common cases that I've seen in the world, and that I want to re-use on projects.

## What are the big changes?

**Account recovery**: this version includes an account recovery mechanism, which sends a code via email. There is a form to request the recovery code, and a separate form to login with the code.

**User model fields**: the User model includes fields for the recovery code.

**Default settings**: the default configuration provides NO CHECKS on authentication for routes. There are comments in the routing file about how to enable them. (I find it confusing that most routes, including homepage, were protected by default.)

## Installation

```console
$ buffalo plugins install github.com/gobuffalo/buffalo-auth
```

**Email**: email delivery is required for the plugin to work properly. The default configuration includes a skeleton `sender` service which should be customized to work with your email system.

## Usage

To generate a basic username / password authentication you can run:

```console
$ buffalo generate auth
```

This will do:

- Generate User authentication actions in `actions/auth.go`:
  - AuthNew
  - AuthCreate
  - AuthDestroy

- Generate User signup actions in `actions/users.go`:
  - UsersNew
  - UsersCreate

- Generate User model and migration ( model will be in `models/user.go`):

- Generate Auth Middlewares
  - SetCurrentUser
  - Authorize

- Add actions and middlewares in `app.go`:
  - [GET] /users/new -> UsersNew
  - [POST] /users -> UsersCreate
  - [GET] /signin -> AuthNew
  - [POST] /signin -> AuthCreate
  - [DELETE] /signout -> AuthDestroy

- Use middlewares for all your actions and skip
  - HomeHandler
  - UsersNew
  - UsersCreate
  - AuthNew
  - AuthCreate

### User model Fields

Sometimes you would want to add extra fields to the user model, to do so, you can pass those to the auth command and use the pop notation for those fields, for example:

```console
$ buffalo generate auth first_name last_name notes:text
```

Will generate a User model (inside `models/user.go`) that looks like:

```go
type User struct {
  ID                   uuid.UUID `json:"id" db:"id"`
  CreatedAt            time.Time `json:"created_at" db:"created_at"`
  UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
  Email                string    `json:"email" db:"email"`
  PasswordHash         string    `json:"password_hash" db:"password_hash"`
  Password             string    `json:"-" db:"-"`
  PasswordConfirmation string    `json:"-" db:"-"`
  RecoveryCode         nulls.String `json:"-" db:"recovery_code"`
  RecoveryExp          nulls.Time   `json:"-" db:"recovery_expiration"`

  FirstName            string    `json:"first_name" db:"first_name"`
  LastName             string    `json:"last_name" db:"last_name"`
  Notes                string    `json:"notes" db:"notes"`
}
```
