package product

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/service"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/i_crud"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/storage"
)

// ProductService define las operaciones de negocio para Product.
type ProductService interface {
	i_crud.CRUDService[Product]
	UploadImages(ctx context.Context, form *multipart.Form, productID string) ([]storage.FileObject, error)
	DeleteImages(ctx context.Context, productID string, images []string) error
}

type productService struct {
	i_crud.CRUDService[Product]
	repo     ProductRepository
	uploader storage.UploadService
	deleter  storage.DeleteService
}

// NewProductService crea una nueva instancia de ProductService inyectando el repositorio.
func NewProductService(
	repo ProductRepository,
	uploader storage.UploadService,
	deleter storage.DeleteService,
) ProductService {

	crudService := service.NewCRUDService(repo)

	return &productService{
		CRUDService: crudService,
		repo:        repo,
		uploader:    uploader,
		deleter:     deleter,
	}
	// return &productService{CRUDService: genService}
}

func (s *productService) UploadImages(ctx context.Context, form *multipart.Form, productID string) ([]storage.FileObject, error) {
	files := form.File["images"]
	var images []storage.FileObject

	for _, f := range files {
		log.Println("📸 Subiendo imagenes...", f.Filename)
		upContext := context.Background()
		file, _ := f.Open()
		fileObject, err := s.uploader.Upload(upContext, file, f, productID)
		if err != nil {
			return nil, err
		}
		images = append(images, fileObject)
	}

	// Obtener producto actual
	product, err := s.repo.FindByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	// Asignar imágenes
	product.Images = append(product.Images, images...)

	// Actualizar en base de datos
	if err := s.Update(ctx, productID, product); err != nil {
		return nil, err
	}

	return images, nil
}

func (s *productService) DeleteImages(ctx context.Context, productID string, filesIds []string) error {

	for _, id := range filesIds {
		if err := s.deleter.DeleteObject(ctx, id); err != nil {
			return err
		}
	}

	return nil

}
