package data

import "time"

type User struct {
	ID        int       `db:"id,omitempty"`
	FirstName string    `db:"first_name"`
	Email     string    `db:"email"`
	LastName  string    `db:"last_name"`
	Active    int       `db:"user_active"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Token     Token     `db:"-"`
}
