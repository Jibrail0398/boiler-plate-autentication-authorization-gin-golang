package db

import (
	_"github.com/lib/pq"
	"database/sql"
	"fmt"
)

type Credential struct{
	Host 	string
	Username string
	Password string
	DatabaseName string
	Port int
}

type postgres struct{
	DB *sql.DB
}

func NewDatabase() *postgres {
	return &postgres{}
}

func (p *postgres) Connect(credential Credential) (*sql.DB,error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        credential.Host, credential.Port, credential.Username, credential.Password, credential.DatabaseName)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }
    
	p.DB = db
    return db, nil
}

func (p *postgres) Up() error{
	_,err := p.DB.Exec(
        `CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			name VARCHAR(200) NOT NULL,
			email VARCHAR(100) NOT NULL,
			password VARCHAR(100) NULL,
			oauth_provider VARCHAR(255) NULL,
			oauth_id VARCHAR(255) NULL,
			verified BOOLEAN NOT NULL,
			created timestamp DEFAULT NOW(),
			updated timestamp DEFAULT NOW()
		
		)
		`)
    if err!=nil{
        return fmt.Errorf("failed to create tabel users: %w", err)
    }

	return nil
}

