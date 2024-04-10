package db

import "context"

//CreateMaterialParam contains the input parameters of  creating Material data
type CreateMaterialParam struct{
	CourseID    int64  `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//CreateMaterialResponse contains the result of the creating Material data 
type CreateMaterialReponse struct {
	Material Material `json:"Material"`
}

//CreateMaterial db handler for api call to create Material data in databse
func(store *Store) CreateMaterial(ctx context.Context, arg CreateMaterialParam)(CreateMaterialReponse, error){
	var result CreateMaterialReponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Material, err = q.CreateMaterial(ctx, CreateMaterialParams{
			CourseID: arg.CourseID,
			Title: arg.Title,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

//DeleteMaterialParam contains the input parameters of the delete the data
type DeleteMaterialParam struct {
	MaterialID  int64  `json:"material_id"`
}


//DeleteMatirila db handler for api call to delete Material data in database
func(store *Store) DeleteMaterial(ctx context.Context, arg DeleteMaterialParam)error {
	return store.Queries.DeleteMaterial(ctx, arg.MaterialID )
}

//GetMaterialparam contains the input parameters of the get Material data
type GetMaterialParam struct {
	MaterialID int64 `json:"material_id"`
}

//GetMaterialResponse contains the result of the get matrial data
type GetMaterialResponse struct {
	Material Material `json:"Material"`
}

//GetMaterial db handler for api call to get Material data in database
func(store *Store) GetMaterial(ctx context.Context, arg GetMaterialParam)(GetMaterialResponse, error){
	var result GetMaterialResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Material, err = q.GetMaterial(ctx, arg.MaterialID)
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

//ListMaterial db handler fro api call to list Material data in database
func(store *Store) ListMaterial(ctx context.Context, params ListMaterialParams)([]Material, error){
	return store.Queries.ListMaterial(ctx, params)
}

//UpdateMaterialParam contains the input parameters of the Update Material data
type UpdateMaterialParam struct {
	MaterialID  int64         `json:"Material_id"`
	CourseID    int64 `json:"course_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
}

//UpdateMaterialResponse contains the result of the Updated Material data
type UpdateMaterialResponse struct {
	Material Material `json:"Material"`
}

//UpdateMatririal db handler for api call to the update Material data in database
func(store *Store) UpdateMaterials(ctx context.Context, arg UpdateMaterialParam)(UpdateMaterialResponse, error){
	var result UpdateMaterialResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Material, err = q.UpdateMaterial(ctx, UpdateMaterialParams{
			MaterialID: arg.MaterialID,
			CourseID: arg.CourseID,
			Title: arg.Title,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}
		return err
	})
	return result, err
}