package product

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/constant"
	"github.com/zakiyalmaya/online-store/model"
)

func getProductParam(ctx *fiber.Ctx) (*model.GetProductRequest, error) {
	categoryID := ctx.Query("category_id")
	limit := ctx.Query("limit")
	page := ctx.Query("page")

	var categoryIDInt *int
	var limitInt, pageInt int
	if categoryID != "" {
		categoryIDParsed, err := strconv.Atoi(categoryID)
		if err != nil {
			return nil, fmt.Errorf("invalid category id")
		}

		categoryIDInt = &categoryIDParsed
	}

	if limit != "" {
		limitParsed, err := strconv.Atoi(limit)
		if err != nil {
			return nil, fmt.Errorf("invalid limit")
		}

		limitInt = limitParsed
	}

	if page != "" {
		pageParsed, err := strconv.Atoi(page)
		if err != nil {
			return nil, fmt.Errorf("invalid page")
		}

		pageInt = pageParsed
	}

	// set default value
	if limit == "" {
		limitInt = constant.DefaultLimit
	}
	
	if page == "" {
		pageInt = constant.DefaultPage
	}

	return &model.GetProductRequest{
		CategoryID: categoryIDInt,
		Limit:      limitInt,
		Page:       pageInt,
	}, nil
}
