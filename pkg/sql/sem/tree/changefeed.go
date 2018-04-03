// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package tree

// CreateChangefeed represents a CREATE CHANGEFEED statement.
type CreateChangefeed struct {
	Targets  TargetList
	SinkType string
	AsOf     AsOfClause
	Options  KVOptions
}

var _ Statement = &CreateChangefeed{}

// Format implements the NodeFormatter interface.
func (node *CreateChangefeed) Format(ctx *FmtCtx) {
	ctx.WriteString("CREATE EXPERIMENTAL_CHANGEFEED EMIT ")
	ctx.FormatNode(&node.Targets)
	ctx.WriteString(" TO ")
	ctx.WriteString(node.SinkType)
	if node.AsOf.Expr != nil {
		ctx.WriteString(" ")
		ctx.FormatNode(&node.AsOf)
	}
	if node.Options != nil {
		ctx.WriteString(" WITH ")
		ctx.FormatNode(&node.Options)
	}
}
