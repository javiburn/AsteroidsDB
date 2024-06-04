package models

type ExampleResponse struct {
  Id          string      `json:"id" form:"id" query:"id" validate:"required"`
  Name          string      `json:"name" form:"name" query:"name" validate:"required"`
  Diameter      int         `json:"diameter" form:"diameter" query:"diameter" validate:"required"`
  Discovery_date string      `json:"discovery_date" form:"discovery_date" query:"discovery_date" validate:"required"`
  Observations  string      `json:"observations" form:"observations" query:"observations"`
  Distances     []struct {
    Date     string  `json:"date"`
    Distance float64 `json:"distance"`
  } `json:"distances"`
}

type ExampleRequest struct {
  Name          string      `json:"name" form:"name" query:"name" validate:"required"`
  Diameter      int         `json:"diameter" form:"diameter" query:"diameter" validate:"required"`
  Discovery_date string      `json:"discovery_date" form:"discovery_date" query:"discovery_date" validate:"required"`
  Observations  string      `json:"observations" form:"observations" query:"observations"`
  Distances     []struct {
      Date     string  `json:"date"`
      Distance float64 `json:"distance"`
    } `json:"distances"`
  }
