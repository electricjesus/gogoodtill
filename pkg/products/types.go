package products

type ProductCreateUpdateRequest struct {
	// REQUIRED FOR CREATE
	Name            string      `json:"product_name"`      // REQUIRED FOR CREATE Product Name
	ParentProductID interface{} `json:"parent_product_id"` // REQUIRED FOR CREATE Parent product ID. null if no parent
	VATCodeID       string      `json:"vat_code_id"`       // REQUIRED FOR CREATE VAT rate ID.
	Price           string      `json:"selling_price"`     // REQUIRED FOR CREATE Product Selling Price
	// ATTRIBUTE
	ChildAttributes ProductChildAttributes `json:"child_attributes"` // Child attributes if any.
	// INFORMATION
	Active        uint   `json:"active,omitempty"`                // Default: 1, Allowed Values: [0, 1]
	Shareable     uint   `json:"sharable,omitempty"`              // Default: 1, Allowed Values: [0, 1]
	CategoryID    string `json:"category_i,omitempty"`            // Category ID. Must exists in the system if provided
	BrandID       string `json:"brand_id,omitempty"`              // Brand ID. Must exists in the system if provided
	SupplierID    string `json:"supplier_id,omitempty"`           // Supplier ID. Must exists in the system if provided
	SKU           string `json:"product_sku,omitempty"`           // Product SKU
	DisplayName   string `json:"display_name,omitempty"`          // Display Name on the POS
	Description   string `json:"product_desc,omitempty"`          // Product Description
	Barcode       string `json:"barcode,omitempty"`               // Barcode
	LabelAsMain   uint   `json:"label_as_main,omitempty"`         // Label this product as main. Used in calculating covers in restaurant scenarios. Default: 0, Values: 0, 1
	IsLoose       uint   `json:"is_loose,omitempty"`              // If the product is a loose product. 1=Yes, 0=No, Default: 0, Values: 0, 1
	IsOpenPrice   uint   `json:"is_open_price_product,omitempty"` // If the product is a open price product. 1=Yes, 0=No, Default: 0, Values: 0, 1
	Weight        uint64 `json:"weight,omitempty"`                // Product Weight. Amount in grams.
	DeliveryCode  string `json:"delivery_code,omitempty"`         // Product Delivery Code. Used in Stock module if provided.
	DisplayColour string `json:"display_colour,omitempty"`        // colour to be shown in POS. Allowed values: 'cc0000', '000000', '333333', '2299ee', '8b8500', 'c63d0f', '7e8f7c', '005a31', 'ff9009', 'df3d82'
	AccountCode   string `json:"account_code,omitempty"`          // Account Code. To be used when sending data to Accounting platform such as Xero or Exact.
	Tags          string `json:"tags,omitempty"`                  // Array of Tag names. See Example.
	// STOCK CONTROL
	Inventory            uint64 `json:"inventory"`
	MinStock             uint64 `json:"min_stock"`
	StockQuantifier      uint64 `json:"stock_quantifier"`       // Default: 1
	TrackInventory       uint   `json:"track_inventory"`        // Default: 1, Allowed Values: [0, 1]
	AlertOn              uint   `json:"alert_on"`               // Default: 0, Allowed Values: [0, 1]
	AlertBelow           uint   `json:"alert_below"`            // Default: 0, Allowed Values: [0, 1]
	DeliveryReorder      uint   `json:"delivery_reorder_on"`    // Default: 0, Allowed Values: [0, 1]
	DeliveryReorderPoint uint   `json:"delivery_reorder_point"` // Default: 0, Allowed Values: [0, 1]
	TakeStockFromParent  uint   `json:"take_stock_from_parent"` // Default: 0, Allowed Values: [0, 1]
	HasIngredientStock   uint   `json:"has_ingredient_stock"`   // Default: 0, Allowed Values: [0, 1]
	// PRICING
	PurchasePrice                 float64 `json:"purchase_price"`                    // Supplier purchase price for a single unit.
	SupplierPurchasePrice         float64 `json:"supplier_purchase_price"`           //  Supplier purchase price for supplier unit.
	SupplierPurchaseUnit          string  `json:"supplier_unit"`                     // Supplier unit.
	StoreToSupplierUnitConversion string  `json:"store_to_supplier_unit_conversion"` // Conversion from store to supplier units.
	// TODO: Printing and KDS
	// TODO: Modifiers
}

type ProductChildAttributes map[string]string

type ProductCreateCurrentVariantsRequest struct {
	ID                  string                 `json:"id"`
	ProductName         string                 `json:"product_name"`
	ProductSku          string                 `json:"product_sku"`
	DisplayName         string                 `json:"display_name"`
	SellingPrice        string                 `json:"selling_price"`
	ChildAttributes     ProductChildAttributes `json:"child_attributes"`
	TakeStockFromParent int                    `json:"take_stock_from_parent"`
	StockQuantifier     string                 `json:"stock_quantifier"`
	Inventory           string                 `json:"inventory"`
}

type ProductCreateVariantsRequest struct {
	ProductName         string                 `json:"product_name"`
	ProductSku          string                 `json:"product_sku"`
	DisplayName         string                 `json:"display_name"`
	SellingPrice        string                 `json:"selling_price"`
	ChildAttributes     ProductChildAttributes `json:"child_attributes"`
	TakeStockFromParent int                    `json:"take_stock_from_parent"`
	StockQuantifier     int                    `json:"stock_quantifier"`
	Inventory           int                    `json:"inventory"`
}

type ProductCreateModifierSets struct {
	ID string `json:"id"`
}

type ProductDeleteRequest struct{}

type ProductDataResponse struct {
	OutletID                      string      `json:"outlet_id"`
	CategoryID                    string      `json:"category_id"`
	BrandID                       string      `json:"brand_id"`
	SupplierID                    string      `json:"supplier_id"`
	VatCodeID                     string      `json:"vat_code_id"`
	ProductName                   string      `json:"product_name"`
	ProductSku                    string      `json:"product_sku"`
	ProductDesc                   string      `json:"product_desc"`
	DisplayName                   string      `json:"display_name"`
	Barcode                       string      `json:"barcode"`
	Weight                        int         `json:"weight"`
	IsLoose                       int         `json:"is_loose"`
	IsOpenPriceProduct            int         `json:"is_open_price_product"`
	ProductLabel                  string      `json:"product_label"`
	LabelAsMain                   int         `json:"label_as_main"`
	Image                         interface{} `json:"image"`
	PurchasePrice                 string      `json:"purchase_price"`
	SupplierPurchasePrice         string      `json:"supplier_purchase_price"`
	SupplierUnit                  interface{} `json:"supplier_unit"`
	StoreToSupplierUnitConversion string      `json:"store_to_supplier_unit_conversion"`
	SellingPrice                  string      `json:"selling_price"`
	TrackInventory                int         `json:"track_inventory"`
	MinStock                      string      `json:"min_stock"`
	HasIngredientStock            int         `json:"has_ingredient_stock"`
	TakeStockFromParent           int         `json:"take_stock_from_parent"`
	StockQuantifier               string      `json:"stock_quantifier"`
	AlertOn                       int         `json:"alert_on"`
	AlertBelow                    string      `json:"alert_below"`
	DeliveryReorderOn             int         `json:"delivery_reorder_on"`
	DeliveryReorderPoint          string      `json:"delivery_reorder_point"`
	DeliveryCode                  string      `json:"delivery_code"`
	HasVariant                    int         `json:"has_variant"`
	Active                        int         `json:"active"`
	Shareable                     int         `json:"shareable"`
	PrintOnReceipt                int         `json:"print_on_receipt"`
	PrintOnKitchen                int         `json:"print_on_kitchen"`
	PrintOnDrink                  int         `json:"print_on_drink"`
	PrintOnOther                  int         `json:"print_on_other"`
	LockPrice                     int         `json:"lock_price"`
	DisplayColour                 string      `json:"display_colour"`
	AccountCode                   interface{} `json:"account_code"`
	DisplayOrder                  int         `json:"display_order"`
	HasModifier                   int         `json:"has_modifier"`
	TicketPrinter1                int         `json:"ticket_printer_1"`
	TicketPrinter2                int         `json:"ticket_printer_2"`
	TicketPrinter3                int         `json:"ticket_printer_3"`
	TicketPrinter4                int         `json:"ticket_printer_4"`
	HasAttributes                 int         `json:"has_attributes"`
	DisplayAttrPos                int         `json:"display_attr_pos"`
	LastSyncDate                  int         `json:"last_sync_date"`
	ID                            string      `json:"id"`
	ClientID                      string      `json:"client_id"`
	UpdatedAt                     string      `json:"updated_at"`
	CreatedAt                     string      `json:"created_at"`
}

type ProductCreateUpdateResponse struct {
	Status  bool                `json:"status"`
	Data    ProductDataResponse `json:"data"`
	Message string              `json:"message"`
}

type ProductListResponse struct {
	Status bool `json:"status"`
	Data   []struct {
		ID                    string      `json:"id"`
		OutletID              string      `json:"outlet_id"`
		ParentProductID       interface{} `json:"parent_product_id"`
		CategoryID            string      `json:"category_id"`
		BrandID               string      `json:"brand_id"`
		SupplierID            string      `json:"supplier_id"`
		ProductName           string      `json:"product_name"`
		ProductSku            string      `json:"product_sku"`
		DisplayName           string      `json:"display_name"`
		PurchasePrice         string      `json:"purchase_price"`
		SupplierPurchasePrice string      `json:"supplier_purchase_price"`
		SupplierUnit          interface{} `json:"supplier_unit"`
		SellingPrice          string      `json:"selling_price"`
		HasVariant            int         `json:"has_variant"`
		TakeStockFromParent   int         `json:"take_stock_from_parent"`
		StockQuantifier       string      `json:"stock_quantifier"`
		HasIngredientStock    int         `json:"has_ingredient_stock"`
		Active                int         `json:"active"`
		Shareable             int         `json:"shareable"`
		Image                 interface{} `json:"image"`
		UnitConversion        string      `json:"unit_conversion"`
		StoreUnit             interface{} `json:"store_unit"`
		TopLevelProduct       bool        `json:"top_level_product"`
		CanChangeProduct      bool        `json:"can_change_product"`
		TrackInventory        int         `json:"track_inventory"`
		Inventory             string      `json:"inventory"`
		MinStock              string      `json:"min_stock"`
		AlertOn               int         `json:"alert_on"`
		AlertBelow            string      `json:"alert_below"`
		DeliveryReorderOn     int         `json:"delivery_reorder_on"`
		DeliveryReorderPoint  string      `json:"delivery_reorder_point"`
		PrintOnReceipt        int         `json:"print_on_receipt"`
		PrintOnKitchen        int         `json:"print_on_kitchen"`
		PrintOnDrink          int         `json:"print_on_drink"`
		PrintOnOther          int         `json:"print_on_other"`
		TicketPrinter1        int         `json:"ticket_printer_1"`
		TicketPrinter2        int         `json:"ticket_printer_2"`
		TicketPrinter3        int         `json:"ticket_printer_3"`
		TicketPrinter4        int         `json:"ticket_printer_4"`
		TakeawayOverridePrice int         `json:"takeaway_override_price"`
		TakeawaySellingPrice  string      `json:"takeaway_selling_price"`
		OcOverridePrice       int         `json:"oc_override_price"`
		OcSellingPrice        string      `json:"oc_selling_price"`
	} `json:"data"`
	Message string `json:"message"`
}

type ProductDetailsReponse struct {
	Success          int  `json:"success"`
	CanChangeProduct bool `json:"can_change_product"`
	Data             struct {
		ID                            string      `json:"id"`
		OutletID                      string      `json:"outlet_id"`
		CategoryID                    string      `json:"category_id"`
		BrandID                       string      `json:"brand_id"`
		SupplierID                    string      `json:"supplier_id"`
		VatCodeID                     string      `json:"vat_code_id"`
		VatRate                       string      `json:"vat_rate"`
		ParentProductID               interface{} `json:"parent_product_id"`
		ProductName                   string      `json:"product_name"`
		ProductSku                    string      `json:"product_sku"`
		ProductDesc                   string      `json:"product_desc"`
		DisplayName                   string      `json:"display_name"`
		Barcode                       string      `json:"barcode"`
		Weight                        int         `json:"weight"`
		IsLoose                       int         `json:"is_loose"`
		IsOpenPriceProduct            int         `json:"is_open_price_product"`
		ProductLabel                  string      `json:"product_label"`
		LabelAsMain                   int         `json:"label_as_main"`
		Image                         interface{} `json:"image"`
		PurchasePrice                 string      `json:"purchase_price"`
		SupplierPurchasePrice         string      `json:"supplier_purchase_price"`
		SupplierUnit                  interface{} `json:"supplier_unit"`
		StoreToSupplierUnitConversion string      `json:"store_to_supplier_unit_conversion"`
		SellingPrice                  string      `json:"selling_price"`
		TrackInventory                int         `json:"track_inventory"`
		InventoryID                   string      `json:"inventory_id"`
		Inventory                     string      `json:"inventory"`
		MinStock                      string      `json:"min_stock"`
		HasIngredientStock            int         `json:"has_ingredient_stock"`
		TakeStockFromParent           int         `json:"take_stock_from_parent"`
		StockQuantifier               string      `json:"stock_quantifier"`
		CurrentVariants               []struct {
			ID                            string      `json:"id"`
			ClientID                      string      `json:"client_id"`
			OutletID                      string      `json:"outlet_id"`
			ParentProductID               string      `json:"parent_product_id"`
			CategoryID                    string      `json:"category_id"`
			BrandID                       interface{} `json:"brand_id"`
			SupplierID                    interface{} `json:"supplier_id"`
			VatCodeID                     string      `json:"vat_code_id"`
			ProductName                   string      `json:"product_name"`
			ProductHandle                 interface{} `json:"product_handle"`
			ProductDesc                   string      `json:"product_desc"`
			ProductSku                    string      `json:"product_sku"`
			DisplayName                   string      `json:"display_name"`
			ProductImage                  interface{} `json:"product_image"`
			Barcode                       string      `json:"barcode"`
			IsLoose                       int         `json:"is_loose"`
			IsOpenPriceProduct            int         `json:"is_open_price_product"`
			ProductLabel                  string      `json:"product_label"`
			Weight                        int         `json:"weight"`
			DeliveryCode                  string      `json:"delivery_code"`
			PurchasePrice                 string      `json:"purchase_price"`
			SupplierPurchasePrice         string      `json:"supplier_purchase_price"`
			SupplierUnit                  interface{} `json:"supplier_unit"`
			StoreToSupplierUnitConversion string      `json:"store_to_supplier_unit_conversion"`
			SellingPrice                  string      `json:"selling_price"`
			TrackInventory                int         `json:"track_inventory"`
			Inventory                     string      `json:"inventory"`
			MinStock                      string      `json:"min_stock"`
			StockLevel1                   interface{} `json:"stock_level_1"`
			StockLevel2                   interface{} `json:"stock_level_2"`
			AlertOn                       int         `json:"alert_on"`
			AlertBelow                    string      `json:"alert_below"`
			DeliveryReorderOn             int         `json:"delivery_reorder_on"`
			DeliveryReorderPoint          string      `json:"delivery_reorder_point"`
			HasVariant                    int         `json:"has_variant"`
			HasAttributes                 int         `json:"has_attributes"`
			DisplayAttrPos                int         `json:"display_attr_pos"`
			TakeStockFromParent           int         `json:"take_stock_from_parent"`
			StockQuantifier               string      `json:"stock_quantifier"`
			HasIngredientStock            int         `json:"has_ingredient_stock"`
			LabelAsMain                   int         `json:"label_as_main"`
			PrintOnReceipt                int         `json:"print_on_receipt"`
			PrintOnKitchen                int         `json:"print_on_kitchen"`
			PrintOnDrink                  int         `json:"print_on_drink"`
			PrintOnOther                  int         `json:"print_on_other"`
			TicketPrinter1                int         `json:"ticket_printer_1"`
			TicketPrinter2                int         `json:"ticket_printer_2"`
			TicketPrinter3                int         `json:"ticket_printer_3"`
			TicketPrinter4                int         `json:"ticket_printer_4"`
			DisplayOrder                  int         `json:"display_order"`
			AccountCode                   interface{} `json:"account_code"`
			Active                        int         `json:"active"`
			IsDeleted                     int         `json:"is_deleted"`
			DeletedAt                     interface{} `json:"deleted_at"`
			Shareable                     int         `json:"shareable"`
			LockPrice                     int         `json:"lock_price"`
			HasModifier                   int         `json:"has_modifier"`
			DisplayColour                 interface{} `json:"display_colour"`
			Image                         interface{} `json:"image"`
			CustomField1                  string      `json:"custom_field_1"`
			CustomField2                  string      `json:"custom_field_2"`
			CustomField3                  string      `json:"custom_field_3"`
			CustomField4                  string      `json:"custom_field_4"`
			LastSyncDate                  int         `json:"last_sync_date"`
			CreatedAt                     string      `json:"created_at"`
			UpdatedAt                     string      `json:"updated_at"`
			ProductDetails                struct {
				SellingPrice         string `json:"selling_price"`
				TrackInventory       int    `json:"track_inventory"`
				Inventory            string `json:"inventory"`
				MinStock             string `json:"min_stock"`
				AlertOn              int    `json:"alert_on"`
				AlertBelow           string `json:"alert_below"`
				DeliveryReorderOn    int    `json:"delivery_reorder_on"`
				DeliveryReorderPoint string `json:"delivery_reorder_point"`
				PrintOnReceipt       int    `json:"print_on_receipt"`
				PrintOnKitchen       int    `json:"print_on_kitchen"`
				PrintOnDrink         int    `json:"print_on_drink"`
				PrintOnOther         int    `json:"print_on_other"`
				TicketPrinter1       int    `json:"ticket_printer_1"`
				TicketPrinter2       int    `json:"ticket_printer_2"`
				TicketPrinter3       int    `json:"ticket_printer_3"`
				TicketPrinter4       int    `json:"ticket_printer_4"`
			} `json:"product_details"`
			Tags         []interface{} `json:"tags"`
			ModifierSets []interface{} `json:"modifier_sets"`
			Attributes   []struct {
				ID            string   `json:"id"`
				ClientID      string   `json:"client_id"`
				AttributeName string   `json:"attribute_name"`
				DefaultValue  []string `json:"default_value"`
				IsGlobal      int      `json:"is_global"`
				Active        int      `json:"active"`
				Pivot         struct {
					ProductID      string `json:"product_id"`
					AttributeID    string `json:"attribute_id"`
					AttributeValue string `json:"attribute_value"`
					DisplayOrder   int    `json:"display_order"`
				} `json:"pivot"`
			} `json:"attributes"`
			ProductAttributes []interface{} `json:"product_attributes"`
		} `json:"current_variants"`
		AlertOn              int         `json:"alert_on"`
		AlertBelow           string      `json:"alert_below"`
		DeliveryReorderOn    int         `json:"delivery_reorder_on"`
		DeliveryReorderPoint string      `json:"delivery_reorder_point"`
		DeliveryCode         string      `json:"delivery_code"`
		HasVariant           int         `json:"has_variant"`
		Active               int         `json:"active"`
		Status               string      `json:"status"`
		Shareable            int         `json:"shareable"`
		PrintOnReceipt       int         `json:"print_on_receipt"`
		PrintOnKitchen       int         `json:"print_on_kitchen"`
		PrintOnDrink         int         `json:"print_on_drink"`
		PrintOnOther         int         `json:"print_on_other"`
		LockPrice            int         `json:"lock_price"`
		DisplayColour        string      `json:"display_colour"`
		AccountCode          interface{} `json:"account_code"`
		DisplayOrder         int         `json:"display_order"`
		Tags                 []struct {
			ID        string `json:"id"`
			ClientID  string `json:"client_id"`
			TagName   string `json:"tag_name"`
			Active    int    `json:"active"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
			Pivot     struct {
				ProductID string `json:"product_id"`
				TagID     string `json:"tag_id"`
			} `json:"pivot"`
		} `json:"tags"`
		TopLevelProduct  bool `json:"top_level_product"`
		CanChangeProduct bool `json:"can_change_product"`
		IsDeleted        int  `json:"is_deleted"`
		ProductDetails   struct {
			SellingPrice         string `json:"selling_price"`
			TrackInventory       int    `json:"track_inventory"`
			Inventory            string `json:"inventory"`
			MinStock             string `json:"min_stock"`
			AlertOn              int    `json:"alert_on"`
			AlertBelow           string `json:"alert_below"`
			DeliveryReorderOn    int    `json:"delivery_reorder_on"`
			DeliveryReorderPoint string `json:"delivery_reorder_point"`
			PrintOnReceipt       int    `json:"print_on_receipt"`
			PrintOnKitchen       int    `json:"print_on_kitchen"`
			PrintOnDrink         int    `json:"print_on_drink"`
			PrintOnOther         int    `json:"print_on_other"`
			TicketPrinter1       int    `json:"ticket_printer_1"`
			TicketPrinter2       int    `json:"ticket_printer_2"`
			TicketPrinter3       int    `json:"ticket_printer_3"`
			TicketPrinter4       int    `json:"ticket_printer_4"`
		} `json:"product_details"`
		HasModifier    int `json:"has_modifier"`
		TicketPrinter1 int `json:"ticket_printer_1"`
		TicketPrinter2 int `json:"ticket_printer_2"`
		TicketPrinter3 int `json:"ticket_printer_3"`
		TicketPrinter4 int `json:"ticket_printer_4"`
		HasAttributes  int `json:"has_attributes"`
		DisplayAttrPos int `json:"display_attr_pos"`
		Attributes     []struct {
			ID            string   `json:"id"`
			ClientID      string   `json:"client_id"`
			AttributeName string   `json:"attribute_name"`
			DefaultValue  []string `json:"default_value"`
			IsGlobal      int      `json:"is_global"`
			Active        int      `json:"active"`
			Pivot         struct {
				ProductID      string      `json:"product_id"`
				AttributeID    string      `json:"attribute_id"`
				AttributeValue interface{} `json:"attribute_value"`
				DisplayOrder   int         `json:"display_order"`
			} `json:"pivot"`
		} `json:"attributes"`
		CourseNo              int    `json:"course_no"`
		TakeawayNoVat         int    `json:"takeaway_no_vat"`
		TakeawayOverridePrice int    `json:"takeaway_override_price"`
		TakeawaySellingPrice  string `json:"takeaway_selling_price"`
		OcDisabled            int    `json:"oc_disabled"`
		OcNoVat               int    `json:"oc_no_vat"`
		OcOverridePrice       int    `json:"oc_override_price"`
		OcSellingPrice        string `json:"oc_selling_price"`
		CostFromIngredient    int    `json:"cost_from_ingredient"`
		CustomField1          string `json:"custom_field_1"`
		CustomField2          string `json:"custom_field_2"`
		CustomField3          string `json:"custom_field_3"`
		CustomField4          string `json:"custom_field_4"`
	} `json:"data"`
}

type ProductDeleteResponse struct {
	Status bool `json:"status"`
	Data   struct {
		Success int `json:"success"`
	} `json:"data"`
	Message string `json:"message"`
}
