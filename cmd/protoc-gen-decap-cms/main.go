package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	cmsv1 "go.einride.tech/protobuf-decap-cms/proto/gen/go/einride/decap/cms/v1"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, file := range gen.Files {
			if !file.Generate {
				continue
			}
			config := proto.GetExtension(file.Desc.Options(), cmsv1.E_Config).(*cmsv1.Config)
			if config == nil {
				continue
			}
			collectMessages(config, file.Desc.Package(), gen.Files)
			g := &generatedYAMLFile{
				GeneratedFile: gen.NewGeneratedFile(file.GeneratedFilenamePrefix+".yml", file.GoImportPath),
			}
			g.Y("# Generated by protoc-gen-decap-cms. DO NOT EDIT.")
			g.Y("backend:")
			g.Up()
			genBackend(g, config.GetBackend())
			g.Down()
			if config.GetLocalBackend() != nil {
				g.Y()
				g.Y("local_backend:")
				g.Up()
				genLocalBackend(g, config.GetLocalBackend())
				g.Down()
			}
			g.Y()
			g.Y("slug:")
			g.Up()
			genSlug(g, config.GetSlug())
			g.Down()
			if config.GetPublishMode() == cmsv1.Config_EDITORIAL_WORKFLOW {
				g.Y()
				g.Y("publish_mode: editorial_workflow")
			}
			if config.GetMediaFolder() != "" {
				g.Y()
				g.Y("media_folder: ", strconv.Quote(config.GetMediaFolder()))
			}
			if config.GetLogoUrl() != "" {
				g.Y()
				g.Y("logo_url: ", strconv.Quote(config.GetLogoUrl()))
			}
			g.Y()
			g.Y("collections:")
			for _, collection := range config.GetCollections() {
				g.Up()
				genCollection(g, collection)
				g.Down()
			}
		}
		return nil
	})
}

func genBackend(g *generatedYAMLFile, backend *cmsv1.Config_Backend) {
	if backend.GetName() != "" {
		g.Y("name: ", strconv.Quote(backend.GetName()))
	}
	if backend.GetRepo() != "" {
		g.Y("repo: ", strconv.Quote(backend.GetRepo()))
	}
	if backend.GetBranch() != "" {
		g.Y("branch: ", strconv.Quote(backend.GetBranch()))
	}
	if backend.GetSquashMerges() {
		g.Y("squash_merges: true")
	}
	if backend.GetBaseUrl() != "" {
		g.Y("base_url: ", strconv.Quote(backend.GetBaseUrl()))
	}
	if backend.GetSiteDomain() != "" {
		g.Y("site_domain: ", strconv.Quote(backend.GetSiteDomain()))
	}
	if backend.GetCommitMessages() != nil {
		g.Y("commit_messages:")
		g.Up()
		if backend.GetCommitMessages().GetCreate() != "" {
			g.Y("create: ", strconv.Quote(backend.GetCommitMessages().GetCreate()))
		}
		if backend.GetCommitMessages().GetUpdate() != "" {
			g.Y("update: ", strconv.Quote(backend.GetCommitMessages().GetUpdate()))
		}
		if backend.GetCommitMessages().GetDelete() != "" {
			g.Y("delete: ", strconv.Quote(backend.GetCommitMessages().GetDelete()))
		}
		if backend.GetCommitMessages().GetUploadMedia() != "" {
			g.Y("uploadMedia: ", strconv.Quote(backend.GetCommitMessages().GetUploadMedia()))
		}
		if backend.GetCommitMessages().GetDeleteMedia() != "" {
			g.Y("deleteMedia: ", strconv.Quote(backend.GetCommitMessages().GetDeleteMedia()))
		}
		g.Down()
	}
}

func genLocalBackend(g *generatedYAMLFile, localBackend *cmsv1.Config_LocalBackend) {
	if localBackend.GetUrl() != "" {
		g.Y("url: ", strconv.Quote(localBackend.GetUrl()))
	}
}

func genSlug(g *generatedYAMLFile, slug *cmsv1.Config_Slug) {
	switch slug.GetEncoding() {
	case cmsv1.Config_Slug_UNICODE:
		g.Y("encoding: ", strconv.Quote("unicode"))
	case cmsv1.Config_Slug_ASCII:
		g.Y("encoding: ", strconv.Quote("ascii"))
	default:
		g.Y("encoding: ", strconv.Quote("unicode"))
	}
	g.Y("clean_accents: ", strconv.FormatBool(slug.GetCleanAccents()))
	if slug.GetSanitizeReplacement() != "" {
		g.Y("sanitize_replacement: ", strconv.Quote(slug.GetSanitizeReplacement()))
	}
}

func genCollection(g *generatedYAMLFile, collection *cmsv1.Collection) {
	g.Y()
	g.Y("- name: ", strconv.Quote(collection.GetName()))
	g.Up()
	defer g.Down()
	if collection.GetLabel() != "" {
		g.Y("label: ", strconv.Quote(collection.GetLabel()))
	}
	if collection.GetLabelSingular() != "" {
		g.Y("label_singular: ", strconv.Quote(collection.GetLabelSingular()))
	}
	if collection.GetFolder() != "" {
		g.Y("folder: ", strconv.Quote(collection.GetFolder()))
	}
	g.Y("create: ", strconv.FormatBool(collection.GetCreate()))
	if collection.GetIdentifierField() != "" {
		g.Y("identifier_field: ", strconv.Quote(collection.GetIdentifierField()))
	}
	if collection.GetFormat() != "" {
		g.Y("format: ", strconv.Quote(collection.GetFormat()))
	}
	if collection.GetDescription() != "" {
		g.Y("description: ", strconv.Quote(collection.GetDescription()))
	}
	if collection.GetSummary() != "" {
		g.Y("summary: ", strconv.Quote(collection.GetSummary()))
	}
	g.Y("editor:")
	g.Up()
	g.Y("preview: ", strconv.FormatBool(collection.GetEditor().GetPreview()))
	g.Down()
	g.Y("fields:")
	for _, field := range collection.GetFields() {
		g.Up()
		genField(g, field)
		g.Down()
	}
}

func genField(g *generatedYAMLFile, field *cmsv1.Field) {
	g.Y()
	g.Y("- name: ", strconv.Quote(field.GetName()))
	g.Up()
	defer g.Down()
	if field.GetLabel() != "" {
		g.Y("label: ", strconv.Quote(field.GetLabel()))
	}
	if field.GetComment() != "" {
		g.Y("comment: ", strconv.Quote(field.GetComment()))
	}
	// Number widget required is handled in switch case
	if _, isNumberWidget := field.GetWidget().GetWidgetType().(*cmsv1.Widget_NumberWidget); !isNumberWidget {
		g.Y("required: ", strconv.FormatBool(field.GetWidget().GetRequiredValue()))
	}
	if field.GetWidget().GetHint() != "" {
		g.Y("hint: ", strconv.Quote(strings.TrimSpace(field.GetWidget().GetHint())))
	}
	if field.GetWidget().GetPattern() != nil {
		g.Y("pattern:")
		g.Up()
		g.Y("- ", strconv.Quote(field.GetWidget().GetPattern().GetRegexp()))
		g.Y("- ", strconv.Quote(field.GetWidget().GetPattern().GetErrorMessage()))
		g.Down()
	}
	switch widget := field.GetWidget().GetWidgetType().(type) {
	case *cmsv1.Widget_StringWidget:
		g.Y("widget: ", strconv.Quote("string"))
		g.Y("default: ", strconv.Quote(widget.StringWidget.GetDefaultValue()))
	case *cmsv1.Widget_TextWidget:
		g.Y("widget: ", strconv.Quote("text"))
		g.Y("default: ", strconv.Quote(widget.TextWidget.GetDefaultValue()))
	case *cmsv1.Widget_BooleanWidget:
		g.Y("widget: ", strconv.Quote("boolean"))
	case *cmsv1.Widget_SelectWidget:
		g.Y("widget: ", strconv.Quote("select"))
		if len(widget.SelectWidget.GetDefaultValue()) == 1 {
			g.Y("default: ", strconv.Quote(widget.SelectWidget.GetDefaultValue()[0]))
		}
		g.Y("multiple: ", widget.SelectWidget.GetMultiple())
		g.Y("options:")
		g.Up()
		for _, option := range widget.SelectWidget.GetOptions() {
			g.Y("- label: ", strconv.Quote(option.GetLabel()))
			g.Up()
			g.Y("value: ", strconv.Quote(option.GetValue()))
			g.Down()
		}
		g.Down()
	case *cmsv1.Widget_DateTimeWidget:
		g.Y("widget: ", strconv.Quote("datetime"))
		if widget.DateTimeWidget.GetDateFormat() != "" {
			g.Y("date_format: ", strconv.Quote(widget.DateTimeWidget.GetDateFormat()))
		}
		if widget.DateTimeWidget.GetTimeFormat() != "" {
			g.Y("time_format: ", strconv.Quote(widget.DateTimeWidget.GetTimeFormat()))
		}
	case *cmsv1.Widget_ObjectWidget:
		g.Y("widget: ", strconv.Quote("object"))
		g.Y("collapsed: ", strconv.FormatBool(widget.ObjectWidget.GetCollapsed()))
		if widget.ObjectWidget.GetSummary() != "" {
			g.Y("summary: ", strconv.Quote(widget.ObjectWidget.GetSummary()))
		}
		g.Y("fields:")
		g.Up()
		for _, objectField := range widget.ObjectWidget.GetFields() {
			genField(g, objectField)
		}
		g.Down()
	case *cmsv1.Widget_ListWidget:
		g.Y("widget: ", strconv.Quote("list"))
		g.Y("collapsed: ", strconv.FormatBool(widget.ListWidget.GetCollapsed()))
		g.Y("minimize_collapsed: ", strconv.FormatBool(widget.ListWidget.GetMinimizeCollapsed()))
		if widget.ListWidget.GetSummary() != "" {
			g.Y("summary: ", strconv.Quote(widget.ListWidget.GetSummary()))
		}
		if len(widget.ListWidget.GetFields()) > 0 {
			g.Y("fields:")
			g.Up()
			for _, objectField := range widget.ListWidget.GetFields() {
				genField(g, objectField)
			}
			g.Down()
		}
	case *cmsv1.Widget_NumberWidget:
		g.Y("widget: ", strconv.Quote("number"))
		g.Y("value_type: ", strconv.Quote(strings.ToLower(widget.NumberWidget.GetValueType().String())))
		g.Y("required: ", true) // required since Decap uses empty string for no value instead of 0
		if !field.GetWidget().GetRequiredValue() {
			g.Y("default: ", widget.NumberWidget.GetDefaultValue())
		}
	case *cmsv1.Widget_RelationWidget:
		g.Y("widget: ", strconv.Quote("relation"))
		g.Y("collection: ", strconv.Quote(widget.RelationWidget.GetCollection()))
		g.Y("value_field: ", strconv.Quote(widget.RelationWidget.GetValueField()))
		g.Y("search_fields:")
		g.Up()
		for _, searchField := range widget.RelationWidget.GetSearchFields() {
			g.Y("- ", strconv.Quote(searchField))
		}
		g.Down()
		if len(widget.RelationWidget.GetDisplayFields()) > 0 {
			g.Y("display_fields:")
			g.Up()
			for _, displayField := range widget.RelationWidget.GetDisplayFields() {
				g.Y("- ", strconv.Quote(displayField))
			}
			g.Down()
		}
		g.Y("multiple: ", strconv.FormatBool(widget.RelationWidget.GetMultiple()))
		if len(widget.RelationWidget.GetFilters()) > 0 {
			g.Y("filters:")
			g.Up()
			for _, filter := range widget.RelationWidget.GetFilters() {
				g.Y("- field: ", strconv.Quote(filter.GetField()))
				g.Up()
				g.Y("values:")
				g.Up()
				for _, value := range filter.GetValues() {
					g.Y("- ", strconv.Quote(value))
				}
				g.Down()
				g.Down()
			}
			g.Down()
		}
	case *cmsv1.Widget_CustomWidget:
		g.Y("widget: ", strconv.Quote(widget.CustomWidget.GetWidget()))
		for _, option := range widget.CustomWidget.GetOptions() {
			g.Y(option)
		}
	}
}

func collectMessages(config *cmsv1.Config, pkg protoreflect.FullName, files []*protogen.File) {
	for _, file := range files {
		if file.Desc.Package() != pkg {
			continue
		}
		for _, message := range file.Messages {
			collection := proto.GetExtension(
				message.Desc.Options(),
				cmsv1.E_Collection,
			).(*cmsv1.Collection)
			if collection == nil {
				continue
			}
			collection = proto.Clone(collection).(*cmsv1.Collection)
			if collection.GetDescription() == "" {
				collection.Description = strings.TrimSpace(string(message.Comments.Leading))
			}
			if collection.GetOwner() != nil {
				if collection.GetDescription() != "" {
					collection.Description += " "
				}
				collection.Description += fmt.Sprintf("[%s]", collection.GetOwner().GetDisplayName())
			}
			collectFields(collection, message)
			config.Collections = append(config.Collections, collection)
		}
	}
}

func collectFields(collection *cmsv1.Collection, message *protogen.Message) {
	for _, protoField := range message.Fields {
		if field, ok := inferField(message, protoField, nil); ok {
			collection.Fields = append(collection.Fields, field)
		}
	}
}

func resolveFieldOwner(fields []*protogen.Field) (*cmsv1.Owner, bool) {
	var owner *cmsv1.Owner
	for i, field := range fields {
		if i == 0 && field.Parent != nil {
			if collectionAnnotation := proto.GetExtension(
				fields[0].Parent.Desc.Options(),
				cmsv1.E_Collection,
			).(*cmsv1.Collection); collectionAnnotation.GetOwner() != nil {
				owner = collectionAnnotation.GetOwner()
			}
		}
		if fieldAnnotation := proto.GetExtension(
			field.Desc.Options(),
			cmsv1.E_Field,
		).(*cmsv1.Field); fieldAnnotation.GetOwner() != nil {
			owner = fieldAnnotation.GetOwner()
		}
	}
	return owner, owner != nil
}

func inferFieldLabel(field *protogen.Field) string {
	result := string(field.Desc.Name())
	result = strings.ReplaceAll(result, "_", " ")
	result = strings.ToUpper(result)
	return result
}

func inferField(
	protoMessage *protogen.Message,
	protoField *protogen.Field,
	parentFields []*protogen.Field,
) (*cmsv1.Field, bool) {
	field := &cmsv1.Field{
		Name:    string(protoField.Desc.Name()),
		Label:   inferFieldLabel(protoField),
		Comment: strings.TrimSpace(string(protoField.Comments.Leading)),
	}
	field.Widget = &cmsv1.Widget{
		Hint:          strings.TrimSpace(string(protoField.Comments.Leading)),
		RequiredValue: inferRequired(protoField),
	}
	if fieldAnnotation := proto.GetExtension(
		protoField.Desc.Options(),
		cmsv1.E_Field,
	).(*cmsv1.Field); fieldAnnotation != nil {
		if fieldAnnotation.GetIgnore() {
			return nil, false
		}
		proto.Merge(field, fieldAnnotation)
	}

	// special handling for the name field - which needs to be a proto string field
	if resource := proto.GetExtension(
		protoMessage.Desc.Options(),
		annotations.E_Resource,
	).(*annotations.ResourceDescriptor); resource != nil && protoField.Desc.Name() == "name" {
		if !(protoField.Desc.Kind() == protoreflect.StringKind && !protoField.Desc.IsList()) {
			log.Fatalf("name field must be proto string type")
		}

		field.Label = "RESOURCE NAME"
		field.Widget.RequiredValue = true
		if field.Widget.WidgetType == nil {
			field.Widget.WidgetType = &cmsv1.Widget_StringWidget{
				StringWidget: &cmsv1.StringWidget{},
			}
		}
		if len(resource.GetPattern()) > 0 {
			pattern := resource.GetPattern()[0]
			exp := regexp.MustCompile(`{.*}`).ReplaceAllString(pattern, `[a-z0-9][a-z0-9-]{0,61}[a-z0-9]`)
			exp = "^" + exp + "$"
			field.Widget.Pattern = &cmsv1.Widget_Pattern{
				Regexp:       exp,
				ErrorMessage: "Must match " + exp,
			}
			defaultValue := pattern[:strings.Index(pattern, "/")+1]
			if sw, ok := field.GetWidget().GetWidgetType().(*cmsv1.Widget_StringWidget); ok {
				sw.StringWidget.DefaultValue = defaultValue
			}
		}
		return field, true
	}

	// if a widget is specified and is a type that is not able to do more decoration, no further inference
	if field.Widget.WidgetType != nil && isUnDecoratableWidgetType(field.GetWidget().GetWidgetType()) {
		return field, true
	}

	if owner, ok := resolveFieldOwner(append(parentFields, protoField)); ok {
		field.Widget.Hint += fmt.Sprintf(" **[[%s]](%s)**", owner.GetDisplayName(), owner.GetUri())
	}
	switch protoField.Desc.Name() {
	case "revision_id", "revision_create_time":
		return nil, false
	}
	switch {
	case protoField.Desc.Kind() == protoreflect.MessageKind &&
		!protoField.Desc.IsList() &&
		protoField.Desc.Message().FullName() == "google.protobuf.Timestamp":
		field.Widget.WidgetType = &cmsv1.Widget_DateTimeWidget{
			DateTimeWidget: &cmsv1.DateTimeWidget{
				DateFormat: "YYYY-MM-DD",
				TimeFormat: "HH:mmZ",
			},
		}
		if protoField.Desc.Name() == "create_time" {
			field.Widget.RequiredValue = true
		}
		return field, true
	case protoField.Desc.Kind() == protoreflect.BoolKind && !protoField.Desc.IsList():
		field.Widget.WidgetType = &cmsv1.Widget_BooleanWidget{
			BooleanWidget: &cmsv1.BooleanWidget{},
		}
		return field, true
	case protoField.Desc.Kind() == protoreflect.StringKind && protoField.Desc.IsList():
		field.Widget.WidgetType = &cmsv1.Widget_ListWidget{
			ListWidget: &cmsv1.ListWidget{
				AllowAdd: true,
			},
		}
		return field, true
	case protoField.Desc.Kind() == protoreflect.StringKind && !protoField.Desc.IsList():
		field.Widget.WidgetType = &cmsv1.Widget_StringWidget{
			StringWidget: &cmsv1.StringWidget{},
		}
		return field, true
	case protoField.Desc.Kind() == protoreflect.EnumKind:
		field.Widget.WidgetType = &cmsv1.Widget_SelectWidget{
			SelectWidget: &cmsv1.SelectWidget{
				Multiple: protoField.Desc.IsList(),
			},
		}
		var options []*cmsv1.SelectWidget_Option
		for i := 0; i < protoField.Desc.Enum().Values().Len(); i++ {
			value := protoField.Desc.Enum().Values().Get(i)
			if inferRequired(protoField) && strings.HasSuffix(string(value.Name()), "_UNSPECIFIED") {
				continue
			}
			options = append(options, &cmsv1.SelectWidget_Option{
				Label: strings.ReplaceAll(string(value.Name()), "_", " "),
				Value: string(value.Name()),
			})
		}
		field.Widget.WidgetType.(*cmsv1.Widget_SelectWidget).SelectWidget.Options = options
		return field, true
	case protoField.Desc.Kind() == protoreflect.MessageKind && !protoField.Desc.IsList() && !protoField.Desc.IsMap():
		objectFields := make([]*cmsv1.Field, 0, len(protoField.Message.Fields))
		for _, protoObjectField := range protoField.Message.Fields {
			if objectField, ok := inferField(
				protoField.Message,
				protoObjectField,
				append(parentFields, protoField),
			); ok {
				objectFields = append(objectFields, objectField)
			}
		}
		field.Widget.WidgetType = &cmsv1.Widget_ObjectWidget{
			ObjectWidget: &cmsv1.ObjectWidget{
				Collapsed: !inferRequired(protoField),
				Summary:   "", // TODO
				Fields:    objectFields,
			},
		}
		return field, len(objectFields) > 0
	case protoField.Desc.Kind() == protoreflect.MessageKind && protoField.Desc.IsList():
		objectFields := make([]*cmsv1.Field, 0, len(protoField.Message.Fields))
		for _, protoObjectField := range protoField.Message.Fields {
			if objectField, ok := inferField(
				protoField.Message,
				protoObjectField,
				append(parentFields, protoField),
			); ok {
				objectFields = append(objectFields, objectField)
			}
		}

		listWidget := &cmsv1.ListWidget{
			AllowAdd:          true,
			Collapsed:         !inferRequired(protoField),
			MinimizeCollapsed: true,
			Fields:            objectFields,
		}
		if field.Widget.WidgetType != nil {
			userDefinedListWidget := field.GetWidget().GetWidgetType().(*cmsv1.Widget_ListWidget).ListWidget
			proto.Merge(listWidget, userDefinedListWidget)
		}

		field.Widget.WidgetType = &cmsv1.Widget_ListWidget{
			ListWidget: listWidget,
		}
		return field, len(objectFields) > 0
	case (protoField.Desc.Kind() == protoreflect.DoubleKind ||
		protoField.Desc.Kind() == protoreflect.FloatKind) && !protoField.Desc.IsList():
		field.Widget.WidgetType = &cmsv1.Widget_NumberWidget{
			NumberWidget: &cmsv1.NumberWidget{
				ValueType: cmsv1.NumberWidget_FLOAT,
			},
		}
		return field, true
	case (protoField.Desc.Kind() == protoreflect.Int64Kind ||
		protoField.Desc.Kind() == protoreflect.Int32Kind) && !protoField.Desc.IsList():
		field.Widget.WidgetType = &cmsv1.Widget_NumberWidget{
			NumberWidget: &cmsv1.NumberWidget{
				ValueType: cmsv1.NumberWidget_INT,
			},
		}
		return field, true
	}
	return nil, false
}

func inferRequired(field *protogen.Field) bool {
	for _, fieldBehavior := range proto.GetExtension(
		field.Desc.Options(),
		annotations.E_FieldBehavior,
	).([]annotations.FieldBehavior) {
		if fieldBehavior == annotations.FieldBehavior_REQUIRED {
			return true
		}
	}
	return false
}

type generatedYAMLFile struct {
	*protogen.GeneratedFile
	level int
}

func (g *generatedYAMLFile) I(level int, v ...any) {
	if level < 1 || len(v) == 0 {
		g.P(v...)
	} else {
		g.P(append([]any{strings.Repeat("  ", level)}, v...)...)
	}
}

func (g *generatedYAMLFile) Y(v ...any) {
	g.I(g.level, v...)
}

func (g *generatedYAMLFile) Up() {
	g.level++
}

func (g *generatedYAMLFile) Down() {
	g.level--
}

func isUnDecoratableWidgetType(t interface{}) bool {
	switch t.(type) {
	case *cmsv1.Widget_ListWidget:
		return false
	default:
		return true
	}
}