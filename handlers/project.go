package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	"github.com/lib/pq"

	"github.com/farahsrw/manageProject/config"
	"github.com/farahsrw/manageProject/models"
	"github.com/farahsrw/manageProject/utils"
)

func CreateProject(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var project models.Project

		err := json.NewDecoder(r.Body).Decode(&project)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Error generating token")
			return
		}

		claims := r.Context().Value("claims").(*models.Claims)
		userID := claims.XataID

		var xataID string
		var exists bool
		err = app.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE xata_id = $1)", userID).Scan(&exists)
		if err != nil || !exists {
			log.Printf("User %s does not exist or error checking existence: %v", userID, err)
			utils.RespondError(w, http.StatusBadRequest, "Invalid user")
			return
		}

		err = app.DB.QueryRow(`INSERT INTO projects ("user", name, repo_url, site_url, description, dependencies, dev_dependencies, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING xata_id`, userID, project.Name, project.RepoURL, project.SiteURL, project.Description, pq.Array(project.Dependencies), pq.Array(project.DevDependencies), project.Status).Scan(&xataID)
		if err != nil {
			log.Printf("Error inserting project: %v\n", err)
			utils.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		project.XataID = xataID
		project.UserID = userID
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(project)
	}
}

func UpdateProject(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var project models.Project

		err := json.NewDecoder(r.Body).Decode(&project)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Error generating token")
			return
		}
		vars := mux.Vars(r)
		xataID := vars["xata_id"]

		claims := r.Context().Value("claims").(*models.Claims)
		userID := claims.XataID

		var savedUserID string
		err = app.DB.QueryRow("SELECT \"user\" FROM projects WHERE xata_id=$1", xataID).Scan(&savedUserID)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.RespondError(w, http.StatusNotFound, "Projects not found")
				return
			}
			utils.RespondError(w, http.StatusInternalServerError, "Error fetching projects")
			return
		}

		if savedUserID != userID {
			utils.RespondError(w, http.StatusForbidden, "You do not have permission to update this project")
			return
		}

		_, err = app.DB.Exec("UPDATE projects SET name=$1, repo_url=$2, site_url=$3, description=$4, dependencies=$5, dev_dependencies=$6, status=$7 WHERE xata_id=$8 AND \"user\"=$9", project.Name, project.RepoURL, project.SiteURL, project.Description, pq.Array(project.Dependencies), pq.Array(project.DevDependencies), project.Status, xataID, userID)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error updating project")
			return
		}

		project.XataID = xataID
		project.UserID = userID

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(project)
	}
}

func GetProjects(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*models.Claims)
		userID := claims.XataID

		rows, err := app.DB.Query("SELECT xata_id, \"user\", name, repo_url, site_url, description, dependencies, dev_dependencies, status FROM projects WHERE \"user\" = $1", userID)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error fetching project")
			return
		}
		defer rows.Close()

		var projects []models.Project
		for rows.Next() {
			var project models.Project
			var dependencies, devDependencies []string

			err = rows.Scan(&project.XataID, &project.UserID, &project.Name, &project.RepoURL, &project.SiteURL, &project.Description, pq.Array(&dependencies), pq.Array(&devDependencies), &project.Status)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Error scanning project")
				return
			}

			project.Dependencies = dependencies
			project.DevDependencies = devDependencies
			projects = append(projects, project)
		}

		err = rows.Err()
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error fetching project")
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(projects)
	}
}

func GetProject(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		xataID := vars["xata_id"]

		claims := r.Context().Value("claims").(*models.Claims)
		userID := claims.XataID

		var project models.Project
		var dependencies, devDependencies []string

		err := app.DB.QueryRow("SELECT xata_id, \"user\", name, repo_url, site_url, description, dependencies,dev_dependencies, status FROM projects WHERE xata_id=$1 AND \"user\"=$2", xataID, userID).Scan(&project.XataID, &project.UserID, &project.Name, &project.RepoURL, &project.SiteURL, &project.Description, pq.Array(&dependencies), pq.Array(&devDependencies), &project.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.RespondError(w, http.StatusNotFound, "Project not found")
				return
			}
			utils.RespondError(w, http.StatusInternalServerError, "Error fetching project")
			return
		}
		project.Dependencies = dependencies
		project.DevDependencies = devDependencies

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(project)
	}
}

func DeleteProject(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		xataID := vars["xata_id"]

		claims := r.Context().Value("claims").(*models.Claims)
		userID := claims.XataID

		var savedUserID string
		err := app.DB.QueryRow("SELECT \"user\" FROM projects WHERE xata_id=$1", xataID).Scan(&savedUserID)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.RespondError(w, http.StatusNotFound, "Project not found")
				return
			}
			utils.RespondError(w, http.StatusInternalServerError, "Error fetching project")
			return
		}

		if savedUserID != userID {
			utils.RespondError(w, http.StatusForbidden, "You do not have permission to delete this project")
			return
		}

		_, err = app.DB.Exec("DELETE FROM projects WHERE xata_id=$1 AND \"user\"=$2", xataID, userID)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error deleting project")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
