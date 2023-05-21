package alias_client

import (
	"time"

	tls_client "github.com/bogdanfinn/tls-client"
)

type AliasSession struct {
	Client        tls_client.HttpClient
	Username      string
	Password      string
	LoginResponse *LoginResponse
}
type CashoutResponse struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	AmountCents     string    `json:"amount_cents"`
	Date            time.Time `json:"date"`
	TransactionType string    `json:"transaction_type"`
	CashOut         struct {
		Method                       string  `json:"method"`
		FeeRevenueCents              string  `json:"fee_revenue_cents"`
		TransferFeeRate              float64 `json:"transfer_fee_rate"`
		TransferAmountCents          string  `json:"transfer_amount_cents"`
		FeeNote                      string  `json:"fee_note"`
		ExchangeRate                 float64 `json:"exchange_rate"`
		LocalizedTransferAmountCents struct {
			AmountCents string `json:"amount_cents"`
			Currency    string `json:"currency"`
		} `json:"localized_transfer_amount_cents"`
	} `json:"cash_out"`
}
type LoginResponse struct {
	AuthToken struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    string `json:"expires_in"`
		TokenType    string `json:"token_type"`
	} `json:"auth_token"`
	User struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		Email         string `json:"email"`
		Name          string `json:"name"`
		ReturnAddress struct {
			Name             string `json:"name"`
			Address1         string `json:"address1"`
			City             string `json:"city"`
			State            string `json:"state"`
			PostalCode       string `json:"postal_code"`
			Country          string `json:"country"`
			CountryCode      string `json:"country_code"`
			Phone            string `json:"phone"`
			RegionID         string `json:"region_id"`
			PhoneCountryCode string `json:"phone_country_code"`
		} `json:"return_address"`
		SellerScore                int      `json:"seller_score"`
		InstantPayoutStatus        string   `json:"instant_payout_status"`
		Status                     string   `json:"status"`
		PhoneNumber                string   `json:"phone_number"`
		BulkListingEligible        bool     `json:"bulk_listing_eligible"`
		EligibleFulfillmentMethods []string `json:"eligible_fulfillment_methods"`
		PhoneVerified              bool     `json:"phone_verified"`
	} `json:"user"`
}
type EarningsResponseSale struct {
	TotalSalesCents                  string  `json:"total_sales_cents"`
	FeesCents                        string  `json:"fees_cents"`
	EarningsCents                    string  `json:"earnings_cents"`
	CommissionCents                  string  `json:"commission_cents"`
	CommissionRate                   float64 `json:"commission_rate"`
	SellerFeeCents                   string  `json:"seller_fee_cents"`
	CashOutFeeCents                  string  `json:"cash_out_fee_cents"`
	CashOutFeeRate                   float64 `json:"cash_out_fee_rate"`
	FinalCashOutAmountCents          string  `json:"final_cash_out_amount_cents"`
	LocalizedFinalCashOutAmountCents struct {
		AmountCents string `json:"amount_cents"`
		Currency    string `json:"currency"`
	} `json:"localized_final_cash_out_amount_cents"`
}
type SalesResponse struct {
	Items []Item `json:"items"`
}
type Item struct {
	ID              string    `json:"id"`
	CreatedDate     time.Time `json:"created_date"`
	Type            string    `json:"type"`
	AmountCents     string    `json:"amount_cents"`
	TransactionType string    `json:"transaction_type"`
}
type Cashout struct {
	AmountUsd       float64 `json:"amount_usd"`
	AmountLocalized float64 `json:"amount_localized"`
	Currency        string  `json:"currency"`
	Fee             float64 `json:"fee"`
	Date            string  `json:"date"`
	ID              string  `json:"id"`
}
type TransactionResponse struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	AmountCents     string    `json:"amount_cents"`
	Date            time.Time `json:"date"`
	TransactionType string    `json:"transaction_type"`
	Sale            struct {
		OrderNumber    string  `json:"order_number"`
		CommissionRate float64 `json:"commission_rate"`
	} `json:"sale"`
}
type EarningsResponse struct {
	AmountCents          string `json:"amount_cents"`
	LocalizedAmountCents struct {
		AmountCents string `json:"amount_cents"`
		Currency    string `json:"currency"`
	} `json:"localized_amount_cents"`
}

type SaleResponse struct {
	PurchaseOrder struct {
		Number                   string    `json:"number"`
		Status                   string    `json:"status"`
		AmountMadeCents          string    `json:"amount_made_cents"`
		Completed                bool      `json:"completed"`
		TakeActionBy             time.Time `json:"take_action_by"`
		CreatedAt                time.Time `json:"created_at"`
		ShipBackCostCents        string    `json:"ship_back_cost_cents"`
		LocalizedAmountMadeCents struct {
			AmountCents string `json:"amount_cents"`
			Currency    string `json:"currency"`
		} `json:"localized_amount_made_cents"`
		Description string `json:"description"`
		Listing     struct {
			Size          float64 `json:"size"`
			ShoeCondition string  `json:"shoe_condition"`
			BoxCondition  string  `json:"box_condition"`
			PriceCents    string  `json:"price_cents"`
			SizeOption    struct {
				Name  string  `json:"name"`
				Value float64 `json:"value"`
			} `json:"size_option"`
			Product struct {
				Sku             string `json:"sku"`
				Name            string `json:"name"`
				Gender          string `json:"gender"`
				HasPicture      bool   `json:"has_picture"`
				SizeUnit        string `json:"size_unit"`
				Type            string `json:"type"`
				PresentationSku struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"presentation_sku"`
				GridPictureURL string `json:"grid_picture_url"`
				MainPictureURL string `json:"main_picture_url"`
			} `json:"product"`
		} `json:"listing"`
		ShippingInfo struct {
			Carrier      string `json:"carrier"`
			TrackingCode string `json:"tracking_code"`
			ServiceLevel string `json:"service_level"`
			TrackingURL  string `json:"tracking_url"`
			LabelURL     string `json:"label_url"`
		} `json:"shipping_info"`
		Progress []struct {
			Completed bool   `json:"completed"`
			Step      string `json:"step"`
			StepEnum  string `json:"step_enum"`
		} `json:"progress"`
	} `json:"purchase_order"`
}
