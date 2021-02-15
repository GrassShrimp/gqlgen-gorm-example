package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/grassshrimp/gqlgen-gorm-example/graph/generated"
	"github.com/grassshrimp/gqlgen-gorm-example/graph/model"
)

func (r *customerResolver) Employee(ctx context.Context, obj *model.Customer) (*model.Employee, error) {
	employee := model.Employee{}

	result := r.DB.Take(&employee, "EmployeeNumber = ?", obj.EmployeeID)

	if result.Error != nil {
		return &model.Employee{}, result.Error
	}

	return &employee, nil
}

func (r *employeeResolver) Office(ctx context.Context, obj *model.Employee) (*model.Office, error) {
	office := model.Office{}

	result := r.DB.Take(&office, "OfficeID = ?", obj.OfficeID)

	if result.Error != nil {
		return &model.Office{}, result.Error
	}

	return &office, nil
}

func (r *employeeResolver) Employee(ctx context.Context, obj *model.Employee) (*model.Employee, error) {
	employee := model.Employee{}

	result := r.DB.Take(&employee, "EmployeeNumber = ?", obj.EmployeeID)

	if result.Error != nil {
		return &model.Employee{}, result.Error
	}

	return &employee, nil
}

func (r *mutationResolver) CreateOffice(ctx context.Context, input model.OfficeCreate) (*model.Office, error) {
	office := model.Office(input)

	result := r.DB.Create(&office)

	if result.Error != nil {
		return &model.Office{}, result.Error
	}

	return &office, nil
}

func (r *mutationResolver) UpdateOffice(ctx context.Context, id string, input model.OfficeUpdate) (*model.Office, error) {
	office := model.Office{}

	result := r.DB.Take(&office, "OfficeCode like ?", id)

	if result.Error != nil {
		return &model.Office{}, result.Error
	}

	result = r.DB.Model(&office).Where("OfficeCode like ?", id).Updates(input)

	if result.Error != nil {
		return &model.Office{}, result.Error
	}

	return &office, nil
}

func (r *mutationResolver) DeleteOffice(ctx context.Context, id string) (bool, error) {
	result := r.DB.Delete(&model.Office{}, "OfficeCode like ?", id)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (r *orderResolver) Customer(ctx context.Context, obj *model.Order) (*model.Customer, error) {
	customer := model.Customer{}

	result := r.DB.Take(&customer, "CustomerNumber = ?", obj.CustomerID)

	if result.Error != nil {
		return &model.Customer{}, result.Error
	}

	return &customer, nil
}

func (r *orderdetailResolver) Order(ctx context.Context, obj *model.Orderdetail) (*model.Order, error) {
	order := model.Order{}

	result := r.DB.Take(&order, "OrderNumber = ?", obj.OrderID)

	if result.Error != nil {
		return &model.Order{}, result.Error
	}

	return &order, nil
}

func (r *orderdetailResolver) Product(ctx context.Context, obj *model.Orderdetail) (*model.Product, error) {
	product := model.Product{}

	result := r.DB.Take(&product, "ProductCode like ?", obj.ProductID)

	if result.Error != nil {
		return &model.Product{}, result.Error
	}

	return &product, nil
}

func (r *productResolver) ProductLine(ctx context.Context, obj *model.Product) (*model.Productline, error) {
	productLine := model.Productline{}

	result := r.DB.Take(&productLine, "ProductLine like ?", obj.ProductLineID)

	if result.Error != nil {
		return &model.Productline{}, result.Error
	}

	return &productLine, nil
}

func (r *queryResolver) Product(ctx context.Context, productCode *string) (*model.Product, error) {
	product := model.Product{}

	result := r.DB.Take(&product, "productCode like ?", productCode)

	if result.Error != nil {
		return &model.Product{}, result.Error
	}

	return &product, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	products := []*model.Product{}
	result := r.DB.Find(&products)

	if result.Error != nil {
		return []*model.Product{}, result.Error
	}

	return products, nil
}

func (r *queryResolver) Offices(ctx context.Context) ([]*model.Office, error) {
	offices := []*model.Office{}
	result := r.DB.Find(&offices)

	if result.Error != nil {
		return []*model.Office{}, result.Error
	}

	return offices, nil
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	customers := []*model.Customer{}

	result := r.DB.Find(&customers)

	if result.Error != nil {
		return []*model.Customer{}, result.Error
	}

	return customers, nil
}

func (r *queryResolver) Employees(ctx context.Context) ([]*model.Employee, error) {
	employees := []*model.Employee{}

	result := r.DB.Find(&employees)

	if result.Error != nil {
		return []*model.Employee{}, result.Error
	}

	return employees, nil
}

func (r *queryResolver) Orderdetails(ctx context.Context) ([]*model.Orderdetail, error) {
	orderdetails := []*model.Orderdetail{}

	result := r.DB.Find(&orderdetails)

	if result.Error != nil {
		return []*model.Orderdetail{}, result.Error
	}

	return orderdetails, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	orders := []*model.Order{}

	result := r.DB.Find(&orders)

	if result.Error != nil {
		return []*model.Order{}, result.Error
	}

	return orders, nil
}

func (r *queryResolver) Payments(ctx context.Context) ([]*model.Payment, error) {
	payments := []*model.Payment{}

	result := r.DB.Joins("Customer").Find(&payments)

	if result.Error != nil {
		return []*model.Payment{}, result.Error
	}

	return payments, nil
}

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

// Employee returns generated.EmployeeResolver implementation.
func (r *Resolver) Employee() generated.EmployeeResolver { return &employeeResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// Orderdetail returns generated.OrderdetailResolver implementation.
func (r *Resolver) Orderdetail() generated.OrderdetailResolver { return &orderdetailResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type customerResolver struct{ *Resolver }
type employeeResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type orderdetailResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
