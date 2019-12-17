package data

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (user *User) Create() error {
	stmt, err := Db.Prepare("insert into users (uuid, name, email, password, created_at) values(?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	uuid, err := createUUID()
	if err != nil {
		return err
	}

	now := time.Now()
	result, err := stmt.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), now)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	//return stmt.QueryRow(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(
	// &user.Id, &user.Uuid, &user.CreatedAt
	// )
	if err != nil {
		return err
	}
	user.Id = int(id)
	user.Uuid = uuid
	user.CreatedAt = now
	return nil
}

func (user *User) CreateSession() (Session, error) {
	stmt, err := Db.Prepare("insert into sessions (uuid, email, user_id, created_at) values(?, ?, ?, ?)")
	sess := Session{
		Email:  user.Email,
		UserId: user.Id,
	}
	defer stmt.Close()

	if err != nil {
		return sess, err
	}
	uuid, err := createUUID()
	if err != nil {
		return sess, err
	}
	now := time.Now()
	result, err := stmt.Exec(uuid, user.Email, user.Id, now)
	if err != nil {
		return Session{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Session{}, err
	}
	//err = stmt.QueryRow(uuid, user.Email, user.Id, time.Now()).Scan(&sess.Id, &sess.Uuid, &sess.CreatedAt)
	sess.Id = int(id)
	sess.Uuid = uuid
	sess.CreatedAt = now
	return sess, nil
}

func (user *User) GetSession() (Session, error) {
	sess := Session{}
	err := Db.QueryRow("select id,uuid,email,user_id,created_at from sessions where user_id=?", user.Id).
		Scan(&sess.Id, &sess.Uuid, &sess.Email, &sess.UserId, &sess.CreatedAt)
	return sess, err
}

func (user *User) CreateThread(topic string) (Thread, error) {
	stmt, err := Db.Prepare("insert into threads (uuid, topic, user_id, created_at) values(?, ?, ?, ?)")
	t := Thread{}
	if err != nil {
		return t, err
	}
	defer stmt.Close()
	uuid, err := createUUID()
	if err != nil {
		return t, err
	}
	now := time.Now()
	result, err := stmt.Exec(uuid, topic, user.Id, now)
	if err != nil {
		return t, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return t, err
	}
	t.Id = int(id)
	t.Uuid = uuid
	t.Topic = topic
	t.UserId = user.Id
	t.CreatedAt = now
	//err = stmt.QueryRow(uuid, topic, user.Id, time.Now()).Scan(&t.Id, &t.Uuid, &t.Topic, &t.UserId, &t.CreatedAt)
	return t, nil
}

func (user *User) CreatePost(t Thread, body string) (Post, error) {
	stmt, err := Db.Prepare("insert into posts (uuid, body, user_id, thread_id, created_at) values(?, ?, ?, ?, ?)")
	defer stmt.Close()
	post := Post{}
	if err != nil {
		return post, err
	}

	uuid, err := createUUID()
	if err != nil {
		return post, err
	}

	now := time.Now()
	result, err := stmt.Exec(uuid, body, user.Id, t.Id, now)
	if err != nil {
		return post, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return post, err
	}
	post.Id = int(id)
	post.Uuid = uuid
	post.UserId = user.Id
	post.ThreadId = t.Id
	post.CreatedAt = now
	//err = stmt.QueryRow(uuid, body, user.Id, t.Id, time.Now()).
	//	Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return post, nil
}

func GetUserByEmail(email string) (User, error) {
	user := User{}
	err := Db.QueryRow("select id, uuid, name, email, password, created_at from users where email = ?;", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

func (sess *Session) Check() (bool, error) {
	err := Db.QueryRow("select id, uuid, email, user_id, created_at from sessions where uuid=?", sess.Uuid).
		Scan(&sess.Id, &sess.Uuid, &sess.Email, &sess.UserId, &sess.CreatedAt)
	if err != nil {
		return false, err
	}
	if sess.Id != 0 {
		return true, nil
	}
	return false, nil
}

func (sess *Session) DeleteByUUID() error {
	stmt, err := Db.Prepare("delete from sessions where uuid=?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sess.Uuid)
	return err
}

func (sess *Session) GetUser() (User, error) {
	user := User{}
	err := Db.QueryRow("select id, uuid, name, email, created_at from users where id=?", sess.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}
