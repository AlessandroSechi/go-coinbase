package coinbase

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	APIBase = "https://api.coinbase.com/v2"
)

type (

	Address struct {
		ID					string			`json:"id,omitempty"`
		Address				string			`json:"address,omitempty"`
		Name				string			`json:"name,omitempty"`
		Network				string			`json:"network,omitempty"`
		CreatedAt			time.Time		`json:"created_at,omitempty"`
		UpdatedAt			time.Time		`json:"updated_at,omitempty"`
		Resource			string			`json:"resource,omitempty"`
		ResourcePath		string			`json:"resource_path,omitempty"`
	}

	Account struct {
		ID					string			`json:"id,omitempty"`
		Name				string			`json:"name,omitempty"`
		Primary				bool			`json:"primary,omitempty"`
		Type				string			`json:"type,omitempty"`
		Currency			string			`json:"currency,omitempty"`
		Balance				string			`json:"balance,omitempty"`
		CreatedAt			time.Time		`json:"created_at,omitempty"`
		UpdatedAt			time.Time		`json:"updated_at,omitempty"`
		Resource			string			`json:"resource,omitempty"`
		ResourcePath		string			`json:"resource_path,omitempty"`
	}

	Buy struct {
		ID					string			`json:"id,omitempty"`
		Status				string			`json:"status,omitempty"`
		PaymentMethod struct {
			ID				string			`json:"id,omitempty"`
			Resource		string			`json:"resource,omitempty"`
			ResourcePath	string			`json:"resource_path,omitempty"`
		}									`json:"payment_method,omitempty"`
		Transaction struct {
			ID				string			`json:"id,omitempty"`
			Resource		string			`json:"resource,omitempty"`
			ResourcePath	string			`json:"resource_path,omitempty"`
		}									`json:"transaction,omitempty"`
		Amount struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"amount,omitempty"`
		Total struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"total,omitempty"`
		SubTotal struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"subtotal,omitempty"`
		Fee struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"fee,omitempty"`
		CreatedAt			time.Time		`json:"created_at,omitempty"`
		UpdatedAt			time.Time		`json:"updated_at,omitempty"`
		PayoutAt			time.Time		`json:"payout_at,omitempty"`
		Resource			string			`json:"resource,omitempty"`
		ResourcePath		string			`json:"resource_path,omitempty"`
		Committed			bool			`json:"committed,omitempty"`
		Instant				bool			`json:"instant,omitempty"`
	}

	// Client represents a Coinbase REST API Client
	Client struct {
		HTTPClient           *http.Client
		APIKey				 string
		APISecret			 string
		APIBase              string
		Log                  io.Writer 	// If set, all request will be logged there
		Localization		 string		// Preferred language for Error Messages
	}

	CreateAddress struct {
		Name				string				  `json:"name,omitempty"`
	}

	Currency struct {
		ID					string				  `json:"id,omitempty"`
		Name				string				  `json:"name,omitempty"`
		MinSize				string				  `json:"min_size,omitempty"`
	}

	Deposit struct {
		ID					string			`json:"id,omitempty"`
		Status				string			`json:"status,omitempty"`
		PaymentMethod struct {
			ID				string			`json:"id,omitempty"`
			Resource		string			`json:"resource,omitempty"`
			ResourcePath	string			`json:"resource_path,omitempty"`
		}									`json:"payment_method,omitempty"`
		Transaction struct {
			ID				string			`json:"id,omitempty"`
			Resource		string			`json:"resource,omitempty"`
			ResourcePath	string			`json:"resource_path,omitempty"`
		}									`json:"transaction,omitempty"`
		Amount struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"amount,omitempty"`
		SubTotal struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"subtotal,omitempty"`
		Fee struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"fee,omitempty"`
		CreatedAt			time.Time		`json:"created_at,omitempty"`
		UpdatedAt			time.Time		`json:"updated_at,omitempty"`
		PayoutAt			time.Time		`json:"payout_at,omitempty"`
		Resource			string			`json:"resource,omitempty"`
		ResourcePath		string			`json:"resource_path,omitempty"`
		Committed			bool			`json:"committed,omitempty"`
	}

	DepositFunds struct {
		Amount							string			`json:"amount,omitempty"`
		Currency						string			`json:"currency,omitempty"`
		PaymentMethod					string			`json:"payment_method,omitempty"`
		Commit							bool			`json:"commit,omitempty"`
	}

	// ErrorResponse represents a Coinbase REST API Error Response
	ErrorResponse struct {
		Response        	*http.Response        `json:"-,omitempty"`
		Errors          	[]Errors              `json:"errors,omitempty"`
	}

	// Error represents a Coinbase REST API Error Detail
	Errors struct {
		ID           		string                `json:"id,omitempty"`
		Message      		string                `json:"message,omitempty"`
		Url			 		string                `json:"url,omitempty"`
	}

	ExchangeRates struct {
		Currency 			string				  `json:"currency,omitempty"`
		Rates				map[string]string	  `json:"rate,omitempty"`
	}

	From struct {
		ID					string				  `json:"id,omitempty"`
		Resource			string				  `json:",omitempty"`
	}

	Network struct {
		Status				string				  `json:"status,omitempty"`
		Name				string                `json:"name,omitempty"`
	}

	Pagination struct {
		EndingBefore		string				 `json:"ending_before,omitempty"`
		StartingAfter		string				 `json:"starting_after,omitempty"`
		Limit				uint8				 `json:"limit,omitempty"`
		Order				string				 `json:"order,omitempty"`
		PreviousUri			string				 `json:"previous_uri,omitempty"`
		NextUri				string				 `json:"next_uri,omitempty"`
	}

	PlaceBuy struct {
		Amount							string			`json:"amount,omitempty"`
		Total							string			`json:"total,omitempty"`
		Currency						string			`json:"currency,omitempty"`
		PaymentMethod					string			`json:"payment_method,omitempty"`
		AgreeBtcAmountVaries			bool			`json:"agree_btc_amount_varies,omitempty"`
		Commit							bool			`json:"commit,omitempty"`
		Quote							bool			`json:"quote,omitempty"`
	}

	PaymentMethod struct {
		ID								string			`json:"id,omitempty"`
		Type							string			`json:"type,omitempty"`
		Name							string			`json:"name,omitempty"`
		Currency						string			`json:"currency,omitempty"`
		PrimaryBuy						bool			`json:"primary_buy,omitempty"`
		PrimarySell						bool			`json:"primary_sell,omitempty"`
		AllowBuy						bool			`json:"allow_buy,omitempty"`
		AllowSell						bool			`json:"allow_sell,omitempty"`
		InstantBuy						bool			`json:"instant_buy,omitempty"`
		InstantSell						bool			`json:"instant_sell,omitempty"`
		CreatedAt						time.Time		`json:"created_at,omitempty"`
		UpdatedAt						time.Time		`json:"updated_at,omitempty"`
		Resource						string			`json:"resource,omitempty"`
		ResourcePath					string			`json:"resource_path,omitempty"`
	}

	PlaceSell struct {
		Amount							string			`json:"amount,omitempty"`
		Total							string			`json:"total,omitempty"`
		Currency						string			`json:"currency,omitempty"`
		PaymentMethod					string			`json:"payment_method,omitempty"`
		AgreeBtcAmountVaries			bool			`json:"agree_btc_amount_varies,omitempty"`
		Commit							bool			`json:"commit,omitempty"`
		Quote							bool			`json:"quote,omitempty"`
	}

	Price struct {
		Amount							string			`json:"amount,omitempty"`
		Currency						string			`json:"currency,omitempty"`
	}

	Response struct {
		Pagination			interface{}			  `json:"pagination"`
		Data				interface{}			  `json:"data"`
	}

	RequestMoney struct {
		Type							string			`json:"type,omitempty"`
		To								string			`json:"to,omitempty"`
		Amount							string			`json:"amount,omitempty"`
		Currency						string			`json:"currency,omitempty"`
		Description						string			`json:"description,omitempty"`
	}

	Sell struct {
		ID					string			`json:"id,omitempty"`
		Status				string			`json:"status,omitempty"`
		PaymentMethod struct {
			ID				string			`json:"id,omitempty"`
			Resource		string			`json:"resource,omitempty"`
			ResourcePath	string			`json:"resource_path,omitempty"`
		}									`json:"payment_method,omitempty"`
		Transaction struct {
			ID				string			`json:"id,omitempty"`
			Resource		string			`json:"resource,omitempty"`
			ResourcePath	string			`json:"resource_path,omitempty"`
		}									`json:"transaction,omitempty"`
		Amount struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"amount,omitempty"`
		Total struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"total,omitempty"`
		SubTotal struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"subtotal,omitempty"`
		Fee struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"fee,omitempty"`
		CreatedAt			time.Time		`json:"created_at,omitempty"`
		UpdatedAt			time.Time		`json:"updated_at,omitempty"`
		PayoutAt			time.Time		`json:"payout_at,omitempty"`
		Resource			string			`json:"resource,omitempty"`
		ResourcePath		string			`json:"resource_path,omitempty"`
		Committed			bool			`json:"committed,omitempty"`
		Instant				bool			`json:"instant,omitempty"`
	}

	SendMoney struct {
		Type							string			`json:"type,omitempty"`
		To								string			`json:"to,omitempty"`
		Amount							string			`json:"amount,omitempty"`
		Currency						string			`json:"currency,omitempty"`
		Description						string			`json:"description,omitempty"`
		SkipNotifications				bool			`json:"skip_notifications,omitempty"`
		Fee								string			`json:"fee,omitempty"`
		Idem							string			`json:"idem,omitempty"`
		ToFinancialInstitution			bool			`json:"to_financial_institution,omitempty"`
		FinancialInstitutionWebsite		string			`json:"financial_institution_website,omitempty"`
	}

	Time struct {
		Iso					string			`json:"iso,omitempty"`
		Epoch				int64			`json:"epoch,omitempty"`
	}

	Transaction struct {
		ID					string			`json:"id,omitempty"`
		Type				string			`json:"type,omitempty"`
		Status				string			`json:"status,omitempty"`
		Amount struct {
			Amount				string		`json:"amount,omitempty"`
			Currency			string		`json:"currency,omitempty"`
		}									`json:"amount,omitempty"`
		NativeAmount struct {
			Amount				string		`json:"amount,omitempty"`
			Currency			string		`json:"currency,omitempty"`
		}									`json:"native_amount,omitempty"`
		Description			string			`json:"description,omitempty"`
		CreatedAt			time.Time		`json:"created_at,omitempty"`
		UpdatedAt			time.Time		`json:"updated_at,omitempty"`
		Resource			string			`json:"resource,omitempty"`
		ResourcePath		string			`json:"resource_path,omitempty"`
		Network				Network			`json:"network,omitempty"`
		From				From			`json:"from,omitempty"`
	}

	TransferMoney struct {
		Type							string			`json:"type,omitempty"`
		To								string			`json:"to,omitempty"`
		Amount							string			`json:"amount,omitempty"`
		Currency						string			`json:"currency,omitempty"`
		Description						string			`json:"description,omitempty"`
	}

	UpdateAccount struct {
		Name				string				  `json:"name,omitempty"`
	}

	UpdateCurrentUser struct {
		Name				string				  `json:"name,omitempty"`
		TimeZone			string				  `json:"time_zone,omitempty"`
		NativeCurrency		string				  `json:"native_currency,omitempty"`
	}

	User struct {
		ID					string				  `json:"id,omitempty"`
		Name				string				  `json:"name,omitempty"`
		Username			string				  `json:"username,omitempty"`
		ProfileLocation		string				  `json:"profile_location,omitempty"`
		ProfileBio			string				  `json:"profile_bio,omitempty"`
		ProfileUrl			string				  `json:"profile_url,omitempty"`
		AvatarUrl			string				  `json:"avatar_url,omitempty"`
		Resource			string				  `json:"resource,omitempty"`
		ResourcePath		string				  `json:"resource_path,omitempty"`
		SendsDisabled		bool				  `json:"sends_disabled,omitempty"`
	}

	Withdrawal struct {
		ID					string			`json:"id,omitempty"`
		Type				string			`json:"type,omitempty"`
		Status				string			`json:"status,omitempty"`
		Amount struct {
			Amount				string		`json:"amount,omitempty"`
			Currency			string		`json:"currency,omitempty"`
		}									`json:"amount,omitempty"`
		PaymentMethod struct {
			ID				string			`json:"id,omitempty"`
			Resource		string			`json:"resource,omitempty"`
			ResourcePath	string			`json:"resource_path,omitempty"`
		}									`json:"payment_method,omitempty"`
		Transaction struct {
			ID				string			`json:"id,omitempty"`
			Resource		string			`json:"resource,omitempty"`
			ResourcePath	string			`json:"resource_path,omitempty"`
		}									`json:"transaction,omitempty"`
		SubTotal struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"subtotal,omitempty"`
		Fee struct {
			Amount			string			`json:"amount,omitempty"`
			Currency		string			`json:"currency,omitempty"`
		}									`json:"fee,omitempty"`
		CreatedAt			time.Time		`json:"created_at,omitempty"`
		UpdatedAt			time.Time		`json:"updated_at,omitempty"`
		PayoutAt			time.Time		`json:"payout_at,omitempty"`
		Resource			string			`json:"resource,omitempty"`
		ResourcePath		string			`json:"resource_path,omitempty"`
		Committed			bool			`json:"committed,omitempty"`
	}

	Withdraw struct {
		Amount							string			`json:"amount,omitempty"`
		Currency						string			`json:"currency,omitempty"`
		PaymentMethod					string			`json:"payment_method,omitempty"`
	}
)

// Error method implementation for ErrorResponse struct
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Errors)
}
