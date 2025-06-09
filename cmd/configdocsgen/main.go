// Known Limitations/Further Improvements:
// Complex Slice/Map YAML Formatting: Default value formatting for slices is basic. Maps are not explicitly formatted for YAML beyond their default string representation.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hookdeck/outpost/internal/config" // Import your project's config package
)

var (
	// inputDir is no longer used for AST parsing but kept for potential future use or to avoid breaking existing scripts if any.
	inputDir   string
	outputFile string
)

// ReflectionFieldInfo is an intermediate struct to hold data extracted via reflection.
type ReflectionFieldInfo struct {
	Path            string // Full dot-notation path from root config. e.g. "Redis.Host"
	FieldName       string // Original Go field name. e.g. "Host"
	FieldTypeStr    string // String representation of field type. e.g. "string", "int", "config.RedisConfig"
	StructName      string // Name of the immediate struct this field belongs to. e.g. "Config", "RedisConfig"
	StructPkgPath   string // Package path of the struct this field belongs to.
	YAMLName        string
	EnvName         string
	EnvSeparator    string
	Description     string
	Required        string      // "true", "false", or "Y", "N", "C" from tag
	DefaultValue    interface{} // Actual default value
	IsEmbeddedField bool        // True if this field comes from an embedded struct and should be inlined in YAML
}

// ConfigField represents a field in a configuration struct (used by generateDocs)
type ConfigField struct {
	Name         string
	Type         string
	YAMLName     string
	EnvName      string
	EnvSeparator string
	Description  string
	Required     string // Y, N, C
	DefaultValue string // String representation
}

// ParsedConfig represents a parsed configuration struct (used by generateDocs)
type ParsedConfig struct {
	FileName     string // Can be set to "reflection" or similar
	Name         string // Go struct name (e.g., "Config", "RedisConfig")
	Fields       []ConfigField
	IsTopLevel   bool   // Flag to identify the root config.Config struct for pathing
	GoStructName string // Same as Name
}

func main() {
	defaultInputDir := "internal/config" // Retained for consistency, though not directly used for parsing config.Config
	defaultOutputFile := "docs/pages/references/configuration.mdx"

	flag.StringVar(&inputDir, "input-dir", defaultInputDir, "Directory containing the Go configuration source files (usage changed with reflection).")
	flag.StringVar(&outputFile, "output-file", defaultOutputFile, "Path to the output Markdown file.")
	flag.Parse()

	fmt.Println("Configuration Documentation Generator (Reflection-based)")
	log.Printf("Output file: %s", outputFile)

	allFieldInfos, err := parseConfigWithReflection()
	if err != nil {
		log.Fatalf("Error parsing config with reflection: %v", err)
	}

	parsedConfigs := transformToParsedConfigs(allFieldInfos)

	err = generateDocs(parsedConfigs, outputFile)
	if err != nil {
		log.Fatalf("Error generating docs: %v", err)
	}

	fmt.Printf("Successfully generated documentation to %s\n", outputFile)
}

func parseConfigWithReflection() ([]ReflectionFieldInfo, error) {
	var infos []ReflectionFieldInfo
	cfg := config.Config{}
	cfg.InitDefaults()

	cfgType := reflect.TypeOf(cfg)
	cfgValue := reflect.ValueOf(cfg)
	targetPkgPath := cfgType.PkgPath()

	err := extractFieldsRecursive(cfgValue, cfgType, "", cfgType.Name(), cfgType.PkgPath(), targetPkgPath, &infos, false)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func extractFieldsRecursive(currentVal reflect.Value, currentType reflect.Type, pathPrefix string, structName string, structPkgPath string, targetPkgPath string, infos *[]ReflectionFieldInfo, isEmbedded bool) error {
	if currentType.Kind() == reflect.Ptr {
		if currentVal.IsNil() {
			// If the pointer is nil, create a zero value of the element type to inspect its structure
			// This ensures fields of this struct are documented, albeit with zero/nil defaults.
			currentType = currentType.Elem()
			currentVal = reflect.Zero(currentType)
		} else {
			currentVal = currentVal.Elem()
			currentType = currentType.Elem()
		}
	}

	if currentType.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct or pointer to a struct, got %s", currentType.Kind())
	}

	for i := 0; i < currentType.NumField(); i++ {
		fieldSpec := currentType.Field(i)
		fieldVal := currentVal.Field(i)

		// Skip unexported fields
		if !fieldVal.CanInterface() {
			continue
		}

		fieldName := fieldSpec.Name
		fullPath := fieldName
		if pathPrefix != "" {
			fullPath = pathPrefix + "." + fieldName
		}

		// Handle anonymous/embedded structs
		if fieldSpec.Anonymous {
			// For embedded structs, recurse with the same pathPrefix (or adjusted if needed)
			// and mark fields as embedded so they can be inlined in YAML.
			err := extractFieldsRecursive(fieldVal, fieldSpec.Type, pathPrefix, structName, structPkgPath, targetPkgPath, infos, true)
			if err != nil {
				return fmt.Errorf("error recursing into embedded struct %s: %w", fieldName, err)
			}
			continue // Skip adding the embedded struct itself as a field
		}

		yamlTag := fieldSpec.Tag.Get("yaml")
		yamlName := strings.Split(yamlTag, ",")[0] // Get name part, ignore omitempty etc.
		if yamlName == "-" {                       // Skip fields explicitly ignored by YAML
			continue
		}
		if yamlName == "" { // Default to field name if yaml tag is missing or empty
			// This behavior might need adjustment based on how YAML typically marshals
			// For documentation, often we want to show the Go field name if no YAML name.
			// However, for config, usually explicit YAML names are preferred.
			// Let's assume if yaml tag is missing, it's not a primary config field for YAML docs.
			// Or, for full documentation, one might want to include it.
			// For now, if no yamlName, we might skip or use Go name.
			// Let's use Go field name if yamlName is empty after split (e.g. tag was just ",omitempty")
			// but if tag was entirely missing, it's less clear.
			// The current AST parser implies fields without YAML tags are still processed.
			// Let's default to Go field name if yamlName is empty.
			yamlName = fieldName
		}

		fieldTypeStr := formatFieldType(fieldSpec.Type)

		// Create info for the current field itself
		// This applies whether it's a basic type, a struct from another package, or a struct we'll recurse into.
		currentFieldInfo := ReflectionFieldInfo{
			Path:            fullPath,
			FieldName:       fieldName,
			FieldTypeStr:    fieldTypeStr,
			StructName:      structName, // The struct this field belongs to (e.g., "Config")
			StructPkgPath:   structPkgPath,
			YAMLName:        yamlName,
			EnvName:         fieldSpec.Tag.Get("env"),
			EnvSeparator:    fieldSpec.Tag.Get("envSeparator"),
			Description:     fieldSpec.Tag.Get("desc"),
			Required:        fieldSpec.Tag.Get("required"),
			DefaultValue:    fieldVal.Interface(), // Capture the default value of this field
			IsEmbeddedField: isEmbedded,           // This is about whether *this field* is part of an embedding context
		}
		// If fieldSpec.Anonymous was true, we already 'continue'd earlier.
		// So, this currentFieldInfo is for a named field.
		*infos = append(*infos, currentFieldInfo)

		// If the field's type is a struct from the target package (and not an interface), recurse into its members.
		// These members will have their StructName as fieldSpec.Type.Name() (e.g., "RedisConfig")
		// and their Path will be prefixed with fullPath (e.g., "Redis.Host").
		actualFieldType := fieldSpec.Type
		if actualFieldType.Kind() == reflect.Ptr {
			actualFieldType = actualFieldType.Elem()
		}

		if fieldSpec.Type.Kind() != reflect.Interface && actualFieldType.Kind() == reflect.Struct && actualFieldType.PkgPath() == targetPkgPath && actualFieldType.Name() != "Time" {
			// Use actualFieldType.Name() for the StructName of the members
			err := extractFieldsRecursive(fieldVal, fieldSpec.Type, fullPath, actualFieldType.Name(), actualFieldType.PkgPath(), targetPkgPath, infos, false) // isEmbedded is false for members of a regularly named struct field
			if err != nil {
				return fmt.Errorf("error recursing into members of struct field %s (type %s): %w", fullPath, actualFieldType.Name(), err)
			}
		}
	}
	return nil
}

func formatFieldType(t reflect.Type) string {
	// This produces types like "string", "int", "[]string", "map[string]int", "config.RedisConfig", "*config.MQsConfig"
	// which is generally good for documentation.
	return t.String()
}

func isTargetPkgStruct(t reflect.Type, targetPkgPath string) bool {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Kind() == reflect.Struct && t.PkgPath() == targetPkgPath && t.Name() != "Time" // Exclude time.Time
}

func formatDefaultValueToString(value interface{}, goType string) string {
	if value == nil {
		return "" // generateDocs will convert this to `nil`
	}

	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return ""
	}

	// Handle time.Time specifically if needed, otherwise %v is usually okay.
	if t, ok := value.(time.Time); ok {
		if t.IsZero() {
			return "" // Represent zero time as empty for cleaner defaults
		}
		return t.Format(time.RFC3339) // Or another suitable format
	}

	// Handle slices
	if val.Kind() == reflect.Slice {
		if val.IsNil() { // Explicitly nil slice
			return ""
		}
		if val.Len() == 0 { // Empty slice []
			return "[]"
		}
		// For non-empty slices, fmt.Sprintf("%v", value) gives "[elem1 elem2 ...]"
		// We might want commas: "[elem1, elem2]"
		var parts []string
		for i := 0; i < val.Len(); i++ {
			parts = append(parts, fmt.Sprintf("%v", val.Index(i).Interface()))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	}

	// Handle booleans to ensure "true" or "false"
	if val.Kind() == reflect.Bool {
		return strconv.FormatBool(val.Bool())
	}

	// Default string representation
	strVal := fmt.Sprintf("%v", value)

	// Avoid Go's default "<nil>" string for non-interface nils (e.g. nil map/slice that wasn't caught above)
	if strVal == "<nil>" {
		return ""
	}

	return strVal
}

func transformToParsedConfigs(infos []ReflectionFieldInfo) []ParsedConfig {
	groupedByStruct := make(map[string][]ReflectionFieldInfo)
	structOrder := []string{} // To maintain an order, e.g., "Config" first

	for _, info := range infos {
		if _, exists := groupedByStruct[info.StructName]; !exists {
			groupedByStruct[info.StructName] = []ReflectionFieldInfo{}
			if info.StructName == "Config" { // Ensure "Config" is first if present
				structOrder = append([]string{info.StructName}, structOrder...)
			} else {
				structOrder = append(structOrder, info.StructName)
			}
		}
		groupedByStruct[info.StructName] = append(groupedByStruct[info.StructName], info)
	}

	// Sort structOrder to have "Config" first, then alphabetically for others
	sort.SliceStable(structOrder, func(i, j int) bool {
		if structOrder[i] == "Config" {
			return true
		}
		if structOrder[j] == "Config" {
			return false
		}
		return structOrder[i] < structOrder[j]
	})

	var parsedConfigs []ParsedConfig
	for _, structName := range structOrder {
		fieldsInfo := groupedByStruct[structName]
		var configFields []ConfigField
		var goStructPkgPath string // Take from the first field, should be consistent

		// Sort fields by their original Go field name for consistent order
		// This is important if ReflectionFieldInfo Path was used for sorting,
		// but here we sort by FieldName within each struct.
		sort.Slice(fieldsInfo, func(i, j int) bool {
			// A simple path sort might be better if fields from embedded structs are mixed.
			// For now, FieldName within the current struct context.
			return fieldsInfo[i].FieldName < fieldsInfo[j].FieldName
		})

		for _, info := range fieldsInfo {
			if goStructPkgPath == "" {
				goStructPkgPath = info.StructPkgPath
			}
			configFields = append(configFields, ConfigField{
				Name:         info.FieldName, // Use the Go field name
				Type:         info.FieldTypeStr,
				YAMLName:     info.YAMLName,
				EnvName:      info.EnvName,
				EnvSeparator: info.EnvSeparator,
				Description:  info.Description,
				Required:     info.Required,
				DefaultValue: formatDefaultValueToString(info.DefaultValue, info.FieldTypeStr),
			})
		}

		// Determine if this struct is the top-level "Config"
		// The main config.Config struct from internal/config
		isTopLevel := (structName == "Config" && goStructPkgPath == "github.com/hookdeck/outpost/internal/config")

		parsedConfigs = append(parsedConfigs, ParsedConfig{
			FileName:     "reflection-generated",
			Name:         structName,
			GoStructName: structName,
			Fields:       configFields,
			IsTopLevel:   isTopLevel,
		})
	}
	return parsedConfigs
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
	yamlBuilder.WriteString("```yaml\n")
	yamlBuilder.WriteString("# Outpost Configuration Example (Generated)\n")
	yamlBuilder.WriteString("# This example shows all available keys with their default values where applicable.\n\n")

	var mainConfigInfo *ParsedConfig
	configInfoMap := make(map[string]*ParsedConfig)
	for i := range parsedConfigs {
		pc := &parsedConfigs[i] // operate on a copy to avoid modifying the slice
		configInfoMap[pc.GoStructName] = pc
		if pc.IsTopLevel { // Use IsTopLevel flag
			mainConfigInfo = pc
		}
	}

	if mainConfigInfo != nil {
		generateYAMLPart(&yamlBuilder, mainConfigInfo, configInfoMap, 0, true)
	} else {
		// Fallback if IsTopLevel wasn't set correctly, try finding "Config" by name
		if cfgByName, ok := configInfoMap["Config"]; ok {
			log.Println("Warning: Main 'Config' struct not found by IsTopLevel flag, falling back to name 'Config'.")
			generateYAMLPart(&yamlBuilder, cfgByName, configInfoMap, 0, true)
		} else {
			yamlBuilder.WriteString("# ERROR: Main 'Config' struct not found. Cannot generate YAML structure.\n")
			log.Println("Error: Main 'Config' struct not found by IsTopLevel flag or by name 'Config'.")
		}
	}
	yamlBuilder.WriteString("```\n")
	yamlContent := yamlBuilder.String()

	// --- Read existing MDX file ---
	mdxBytes, err := os.ReadFile(outputPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Warning: Output file %s does not exist. A new file will be created with placeholders.", outputPath)
			// Create a template with placeholders if file doesn't exist
			templateContent := fmt.Sprintf(`---
title: Configuration Reference
description: Detailed configuration options for Outpost.
---

This document outlines all the configuration options available for Outpost, settable via environment variables or a YAML configuration file.

## Environment Variables

%s

%s

%s

## YAML Configuration

Below is an example YAML configuration file showing all available options and their default values.

%s

%s

%s
`, envVarsStartPlaceholder, envVarsContent, envVarsEndPlaceholder, yamlStartPlaceholder, yamlContent, yamlEndPlaceholder)
			mdxBytes = []byte(templateContent)
		} else {
			return fmt.Errorf("failed to read output file %s: %w", outputPath, err)
		}
	}
	mdxContent := string(mdxBytes)

	// --- Replace placeholders ---
	envStartIndex := strings.Index(mdxContent, envVarsStartPlaceholder)
	envEndIndex := strings.Index(mdxContent, envVarsEndPlaceholder)

	if envStartIndex != -1 && envEndIndex != -1 && envEndIndex > envStartIndex {
		newMdxContent, changed := replacePlaceholder(mdxContent, envVarsStartPlaceholder, envVarsEndPlaceholder, "\n"+envVarsContent+"\n")
		if !changed {
			// Placeholder found, but content was already up-to-date.
			log.Printf("Info: The content for the ENV vars placeholder '%s' was already up-to-date. No changes made to this block.", envVarsStartPlaceholder)
		} else {
			mdxContent = newMdxContent // Content was updated.
		}
	} else {
		// Placeholder not found or in wrong order.
		log.Printf("Warning: The ENV vars placeholder '%s' (and/or its corresponding end tag '%s') was not found or is in the wrong order in the output file. The ENV vars block will not be updated.", envVarsStartPlaceholder, envVarsEndPlaceholder)
	}

	yamlStartIndex := strings.Index(mdxContent, yamlStartPlaceholder)
	yamlEndIndex := strings.Index(mdxContent, yamlEndPlaceholder)

	if yamlStartIndex != -1 && yamlEndIndex != -1 && yamlEndIndex > yamlStartIndex {
		newMdxContent, changed := replacePlaceholder(mdxContent, yamlStartPlaceholder, yamlEndPlaceholder, "\n"+yamlContent+"\n")
		if !changed {
			// Placeholder found, but content was already up-to-date.
			log.Printf("Info: The content for the YAML placeholder '%s' was already up-to-date. No changes made to this block.", yamlStartPlaceholder)
		} else {
			mdxContent = newMdxContent // Content was updated.
		}
	} else {
		// Placeholder not found or in wrong order.
		log.Printf("Warning: The YAML placeholder '%s' (and/or its corresponding end tag '%s') was not found or is in the wrong order in the output file. The YAML block will not be updated.", yamlStartPlaceholder, yamlEndPlaceholder)
	}

	// --- Write updated content back to MDX file ---
	err = os.WriteFile(outputPath, []byte(mdxContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated content to %s: %w", outputPath, err)
	}

	return nil
}

func replacePlaceholder(content, startPlaceholder, endPlaceholder, newBlockContent string) (string, bool) {
	startIndex := strings.Index(content, startPlaceholder)
	endIndex := strings.Index(content, endPlaceholder)

	if startIndex != -1 && endIndex != -1 && endIndex > startIndex {
		// Include the start placeholder, replace content until end placeholder, then add end placeholder
		newContent := content[:startIndex+len(startPlaceholder)] +
			newBlockContent +
			content[endIndex:]
		return newContent, newContent != content
	}
	return content, false // Placeholder not found or content is the same
}

func generateYAMLPart(builder *strings.Builder, configInfo *ParsedConfig, allConfigs map[string]*ParsedConfig, indentLevel int, isRoot bool) {
	indent := strings.Repeat("  ", indentLevel)

	// Sort fields by YAMLName for consistent YAML output
	// This is important because map iteration order is not guaranteed in Go for allConfigs
	// and field order from reflection is Go struct order, not necessarily desired YAML order.
	// However, for struct fields, we use the order from ParsedConfig.Fields which is
	// now sorted by Go FieldName in transformToParsedConfigs.
	// For truly aesthetic YAML, one might want a custom sort order.
	// For now, using the order from ParsedConfig.Fields.

	for _, field := range configInfo.Fields {
		if field.YAMLName == "" || field.YAMLName == "-" { // Skip if no YAML name or explicitly ignored
			continue
		}
		// Field description as a comment
		if field.Description != "" {
			// Ensure multi-line descriptions are commented correctly
			descLines := strings.Split(field.Description, "\n")
			for _, line := range descLines {
				builder.WriteString(fmt.Sprintf("%s# %s\n", indent, strings.TrimSpace(line)))
			}
		}

		// Removed default value comments as per feedback. The value itself is shown.
		// if field.DefaultValue != "" && !isStructTypeField {
		// 	builder.WriteString(fmt.Sprintf("%s# Default: %s\n", indent, field.DefaultValue))
		// } else if field.DefaultValue == "" && !isStructTypeField && field.Required != "Y" && field.Required != "true" {
		// 	// Indicate if no default and not required (optional)
		// 	builder.WriteString(fmt.Sprintf("%s# Default: (none)\n", indent))
		// }

		// Required status as a comment
		if field.Required != "" && field.Required != "N" && field.Required != "false" {
			builder.WriteString(fmt.Sprintf("%s# Required: %s\n", indent, field.Required))
		}

		// Field line
		builder.WriteString(fmt.Sprintf("%s%s:", indent, field.YAMLName))

		var nestedStructShortName string
		fieldTypeForLookup := field.Type // This is string like "config.RedisConfig" or "*config.MQsConfig"
		if strings.HasPrefix(fieldTypeForLookup, "*") {
			fieldTypeForLookup = fieldTypeForLookup[1:] // remove "*"
		}
		// Further strip package path if present, e.g. "config.RedisConfig" -> "RedisConfig"
		// or "time.Time" -> "Time"
		parts := strings.Split(fieldTypeForLookup, ".")
		if len(parts) > 0 {
			nestedStructShortName = parts[len(parts)-1] // Get the last part
		}

		if nestedStructShortName != "" {
			if nestedConfig, ok := allConfigs[nestedStructShortName]; ok && field.Type != "string" && nestedStructShortName != "Time" { // also ensure we don't try to expand time.Time as a custom struct
				// It's a nested struct we should expand
				builder.WriteString("\n")
				generateYAMLPart(builder, nestedConfig, allConfigs, indentLevel+1, false)
			} else {
				// Scalar, slice, map, or a struct from a different package (or time.Time)
				valueStr := field.DefaultValue
				if valueStr == "" || (valueStr == "[]" && strings.HasPrefix(field.Type, "[]")) {
					valueStr = getYAMLPlaceholderForType(field.Type)
				} else {
					valueStr = formatSimpleValueToYAML(valueStr, field.Type)
				}
				builder.WriteString(fmt.Sprintf(" %s\n", valueStr))
			}
		} else {
			// Fallback if short name extraction failed (should not happen for valid types)
			valueStr := field.DefaultValue
			if valueStr == "" || (valueStr == "[]" && strings.HasPrefix(field.Type, "[]")) {
				valueStr = getYAMLPlaceholderForType(field.Type)
			} else {
				valueStr = formatSimpleValueToYAML(valueStr, field.Type)
			}
			builder.WriteString(fmt.Sprintf(" %s\n", valueStr))
		}
		builder.WriteString("\n") // Add a blank line after each top-level entry in a struct for readability
	}
}

func getYAMLPlaceholderForType(goType string) string {
	switch {
	case strings.HasPrefix(goType, "[]string"):
		return "[item1, item2]"
	case strings.HasPrefix(goType, "[]"): // Generic slice
		return "[]"
	case strings.HasPrefix(goType, "map["):
		return "{key: value}"
	case goType == "string":
		return "\"\""
	case goType == "int", goType == "int64", goType == "float64":
		return "0"
	case goType == "bool":
		return "false"
	default:
		return "# <" + goType + ">"
	}
}

func formatSimpleValueToYAML(value, goType string) string {
	// If it's a string, ensure it's quoted if it contains special chars or is empty
	if goType == "string" {
		if value == "" {
			return "\"\"" // Explicitly empty string
		}
		// Always quote non-empty strings to handle all special characters correctly for YAML.
		return strconv.Quote(value)
	}
	if value == "[]" && strings.HasPrefix(goType, "[]") { // Empty slice from default
		return "[]"
	}
	// For non-string types (numbers, booleans), the default fmt.Sprintf("%v") representation is usually fine for YAML.
	return value
}

// collectEnvVarFields gathers all fields that have an EnvName, from all ParsedConfig structs.
func collectEnvVarFields(parsedConfigs []ParsedConfig) []ConfigField {
	var allFields []ConfigField
	seenEnvVars := make(map[string]bool) // To avoid duplicates if structs are processed multiple times or nested weirdly

	// Process top-level "Config" first if available, then others.
	// This helps in establishing a somewhat predictable order if paths were involved.
	// With the current flat list from reflection, order of parsedConfigs matters less here.

	sortedParsedConfigs := make([]ParsedConfig, len(parsedConfigs))
	copy(sortedParsedConfigs, parsedConfigs)
	sort.Slice(sortedParsedConfigs, func(i, j int) bool {
		if sortedParsedConfigs[i].IsTopLevel {
			return true
		}
		if sortedParsedConfigs[j].IsTopLevel {
			return false
		}
		return sortedParsedConfigs[i].Name < sortedParsedConfigs[j].Name
	})

	for _, pc := range sortedParsedConfigs {
		for _, field := range pc.Fields {
			if field.EnvName != "" && !seenEnvVars[field.EnvName] {
				allFields = append(allFields, field)
				seenEnvVars[field.EnvName] = true
			}
		}
	}
	// Sort by EnvName for consistent table output
	sort.Slice(allFields, func(i, j int) bool {
		return allFields[i].EnvName < allFields[j].EnvName
	})
	return allFields
}

func formatRequiredText(reqStatus string) string {
	switch strings.ToUpper(reqStatus) {
	case "Y", "TRUE":
		return "Yes"
	case "N", "FALSE":
		return "No"
	case "C":
		return "Conditional"
	default:
		if reqStatus != "" {
			return reqStatus // Show as is if not recognized
		}
		return "No" // Default to No if empty
	}
}
