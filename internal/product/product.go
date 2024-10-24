package product

// This is a product a multiples products in a package
type Product struct {
	ID          int      `json:"id"`
	UserID      int      `json:"user_id"`
	WarehouseID int      `json:"warehouse_id"`
	Status      Status   `json:"status"`
	Invoices    []string `json:"invoices"` //Aws s3 urls
}

type Status string

const (
	TransitToWarehouse     Status = "TRANSITTOWAREHOUSE"
	InWarehouse            Status = "INWAREHOUSE"
	TransitToUserDirection Status = "TRANSITTOUSERDIRECTION"
	ArrivedToUser          Status = "ARRIVEDTOUSER"
)
