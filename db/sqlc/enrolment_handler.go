package db

import "context"

//ListEnrolments db handler for api call to list enrolment data of the database
func (store *Store) ListEnrolments(ctx context.Context, params ListEnrolmentsParams) ([]CourseEnrolment, error) {
	return store.Queries.ListEnrolments(ctx, params)
}