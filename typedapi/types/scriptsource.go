package types

type ScriptSource any

type ScriptSourceVariant interface {
	ScriptSourceCaster() *ScriptSource
}
