package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

func (m *AttendeeModel) Insert(attendee Attendee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		INSERT INTO attendees (user_id, event_id)
		VALUES (?, ?)
	`

	_, err := m.DB.ExecContext(ctx, query, attendee.UserID, attendee.EventID)
	if err != nil {
		return fmt.Errorf("failed to insert attendee: %w", err)
	}

	return nil
}

func (m *AttendeeModel) GetAll() ([]Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `SELECT id, user_id, event_id FROM attendees`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query attendees: %w", err)
	}
	defer rows.Close()

	var attendees []Attendee
	for rows.Next() {
		var attendee Attendee
		err := rows.Scan(&attendee.ID, &attendee.UserID, &attendee.EventID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan attendee: %w", err)
		}
		attendees = append(attendees, attendee)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over attendees: %w", err)
	}

	return attendees, nil
}

func (m *AttendeeModel) Get(id string) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `SELECT id, user_id, event_id FROM attendees WHERE id = ?`

	var attendee Attendee
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&attendee.ID, &attendee.UserID, &attendee.EventID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("attendee not found")
		}
		return nil, fmt.Errorf("failed to get attendee: %w", err)
	}

	return &attendee, nil
}

func (m *AttendeeModel) Update(id string, attendee Attendee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		UPDATE attendees 
		SET user_id = ?, event_id = ?
		WHERE id = ?
	`

	result, err := m.DB.ExecContext(ctx, query, attendee.UserID, attendee.EventID, id)
	if err != nil {
		return fmt.Errorf("failed to update attendee: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("attendee not found")
	}

	return nil
}

func (m *AttendeeModel) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `DELETE FROM attendees WHERE id = ?`

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete attendee: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("attendee not found")
	}

	return nil
}

func (m *AttendeeModel) GetByEventID(eventID int) ([]Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `SELECT id, user_id, event_id FROM attendees WHERE event_id = ?`

	rows, err := m.DB.QueryContext(ctx, query, eventID)
	if err != nil {
		return nil, fmt.Errorf("failed to query attendees by event: %w", err)
	}
	defer rows.Close()

	var attendees []Attendee
	for rows.Next() {
		var attendee Attendee
		err := rows.Scan(&attendee.ID, &attendee.UserID, &attendee.EventID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan attendee: %w", err)
		}
		attendees = append(attendees, attendee)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over attendees: %w", err)
	}

	return attendees, nil
}
