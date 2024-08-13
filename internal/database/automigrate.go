package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// AutoMigrate creates or updates the database schema based on the provided models.
func AutoMigrate(db *sql.DB, models ...interface{}) error {
	for _, model := range models {
		tableName := getTableName(model)
		fields := getFields(model)

		// Check if table already exists and get existing columns
		existingFields, err := getExistingColumns(db, tableName)
		if err != nil {
			return fmt.Errorf("failed to check existing table %s: %w", tableName, err)
		}

		if len(existingFields) == 0 {
			// Table does not exist, create it
			query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, fields)
			if _, err := db.Exec(query); err != nil {
				return fmt.Errorf("failed to auto migrate table %s: %w", tableName, err)
			}
		} else {
			// Table exists, check for missing columns
			for _, field := range strings.Split(fields, ", ") {
				columnName := strings.Split(field, " ")[0]
				if _, exists := existingFields[columnName]; !exists {
					// Add missing column
					query := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", tableName, field)
					if _, err := db.Exec(query); err != nil {
						return fmt.Errorf("failed to add column %s to table %s: %w", columnName, tableName, err)
					}
				}
			}
		}
	}
	return nil
}

// getTableName converts the model's struct name to a snake_case table name.
func getTableName(model interface{}) string {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	return strings.ToLower(modelType.Name()) + "s"
}

// getFields generates the SQL columns definition from the model's struct fields.
func getFields(model interface{}) string {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	var fields []string

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		// Skip if the field is a struct or a pointer to a struct (likely a relation)
		if field.Type.Kind() == reflect.Struct || (field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) {
			continue
		}

		// Check for json tag
		columnName := field.Tag.Get("json")
		if columnName == "" || columnName == "-" {
			continue // Ignore fields without json tag or explicitly ignored fields
		}

		columnType := getSQLType(field.Type)
		fields = append(fields, fmt.Sprintf("%s %s", columnName, columnType))
	}

	return strings.Join(fields, ", ")
}

// getSQLType maps Go types to SQL types.
func getSQLType(fieldType reflect.Type) string {
	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}

	switch fieldType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "INTEGER"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "INTEGER"
	case reflect.Float32, reflect.Float64:
		return "REAL"
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.String:
		return "TEXT"
	case reflect.Struct:
		if fieldType == reflect.TypeOf(time.Time{}) {
			return "TIMESTAMP"
		}
	}

	return "TEXT"
}

// getExistingColumns retrieves the columns of an existing table in PostgreSQL.
func getExistingColumns(db *sql.DB, tableName string) (map[string]bool, error) {
	query := fmt.Sprintf(`
		SELECT column_name 
		FROM information_schema.columns 
		WHERE table_name = '%s';`, tableName)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := make(map[string]bool)
	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			return nil, err
		}
		columns[columnName] = true
	}
	return columns, nil
}
