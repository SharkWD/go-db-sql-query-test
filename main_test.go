package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	res, err := selectClient(db, clientID)

	require.NoError(t, err)

	assert.Equal(t, clientID, res.ID)
	assert.NotEmpty(t, res.FIO)
	assert.NotEmpty(t, res.Login)
	assert.NotEmpty(t, res.Birthday)
	assert.NotEmpty(t, res.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	clientID := -1
	defer db.Close()

	// напиши тест здесь
	res, err := selectClient(db, clientID)

	require.Equal(t, sql.ErrNoRows, err)

	assert.Empty(t, res.ID)
	assert.Empty(t, res.FIO)
	assert.Empty(t, res.Login)
	assert.Empty(t, res.Birthday)
	assert.Empty(t, res.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.ID, err = insertClient(db, cl)

	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)

	res, err := selectClient(db, cl.ID)

	require.NoError(t, err)
	assert.Equal(t, cl, res)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)

	require.NoError(t, err)
	require.NotEmpty(t, id)

	_, err = selectClient(db, id)
	require.NoError(t, err)

	err = deleteClient(db, id)
	require.NoError(t, err)

	_, err = selectClient(db, id)

	require.Equal(t, sql.ErrNoRows, err)

}
