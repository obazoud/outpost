// Known Limitations/Further Improvements:
// Embedded Structs (AST): Fields from anonymously embedded structs are not fully resolved during AST parsing for documentation as part of the parent struct.
// Complex Slice/Map YAML Formatting: Default value formatting for slices is basic. Maps are not explicitly formatted for YAML beyond their default string representation.

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/hookdeck/outpost/internal/config" // Import your project's config package
)

var (
	inputDir   string
	outputFile string
)

func main() {
	// Default paths are relative to the project root, where `go run cmd/configdocsgen/main.go` is expected to be executed.
	defaultInputDir := "internal/config"
	defaultOutputFile := "docs/pages/references/configuration.mdx"

	flag.StringVar(&inputDir, "input-dir", defaultInputDir, "Directory containing the Go configuration source files.")
	flag.StringVar(&outputFile, "output-file", defaultOutputFile, "Path to the output Markdown file.")
	flag.Parse()

	fmt.Println("Configuration Documentation Generator")
	log.Printf("Input directory: %s", inputDir)
	log.Printf("Output file: %s", outputFile)

	// TODO:
	// 2. Implement parsing of Go files in inputDir
	//    - Use go/parser and go/ast
	// 3. Identify config structs
	// 4. Extract field information (name, yaml, env, desc, required tags, data type)
	// 5. Instantiate config.Config and call InitDefaults()
	//    - Use reflection to get default values
	// 6. Handle "one-of" types (e.g., MQsConfig)
	// 7. Generate Markdown:
	//    - Environment Variables Table
	//    - YAML Configuration Section
	// 8. Write to output file

	parsedConfigs, err := parseConfigFiles(inputDir)
	if err != nil {
		log.Fatalf("Error parsing config files: %v", err)
	}

	// For now, just print what was found (later this will be processed)
	for _, cfg := range parsedConfigs {
		log.Printf("Found config struct: %s in file %s", cfg.Name, cfg.FileName)
	}

	// Attempt to get default values and integrate them BEFORE generating docs
	log.Println("Attempting to load and reflect on config.Config for default values...")
	defaults, err := getConfigDefaults()
	if err != nil {
		log.Printf("Warning: Could not get config defaults: %v", err)
		log.Println("Default values will be missing or incorrect in the generated documentation.")
	} else {
		log.Printf("Successfully reflected on config.Config. Total default keys found: %d", len(defaults))
		// Integrate defaults into parsedConfigs
		integrateDefaults(parsedConfigs, defaults) // This modifies parsedConfigs in place
	}

	// Now generate docs with populated defaults
	err = generateDocs(parsedConfigs, outputFile)
	if err != nil {
		log.Fatalf("Error generating docs: %v", err)
	}

	fmt.Printf("Successfully generated documentation to %s\n", outputFile)
}

// ConfigField represents a field in a configuration struct
type ConfigField struct {
	Name         string
	Type         string
	YAMLName     string
	EnvName      string
	EnvSeparator string
	Description  string
	Required     string // Y, N, C
	DefaultValue string
}

// ParsedConfig represents a parsed configuration struct
type ParsedConfig struct {
	FileName     string
	Name         string // Go struct name
	Fields       []ConfigField
	IsTopLevel   bool   // Flag to identify the root config.Config struct for pathing
	GoStructName string // To help build paths for defaults map
	// We might need to store the original ast.StructType for default value resolution later
}

func parseConfigFiles(dirPath string) ([]ParsedConfig, error) {
	var configs []ParsedConfig
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dirPath, func(fi os.FileInfo) bool {
		return !fi.IsDir() && strings.HasSuffix(fi.Name(), ".go") && !strings.HasSuffix(fi.Name(), "_test.go")
	}, parser.ParseComments)

	if err != nil {
		return nil, fmt.Errorf("failed to parse directory %s: %w", dirPath, err)
	}

	for _, pkg := range pkgs {
		for fileName, file := range pkg.Files {
			log.Printf("Processing file: %s", fileName)
			ast.Inspect(file, func(n ast.Node) bool {
				ts, ok := n.(*ast.TypeSpec)
				if !ok || ts.Type == nil {
					return true // Continue traversal
				}

				s, ok := ts.Type.(*ast.StructType)
				if !ok {
					return true // Continue traversal
				}

				// Found a struct
				structName := ts.Name.Name
				log.Printf("Found struct: %s in file %s", structName, fileName)
				isTopLevel := (structName == "Config") // Simple check for the main Config

				var fields []ConfigField
				for _, field := range s.Fields.List {
					if len(field.Names) > 0 { // Ensure it's a named field
						fieldName := field.Names[0].Name
						var fieldTypeStr string
						switch typeExpr := field.Type.(type) {
						case *ast.Ident:
							fieldTypeStr = typeExpr.Name
						case *ast.SelectorExpr: // For types like pkg.Type
							if pkgIdent, ok := typeExpr.X.(*ast.Ident); ok {
								fieldTypeStr = pkgIdent.Name + "." + typeExpr.Sel.Name
							} else {
								fieldTypeStr = fmt.Sprintf("%s", field.Type) // Fallback
							}
						case *ast.ArrayType:
							// Further inspect Elt for element type if needed
							if eltIdent, ok := typeExpr.Elt.(*ast.Ident); ok {
								fieldTypeStr = "[]" + eltIdent.Name
							} else {
								fieldTypeStr = fmt.Sprintf("%s", field.Type) // Fallback for complex array/slice types
							}
						case *ast.MapType:
							// Further inspect Key and Value for their types
							fieldTypeStr = fmt.Sprintf("%s", field.Type) // Fallback for complex map types
						default:
							fieldTypeStr = fmt.Sprintf("%s", field.Type) // Fallback
						}

						var yamlName, envName, envSeparator, description, requiredValue string

						if field.Tag != nil {
							tagValue := field.Tag.Value // Raw string like `yaml:"name" env:"VAR"`
							// Remove backticks
							tagValue = strings.Trim(tagValue, "`")

							// The line below was unused and caused a compiler error.
							// tags := strings.Fields(tagValue)
							// A more robust way is to use reflect.StructTag, but that requires an instance.
							// For AST parsing, manual parsing or a dedicated tag parser is common.
							// Let's do a simplified manual parse for now.

							parsedTag := parseStructTag(tagValue)
							yamlName = parsedTag["yaml"]
							envName = parsedTag["env"]
							envSeparator = parsedTag["envSeparator"]
							description = parsedTag["desc"]
							requiredValue = parsedTag["required"]
						}

						// TODO: Handle embedded structs (field.Names would be empty)

						log.Printf("  Field: %s, Type: %s, YAML: '%s', Env: '%s', Desc: '%s', Required: '%s'",
							fieldName, fieldTypeStr, yamlName, envName, description, requiredValue)

						fields = append(fields, ConfigField{
							Name:         fieldName,
							Type:         fieldTypeStr,
							YAMLName:     yamlName,
							EnvName:      envName,
							EnvSeparator: envSeparator,
							Description:  description,
							Required:     requiredValue,
						})
					} else {
						// This could be an embedded struct
						// field.Type would give its type.
						// Example: `MyEmbeddedStruct` or `pkg.MyEmbeddedStruct`
						// We need to decide how to represent these.
						// For now, we are only processing fields with names.
						// If an embedded struct's fields should be "flattened" into the parent,
						// this logic needs to recursively call a part of parseConfigFiles or similar.
						// The current reflection for defaults *will* see flattened fields.
						// The AST parsing needs to align if we want to document them as part of the parent.
						if typeIdent, ok := field.Type.(*ast.Ident); ok {
							log.Printf("  Found embedded-like field (no name): Type: %s", typeIdent.Name)
							// Here, you might look up `typeIdent.Name` in `allConfigs` (if populated first pass)
							// and then merge its fields. This gets complex with ordering and paths.
							// For now, we'll skip documenting fields of embedded structs directly here,
							// assuming they are separate `ParsedConfig` entries if they are config structs themselves.
						}
					}
				}
				configs = append(configs, ParsedConfig{
					FileName:     filepath.Base(fileName),
					Name:         structName, // This is the Go struct name
					GoStructName: structName,
					Fields:       fields,
					IsTopLevel:   isTopLevel,
				})
				return false // Stop traversal for this struct, already processed
			})
		}
	}
	return configs, nil
}

// parseStructTag parses a struct tag string and returns a map of key-value pairs.
// It handles quoted values that may contain spaces.
func parseStructTag(tagStr string) map[string]string {
	tags := make(map[string]string)
	for tagStr != "" {
		// Skip leading spaces.
		i := 0
		for i < len(tagStr) && tagStr[i] == ' ' {
			i++
		}
		tagStr = tagStr[i:]
		if tagStr == "" {
			break
		}

		// Find the key.
		i = 0
		for i < len(tagStr) && tagStr[i] != ' ' && tagStr[i] != ':' {
			i++
		}
		if i == 0 || i+1 >= len(tagStr) || tagStr[i] != ':' || tagStr[i+1] != '"' {
			// Malformed tag or no value, skip to next potential tag part
			// This might happen if a tag is just `key` without `:"value"`
			// Or if it's not a quoted value. For simplicity, we assume all values are quoted.
			nextSpace := strings.Index(tagStr, " ")
			if nextSpace == -1 {
				break // No more tags
			}
			tagStr = tagStr[nextSpace:]
			continue
		}
		key := tagStr[:i]
		tagStr = tagStr[i+1:] // Skip key and ':'

		// Find the value (quoted).
		if tagStr[0] != '"' {
			// Malformed tag: value not quoted
			nextSpace := strings.Index(tagStr, " ")
			if nextSpace == -1 {
				break
			}
			tagStr = tagStr[nextSpace:]
			continue
		}
		i = 1 // Skip leading quote
		for i < len(tagStr) && tagStr[i] != '"' {
			if tagStr[i] == '\\' { // Handle escaped quotes
				i++
			}
			i++
		}
		if i >= len(tagStr) {
			// Malformed tag: unclosed quote
			break
		}
		value := tagStr[1:i]
		tags[key] = value

		// Move to next tag.
		tagStr = tagStr[i+1:]
	}
	return tags
}

// getConfigDefaults attempts to instantiate config.Config, run InitDefaults,
// and extract default values using reflection.
func getConfigDefaults() (map[string]interface{}, error) {
	cfg := config.Config{} // Create an instance of the actual Config struct
	cfg.InitDefaults()     // Initialize it with default values

	defaults := make(map[string]interface{})
	extractDefaultsRecursive(reflect.ValueOf(cfg), defaults, "")

	// TODO: Map these defaults back to the ParsedConfig fields,
	// potentially using a path like "Redis.Host" or by matching struct and field names.

	return defaults, nil
}

func extractDefaultsRecursive(val reflect.Value, defaultsMap map[string]interface{}, prefix string) {
	// Dereference pointers if any
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Ensure we are dealing with a struct
	if val.Kind() != reflect.Struct {
		return
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		// Skip unexported fields
		if !fieldVal.CanInterface() {
			continue
		}

		currentPath := field.Name
		if prefix != "" {
			currentPath = prefix + "." + field.Name
		}

		if fieldVal.Kind() == reflect.Struct {
			// Check if it's a time.Time struct or other common struct we don't want to recurse into deeply for this purpose
			// For example, time.Time has unexported fields that would cause issues or are not relevant as "defaults"
			if fieldVal.Type().PkgPath() == "time" && fieldVal.Type().Name() == "Time" {
				defaultsMap[currentPath] = fieldVal.Interface()
			} else {
				// Recurse for nested structs
				extractDefaultsRecursive(fieldVal, defaultsMap, currentPath)
			}
		} else if fieldVal.Kind() == reflect.Slice || fieldVal.Kind() == reflect.Array {
			// Handle slices/arrays - store them directly
			// For complex slice elements (structs), further handling might be needed if defaults per element are relevant
			defaultsMap[currentPath] = fieldVal.Interface()
		} else {
			defaultsMap[currentPath] = fieldVal.Interface()
		}
	}
}

func integrateDefaults(parsedConfigs []ParsedConfig, defaults map[string]interface{}) {
	// Find the top-level "Config" struct first to establish base for paths
	var topLevelConfig *ParsedConfig
	otherConfigs := make(map[string]*ParsedConfig) // Map by GoStructName

	for i := range parsedConfigs {
		pc := &parsedConfigs[i]          // Get a pointer to modify in place
		if pc.GoStructName == "Config" { // Assuming "Config" is the main one
			topLevelConfig = pc
		}
		otherConfigs[pc.GoStructName] = pc
	}

	if topLevelConfig == nil {
		log.Println("Warning: Top-level 'Config' struct not found in parsed files. Cannot accurately map all defaults.")
		return
	}

	// Recursive function to build paths and assign defaults
	var assignDefaults func(currentFields *[]ConfigField, currentPathPrefix string)
	assignDefaults = func(currentFields *[]ConfigField, currentPathPrefix string) {
		for i := range *currentFields {
			field := &(*currentFields)[i] // Pointer to modify the field

			defaultPathKey := field.Name // Default path is just the field name for top-level
			if currentPathPrefix != "" {
				defaultPathKey = currentPathPrefix + "." + field.Name
			}

			if defaultValue, ok := defaults[defaultPathKey]; ok {
				field.DefaultValue = fmt.Sprintf("%v", defaultValue)
				log.Printf("Assigned default for %s: %s", defaultPathKey, field.DefaultValue)
			} else {
				// If not found directly, it might be a nested struct defined elsewhere.
				// The current `extractDefaultsRecursive` stores nested struct fields as "Parent.Field".
				// The AST parsing gives us field types like "RedisConfig". We need to bridge this.
				// For now, this simple pathing works for direct fields and first-level nesting if paths align.
				// More complex mapping might be needed if AST field type (e.g. "RedisConfig") needs to be part of path.
				// The current `extractDefaultsRecursive` builds paths like "Redis.Host", "Redis.Port".
				// Our AST parsing gives us `Config.Redis` (type `RedisConfig`), then `RedisConfig.Host`.
				// We need to ensure the path construction aligns.

				// If field.Type is a known struct type from our parsed configs, recurse.
				// This part needs careful alignment of how paths are constructed in `extractDefaultsRecursive`
				// and how we look them up here.
				// The current `extractDefaultsRecursive` uses Go field names for paths.
				// So, if `Config` has a field `Redis RedisConfig`, the path for redis host is `Redis.Host`.

				// Let's assume field.Type is the Go type name of the struct (e.g., "RedisConfig")
				// And that our `extractDefaultsRecursive` has built paths like "Redis.Host"
				// where "Redis" is the field name in the parent struct (`Config`)
				// and "Host" is the field name in the child struct (`RedisConfig`).

				// If the field itself is a struct type we parsed, we need to recurse into its fields.
				// The `defaultPathKey` for fields of this nested struct would be `ParentField.NestedField`.
				// Example: Config.Redis (field.Name = "Redis"), its type is "RedisConfig".
				// We need to find the ParsedConfig for "RedisConfig" and then iterate its fields.
				// The path prefix for fields of RedisConfig would be "Redis".

				if nestedStructInfo, isStruct := otherConfigs[field.Type]; isStruct {
					log.Printf("Recursing for nested struct field %s (type %s) with path prefix %s", field.Name, field.Type, defaultPathKey)
					// The prefix for the children of this field is the full path to this field.
					assignDefaults(&nestedStructInfo.Fields, defaultPathKey)
				} else {
					log.Printf("No direct default found for %s (path key: %s)", field.Name, defaultPathKey)
				}
			}
		}
	}

	// Start with the top-level config, no prefix for its direct fields
	assignDefaults(&topLevelConfig.Fields, "")

	// After processing top-level, some nested structs might have had their defaults assigned
	// by the recursion above if their parent field was processed.
	// We might need a more robust way to ensure all parsed structs are processed if they can be standalone.
	// For now, this focuses on defaults reachable from `config.Config`.
}

const (
	envVarsStartPlaceholder = "{/* BEGIN AUTOGENERATED CONFIG ENV VARS */}"
	envVarsEndPlaceholder   = "{/* END AUTOGENERATED CONFIG ENV VARS */}"
	yamlStartPlaceholder    = "{/* BEGIN AUTOGENERATED CONFIG YAML */}"
	yamlEndPlaceholder      = "{/* END AUTOGENERATED CONFIG YAML */}"
)

func generateDocs(parsedConfigs []ParsedConfig, outputPath string) error {
	// --- Generate ENV VARS Table Content (including headers) ---
	var envVarsBuilder strings.Builder
	envVarsBuilder.WriteString("| Variable | Description | Default | Required |\n")
	envVarsBuilder.WriteString("|----------|-------------|---------|----------|\n")
	envVarFields := collectEnvVarFields(parsedConfigs)
	for _, field := range envVarFields {
		if field.EnvName == "" {
			continue
		}
		requiredText := formatRequiredText(field.Required)
		defaultValueText := field.DefaultValue
		if defaultValueText == "" || defaultValueText == "<nil>" { // Handle <nil> from Sprintf as well
			defaultValueText = "`nil`"
		} else {
			// Escape special characters for Markdown table cells
			defaultValueText = strings.ReplaceAll(defaultValueText, "|", "\\|")
			// defaultValueText = strings.ReplaceAll(defaultValueText, "{", "\\{")
			// defaultValueText = strings.ReplaceAll(defaultValueText, "}", "\\}")
			// Enclose in backticks if not already `nil`
			defaultValueText = fmt.Sprintf("`%s`", defaultValueText)
		}

		descriptionText := strings.ReplaceAll(field.Description, "|", "\\|")
		descriptionText = strings.ReplaceAll(descriptionText, "\n", " ") // Ensure description is single line for table
		// Escape curly braces for MDX in descriptions
		descriptionText = strings.ReplaceAll(descriptionText, "{", "\\{")
		descriptionText = strings.ReplaceAll(descriptionText, "}", "\\}")

		envVarsBuilder.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s |\n",
			field.EnvName,
			descriptionText,
			defaultValueText,
			requiredText,
		))
	}
	envVarsContent := strings.TrimRight(envVarsBuilder.String(), "\n")
	// --- Generate YAML Content (including fences) ---
	var yamlBuilder strings.Builder
	// The "## YAML" header should be manually placed in the MDX file.
	yamlBuilder.WriteString("```yaml\n")
	yamlBuilder.WriteString("# Outpost Configuration Example (Generated)\n")
	yamlBuilder.WriteString("# This example shows all available keys with their default values where applicable.\n\n")
	var mainConfigInfo *ParsedConfig
	configInfoMap := make(map[string]*ParsedConfig)
	for i := range parsedConfigs {
		pc := &parsedConfigs[i]
		configInfoMap[pc.GoStructName] = pc
		if pc.GoStructName == "Config" {
			mainConfigInfo = pc
		}
	}
	if mainConfigInfo != nil {
		generateYAMLPart(&yamlBuilder, mainConfigInfo, configInfoMap, 0, true)
	} else {
		yamlBuilder.WriteString("# ERROR: Main 'Config' struct not found. Cannot generate YAML structure.\n")
	}
	yamlBuilder.WriteString("```\n") // Add closing fence
	yamlContent := strings.TrimRight(yamlBuilder.String(), "\n")

	// --- Read existing file and replace placeholder sections ---
	existingContentBytes, err := os.ReadFile(outputPath)
	if err != nil {
		log.Printf("Warning: Could not read existing output file %s. Will create a new one. Error: %v", outputPath, err)
		// If file doesn't exist, create it with placeholders and content
		var newFileBuilder strings.Builder
		newFileBuilder.WriteString("---\ntitle: \"Outpost Configuration\"\n---\n\n") // Basic frontmatter
		newFileBuilder.WriteString("<!-- Placeholder for intro content -->\n\n")
		newFileBuilder.WriteString(envVarsStartPlaceholder + "\n")
		newFileBuilder.WriteString(envVarsContent)
		newFileBuilder.WriteString(envVarsEndPlaceholder + "\n\n")
		newFileBuilder.WriteString("<!-- Placeholder for content between sections -->\n\n")
		newFileBuilder.WriteString(yamlStartPlaceholder + "\n")
		newFileBuilder.WriteString(yamlContent)
		newFileBuilder.WriteString(yamlEndPlaceholder + "\n\n")
		newFileBuilder.WriteString("<!-- Placeholder for outro content -->\n")
		return os.WriteFile(outputPath, []byte(newFileBuilder.String()), 0644)
	}

	existingContent := string(existingContentBytes)

	finalContent, envReplaced := replacePlaceholder(existingContent, envVarsStartPlaceholder, envVarsEndPlaceholder, envVarsContent)
	if !envReplaced {
		log.Printf("Warning: ENV VARS placeholders not found in %s. ENV VARS section not updated.", outputPath)
	}

	finalContent, yamlReplaced := replacePlaceholder(finalContent, yamlStartPlaceholder, yamlEndPlaceholder, yamlContent)
	if !yamlReplaced {
		log.Printf("Warning: YAML placeholders not found in %s. YAML section not updated.", outputPath)
	}

	if !envReplaced && !yamlReplaced {
		log.Printf("Neither ENV VARS nor YAML placeholders were found. File %s was not modified.", outputPath)
		return nil // Or return an error if placeholders are mandatory
	}

	return os.WriteFile(outputPath, []byte(finalContent), 0644)
}

func replacePlaceholder(content, startPlaceholder, endPlaceholder, newBlockContent string) (string, bool) {
	startIndex := strings.Index(content, startPlaceholder)
	endIndex := strings.Index(content, endPlaceholder)

	if startIndex != -1 && endIndex != -1 && endIndex > startIndex {
		// Preserve the placeholders themselves, replace content between them
		newContent := content[:startIndex+len(startPlaceholder)] + "\n" + newBlockContent + "\n" + content[endIndex:]
		return newContent, true
	}
	return content, false // Placeholders not found or in wrong order
}

func generateYAMLPart(builder *strings.Builder, configInfo *ParsedConfig, allConfigs map[string]*ParsedConfig, indentLevel int, isRoot bool) {
	indent := strings.Repeat("  ", indentLevel)

	// Special handling for MQsConfig and PublishMQConfig to show "one of"
	if configInfo.GoStructName == "MQsConfig" || configInfo.GoStructName == "PublishMQConfig" {
		builder.WriteString(fmt.Sprintf("%s# Choose one of the following MQ providers:\n", indent))
	}

	// Sort fields by YAML name for consistent output
	sortedFields := make([]ConfigField, len(configInfo.Fields))
	copy(sortedFields, configInfo.Fields)
	sort.Slice(sortedFields, func(i, j int) bool {
		if sortedFields[i].YAMLName == "" {
			return false
		}
		if sortedFields[j].YAMLName == "" {
			return true
		}
		return sortedFields[i].YAMLName < sortedFields[j].YAMLName
	})

	for _, field := range sortedFields {
		if field.YAMLName == "" {
			continue
		}

		if field.Description != "" {
			descLines := strings.Split(field.Description, "\n")
			for _, line := range descLines {
				builder.WriteString(fmt.Sprintf("%s# %s\n", indent, line))
			}
		}

		if nestedConfigInfo, ok := allConfigs[field.Type]; ok {
			builder.WriteString(fmt.Sprintf("%s%s:\n", indent, field.YAMLName))
			generateYAMLPart(builder, nestedConfigInfo, allConfigs, indentLevel+1, false)
		} else if strings.HasPrefix(field.Type, "[]") { // Handle slices
			builder.WriteString(fmt.Sprintf("%s%s:\n", indent, field.YAMLName))
			// Attempt to use the actual reflected default value for slices
			// This part is tricky because `field.DefaultValue` is just a string.
			// We need a way to get the actual `reflect.Value` of the default.
			// For now, we'll rely on a placeholder or a very simple interpretation of field.DefaultValue
			// A proper solution would involve passing the `defaults` map (from getConfigDefaults)
			// down to here and looking up the field's default by its path.

			// Let's assume field.DefaultValue is a string like "[]" or "[item1 item2]"
			// This is a simplification.
			if field.DefaultValue == "[]" || field.DefaultValue == "" || field.DefaultValue == "<nil>" {
				builder.WriteString(fmt.Sprintf("%s  [] # Empty list\n", indent))
			} else {
				// Attempt to parse a simple string representation of a slice.
				// This is very basic and will likely need improvement.
				trimmed := strings.Trim(field.DefaultValue, "[]")
				if trimmed != "" {
					elements := strings.Fields(trimmed) // Splits by space, assumes simple elements
					for _, elem := range elements {
						// Try to format element based on assumed type (e.g., string if it's not number/bool)
						elemType := strings.TrimPrefix(field.Type, "[]") // e.g. "string" from "[]string"
						builder.WriteString(fmt.Sprintf("%s  - %s\n", indent, formatSimpleValueToYAML(elem, elemType)))
					}
				} else {
					builder.WriteString(fmt.Sprintf("%s  [] # Empty list from default value: %s\n", indent, field.DefaultValue))
				}
			}
		} else { // Simple field
			defaultValueText := field.DefaultValue
			if defaultValueText == "" || defaultValueText == "<nil>" {
				defaultValueText = getYAMLPlaceholderForType(field.Type)
			} else {
				defaultValueText = formatSimpleValueToYAML(defaultValueText, field.Type)
			}
			builder.WriteString(fmt.Sprintf("%s%s: %s\n", indent, field.YAMLName, defaultValueText))
		}
		if indentLevel == 0 && isRoot {
			builder.WriteString("\n")
		}
	}
}

func getYAMLPlaceholderForType(goType string) string {
	if strings.HasPrefix(goType, "[]") { // Slice
		return "[] # Empty list or no default"
	}
	switch goType {
	case "string":
		return `""`
	case "int", "int64", "int32", "uint", "uint64", "uint32":
		return "0"
	case "bool":
		return "false"
	case "map[string]string", "map[string]interface{}": // Basic map types
		return "{} # Empty map or no default"
	default:
		// For unknown or complex struct types not handled as nested.
		// Check if it's a known config struct type (should have been handled by nestedConfigInfo)
		// If not, it's likely a type we don't have specific YAML for.
		return "null # Check type and provide appropriate default"
	}
}

func formatSimpleValueToYAML(value, goType string) string {
	if value == "<nil>" { // Handle explicit nil from reflection
		return "null"
	}
	if goType == "string" {
		// Ensure strings are quoted, unless they are already clearly a YAML string that doesn't need quotes
		// (e.g. simple alphanumeric, or already quoted). This is a simplification.
		if _, err := strconv.ParseFloat(value, 64); err != nil && value != "true" && value != "false" && value != "null" {
			if !(strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`)) &&
				!(strings.HasPrefix(value, `'`) && strings.HasSuffix(value, `'`)) {
				return fmt.Sprintf(`"%s"`, strings.ReplaceAll(value, `"`, `\"`))
			}
		}
		return value // Return as is if it looks like a number, bool, null, or already quoted
	}
	// For non-string simple types, assume DefaultValue is already a good representation
	return value
}

// collectEnvVarFields flattens all fields that have an environment variable name.
// This helps in case fields are defined across multiple ParsedConfig structs but should be in one table.
func collectEnvVarFields(parsedConfigs []ParsedConfig) []ConfigField {
	var allFields []ConfigField
	seenEnvVars := make(map[string]bool) // To avoid duplicates if somehow an env var is on multiple fields

	// It's better to iterate based on the main config structure to get a somewhat logical order
	// For now, a simple iteration. Order might need refinement.
	for _, pc := range parsedConfigs {
		for _, field := range pc.Fields {
			if field.EnvName != "" && !seenEnvVars[field.EnvName] {
				allFields = append(allFields, field)
				seenEnvVars[field.EnvName] = true
			}
		}
	}
	// Sort fields alphabetically by EnvName for consistent output
	sort.Slice(allFields, func(i, j int) bool {
		return allFields[i].EnvName < allFields[j].EnvName
	})
	return allFields
}

func formatRequiredText(reqStatus string) string {
	switch reqStatus {
	case "Y":
		return "Yes"
	case "N":
		return "No"
	case "C":
		return "Conditional (see description)"
	default:
		return reqStatus // Or "Unknown"
	}
}
