package handler

import (
  "context"
  "database/sql"
  "net/http"
  "strconv"

  "github.com/labstack/echo/v4"
  "postgres-go-echo-htmx-bulma/pkg/db"
  "github.com/jackc/pgx/v5/pgxpool"
)

// Hero represents the Hero model.
type Hero struct {
  ID   *int    `json:"id,omitempty"`
  Name string `json:"name"`
}

//  ListHeroHandler handles the listing of all heroes.
//	@Summary		List all heroes
//	@Description	Lists all heroes in the database
//	@Produce		json
//	@Success		200	{array}		handler.Hero
//	@Failure		500	{object}	map[string]string
//	@Router			/heroes   [get]
func ListHeroHandler(dbpool *pgxpool.Pool) echo.HandlerFunc {
  return func(c echo.Context) error {
    ctx := context.Background()
    heroes, err := db.New(dbpool).ListHeros(ctx)
    if err != nil {
      return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to list heroes"})
    }

    return c.JSON(http.StatusOK, heroes)
  }
}

// CreateHeroHandler handles the creation of a new hero.
//	@Summary		Create a new hero
//	@Description	Creates a new hero in the database
//	@Param			hero	body	handler.Hero	true	"Hero information"
//	@Produce		json
//	@Success		201	{object}	handler.Hero
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/heroes   [post]
func CreateHeroHandler(dbpool *pgxpool.Pool) echo.HandlerFunc {
  return func(c echo.Context) error {
    var hero Hero
    if err := c.Bind(&hero); err != nil {
      return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
    }

    ctx := context.Background()

    // TODO: try to DRY this part
    if hero.ID == nil {
      result, err := db.New(dbpool).CreateHero(ctx, hero.Name)
      if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create hero"})
      }
      return c.JSON(http.StatusCreated, result)
    } else {
      arg := db.CreateHeroWithIDParams {
        ID: int32(*hero.ID),
        Name: hero.Name,
      }
      result, err := db.New(dbpool).CreateHeroWithID(ctx, arg)
      if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create hero"})
      }
      return c.JSON(http.StatusCreated, result)
    }
  }
}

// GetHeroHandler retrieves a hero by ID.
//	@Summary		Get a hero by ID
//	@Description	Retrieves a hero from the database by its ID
//	@Param			id	path	int	true	"Hero ID"
//	@Produce		json
//	@Success		200	{object}	handler.Hero
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/heroes/{id}   [get]
func GetHeroHandler(dbpool *pgxpool.Pool) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
      return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID in the request"})
    }

    ctx := context.Background()
    result, err := db.New(dbpool).GetHero(ctx, int32(id))
    if err == sql.ErrNoRows {
      return c.JSON(http.StatusNotFound, map[string]string{"error": "Hero not found"})
    } else if err != nil {
      return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch hero"})
    }

    return c.JSON(http.StatusOK, result)
  }
}

// UpdateHeroHandler updates a hero by ID.
//	@Summary		Update a hero
//	@Description	Updates a hero in the database by its ID
//	@Param			id		path	int				true	"Hero ID"
//	@Param			hero	body	handler.Hero	true	"Hero information"
//	@Produce		json
//	@Success		200	{object}	map[string]string	//	Success		message	can			be	more		descriptive	here	(e.g., "Hero updated successfully")
//	@Failure		400	{object}	map[string]string	//	Detail		error	messages	for	bad			requests
//	@Failure		404	{object}	map[string]string	//	Specific	error	message		for	not			found
//	@Failure		500	{object}	map[string]string	//	Generic		error	message		for	internal	errors
//	@Router			/heroes/{id}   [put]
func UpdateHeroHandler(dbpool *pgxpool.Pool) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
      return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID in the request"})
    }

    var updatedHero Hero
    if err := c.Bind(&updatedHero); err != nil {
      return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
    }

    ctx := context.Background()
    arg := db.UpdateHeroParams{
      ID:   int32(id),
      Name: updatedHero.Name,
    }
    err = db.New(dbpool).UpdateHero(ctx, arg)
    if err == sql.ErrNoRows {
      return c.JSON(http.StatusNotFound, map[string]string{"error": "Hero not found"})
    } else if err != nil {
      return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update hero"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Hero updated successfully"})
  }
}

// DeleteHeroHandler deletes a hero by ID.
//	@Summary		Delete a hero
//	@Description	Deletes a hero from the database by its ID
//	@Param			id	path	int	true	"Hero ID"
//	@Produce		json
//	@Success		204
//	@Failure		400	{object}	map[string]string	//	Detail		error	messages	for	bad			requests
//	@Failure		404	{object}	map[string]string	//	Specific	error	message		for	not			found
//	@Failure		500	{object}	map[string]string	//	Generic		error	message		for	internal	errors
//	@Router			/heroes/{id}   [delete]
func DeleteHeroHandler(dbpool *pgxpool.Pool) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
      return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID in the request"})
    }

    ctx := context.Background()
    err = db.New(dbpool).DeleteHero(ctx, int32(id))
    if err == sql.ErrNoRows {
      return c.JSON(http.StatusNotFound, map[string]string{"error": "Hero not found"})
    } else if err != nil {
      return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete hero"})
    }

    return c.NoContent(http.StatusNoContent)
  }
}

