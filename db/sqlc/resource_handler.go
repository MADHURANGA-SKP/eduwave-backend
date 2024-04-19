package db

import (
	"context"
)

//CreateResourceParam contains the input parameters of data
type CreateResourceParam struct {
	MaterialID  int64          `json:"material_id"`
	Title       string         `json:"title"`
	Type        TypeResource   `json:"type"`
	ContentUrl  string         `json:"content_url"`
}

//CreateResourceResponse contains the result of the creation of data
type CreateResourceResponse struct {
	Resource Resource `json:"resource"`
}

//CreateResource db handler fro api call to update resource data in database
func (store *Store) CreateResource(ctx context.Context, arg CreateResourceParam) (CreateResourceResponse, error) {
	var result CreateResourceResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Resource, err = q.CreateResource(ctx, CreateResourceParams{
			MaterialID: arg.MaterialID,
			Title:      arg.Title,
			Type:       arg.Type,
			ContentUrl: arg.ContentUrl,
		})

		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

//DeleteMaterialParam contains the input parameters of the delete the data
type DeleteResourceParam struct {
	ResourceID  int64  `json:"resource_id"`
}

//DeleteResource db handler for api call to delete exact data from the database
func (store *Store) DeleteResource(ctx context.Context, arg DeleteResourceParam) error {
	return store.Queries.DeleteResource(ctx, arg.ResourceID )
}

//GetResourceParam contains the input paramters of the retriving data
type GetResourceParam struct {
    ResourceID int64         `uri:"resource_id,min=1"`
}

//GetResourceResponse contains the result of the retriving data
type GetResourceResponse struct {
	Resource Resource `json:"resource"`
}

//GetResource db handler for api call to retrive a resource data from teh databse
func (store *Store) GetResource(ctx context.Context, arg GetResourceParam) (GetResourceResponse, error) {
	var result GetResourceResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Resource, err = q.GetResource(ctx, arg.ResourceID)

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

//ListResource db handler for api call to retrive a resource data from teh databse
func (store *Store) ListResource(ctx context.Context, params ListResourceParams) ([]Resource, error) {
	return store.Queries.ListResource(ctx, params)
}

//UpdateResourceParam contains the input parameters of the updating data
type UpdateResourceParam struct {
	ResourceID int64        `json:"resource_id"`
	Title      string       `json:"title"`
	Type       TypeResource `json:"type"`
	ContentUrl string       `json:"content_url"`
}

//UpdateResourceResponse contains the result of the updating data
type UpdateResourceResponse struct {
	Resource Resource `json:"resource"`
}

//UpdateResource db handler for api call to update resource data in the database
func(store *Store) UpdateResource(ctx context.Context, arg UpdateResourceParam)(UpdateResourceResponse, error){
	var result UpdateResourceResponse

	err := store.execTx( ctx, func(q *Queries) error {
		var err error

		result.Resource, err = q.UpdateResource(ctx, UpdateResourceParams{
			ResourceID: arg.ResourceID,
			Title: arg.Title,
			Type: arg.Type,
			ContentUrl:  arg.ContentUrl,
			Files: arg.Files,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}