// Package flagsetscheme provides a runtime.TypeScheme with support
// for command line option sets for the described object types.
// Therefore, the object types (VersionTypedObjectType) have to provide
// a flagsets.ConfigOptionTypeSetHandler. To support CLI help information,
// they should additionally provide a description and structure information.
package flagsetscheme
