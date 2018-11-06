package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"

	mid "github.com/coderbunker/heikenet-backend/middleware"
	"github.com/coderbunker/heikenet-backend/models"
)

func CreateProfile(c echo.Context) error {
	new_profile := new(models.NewProfile)
	if err := c.Bind(new_profile); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "can't read json",
		})
	}

	db, err := mid.GetDB(c)
	if err != nil {
		log.Fatal(err)
	}

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "can't read id",
		})
	}

	profile, err := models.CreateProfile(db, id, new_profile)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusCreated, map[string]string{
			"id": profile.ID,
		})
	}
}

func GetProfile(c echo.Context) error {
	id, err := uuid.FromString(c.Param("pid"))
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "can't read id",
		})
	}

	db, err := mid.GetDB(c)
	if err != nil {
		log.Fatal(err)
	}

	profile, err := models.GetProfile(db, id)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, profile)
	}
}

func UpdateProfile(c echo.Context) error {
	new_profile := new(models.NewProfile)
	if err := c.Bind(new_profile); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "can't read json",
		})
	}

	id, err := uuid.FromString(c.Param("pid"))
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "can't read id",
		})
	}

	db, err := mid.GetDB(c)
	if err != nil {
		log.Fatal(err)
	}

	profile, err := models.UpdateProfile(db, id, new_profile)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, map[string]string{
			"id": profile.ID,
		})
	}
}

func DeleteProfile(c echo.Context) error {
	id, err := uuid.FromString(c.Param("pid"))
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "can't read id",
		})
	}

	db, err := mid.GetDB(c)
	if err != nil {
		log.Fatal(err)
	}

	err = models.DeleteProfile(db, id)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})
	} else {
		return c.NoContent(http.StatusNoContent)
	}
}
